package viewer

import (
	"context"
	"lybbrio/internal/ent"
)

type viewerCtxKeyType string

const viewerCtxKey viewerCtxKeyType = "viewer"

type Viewer interface {
	IsAdmin() bool
	User() *ent.User
}

type UserViewer struct {
	u *ent.User
	p *ent.UserPermissions
}

func (v UserViewer) IsAdmin() bool {
	return v.p.Admin
}

func (v UserViewer) User() *ent.User {
	return v.u
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

func (v SystemAdminViewer) User() *ent.User {
	return nil
}

func NewSystemAdminContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, viewerCtxKey, SystemAdminViewer{})
}
