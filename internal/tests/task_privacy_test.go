package test

import (
	"context"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/schema/task_enums"
	"lybbrio/internal/viewer"
	"testing"

	"github.com/stretchr/testify/require"
)

// TODO: Privacy tests for tasks
func Test_CreateSystemTasks(t *testing.T) {
	tests := []struct {
		name           string
		creatorContext func(testData) context.Context
		shouldCreate   bool
	}{
		{
			name:           "user1 creates system task",
			creatorContext: func(data testData) context.Context { return data.user1ViewerContext },
			shouldCreate:   false,
		},
		{
			name:           "user2 creates system task",
			creatorContext: func(data testData) context.Context { return data.user2ViewerContext },
			shouldCreate:   false,
		},
		{
			name:           "admin creates system task",
			creatorContext: func(data testData) context.Context { return data.adminViewerContext },
			shouldCreate:   true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			tearDown, client, data := setupTest(t, tt.name)
			defer tearDown(t)
			_, err := client.Task.Create().
				SetType(task_enums.TypeNoOp).
				SetIsSystemTask(true).
				Save(tt.creatorContext(data))
			if tt.shouldCreate {
				require.NoError(err, "failed to create system task")
			} else {
				require.Error(err, "created system task")
			}
		})
	}
}

func Test_CreateTaskOfType(t *testing.T) {
	tests := []struct {
		name           string
		creatorContext func(testData) context.Context
		taskType       task_enums.TaskType
		shouldCreate   bool
	}{
		{
			name:           "user1 creates noop task",
			creatorContext: func(data testData) context.Context { return data.user1ViewerContext },
			taskType:       task_enums.TypeNoOp,
			shouldCreate:   true,
		},
		{
			name:           "user2 creates noop task",
			creatorContext: func(data testData) context.Context { return data.user2ViewerContext },
			taskType:       task_enums.TypeNoOp,
			shouldCreate:   true,
		},
		{
			name:           "admin creates noop task",
			creatorContext: func(data testData) context.Context { return data.adminViewerContext },
			taskType:       task_enums.TypeNoOp,
			shouldCreate:   true,
		},
		{
			name:           "user1 creates import task",
			creatorContext: func(data testData) context.Context { return data.user1ViewerContext },
			taskType:       task_enums.TypeCalibreImport,
			shouldCreate:   false,
		},
		{
			name:           "user2 creates import task",
			creatorContext: func(data testData) context.Context { return data.user2ViewerContext },
			taskType:       task_enums.TypeCalibreImport,
			shouldCreate:   false,
		},
		{
			name:           "admin creates import task",
			creatorContext: func(data testData) context.Context { return data.adminViewerContext },
			taskType:       task_enums.TypeCalibreImport,
			shouldCreate:   true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			tearDown, client, data := setupTest(t, tt.name)
			defer tearDown(t)
			_, err := client.Task.Create().
				SetType(tt.taskType).
				Save(tt.creatorContext(data))
			if tt.shouldCreate {
				require.NoError(err, "failed to create task")
			} else {
				require.Error(err, "created task when it should not have been created")
			}
		})
	}
}

func Test_UpdateSystemTask(t *testing.T) {
	tests := []struct {
		name           string
		updaterContext func(testData) context.Context
		shouldUpdate   bool
	}{
		{
			name:           "user1 updates system task",
			updaterContext: func(data testData) context.Context { return data.user1ViewerContext },
			shouldUpdate:   false,
		},
		{
			name:           "user2 updates system task",
			updaterContext: func(data testData) context.Context { return data.user2ViewerContext },
			shouldUpdate:   false,
		},
		{
			name:           "admin updates system task",
			updaterContext: func(data testData) context.Context { return data.adminViewerContext },
			shouldUpdate:   true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			tearDown, client, data := setupTest(t, tt.name)
			defer tearDown(t)
			adminCtx := viewer.NewSystemAdminContext(context.Background())
			task := client.Task.Create().
				SetType(task_enums.TypeNoOp).
				SetIsSystemTask(true).
				SaveX(adminCtx)
			_, err := task.Update().
				SetProgress(0.5).
				Save(tt.updaterContext(data))
			if tt.shouldUpdate {
				require.NoError(err, "failed to update system task")
			} else {
				require.Error(err, "updated system task")
			}
		})
	}

}

