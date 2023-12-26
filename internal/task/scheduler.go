package task

import (
	"context"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/task_enums"
	enttask "lybbrio/internal/ent/task"
	"time"

	"github.com/rs/zerolog/log"
)

type SchedulerConfig struct {
	Ctx       context.Context
	Cadence   time.Duration
	WorkQueue chan<- TaskWrapper
}

type Scheduler struct {
	client    *ent.Client
	ctx       context.Context
	workQueue chan<- TaskWrapper
	cadence   time.Duration
}

func NewScheduler(client *ent.Client, config *SchedulerConfig) *Scheduler {
	schedulerCtx := log.With().Str("component", "scheduler").Logger().WithContext(config.Ctx)
	return &Scheduler{
		client:    client,
		ctx:       schedulerCtx,
		workQueue: config.WorkQueue,
		cadence:   config.Cadence,
	}
}

func (s *Scheduler) Schedule(ctx context.Context) error {
	log := log.Ctx(ctx)
	tasks, err := s.client.Task.Query().
		Where(enttask.StatusEQ(
			task_enums.Status(task_enums.StatusPending),
		)).
		All(ctx)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		taskFunc := taskMap.Get(task.Type)
		if taskFunc == nil {
			log.Error().Interface("task", task).Msg("Unimplemented task type requested")
			task.Status = task_enums.StatusFailure
			task.Error = "task type not implemented"
			if _, err := s.client.Task.UpdateOne(task).Save(ctx); err != nil {
				log.Error().Err(err).Interface("task", task).Msg("Failed to update task")
			}
			continue
		}
		log.Debug().Str("task", task.ID.String()).Msg("Scheduling task")
		s.workQueue <- TaskWrapper{Task: task, Func: taskFunc}
	}
	return nil
}

func (s *Scheduler) Start() {
	log := log.Ctx(s.ctx)
	ticker := time.NewTicker(s.cadence)
	go func() {
		for {
			select {
			case <-s.ctx.Done():
				log.Info().Msg("Scheduler exiting")
				return
			case <-ticker.C:
				log.Trace().Msg("Scheduling tasks")
				if err := s.Schedule(s.ctx); err != nil {
					log.Error().Err(err).Msg("Failed to schedule tasks")
				}
			}
		}
	}()
}
