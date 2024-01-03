package permissions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_String(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	for i := Admin; i < count; i++ {
		require.NotEqual("", Permission(i).String())
	}

	require.Equal("", Permission(0).String())
}

func Test_FromString(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	for i := Admin; i < count; i++ {
		require.Equal(i, FromString(Permission(i).String()))
	}

	require.Equal(Permission(0), FromString("asdf"))
}

func Test_NewPermissions(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	p := NewPermissions(Admin, CanCreatePublic)
	require.True(p.Has(Admin))
	require.True(p.Has(CanCreatePublic))
	require.False(p.Has(CanEdit))
}

func Test_Permissions_Has(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	p := Permissions{
		Admin:           struct{}{},
		CanCreatePublic: struct{}{},
	}

	require.True(p.Has(Admin))
	require.True(p.Has(CanCreatePublic))
	require.False(p.Has(CanEdit))
}

func Test_Permissions_Add(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	p := Permissions{}

	p.Add(Admin)
	require.True(p.Has(Admin))
	require.False(p.Has(CanCreatePublic))
	require.False(p.Has(CanEdit))

	p.Add(CanCreatePublic)
	require.True(p.Has(Admin))
	require.True(p.Has(CanCreatePublic))
	require.False(p.Has(CanEdit))

	p.Add(CanEdit)
	require.True(p.Has(Admin))
	require.True(p.Has(CanCreatePublic))
	require.True(p.Has(CanEdit))
}

func Test_Permissions_StringSlice(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	p := Permissions{
		Admin:           struct{}{},
		CanCreatePublic: struct{}{},
	}

	require.Equal([]string{"Admin", "CanCreatePublic"}, p.StringSlice())
}

func Test_From(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	type testStruct struct {
		Admin           bool
		CanCreatePublic bool
		CanEdit         bool
	}

	p := From(&testStruct{
		Admin:           true,
		CanCreatePublic: true,
		CanEdit:         false,
	})
	require.True(p.Has(Admin))
	require.True(p.Has(CanCreatePublic))
	require.False(p.Has(CanEdit))

	p2 := From(testStruct{
		Admin:           false,
		CanCreatePublic: true,
		CanEdit:         true,
	})
	require.False(p2.Has(Admin))
	require.True(p2.Has(CanCreatePublic))
	require.True(p2.Has(CanEdit))
}

func Test_All(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	p := All()
	for i := Admin; i < count; i++ {
		require.True(p.Has(Permission(i)))
	}

}
