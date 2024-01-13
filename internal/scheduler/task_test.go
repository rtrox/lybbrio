package scheduler

import (
	"context"
	"errors"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestTaskMap(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tm := TaskMap{
		"test": NoOpTask,
	}
	require.NotNil(tm.Get("test"), "Task not registered")
	require.Nil(tm.Get("test2"), "Task registered when it shouldn't be")
}

func TestConcurrentTaskMap(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tm := newConcurrentTaskMap()

	require.NotPanics(func() {
		tm.RegisterTasks(TaskMap{
			"test": NoOpTask,
		})
	})
	require.NotNil(tm.Get("test"), "Task not registered")
	require.Nil(tm.Get("test2"), "Task registered when it shouldn't be")

}

func Test_NoOpTask(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	logger := zerolog.New(nil)
	ctx := logger.WithContext(context.Background())
	m, err := NoOpTask(ctx, nil, func(progress float64) error {
		require.Equal(1.0, progress)
		return nil
	})
	require.NoError(err)
	require.NotEmpty(m)

	expectedErr := errors.New("")
	_, err2 := NoOpTask(ctx, nil, func(progress float64) error {
		return expectedErr
	})
	require.Error(err2)
	require.Equal(expectedErr, err2)
}
