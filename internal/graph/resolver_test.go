package graph

import (
	"lybbrio/internal/db"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewSchema(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	client := db.OpenTest(t, "Test_NewSchema")
	defer client.Close()

	schema := NewSchema(client)
	require.NotNil(schema)
}
