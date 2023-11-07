package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/oat431/api-tutorial-go/resolver"
)

type UserSchema interface {
	GetAllDBUsers() *graphql.Field
	GetAllDBUsersPagination() *graphql.Field
	GetAllDBUsersById() *graphql.Field
	CreateUser() *graphql.Field
	UpdateUser() *graphql.Field
	DeleteUser() *graphql.Field
}

type userResolver struct {
	userResolver resolver.UserResolver
}

func CreateUserSchema(gqlResolver resolver.UserResolver) *userResolver {
	return &userResolver{gqlResolver}
}

