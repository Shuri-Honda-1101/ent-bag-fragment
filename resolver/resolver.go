package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/Shuri-Honda-1101/ent-bug-fragment/ent"
	"github.com/Shuri-Honda-1101/ent-bug-fragment/gql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return gql.NewExecutableSchema(gql.Config{
		Resolvers: &Resolver{client},
	})
}
