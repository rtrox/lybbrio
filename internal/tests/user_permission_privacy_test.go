package test

import (
	"context"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/userpermissions"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CreateUserPermissionsRequiresAdmin(t *testing.T) {
	tests := []struct {
		name           string
		creatorContext func(testData) context.Context
		shouldCreate   bool
	}{
		{
			name:           "user1 creates user permissions",
			creatorContext: func(data testData) context.Context { return data.user1ViewerContext },
			shouldCreate:   false,
		},
		{
			name:           "user2 creates user permissions",
			creatorContext: func(data testData) context.Context { return data.user2ViewerContext },
			shouldCreate:   false,
		},
		{
			name:           "admin creates user permissions",
			creatorContext: func(data testData) context.Context { return data.adminViewerContext },
			shouldCreate:   true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			teardown, client, data := setupTest(t, tt.name)
			defer teardown(t)
			_, err := client.UserPermissions.Create().
				SetAdmin(true).
				Save(tt.creatorContext(data))
			if tt.shouldCreate {
				require.NoError(t, err, "user_permissions creation should have succeeded")
			} else {
				require.Error(t, err, "user_permissions creation should have failed")
				require.Contains(t, err.Error(), "ent/privacy", "user_permissions creation should throw privacy error")
			}
		})
	}
}

func Test_UpdateUserPermissionsRequiresAdmin(t *testing.T) {
	tests := []struct {
		name           string
		updaterContext func(testData) context.Context
		shouldUpdate   bool
	}{
		{
			name:           "user1 updates user permissions",
			updaterContext: func(data testData) context.Context { return data.user1ViewerContext },
			shouldUpdate:   false,
		},
		{
			name:           "user2 updates user permissions",
			updaterContext: func(data testData) context.Context { return data.user2ViewerContext },
			shouldUpdate:   false,
		},
		{
			name:           "admin updates user permissions",
			updaterContext: func(data testData) context.Context { return data.adminViewerContext },
			shouldUpdate:   true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			teardown, client, data := setupTest(t, tt.name)
			defer teardown(t)
			userPermissions, err := client.UserPermissions.Create().
				SetAdmin(false).
				Save(data.adminViewerContext)
			require.NoError(t, err, "failed to create user_permissions")
			_, err = userPermissions.Update().
				SetAdmin(true).
				Save(tt.updaterContext(data))
			if tt.shouldUpdate {
				require.NoError(t, err, "user_permissions update should have succeeded")
			} else {
				require.Error(t, err, "user_permissions update should have failed")
				require.Contains(t, err.Error(), "ent/privacy", "user_permissions update should throw privacy error")
			}
		})
	}
}

func Test_ViewUserPermissionsRequiresUserFilter(t *testing.T) {
	tests := []struct {
		name           string
		viewerContext  func(testData) context.Context
		permissionUser ksuid.ID
		permissionID   ksuid.ID
		shouldView     bool
	}{
		{
			name:           "user1 views user permissions for user1",
			viewerContext:  func(data testData) context.Context { return data.user1ViewerContext },
			permissionUser: USER_ID_1,
			permissionID:   USER_PERM_1,
			shouldView:     true,
		},
		{
			name:           "user1 views user permissions for user2",
			viewerContext:  func(data testData) context.Context { return data.user1ViewerContext },
			permissionUser: USER_ID_2,
			permissionID:   USER_PERM_2,
			shouldView:     false,
		},
		{
			name:           "user2 views user permissions for user2",
			viewerContext:  func(data testData) context.Context { return data.user2ViewerContext },
			permissionUser: USER_ID_2,
			permissionID:   USER_PERM_2,
			shouldView:     true,
		},
		{
			name:           "user2 views user permissions for user1",
			viewerContext:  func(data testData) context.Context { return data.user2ViewerContext },
			permissionUser: USER_ID_1,
			permissionID:   USER_PERM_1,
			shouldView:     false,
		},
		{
			name:           "admin views user permissions for user1",
			viewerContext:  func(data testData) context.Context { return data.adminViewerContext },
			permissionUser: USER_ID_1,
			permissionID:   USER_PERM_1,
			shouldView:     true,
		},
		{
			name:           "admin views user permissions for user2",
			viewerContext:  func(data testData) context.Context { return data.adminViewerContext },
			permissionUser: USER_ID_2,
			permissionID:   USER_PERM_2,
			shouldView:     true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			teardown, client, data := setupTest(t, tt.name)
			defer teardown(t)

			// By User
			_, err := client.UserPermissions.Query().
				Where(userpermissions.UserID(tt.permissionUser)).
				First(tt.viewerContext(data))
			if tt.shouldView {
				require.NoError(t, err, "user_permissions view should have succeeded")
			} else {
				require.Error(t, err, "user_permissions view should have failed")
				require.Contains(t, err.Error(), "not found", "user_permissions view should throw not found")
			}
			// By ID
			_, err = client.UserPermissions.Query().
				Where(userpermissions.ID(tt.permissionID)).
				First(tt.viewerContext(data))
			if tt.shouldView {
				require.NoError(t, err, "user_permissions view should have succeeded")
			} else {
				require.Error(t, err, "user_permissions view should have failed")
				require.Contains(t, err.Error(), "not found", "user_permissions view should throw not found")
			}
		})
	}
}
