package graph

import (
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/argon2id"
	"lybbrio/internal/graph/generated"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the resolver root.
type Resolver struct {
	client         *ent.Client
	argon2idConfig argon2id.Config
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client, argon2idConfig argon2id.Config) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			client:         client,
			argon2idConfig: argon2idConfig,
		},
	})
}
