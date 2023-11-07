package schema

import (
	"github.com/graphql-go/graphql"
)

type Schema interface {
	Query() *graphql.Object
	Mutation() *graphql.Object
	Root() *graphql.Schema
}

type RootSchema struct {
	userSchema UserSchema
}

func CreateRootSchema(userSchema UserSchema) *RootSchema {
	return &RootSchema{
		userSchema: userSchema,
	}
}

func (r *RootSchema) Query() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"GetAllDBUsers" : r.userSchema.GetAllDBUsers(),
			"GetAllDBUsersPagination" : r.userSchema.GetAllDBUsersPagination(),
			"GetAllDBUsersById" : r.userSchema.GetAllDBUsersById(),
		},
	})
}

func (r *RootSchema) Mutation() *graphql.Object {
	mutation := graphql.ObjectConfig{Name: "Mutation"}
	return graphql.NewObject(mutation)
}

func (r *RootSchema) Root() *graphql.Schema {
	rootQuery := r.Query()
	rootMutation := r.Mutation()
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	return &schema
}
