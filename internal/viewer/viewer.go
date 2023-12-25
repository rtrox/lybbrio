package viewer

import (
	"context"
	"lybbrio/internal/ent"
)

type viewerCtxKeyType string

const viewerCtxKey viewerCtxKeyType = "viewer"

type Viewer interface {
	IsAdmin() bool
	User() (*ent.User, bool)
	Permissions() (*ent.UserPermissions, bool)
}

type UserViewer struct {
	u *ent.User
	p *ent.UserPermissions
}

func (v UserViewer) IsAdmin() bool {
	if v.p == nil {
		return false
	}
	return v.p.Admin
}

func (v UserViewer) User() (*ent.User, bool) {
	if v.u != nil {
		return v.u, true
	}
	return nil, false
}

func (v UserViewer) Permissions() (*ent.UserPermissions, bool) {
	if v.p != nil {
		return v.p, true
	}
	return nil, false
}

func NewContext(ctx context.Context, u *ent.User, p *ent.UserPermissions) context.Context {
	return context.WithValue(ctx, viewerCtxKey, UserViewer{u: u, p: p})
}

func FromContext(ctx context.Context) Viewer {
	return ctx.Value(viewerCtxKey).(Viewer)
}

type SystemAdminViewer struct{}

func (v SystemAdminViewer) IsAdmin() bool {
	return true
}

func (v SystemAdminViewer) User() (*ent.User, bool) {
	return nil, false
}

func (v SystemAdminViewer) Permissions() (*ent.UserPermissions, bool) {
	return nil, false
}

func NewSystemAdminContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, viewerCtxKey, SystemAdminViewer{})
}
