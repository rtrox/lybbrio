package graph

import (
	"context"
	"lybbrio/internal/db"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/ent/user"
	"lybbrio/internal/viewer"
	"testing"
)

type testContext struct {
	Client       *ent.Client
	AdminCtx     context.Context
	UserCtx      context.Context
	Resolver     *Resolver
	teardownFunc func()
}

func (tc testContext) User() *ent.User {
	uid, _ := viewer.FromContext(tc.UserCtx).UserID()
	user := tc.Client.User.Query().Where(user.ID(uid)).OnlyX(tc.UserCtx)
	return user
}

func (tc testContext) Teardown() {
	tc.teardownFunc()
	tc.Client.Close()
}

func setupTest(t *testing.T, testName string, teardown ...func()) testContext {
	client := db.OpenTest(t, testName)

	adminCtx := viewer.NewSystemAdminContext(context.Background())
	perms := client.UserPermissions.Create().SaveX(adminCtx)
	user := client.User.Create().
		SetUsername("some_user_name").
		SetEmail("some_email").
		SetUserPermissions(perms).
		SaveX(adminCtx)

	userCtx := viewer.NewContext(
		context.Background(),
		user.ID,
		permissions.NewPermissions(),
	)

	r := &Resolver{client}

	ret := testContext{
		Client:   client,
		AdminCtx: adminCtx,
		UserCtx:  userCtx,
		Resolver: r,
	}

	if len(teardown) > 0 {
		ret.teardownFunc = teardown[0]
	} else {
		ret.teardownFunc = func() {}
	}
	return ret
}
