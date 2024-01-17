package viewer

import (
	"context"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/schema/permissions"
)

type viewerCtxKeyType string

const viewerCtxKey viewerCtxKeyType = "viewer"

type Viewer interface {
	IsAdmin() bool
	UserID() (ksuid.ID, bool)
	Has(permissions.Permission) bool
}

type UserViewer struct {
	uid ksuid.ID
	p   permissions.Permissions
}

func (v UserViewer) IsAdmin() bool {
	if v.p == nil {
		return false
	}
	return v.p.Has(permissions.Admin)
}

func (v UserViewer) UserID() (ksuid.ID, bool) {
	if v.uid != "" {
		return v.uid, true
	}
	return "", false
}

func (v UserViewer) Has(p permissions.Permission) bool {
	if v.p == nil {
		return false
	}
	return v.p.Has(p)
}

func NewContext(ctx context.Context, uid ksuid.ID, permissions permissions.Permissions) context.Context {
	return context.WithValue(ctx, viewerCtxKey, UserViewer{uid: uid, p: permissions})
}

func FromContext(ctx context.Context) Viewer {
	view, ok := ctx.Value(viewerCtxKey).(Viewer)
	if !ok {
		return nil
	}
	return view
}

type SystemAdminViewer struct{}

func (v SystemAdminViewer) IsAdmin() bool {
	return true
}

func (v SystemAdminViewer) UserID() (ksuid.ID, bool) {
	return "", false
}

func (v SystemAdminViewer) Has(p permissions.Permission) bool {
	return p == permissions.Admin
}

func NewSystemAdminContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, viewerCtxKey, SystemAdminViewer{})
}

type AnonymousViewer struct{}

func (v AnonymousViewer) IsAdmin() bool {
	return false
}

func (v AnonymousViewer) UserID() (ksuid.ID, bool) {
	return "", false
}

func (v AnonymousViewer) Has(p permissions.Permission) bool {
	return false
}

func NewAnonymousContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, viewerCtxKey, AnonymousViewer{})
}
