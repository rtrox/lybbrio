package rule

import (
	"context"
	"errors"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/privacy"
	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/ent/schema/task_enums"
	"lybbrio/internal/middleware"
	"lybbrio/internal/viewer"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_DenyIfNoViewer(t *testing.T) {
	tests := []struct {
		name        string
		contextFunc func() context.Context
		wantErr     bool
	}{
		{
			name: "viewer is nil",
			contextFunc: func() context.Context {
				return context.Background()
			},
			wantErr: true,
		},
		{
			name: "viewer is not nil",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "", nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			privacyFunc := DenyIfNoViewer()
			err1 := privacyFunc.EvalQuery(tt.contextFunc(), nil)
			err2 := privacyFunc.EvalMutation(tt.contextFunc(), nil)
			if tt.wantErr {
				require.Error(err1)
				require.Error(err2)
				require.ErrorContains(err1, "viewer-context is missing")
				require.ErrorContains(err2, "viewer-context is missing")
			} else {
				require.Equal(privacy.Skip, err1)
				require.Equal(privacy.Skip, err2)
			}
		})
	}
}

func Test_AllowIfAdmin(t *testing.T) {
	tests := []struct {
		name        string
		contextFunc func() context.Context
		wantErr     bool
	}{
		{
			name: "viewer is not admin",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "", nil)
			},
			wantErr: true,
		},
		{
			name: "viewer is admin",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "", permissions.NewPermissions(permissions.Admin))
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			privacyFunc := AllowIfAdmin()
			err1 := privacyFunc.EvalQuery(tt.contextFunc(), nil)
			err2 := privacyFunc.EvalMutation(tt.contextFunc(), nil)
			if tt.wantErr {
				require.Equal(privacy.Skip, err1)
				require.Equal(privacy.Skip, err2)
			} else {
				require.Equal(privacy.Allow, err1)
				require.Equal(privacy.Allow, err2)
			}
		})
	}
}

func Test_AllowIfSuperRead(t *testing.T) {
	tests := []struct {
		name        string
		contextFunc func() context.Context
		superRead   bool
		wantErr     bool
	}{
		{
			name: "viewer is not admin",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "", nil)
			},
			superRead: true,
			wantErr:   true,
		},
		{
			name: "viewer is admin",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "", permissions.NewPermissions(permissions.Admin))
			},
			superRead: true,
			wantErr:   false,
		},
		{
			name: "viewer is admin but superRead is false",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "", permissions.NewPermissions(permissions.Admin))
			},
			superRead: false,
			wantErr:   true,
		},
		{
			name: "viewer is not admin and superRead is false",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "", nil)
			},
			superRead: false,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			privacyFunc := AllowIfSuperRead()
			ctx := tt.contextFunc()
			if tt.superRead {
				ctx = middleware.WithSuperRead(ctx)
			}

			err1 := privacyFunc.EvalQuery(ctx, nil)
			err2 := privacyFunc.EvalMutation(ctx, nil)
			if tt.wantErr {
				require.Equal(privacy.Skip, err1)
				require.Equal(privacy.Skip, err2)
			} else {
				require.Equal(privacy.Allow, err1)
				require.Equal(privacy.Allow, err2)
			}
		})
	}
}

func Test_FilterUserRule(t *testing.T) {
	tests := []struct {
		name        string
		contextFunc func() context.Context
		wantErr     error
		query       ent.Query
	}{
		{
			name: "viewer does not container user",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "", nil)
			},
			wantErr: privacy.Deny,
		},
		{
			name: "viewer contains user",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "asdf", nil)
			},
			wantErr: privacy.Skip,
			query:   &ent.ShelfQuery{},
		},
		{
			name: "filter does not implement UserFilter",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "asdf", nil)
			},
			wantErr: privacy.Deny,
			query:   &ent.BookQuery{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			var query ent.Query
			if q, ok := tt.query.(*ent.ShelfQuery); ok {
				if q != nil {
					query = *q
				}
			}
			privacyFunc := FilterUserRule()
			err := privacyFunc.EvalQuery(tt.contextFunc(), tt.query)
			if err2 := errors.Unwrap(err); err2 != nil {
				require.Equal(tt.wantErr, err2, "error: %v", err)
			} else {
				require.Equal(tt.wantErr, err)
			}
			if err == privacy.Skip {
				require.NotEqual(query, tt.query)
			}
		})
	}
}

