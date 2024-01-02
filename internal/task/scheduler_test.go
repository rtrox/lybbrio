package task

import (
	"context"
	"lybbrio/internal/db"
	"lybbrio/internal/ent/schema/task_enums"
	enttask "lybbrio/internal/ent/task"
	"lybbrio/internal/viewer"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewScheduler(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	client := db.OpenTest(t, "TestNewScheduler")
	ctx := context.Background()
	ch := make(chan TaskWrapper)
	s := NewScheduler(client, &SchedulerConfig{
		Ctx:       ctx,
		Cadence:   1 * time.Second,
		WorkQueue: ch,
	})

	require.NotNil(s)

	var expected chan<- TaskWrapper = ch
	require.Equal(expected, s.workQueue, "Scheduler has wrong work queue")
	require.Equal(1*time.Second, s.cadence, "Scheduler has wrong cadence")
	require.NotNil(s.ctx, "Scheduler has nil context")

	require.NotPanics(func() {
		require.NotNil(s.taskMap.Get(task_enums.TypeNoOp), "Task not registered")
	})

}

func TestSchedulerSchedule(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	client := db.OpenTest(t, "TestSchedulerSchedule")
	ctx, cancel := context.WithTimeout(
		viewer.NewSystemAdminContext(context.Background()),
		2*time.Second,
	)
	defer cancel()

	ch := make(chan TaskWrapper, 1)
	s := NewScheduler(client, &SchedulerConfig{
		Ctx:       ctx,
		Cadence:   1 * time.Second,
		WorkQueue: ch,
	})

	require.NotNil(s)

	require.NotPanics(func() {
		require.NotNil(s.taskMap.Get(task_enums.TypeNoOp), "Task not registered")
	})

	task, err := client.Task.Create().
		SetType(task_enums.TypeNoOp).
		Save(ctx)
	require.NoError(err)

	require.NotPanics(func() {
		s.Schedule(ctx)
	})

	select {
	case <-ctx.Done():
		require.Fail("Scheduler did not schedule task")
	case tw := <-ch:
		require.Equal(task.ID, tw.Task.ID, "Scheduler scheduled wrong task")
		require.NotNil(tw.Func, "Scheduler scheduled nil task function")
	}
}

func TestSchedulerFailsUnknownTask(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	client := db.OpenTest(t, "TestSchedulerFailsUnknownTask")
	ctx, cancel := context.WithTimeout(
		viewer.NewSystemAdminContext(context.Background()),
		2*time.Second,
	)
	defer cancel()

	ch := make(chan TaskWrapper, 1)
	s := NewScheduler(client, &SchedulerConfig{
		Ctx:       ctx,
		Cadence:   1 * time.Second,
		WorkQueue: ch,
	})

	require.NotNil(s)

	require.NotPanics(func() {
		require.NotNil(s.taskMap.Get(task_enums.TypeNoOp), "Task not registered")
	})

	task, err := client.Task.Create().
		SetType(task_enums.TypeCalibreImport). // Not Registered in this test
		SetStatus(task_enums.StatusPending).
		Save(ctx)
	require.NoError(err)

	err = s.Schedule(ctx)
	require.NoError(err)

	updatedTask, err := client.Task.Query().Where(enttask.IDEQ(task.ID)).Only(ctx)
	require.NoError(err)

	require.Equal(task_enums.StatusFailure, updatedTask.Status, "Task status is not failure")

}

func TestSchedulerStart(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	client := db.OpenTest(t, "TestSchedulerStart")
	ctx, cancel := context.WithTimeout(
		viewer.NewSystemAdminContext(context.Background()),
		100*time.Second,
	)
	defer cancel()

	ch := make(chan TaskWrapper, 1)
	s := NewScheduler(client, &SchedulerConfig{
		Ctx:       ctx,
		Cadence:   1 * time.Millisecond,
		WorkQueue: ch,
	})

	require.NotNil(s)

	require.NotPanics(func() {
		require.NotNil(s.taskMap.Get(task_enums.TypeNoOp), "Task not registered")
	})

	task, err := client.Task.Create().
		SetType(task_enums.TypeNoOp).
		SetStatus(task_enums.StatusPending).
		Save(ctx)
	require.NoError(err)

	s.Start()

	select {
	case <-ctx.Done():
		require.Fail("Scheduler did not schedule task")
	case tw := <-ch:
		require.Equal(task.ID, tw.Task.ID, "Scheduler scheduled wrong task")
		require.NotNil(tw.Func, "Scheduler scheduled nil task function")
	}
}
