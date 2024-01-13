package scheduler

import (
	"context"
	"fmt"
	"lybbrio/internal/db"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/enttest"
	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/ent/schema/task_enums"
	"lybbrio/internal/viewer"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewWorkerPool(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	client := db.OpenTest(t, "TestNewWorkerPool")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wp := NewWorkerPool(client, &WorkerPoolConfig{
		NumWorkers:  2,
		QueueLength: 10,
		Ctx:         ctx,
	})
	require.NotNil(wp, "WorkerPool is nil")
	require.NotNil(wp.ch, "WorkerPool channel is nil")
	require.NotNil(wp.client, "WorkerPool client is nil")
	require.NotNil(wp.ctx, "WorkerPool ctx is nil")
	for _, worker := range wp.workers {
		require.NotNil(worker, "WorkerPool worker is nil")
		require.NotNil(worker.ch, "WorkerPool worker channel is nil")
		require.NotNil(worker.client, "WorkerPool worker client is nil")
		require.NotNil(worker.ctx, "WorkerPool worker ctx is nil")
	}

	require.Equal(2, len(wp.workers), "WorkerPool has wrong number of workers")
	require.Equal(10, cap(wp.ch), "WorkerPool has wrong channel capacity")

	var expected chan<- TaskWrapper = wp.ch
	require.Equal(expected, wp.WorkQueue(), "WorkerPool WorkQueue returns wrong channel")
}

func TestNewWorker(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	client := enttest.Open(t, "sqlite3", "file:TestNewWorker?mode=memory&cache=shared&_fk=1")
	worker := NewWorker(make(chan TaskWrapper), client, context.Background(), 0)
	require.NotNil(worker, "Worker is nil")
	require.NotNil(worker.ch, "Worker channel is nil")
	require.NotNil(worker.client, "Worker client is nil")
	require.NotNil(worker.ctx, "Worker ctx is nil")
}

func TestWorkerPoolRunsTask(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	client := db.OpenTest(t, "TestWorkerPoolRunsTask")

	ctx, cancel := context.WithTimeout(
		viewer.NewSystemAdminContext(context.Background()),
		1*time.Second,
	)
	defer cancel()

	wp := NewWorkerPool(client, &WorkerPoolConfig{
		NumWorkers:  1,
		QueueLength: 10,
		Ctx:         ctx,
	})
	wp.Start()

	task := client.Task.Create().
		SetType(task_enums.TypeNoOp).
		SetIsSystemTask(true).
		SaveX(ctx)

	done := make(chan struct{})
	wp.WorkQueue() <- TaskWrapper{
		Task: task,
		Func: func(ctx context.Context, task *ent.Task, cb ProgressCallback) (msg string, err error) {
			close(done)
			return "done", nil
		},
	}

	select {
	case <-done:
		time.Sleep(100 * time.Millisecond)
	case <-ctx.Done():
		require.Fail("Task did not complete in time")
	}

	task = client.Task.GetX(ctx, task.ID)
	require.Equal(task_enums.StatusSuccess, task.Status, "Task status is not completed")
}

func TestWorkerPoolRunsAsUser(t *testing.T) {
	t.Parallel()

	client := db.OpenTest(t, "TestWorkerPoolRunsAsUser")

	adminCtx := viewer.NewSystemAdminContext(context.Background())

	userPerms := client.UserPermissions.Create().
		SetAdmin(true).
		SaveX(adminCtx)
	user := client.User.Create().
		SetUserPermissions(userPerms).
		SetUsername("user").
		SetEmail("user@notanemail.com").
		SaveX(adminCtx)

	ctx, cancel := context.WithTimeout(
		viewer.NewContext(context.Background(), user.ID, permissions.From(userPerms)),
		1*time.Second,
	)
	defer cancel()

	wp := NewWorkerPool(client, &WorkerPoolConfig{
		NumWorkers:  1,
		QueueLength: 10,
		Ctx:         adminCtx,
	})
	wp.Start()

	task := client.Task.Create().
		SetType(task_enums.TypeNoOp).
		SetUser(user).
		SaveX(ctx)

	done := make(chan struct{})
	wp.WorkQueue() <- TaskWrapper{
		Task: task,
		Func: func(ctx context.Context, task *ent.Task, cb ProgressCallback) (msg string, err error) {
			close(done)
			view := viewer.FromContext(ctx)
			require.NotNil(t, view, "Viewer is nil")
			taskUserID, ok := view.UserID()
			require.True(t, ok, "Viewer does not have user")
			require.Equal(t, user.ID, taskUserID, "Task viewer context is not the same as the user who created the task")
			return "done", nil
		},
	}

	select {
	case <-done:
		time.Sleep(100 * time.Millisecond)
	case <-ctx.Done():
		require.Fail(t, "Task did not complete in time")
	}

	task = client.Task.GetX(ctx, task.ID)
	require.Equal(t, task_enums.StatusSuccess, task.Status, "Task status is not completed")
}

func TestWorkerPoolHandlesError(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	client := db.OpenTest(t, "TestWorkerPoolHandlesError")

	ctx, cancel := context.WithTimeout(
		viewer.NewSystemAdminContext(context.Background()),
		1*time.Second,
	)
	defer cancel()

	wp := NewWorkerPool(client, &WorkerPoolConfig{
		NumWorkers:  1,
		QueueLength: 10,
		Ctx:         ctx,
	})

	wp.Start()

	task := client.Task.Create().
		SetType(task_enums.TypeNoOp).
		SetIsSystemTask(true).
		SaveX(ctx)

	done := make(chan struct{})
	wp.WorkQueue() <- TaskWrapper{
		Task: task,
		Func: func(ctx context.Context, task *ent.Task, cb ProgressCallback) (msg string, err error) {
			close(done)
			return "error", fmt.Errorf("error")
		},
	}

	select {
	case <-done:
		time.Sleep(100 * time.Millisecond)
	case <-ctx.Done():
		require.Fail("Task did not complete in time")
	}

	task = client.Task.GetX(ctx, task.ID)
	require.Equal(task_enums.StatusFailure, task.Status, "Task status is not failure")
}

func TestTaskFuncCantPanic(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	client := db.OpenTest(t, "TestTaskFuncCantPanic")

	ctx, cancel := context.WithTimeout(
		viewer.NewSystemAdminContext(context.Background()),
		1*time.Second,
	)

	defer cancel()

	wp := NewWorkerPool(client, &WorkerPoolConfig{
		NumWorkers:  1,
		QueueLength: 10,
		Ctx:         ctx,
	})

	wp.Start()

	task := client.Task.Create().
		SetType(task_enums.TypeNoOp).
		SetIsSystemTask(true).
		SaveX(ctx)

	done := make(chan struct{})
	wp.WorkQueue() <- TaskWrapper{
		Task: task,
		Func: func(ctx context.Context, task *ent.Task, cb ProgressCallback) (msg string, err error) {
			close(done)
			panic("I'm panicking!")
		},
	}

	select {
	case <-done:
		time.Sleep(100 * time.Millisecond)
	case <-ctx.Done():
		require.Fail("Task did not complete in time")
	}

	task = client.Task.GetX(ctx, task.ID)
	require.Equal(task_enums.StatusFailure, task.Status, "Task status is not failure")
	require.NotEqual("", task.Error, "Task error is empty")
}