func Test_FilterUserOrPublicRule(t *testing.T) {
	tests := []struct {
		name        string
		contextFunc func() context.Context
		wantErr     error
		query       ent.Query
	}{
		{
			name: "viewer does not container user",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "", nil)
			},
			wantErr: privacy.Deny,
			query:   &ent.ShelfQuery{},
		},
		{
			name: "viewer contains user",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "asdf", nil)
			},
			wantErr: privacy.Skip,
			query:   &ent.ShelfQuery{},
		},
		{
			name: "filter does not implement UserOrPublicFilter",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "asdf", nil)
			},
			wantErr: privacy.Deny,
			query:   &ent.BookQuery{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			var query ent.Query
			if q, ok := tt.query.(*ent.ShelfQuery); ok {
				if q != nil {
					query = *q
				}
			}
			privacyFunc := FilterUserOrPublicRule()
			err := privacyFunc.EvalQuery(tt.contextFunc(), tt.query)
			if err2 := errors.Unwrap(err); err2 != nil {
				require.Equal(tt.wantErr, err2, "error: %v", err)
			} else {
				require.Equal(tt.wantErr, err)
			}
			if err == privacy.Skip {
				require.NotEqual(query, tt.query)
			}
		})
	}
}

func Test_DenyPublicWithoutPermissionRule(t *testing.T) {
	tests := []struct {
		name         string
		contextFunc  func() context.Context
		wantErr      error
		mutationFunc func() ent.Mutation
	}{
		{
			name: "viewer does not contain perms",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "", nil)
			},
			wantErr: privacy.Deny,
			mutationFunc: func() ent.Mutation {
				ret := &ent.ShelfMutation{}
				ret.SetPublic(true)
				return ret
			},
		},
		{
			name: "mutation does not implement PublicableMutation",
			contextFunc: func() context.Context {
				return viewer.NewContext(
					context.Background(),
					"",
					permissions.NewPermissions(permissions.CanCreatePublic),
				)
			},
			wantErr: privacy.Deny,
			mutationFunc: func() ent.Mutation {
				return &ent.BookMutation{}
			},
		},
		{
			name: "viewer does not have permission",
			contextFunc: func() context.Context {
				return viewer.NewContext(
					context.Background(),
					"",
					permissions.NewPermissions(permissions.CanEdit),
				)
			},
			wantErr: privacy.Deny,
			mutationFunc: func() ent.Mutation {
				ret := &ent.ShelfMutation{}
				ret.SetPublic(true)
				return ret
			},
		},
		{
			name: "viewer has permission",
			contextFunc: func() context.Context {
				return viewer.NewContext(
					context.Background(),
					"",
					permissions.NewPermissions(permissions.CanCreatePublic),
				)
			},
			wantErr: privacy.Skip,
			mutationFunc: func() ent.Mutation {
				ret := &ent.ShelfMutation{}
				ret.SetPublic(true)
				return ret
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			privacyFunc := DenyPublicWithoutPermissionRule()
			err := privacyFunc.EvalMutation(tt.contextFunc(), tt.mutationFunc())
			if err2 := errors.Unwrap(err); err2 != nil {
				require.Equal(tt.wantErr, err2, "error: %v", err)
			} else {
				require.Equal(tt.wantErr, err)
			}
		})
	}
}

func Test_DenyMismatchedUserRule(t *testing.T) {
	tests := []struct {
		name         string
		contextFunc  func() context.Context
		wantErr      error
		mutationFunc func() ent.Mutation
	}{
		{
			name: "mutation does not implement UserableMutation",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "asdf", nil)
			},
			wantErr: privacy.Deny,
			mutationFunc: func() ent.Mutation {
				return &ent.BookMutation{}
			},
		},
		{
			name: "viewer does not contain user",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "", nil)
			},
			wantErr: privacy.Deny,
			mutationFunc: func() ent.Mutation {
				ret := &ent.ShelfMutation{}
				ret.SetOp(ent.OpCreate)
				ret.SetUserID("asdf")
				return ret
			},
		},
		{
			name: "viewer contains mismatched user - create",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "asdf", nil)
			},
			wantErr: privacy.Deny,
			mutationFunc: func() ent.Mutation {
				ret := &ent.ShelfMutation{}
				ret.SetOp(ent.OpCreate)
				ret.SetUserID("fdsa")
				return ret
			},
		},
		{
			name: "viewer contains matched user - create",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "asdf", nil)
			},
			wantErr: privacy.Skip,
			mutationFunc: func() ent.Mutation {
				ret := &ent.ShelfMutation{}
				ret.SetOp(ent.OpCreate)
				ret.SetUserID("asdf")
				return ret
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			privacyFunc := DenyMismatchedUserRule()
			err := privacyFunc.EvalMutation(tt.contextFunc(), tt.mutationFunc())
			if err2 := errors.Unwrap(err); err2 != nil {
				require.Equal(tt.wantErr, err2, "error: %v", err)
			} else {
				require.Equal(tt.wantErr, err)
			}
		})
	}
}

func Test_FilterSelfRule(t *testing.T) {
	tests := []struct {
		name        string
		contextFunc func() context.Context
		wantErr     error
		query       ent.Query
	}{
		{
			name: "viewer does not container user",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "", nil)
			},
			wantErr: privacy.Deny,
		},
		{
			name: "viewer contains user",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "asdf", nil)
			},
			wantErr: privacy.Skip,
			query:   &ent.UserQuery{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			var query ent.Query
			if q, ok := tt.query.(*ent.ShelfQuery); ok {
				if q != nil {
					query = *q
				}
			}
			privacyFunc := FilterSelfRule()
			err := privacyFunc.EvalQuery(tt.contextFunc(), tt.query)
			if err2 := errors.Unwrap(err); err2 != nil {
				require.Equal(tt.wantErr, err2, "error: %v", err)
			} else {
				require.Equal(tt.wantErr, err)
			}
			if err == privacy.Skip {
				require.NotEqual(query, tt.query)
			}
		})
	}
}

