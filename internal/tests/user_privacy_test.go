package test

import (
	"context"
	"fmt"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/viewer"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_UserCannotMutateOtherUsers(t *testing.T) {
	tests := []struct {
		name           string
		updaterContext func(testData) context.Context
		updatedID      ksuid.ID
		shouldUpdate   bool
	}{
		{
			name:           "user1 updates user2",
			updaterContext: func(data testData) context.Context { return data.user1ViewerContext },
			updatedID:      USER_ID_2,
			shouldUpdate:   false,
		},
		{
			name:           "user1 updates user1",
			updaterContext: func(data testData) context.Context { return data.user1ViewerContext },
			updatedID:      USER_ID_1,
			shouldUpdate:   true,
		},
		{
			name:           "user2 updates user1",
			updaterContext: func(data testData) context.Context { return data.user2ViewerContext },
			updatedID:      USER_ID_1,
			shouldUpdate:   false,
		},
		{
			name:           "user2 updates user2",
			updaterContext: func(data testData) context.Context { return data.user2ViewerContext },
			updatedID:      USER_ID_2,
			shouldUpdate:   true,
		},
		{
			name:           "admin updates user1",
			updaterContext: func(data testData) context.Context { return data.adminViewerContext },
			updatedID:      USER_ID_1,
			shouldUpdate:   true,
		},
		{
			name:           "admin updates user2",
			updaterContext: func(data testData) context.Context { return data.adminViewerContext },
			updatedID:      USER_ID_2,
			shouldUpdate:   true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			teardown, client, data := setupTest(t, tt.name)
			defer teardown(t)

			_, err := client.User.UpdateOneID(tt.updatedID).
				SetUsername("newName").
				Save(tt.updaterContext(data))
			if tt.shouldUpdate {
				require.NoError(t, err, "failed to update user")
			} else {
				require.Error(t, err, "expected error updating user")
			}

		})
	}
}

func Test_UsersCanSeeEachOther(t *testing.T) {
	tests := []struct {
		name          string
		viewerContext func(testData) context.Context
		expectedCount int
	}{
		{
			name:          "user1 views users",
			viewerContext: func(data testData) context.Context { return data.user1ViewerContext },
			expectedCount: 3,
		},
		{
			name:          "user2 views users",
			viewerContext: func(data testData) context.Context { return data.user2ViewerContext },
			expectedCount: 3,
		},
		{
			name:          "admin views users",
			viewerContext: func(data testData) context.Context { return data.adminViewerContext },
			expectedCount: 3,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			teardown, client, data := setupTest(t, tt.name)
			defer teardown(t)

			users, err := client.User.Query().
				All(tt.viewerContext(data))
			require.NoError(t, err, "failed to query users")
			require.Len(t, users, tt.expectedCount, "incorrect number of users")
		})

	}
}

func Test_AnonymousUserCanCreateUser(t *testing.T) {
	t.Parallel()
	teardown, client, _ := setupTest(t, "anonymous user can create user")
	defer teardown(t)

	ctx := viewer.NewAnonymousContext(context.Background())
	pc := client.UserPermissions.Create()
	fmt.Printf("%+v", pc.Mutation())
	perms, err := pc.Save(ctx)
	require.NoError(t, err, "failed to create user permissions")

	_, err = client.User.Create().
		SetUsername("newUser").
		SetEmail("asdf@asdf.com").
		SetUserPermissions(perms).
		Save(ctx)
	require.NoError(t, err, "failed to create user")
}

func Test_AnonymousUserCannotCreateUserWithPermissions(t *testing.T) {
	for p := range permissions.All() {
		p := p
		t.Run(p.String(), func(t *testing.T) {
			t.Parallel()
			teardown, client, _ := setupTest(t, "anonymous user cannot create user with "+p.String())
			defer teardown(t)
			c := client.UserPermissions.Create()
			err := c.Mutation().SetField(p.FieldName(), true)
			require.NoError(t, err, "failed to set field "+p.FieldName())
			_, err = c.Save(viewer.NewAnonymousContext(context.Background()))
			require.Error(t, err, "expected error creating user with permission "+p.String())
		})
	}
}
