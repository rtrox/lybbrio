package viewer

import (
	"context"
	"testing"

	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/ksuid"

	"github.com/stretchr/testify/assert"
)

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
		p           *ent.UserPermissions
		shouldAdmin bool
	}{
		{
			name:        "nil",
			p:           nil,
			shouldAdmin: false,
		},
		{
			name:        "not admin",
			p:           &ent.UserPermissions{Admin: false},
			shouldAdmin: false,
		},
		{
			name:        "admin",
			p:           &ent.UserPermissions{Admin: true},
			shouldAdmin: true,
		},
	}
	assert := assert.New(t)
	for _, tt := range params {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(tt.shouldAdmin, UserViewer{p: tt.p}.IsAdmin())
		})
	}
}

func Test_NewContext(t *testing.T) {
	assert := assert.New(t)
	expectedID := ksuid.MustNew("usr")
	ctx := NewContext(
		context.Background(),
		&ent.User{ID: expectedID},
		&ent.UserPermissions{Admin: true},
	)
	if !assert.NotNil(ctx, "NewContext should not return nil") {
		t.FailNow()
	}
	v := FromContext(ctx)
	if !assert.NotNil(v, "FromContext should not return nil") {
		t.FailNow()
	}
	assert.Implements((*Viewer)(nil), v, "FromContext should return a Viewer")
	u, ok := v.User()
	assert.True(ok, "FromContext should return the user")
	assert.Equal(expectedID, u.ID, "FromContext should return the same user")
	assert.True(v.IsAdmin(), "FromContext should return the same permissions")
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
	_, ok := v.User()
	assert.False(ok, "FromContext should not return a user")
	assert.True(v.IsAdmin(), "FromContext should return admin permissions")
}
