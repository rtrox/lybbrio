package task

import (
	"context"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/task_enums"

	"github.com/rs/zerolog/log"
)

// TODO: Logger Context

var taskMap = taskMapType{
	task_enums.TypeNoOp: NoOpTask,
}

type taskMapType map[task_enums.TaskType]TaskFunc

func (t taskMapType) Get(taskType task_enums.TaskType) TaskFunc {
	ret, ok := t[taskType]
	if !ok {
		return nil
	}
	return ret
}

func NoOpTask(ctx context.Context, task *ent.Task, client *ent.Client) (string, error) {
	log := log.Ctx(ctx)
	log.Info().Interface("task", task).Msg("NoOpTask")
	return "done", nil
}

type TaskFunc func(ctx context.Context, task *ent.Task, client *ent.Client) (msg string, err error)

type TaskWrapper struct {
	Task *ent.Task
	Func TaskFunc
}
