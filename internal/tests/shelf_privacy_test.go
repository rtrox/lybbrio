package test

import (
	"context"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/middleware"
	"lybbrio/internal/viewer"
	"testing"

	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

func Test_CreateShelfRespectsUserFilter(t *testing.T) {
	tests := []struct {
		name           string
		creatorContext func(testData) context.Context
		shelfOwner     ksuid.ID
		shouldCreate   bool
	}{
		{
			name:           "user1 creates shelf for user1",
			creatorContext: func(data testData) context.Context { return data.user1ViewerContext },
			shelfOwner:     USER_ID_1,
			shouldCreate:   true,
		},
		{
			name:           "user1 creates shelf for user2",
			creatorContext: func(data testData) context.Context { return data.user1ViewerContext },
			shelfOwner:     USER_ID_2,
			shouldCreate:   false,
		},
		{
			name:           "user2 creates shelf for user2",
			creatorContext: func(data testData) context.Context { return data.user2ViewerContext },
			shelfOwner:     USER_ID_2,
			shouldCreate:   true,
		},
		{
			name:           "user2 creates shelf for user1",
			creatorContext: func(data testData) context.Context { return data.user2ViewerContext },
			shelfOwner:     USER_ID_1,
			shouldCreate:   false,
		},
		{
			name:           "admin creates shelf for user1",
			creatorContext: func(data testData) context.Context { return data.adminViewerContext },
			shelfOwner:     USER_ID_1,
			shouldCreate:   true,
		},
		{
			name:           "admin creates shelf for user2",
			creatorContext: func(data testData) context.Context { return data.adminViewerContext },
			shelfOwner:     USER_ID_2,
			shouldCreate:   true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			teardown, client, data := setupTest(t, tt.name)
			defer teardown(t)
			_, err := client.Shelf.Create().
				SetUserID(tt.shelfOwner).
				SetName("test").
				Save(tt.creatorContext(data))
			if tt.shouldCreate {
				require.NoError(t, err, "failed to create shelf")
			} else {
				require.Error(t, err, "shelf creation should have failed")
			}
		})
	}
}

func Test_ShelfViewRespectsUserFilter(t *testing.T) {
	tests := []struct {
		name       string
		viewer     func(testData) context.Context
		shelfCount int
	}{
		{
			name:       "user1",
			viewer:     func(data testData) context.Context { return data.user1ViewerContext },
			shelfCount: 3,
		},
		{
			name:       "user2",
			viewer:     func(data testData) context.Context { return data.user2ViewerContext },
			shelfCount: 4,
		},
		{
			name:       "admin",
			viewer:     func(data testData) context.Context { return data.adminViewerContext },
			shelfCount: 2,
		},
		{
			name: "admin with super read",
			viewer: func(data testData) context.Context {
				return middleware.WithSuperRead(data.adminViewerContext)
			},
			shelfCount: 5,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			teardown, client, data := setupTest(t, tt.name)
			defer teardown(t)
			client.Shelf.Create().
				SetUserID(USER_ID_1).
				SetName("test").
				SaveX(data.adminViewerContext)
			client.Shelf.Create().
				SetUserID(USER_ID_1).
				SetName("test2").
				SetPublic(true).
				SaveX(data.adminViewerContext)
			client.Shelf.Create().
				SetUserID(USER_ID_2).
				SetName("test").
				SaveX(data.adminViewerContext)
			client.Shelf.Create().
				SetUserID(USER_ID_2).
				SetName("test2").
				SetPublic(true).
				SaveX(data.adminViewerContext)
			client.Shelf.Create().
				SetUserID(USER_ID_2).
				SetName("test3").
				SaveX(data.adminViewerContext)

			shelves, err := client.Shelf.Query().All(tt.viewer(data))
			require.NoError(t, err, "failed to query shelves")
			require.Len(t, shelves, tt.shelfCount, "incorrect number of shelves")

			view := viewer.FromContext(tt.viewer(data))
			if !view.IsAdmin() {
				uid, ok := view.UserID()
				for _, shelf := range shelves {
					require.True(t, ok, "viewer should have a user")
					if !shelf.Public {
						require.Equal(t, uid, shelf.UserID, "user should not be able to see shelves belonging to other users")
					}
				}
			}
		})
	}
}

func Test_CreatePublicShelfRespectsPermissionsRule(t *testing.T) {
	tests := []struct {
		name           string
		creatorContext func(testData) context.Context
		shouldCreate   bool
	}{
		{
			name:           "user1 creates public shelf",
			creatorContext: func(data testData) context.Context { return data.user1ViewerContext },
			shouldCreate:   false,
		},
		{
			name:           "user2 creates public shelf",
			creatorContext: func(data testData) context.Context { return data.user2ViewerContext },
			shouldCreate:   true,
		},
		{
			name:           "admin creates public shelf",
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

			ctx := tt.creatorContext(data)
			uid, _ := viewer.FromContext(ctx).UserID()

			_, err := client.Shelf.Create().
				SetUserID(uid).
				SetName("test").
				SetPublic(true).
				Save(ctx)
			if tt.shouldCreate {
				require.NoError(t, err, "failed to create shelf")
			} else {
				require.Error(t, err, "shelf creation should have failed")
			}
		})
	}
}
