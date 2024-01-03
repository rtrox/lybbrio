package task

import (
	"context"
	"fmt"
	"lybbrio/internal/ent"
	"lybbrio/internal/viewer"

	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/ent/schema/task_enums"

	"github.com/rs/zerolog/log"
)

type WorkerPoolConfig struct {
	NumWorkers  int
	QueueLength int
	Ctx         context.Context
}

type WorkerPool struct {
	ch      chan TaskWrapper
	workers []*Worker
	client  *ent.Client
	ctx     context.Context
}

func NewWorkerPool(client *ent.Client, config *WorkerPoolConfig) *WorkerPool {
	ctx := log.With().Str("component", "worker-pool").Logger().WithContext(config.Ctx)
	log.Ctx(ctx).Debug().Msg("Creating worker pool")

	ch := make(chan TaskWrapper, config.QueueLength)
	workers := make([]*Worker, config.NumWorkers)
	for i := 0; i < config.NumWorkers; i++ {
		workers[i] = NewWorker(ch, client, ctx, i)
	}

	return &WorkerPool{
		ch:      ch,
		workers: workers,
		client:  client,
		ctx:     ctx,
	}
}

func (wp *WorkerPool) WorkQueue() chan<- TaskWrapper {
	return wp.ch
}

func (wp *WorkerPool) Start() {
	for _, worker := range wp.workers {
		go worker.Start()
	}
}

type Worker struct {
	ch     <-chan TaskWrapper
	idx    int
	client *ent.Client
	ctx    context.Context
}

func NewWorker(ch <-chan TaskWrapper, client *ent.Client, ctx context.Context, idx int) *Worker {
	ctx = log.Ctx(ctx).With().Int("worker", idx).Logger().WithContext(ctx)
	log.Ctx(ctx).Trace().Msg("Creating worker")
	return &Worker{
		ch:     ch,
		idx:    idx,
		client: client,
		ctx:    ctx,
	}
}

func (w *Worker) startTask(ctx context.Context, task TaskWrapper) error {
	ctx = log.Ctx(w.ctx).With().Str("task", task.Task.ID.String()).Logger().WithContext(ctx)
	log := log.Ctx(ctx)
	log.Info().Msg("Starting task")

	var err error
	task.Task, err = task.Task.Update().
		SetStatus(task_enums.StatusInProgress).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update task status: %w", err)
	}

	taskCtx := ctx
	if !task.Task.IsSystemTask {
		user, err := task.Task.QueryUser().First(ctx)
		if err != nil {
			return fmt.Errorf("failed to get task creator: %w", err)
		}
		perms, err := user.QueryUserPermissions().First(ctx)
		if err != nil {
			return fmt.Errorf("failed to get task creator permissions: %w", err)
		}
		taskCtx = viewer.NewContext(ctx, user.ID, permissions.From(perms))
	}

	msg, err := func() (msg string, err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Error().Interface("recover", r).Msg("Recovered from panic")
				msg = "Panic"
				err = fmt.Errorf("panic: %v", r)
			}
		}()
		return task.Func(taskCtx, task.Task, func(progress float64) error {
			var err error
			task.Task, err = task.Task.Update().SetProgress(progress).Save(ctx)
			return err
		})
	}()

	update := task.Task.Update()
	if err != nil {
		log.Error().Err(err).Msg("Task failed")
		update.
			SetStatus(task_enums.StatusFailure).
			SetError(err.Error())
	} else {
		log.Info().Msg("Task succeeded")
		update.
			SetStatus(task_enums.StatusSuccess).
			SetMessage(msg)
	}
	if _, err := update.Save(ctx); err != nil {
		log.Error().Err(err).Msg("Failed to update task")
	}
	return nil
}

func (w *Worker) Start() {
	for {
		log.Ctx(w.ctx).Trace().Int("chanLength", len(w.ch)).Msg("Waiting for task")
		select {
		case task := <-w.ch:
			err := w.startTask(w.ctx, task)
			if err != nil {
				log.Error().Err(err).Msg("Failed to start task")
			}
		case <-w.ctx.Done():
			log.Info().Msg("Worker exiting")
			return
		}
	}
}
