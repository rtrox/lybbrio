package graph

import (
	"lybbrio/internal/db"
	"lybbrio/internal/ent/schema/argon2id"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewSchema(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	client := db.OpenTest(t, "Test_NewSchema")
	argon2idConfig := argon2id.Config{
		Memory:      64,
		Iterations:  1,
		Parallelism: 1,
		SaltLen:     16,
		KeyLen:      32,
	}
	defer client.Close()

	schema := NewSchema(client, argon2idConfig)
	require.NotNil(schema)
}