func Test_FilterUserOrSystemRule(t *testing.T) {
	tests := []struct {
		name        string
		contextFunc func() context.Context
		wantErr     error
		query       ent.Query
	}{
		{
			name: "viewer does not container user",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "", nil)
			},
			wantErr: privacy.Deny,
			query:   &ent.ShelfQuery{},
		},
		{
			name: "viewer contains user",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "asdf", nil)
			},
			wantErr: privacy.Skip,
			query:   &ent.TaskQuery{},
		},
		{
			name: "filter does not implement UserOrPublicFilter",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "asdf", nil)
			},
			wantErr: privacy.Deny,
			query:   &ent.ShelfQuery{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			var query ent.Query
			if q, ok := tt.query.(*ent.ShelfQuery); ok {
				if q != nil {
					query = *q
				}
			}
			privacyFunc := FilterUserOrSystemRule()
			err := privacyFunc.EvalQuery(tt.contextFunc(), tt.query)
			if err2 := errors.Unwrap(err); err2 != nil {
				require.Equal(tt.wantErr, err2, "error: %v", err)
			} else {
				require.Equal(tt.wantErr, err)
			}
			if err == privacy.Skip {
				require.NotEqual(query, tt.query)
			}
		})
	}
}

func Test_AllowTasksOfType(t *testing.T) {
	tests := []struct {
		name         string
		contextFunc  func() context.Context
		wantErr      error
		mutationFunc func() *ent.TaskMutation
	}{
		{
			name: "mutation is missing task type",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "", nil)
			},
			wantErr: privacy.Deny,
			mutationFunc: func() *ent.TaskMutation {
				return &ent.TaskMutation{}
			},
		},
		{
			name: "NoOp allowed - create",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "", nil)
			},
			wantErr: privacy.Allow,
			mutationFunc: func() *ent.TaskMutation {
				ret := &ent.TaskMutation{}
				ret.SetOp(ent.OpCreate)
				ret.SetType(task_enums.TypeNoOp)
				return ret
			},
		},
		{
			name: "CalibreImport denied - create",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "", nil)
			},
			wantErr: privacy.Deny,
			mutationFunc: func() *ent.TaskMutation {
				ret := &ent.TaskMutation{}
				ret.SetOp(ent.OpCreate)
				ret.SetType(task_enums.TypeCalibreImport)
				return ret
			},
		},
		// TODO: figure out how to test update mutations
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			privacyFunc := AllowTasksOfType(task_enums.TypeNoOp)
			err := privacyFunc.EvalMutation(tt.contextFunc(), tt.mutationFunc())
			if err2 := errors.Unwrap(err); err2 != nil {
				require.Equal(tt.wantErr, err2, "error: %v", err)
			} else {
				require.Equal(tt.wantErr, err)
			}
		})
	}
}

func Test_DenySystemTaskForNonAdmin(t *testing.T) {
	tests := []struct {
		name         string
		contextFunc  func() context.Context
		wantErr      error
		mutationFunc func() *ent.TaskMutation
	}{
		{
			name: "viewer is admin system task",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "", permissions.NewPermissions(permissions.Admin))
			},
			wantErr: privacy.Skip,
			mutationFunc: func() *ent.TaskMutation {
				ret := &ent.TaskMutation{}
				ret.SetOp(ent.OpCreate)
				ret.SetIsSystemTask(true)
				return ret
			},
		},
		{
			name: "viewer is admin not system task",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "", permissions.NewPermissions(permissions.Admin))
			},
			wantErr: privacy.Skip,
			mutationFunc: func() *ent.TaskMutation {
				ret := &ent.TaskMutation{}
				ret.SetOp(ent.OpCreate)
				ret.SetIsSystemTask(false)
				return ret
			},
		},
		{
			name: "viewer is not admin system task",
			contextFunc: func() context.Context {
				return viewer.NewContext(context.Background(), "", nil)
			},
			wantErr: privacy.Deny,
			mutationFunc: func() *ent.TaskMutation {
				ret := &ent.TaskMutation{}
				ret.SetOp(ent.OpCreate)
				ret.SetIsSystemTask(true)
				return ret
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			privacyFunc := DenySystemTaskForNonAdmin()
			err := privacyFunc.EvalMutation(tt.contextFunc(), tt.mutationFunc())
			if err2 := errors.Unwrap(err); err2 != nil {
				require.Equal(tt.wantErr, err2, "error: %v", err)
			} else {
				require.Equal(tt.wantErr, err)
			}
		})
	}
}
