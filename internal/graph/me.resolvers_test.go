package graph

import (
	"context"
	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/viewer"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Me(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupTest(t, "Test_Me")
	defer tc.Teardown()

	me, err := tc.Resolver.Query().Me(tc.UserCtx)
	require.NoError(err)

	require.Equal(tc.User().ID, me.ID)
	require.Equal(tc.User().Username, me.Username)
	require.Equal(tc.User().Email, me.Email)
}

func Test_Me_NoViewer(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupTest(t, "Test_Me_NoViewer")
	defer tc.Teardown()

	me, err := tc.Resolver.Query().Me(context.Background())
	require.Nil(me)
	require.Error(err)
}

func Test_Me_NoUserID(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupTest(t, "Test_Me_NoUserID")
	defer tc.Teardown()

	me, err := tc.Resolver.Query().Me(
		viewer.NewContext(context.Background(), "", permissions.NewPermissions()),
	)
	require.Nil(me)
	require.Error(err)
}
