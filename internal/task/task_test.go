package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
