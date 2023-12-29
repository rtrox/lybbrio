package test

import (
	"context"
	"lybbrio/internal/db"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/viewer"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	USER_ID_1 = ksuid.MustNew("usr")
	USER_ID_2 = ksuid.MustNew("usr")
	ADMIN_ID  = ksuid.MustNew("usr")

	USER_PERM_1 = ksuid.MustNew("prm")
	USER_PERM_2 = ksuid.MustNew("prm")
	ADMIN_PERM  = ksuid.MustNew("prm")
)

type testData struct {
	user1ViewerContext context.Context
	user2ViewerContext context.Context
	adminViewerContext context.Context
}

func setupTest(t *testing.T, testName string) (teardownFunc func(t *testing.T), client *ent.Client, data testData) {
	client = db.OpenTest(t, testName)

	adminContext := viewer.NewSystemAdminContext(context.Background())
	user1perms := client.UserPermissions.Create().
		SetID(USER_PERM_1).
		SaveX(adminContext)
	user2perms := client.UserPermissions.Create().
		SetID(USER_PERM_2).
		SetCanCreatePublic(true).
		SaveX(adminContext)
	adminperms := client.UserPermissions.Create().SetID(ADMIN_PERM).SetAdmin(true).SaveX(adminContext)
	user1 := client.User.Create().
		SetID(USER_ID_1).
		SetUsername("user1").
		SetEmail("user1@test.test").
		SetUserPermissions(user1perms).
		SaveX(adminContext)
	user2 := client.User.Create().
		SetID(USER_ID_2).
		SetUsername("user2").
		SetEmail("user2@test.test").
		SetUserPermissions(user2perms).
		SaveX(adminContext)
	admin := client.User.Create().
		SetID(ADMIN_ID).
		SetUsername("admin").
		SetEmail("admin@test.test").
		SetUserPermissions(adminperms).
		SaveX(adminContext)

	data = testData{
		user1ViewerContext: viewer.NewContext(context.Background(), user1, user1perms),
		user2ViewerContext: viewer.NewContext(context.Background(), user2, user2perms),
		adminViewerContext: viewer.NewContext(context.Background(), admin, adminperms),
	}

	teardownFunc = func(t *testing.T) {
		require.NoError(t, client.Close(), "failed to close ent client")
	}
	return teardownFunc, client, data
}
