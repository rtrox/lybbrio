package viewer

import (
	"context"
	"testing"

	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/schema/permissions"

	"github.com/stretchr/testify/assert"
)

func Test_EmptyViewer(t *testing.T) {
	assert := assert.New(t)
	ctx := NewContext(context.Background(), "", nil)
	v := FromContext(ctx)
	assert.False(v.IsAdmin())
	assert.Empty(v.UserID())
	assert.False(v.Has(permissions.Admin))
}

func Test_NoViewer(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	v := FromContext(ctx)
	assert.Nil(v)
}

func Test_UserViewerImplementsViewer(t *testing.T) {
	assert := assert.New(t)
	assert.NotPanics(func() {
		var _ Viewer = UserViewer{}
	}, "UserViewer does not implement Viewer")
}

func Test_SystemAdminViewerImplementsViewer(t *testing.T) {
	assert := assert.New(t)
	assert.NotPanics(func() {
		var _ Viewer = SystemAdminViewer{}
	})
}

func Test_SystemAdminViewerIsAlwaysAdmin(t *testing.T) {
	assert := assert.New(t)
	v := SystemAdminViewer{}
	assert.True(v.IsAdmin())
}

func Test_UserViewerDefaultNoAdmin(t *testing.T) {
	assert := assert.New(t)
	v := UserViewer{}
	assert.False(v.IsAdmin())
}

func Test_UserViewerPermissionsAdmin(t *testing.T) {
	params := []struct {
		name        string
		p           permissions.Permissions
		shouldAdmin bool
	}{
		{
			name:        "not admin",
			p:           permissions.NewPermissions(),
			shouldAdmin: false,
		},
		{
			name:        "admin",
			p:           permissions.NewPermissions(permissions.Admin),
			shouldAdmin: true,
		},
	}
	assert := assert.New(t)
	for _, tt := range params {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(tt.shouldAdmin, UserViewer{p: tt.p}.IsAdmin())
			assert.Equal(tt.shouldAdmin, UserViewer{p: tt.p}.Has(permissions.Admin))
		})
	}
}

func Test_UserViewerHasPermissions(t *testing.T) {
	tests := []struct {
		name          string
		p             permissions.Permissions
		shouldHave    permissions.Permission
		shouldNotHave permissions.Permission
	}{
		{
			name:          "no permissions",
			p:             permissions.NewPermissions(),
			shouldHave:    0,
			shouldNotHave: permissions.Admin,
		},
		{
			name:          "admin",
			p:             permissions.NewPermissions(permissions.Admin),
			shouldHave:    permissions.Admin,
			shouldNotHave: permissions.CanCreatePublic,
		},
		{
			name:          "can create public",
			p:             permissions.NewPermissions(permissions.CanCreatePublic),
			shouldHave:    permissions.CanCreatePublic,
			shouldNotHave: permissions.CanEdit,
		},
		{
			name:          "can edit",
			p:             permissions.NewPermissions(permissions.CanEdit),
			shouldHave:    permissions.CanEdit,
			shouldNotHave: permissions.CanCreatePublic,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert := assert.New(t)
			v := UserViewer{p: tt.p}
			if tt.shouldNotHave != 0 {
				assert.False(v.Has(tt.shouldNotHave))
			}
			if tt.shouldHave != 0 {
				assert.True(v.Has(tt.shouldHave))
			}
		})
	}
}

func Test_NewContext(t *testing.T) {
	assert := assert.New(t)
	expectedID := ksuid.MustNew("usr")
	ctx := NewContext(
		context.Background(),
		expectedID,
		permissions.NewPermissions(permissions.Admin),
	)
	if !assert.NotNil(ctx, "NewContext should not return nil") {
		t.FailNow()
	}
	v := FromContext(ctx)
	if !assert.NotNil(v, "FromContext should not return nil") {
		t.FailNow()
	}
	assert.Implements((*Viewer)(nil), v, "FromContext should return a Viewer")
	uid, ok := v.UserID()
	assert.True(ok, "FromContext should return the user")
	assert.Equal(expectedID, uid, "FromContext should return the same user")
	assert.True(v.IsAdmin(), "FromContext should return the same permissions")
	assert.True(v.Has(permissions.Admin), "FromContext should return the same permissions")
}

func Test_NewSystemAdminContext(t *testing.T) {
	assert := assert.New(t)
	ctx := NewSystemAdminContext(context.Background())
	if !assert.NotNil(ctx, "NewSystemAdminContext should not return nil") {
		t.FailNow()
	}
	v := FromContext(ctx)
	if !assert.NotNil(v, "FromContext should not return nil") {
		t.FailNow()
	}
	assert.Implements((*Viewer)(nil), v, "FromContext should return a Viewer")
	_, ok := v.UserID()
	assert.False(ok, "FromContext should not return a user")
	assert.True(v.IsAdmin(), "FromContext should return admin permissions")
	assert.True(v.Has(permissions.Admin), "FromContext should return admin permissions")
}
