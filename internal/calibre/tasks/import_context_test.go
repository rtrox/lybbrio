package tasks

import (
	"context"
	"testing"

	"lybbrio/internal/ent/schema/ksuid"

	"github.com/stretchr/testify/require"
)

func Test_AddFailedBook(t *testing.T) {
	require := require.New(t)
	c := newImportContext()
	c.AddFailedBook("test")
	require.Equal([]string{"test"}, c.failedBooks)
}

func Test_AuthorVisited(t *testing.T) {
	require := require.New(t)
	c := newImportContext()
	expected := ksuid.MustNew("aut")
	c.AddAuthorVisited(1, expected)
	kid, ok := c.AuthorVisited(1)
	require.True(ok)
	require.Equal(expected, kid)

	_, ok = c.AuthorVisited(2)
	require.False(ok)
}

func Test_TagVisited(t *testing.T) {
	require := require.New(t)
	c := newImportContext()
	expected := ksuid.MustNew("tag")
	c.AddTagVisited(1, expected)
	kid, ok := c.TagVisited(1)
	require.True(ok)
	require.Equal(expected, kid)

	_, ok = c.TagVisited(2)
	require.False(ok)
}

func Test_PublisherVisited(t *testing.T) {
	require := require.New(t)
	c := newImportContext()
	expected := ksuid.MustNew("pub")
	c.AddPublisherVisited(1, expected)
	kid, ok := c.PublisherVisited(1)
	require.True(ok)
	require.Equal(expected, kid)

	_, ok = c.PublisherVisited(2)
	require.False(ok)
}

func Test_LanguageVisited(t *testing.T) {
	require := require.New(t)
	c := newImportContext()
	expected := ksuid.MustNew("lan")
	c.AddLanguageVisited(1, expected)
	kid, ok := c.LanguageVisited(1)
	require.True(ok)
	require.Equal(expected, kid)

	_, ok = c.LanguageVisited(2)
	require.False(ok)
}

func Test_SeriesVisited(t *testing.T) {
	require := require.New(t)
	c := newImportContext()
	expected := ksuid.MustNew("ser")
	c.AddSeriesVisited(1, expected)
	kid, ok := c.SeriesVisited(1)
	require.True(ok)
	require.Equal(expected, kid)

	_, ok = c.SeriesVisited(2)
	require.False(ok)
}

func Test_ImportContext(t *testing.T) {
	require := require.New(t)
	ic := newImportContext()
	ctx := context.Background()
	ctx = importContextTo(ctx, ic)
	require.NotNil(ctx)

	ic2 := importContextFrom(ctx)
	require.NotNil(ic2)
	require.Equal(ic, ic2)
}

func Test_ImportContext_ReturnsNilWhenNotAttached(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()
	ic := importContextFrom(ctx)
	require.Nil(ic)
}

func Test_ImportContext_String(t *testing.T) {
	require := require.New(t)
	ic := newImportContext()
	require.Empty(ic.String())
	ic.AddFailedBook("test")
	ic.AddFailedBook("test2")

	require.Contains(ic.String(), "test")
	require.Contains(ic.String(), "test2")
}