func Test_UpdateTaskOfType(t *testing.T) {
	tests := []struct {
		name           string
		updaterContext func(testData) context.Context
		taskType       task_enums.TaskType
		shouldUpdate   bool
	}{
		{
			name:           "user1 updates noop task",
			updaterContext: func(data testData) context.Context { return data.user1ViewerContext },
			taskType:       task_enums.TypeNoOp,
			shouldUpdate:   true,
		},
		{
			name:           "user2 updates noop task",
			updaterContext: func(data testData) context.Context { return data.user2ViewerContext },
			taskType:       task_enums.TypeNoOp,
			shouldUpdate:   true,
		},
		{
			name:           "admin updates noop task",
			updaterContext: func(data testData) context.Context { return data.adminViewerContext },
			taskType:       task_enums.TypeNoOp,
			shouldUpdate:   true,
		},
		{
			name:           "user1 updates import task",
			updaterContext: func(data testData) context.Context { return data.user1ViewerContext },
			taskType:       task_enums.TypeCalibreImport,
			shouldUpdate:   false,
		},
		{
			name:           "user2 updates import task",
			updaterContext: func(data testData) context.Context { return data.user2ViewerContext },
			taskType:       task_enums.TypeCalibreImport,
			shouldUpdate:   false,
		},
		{
			name:           "admin updates import task",
			updaterContext: func(data testData) context.Context { return data.adminViewerContext },
			taskType:       task_enums.TypeCalibreImport,
			shouldUpdate:   true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			tearDown, client, data := setupTest(t, tt.name)
			defer tearDown(t)
			adminCtx := viewer.NewSystemAdminContext(context.Background())
			updaterCtx := tt.updaterContext(data)
			userID, _ := viewer.FromContext(updaterCtx).UserID()

			task := client.Task.Create().
				SetType(tt.taskType).
				SetUserID(userID).
				SaveX(adminCtx)
			_, err := task.Update().
				SetProgress(0.5).
				Save(tt.updaterContext(data))
			if tt.shouldUpdate {
				require.NoError(err, "failed to update task")
			} else {
				require.Error(err, "updated task when it should not have been updated")
			}
		})
	}
}

func Test_UsersCanOnlyUpdateOwnTasks(t *testing.T) {
	tests := []struct {
		name           string
		updaterContext func(testData) context.Context
		taskOwner      ksuid.ID
		shouldUpdate   bool
	}{
		{
			name:           "user1 updates user1 task",
			updaterContext: func(data testData) context.Context { return data.user1ViewerContext },
			taskOwner:      USER_ID_1,
			shouldUpdate:   true,
		},
		{
			name:           "user1 updates user2 task",
			updaterContext: func(data testData) context.Context { return data.user1ViewerContext },
			taskOwner:      USER_ID_2,
			shouldUpdate:   false,
		},
		{
			name:           "user2 updates user2 task",
			updaterContext: func(data testData) context.Context { return data.user2ViewerContext },
			taskOwner:      USER_ID_2,
			shouldUpdate:   true,
		},
		{
			name:           "user2 updates user1 task",
			updaterContext: func(data testData) context.Context { return data.user2ViewerContext },
			taskOwner:      USER_ID_1,
			shouldUpdate:   false,
		},
		{
			name:           "admin updates user1 task",
			updaterContext: func(data testData) context.Context { return data.adminViewerContext },
			taskOwner:      USER_ID_1,
			shouldUpdate:   true,
		},
		{
			name:           "admin updates user2 task",
			updaterContext: func(data testData) context.Context { return data.adminViewerContext },
			taskOwner:      USER_ID_2,
			shouldUpdate:   true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			tearDown, client, data := setupTest(t, tt.name)
			defer tearDown(t)
			adminCtx := viewer.NewSystemAdminContext(context.Background())
			task := client.Task.Create().
				SetType(task_enums.TypeNoOp).
				SetUserID(tt.taskOwner).
				SaveX(adminCtx)
			_, err := task.Update().
				SetProgress(0.5).
				Save(tt.updaterContext(data))
			if tt.shouldUpdate {
				require.NoError(err, "failed to update task")
			} else {
				require.Error(err, "updated task when it should not have been updated")
			}
		})
	}
}

func Test_ViewTask(t *testing.T) {
	tests := []struct {
		name      string
		viewer    func(testData) context.Context
		taskCount int
	}{
		{
			name: "user1",
			viewer: func(data testData) context.Context {
				return data.user1ViewerContext
			},
			taskCount: 2,
		},
		{
			name: "user2",
			viewer: func(data testData) context.Context {
				return data.user2ViewerContext
			},
			taskCount: 2,
		},
		{
			name: "admin",
			viewer: func(data testData) context.Context {
				return data.adminViewerContext
			},
			taskCount: 4,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			tearDown, client, data := setupTest(t, tt.name)
			adminCtx := viewer.NewSystemAdminContext(context.Background())
			defer tearDown(t)
			client.Task.Create().
				SetUserID(USER_ID_1).
				SaveX(adminCtx)
			client.Task.Create().
				SetUserID(USER_ID_2).
				SaveX(adminCtx)
			client.Task.Create().
				SetUserID(ADMIN_ID).
				SaveX(adminCtx)
			client.Task.Create().
				SetIsSystemTask(true).
				SaveX(adminCtx)

			tasks, err := client.Task.Query().
				All(tt.viewer(data))
			require.NoError(err, "failed to query tasks")
			require.Equal(tt.taskCount, len(tasks), "wrong number of tasks returned")

			view := viewer.FromContext(tt.viewer(data))
			uid, _ := view.UserID()
			for _, task := range tasks {
				if !view.IsAdmin() && !task.IsSystemTask {
					require.Equal(uid, task.UserID, "task does not belong to user")
				}
			}
		})
	}
}
