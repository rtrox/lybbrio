package task

import (
	"context"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/task_enums"
	"sync"

	"github.com/rs/zerolog/log"
)

type ProgressCallback func(progress float64) error
type TaskFunc func(ctx context.Context, task *ent.Task, cb ProgressCallback) (msg string, err error)

// TODO: TaskWrapper is opaque, hiding the ent.Task. Give it an interface to allow easier testing.
type TaskWrapper struct {
	Task *ent.Task
	Func TaskFunc
}

type TaskMap map[task_enums.TaskType]TaskFunc

func (t TaskMap) Get(taskType task_enums.TaskType) TaskFunc {
	ret, ok := t[taskType]
	if !ok {
		return nil
	}
	return ret
}

type concurrentTaskMap struct {
	m   TaskMap
	mut sync.RWMutex
}

func newConcurrentTaskMap() concurrentTaskMap {
	return concurrentTaskMap{
		m: make(TaskMap),
	}
}

func (t *concurrentTaskMap) Get(taskType task_enums.TaskType) TaskFunc {
	t.mut.RLock()
	defer t.mut.RUnlock()
	return t.m[taskType]
}

func (t *concurrentTaskMap) RegisterTasks(taskMap TaskMap) {
	t.mut.Lock()
	defer t.mut.Unlock()
	for k, v := range taskMap {
		t.m[k] = v
	}
}

func NoOpTask(ctx context.Context, task *ent.Task, cb ProgressCallback) (msg string, err error) {
	log := log.Ctx(ctx)
	log.Info().Interface("task", task).Msg("NoOpTask")
	cb(1)
	return "done", nil
}
