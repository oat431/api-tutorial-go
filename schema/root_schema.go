package schema

import (
	"github.com/graphql-go/graphql"
)

type Schema interface {
	Query() *graphql.Object
	Mutation() *graphql.Object
	Root() *graphql.Schema
}

type rootSchema struct {
}

func CreateRootSchema() *rootSchema {
	return &rootSchema{}
}

func (r *rootSchema) Query() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
	})
}

func (r *rootSchema) Mutation() *graphql.Object {
	mutation := graphql.ObjectConfig{Name: "Mutation"}
	return graphql.NewObject(mutation)
}

func (r *rootSchema) Root() *graphql.Schema {
	rootQuery := r.Query()
	rootMutation := r.Mutation()
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	return &schema
}
