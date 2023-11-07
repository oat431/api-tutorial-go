package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/oat431/api-tutorial-go/resolver"
)

type UserSchema interface {
	GetAllDBUsers() *graphql.Field
	GetAllDBUsersPagination() *graphql.Field
	GetAllDBUsersById() *graphql.Field
	// CreateUser() *graphql.Field
	// UpdateUser() *graphql.Field
	// DeleteUser() *graphql.Field
}

type userResolver struct {
	userResolver resolver.UserResolver
}

func CreateUserSchema(gqlResolver resolver.UserResolver) *userResolver {
	return &userResolver{gqlResolver}
}

var userDtoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserDto",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name" : &graphql.Field{
			Type: graphql.String,
		},
		"email" : &graphql.Field{
			Type: graphql.String,
		},
		"age" : &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var pageUserDtoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PageUserDto",
	Fields: graphql.Fields{
		"items": &graphql.Field{
			Type: graphql.NewList(userDtoType),
		},
		"page" : &graphql.Field{
			Type: graphql.Int,
		},
		"size" : &graphql.Field{
			Type: graphql.Int,
		},
		"maxPage" : &graphql.Field{
			Type: graphql.Int,
		},
		"totalPages" : &graphql.Field{
			Type: graphql.Int,
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
		"last": &graphql.Field{
			Type: graphql.Boolean,
		},
		"first": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

func (u *userResolver) GetAllDBUsers() *graphql.Field {
	field := graphql.Field{
		Type: graphql.NewList(userDtoType),
		Description: "Get all DB users",
		Resolve: u.userResolver.GetAllDBUsers,
	}
	return &field
}

func (u *userResolver) GetAllDBUsersPagination() *graphql.Field {
	field := graphql.Field{
		Type: pageUserDtoType,
		Description: "Get all DB users as Pagination",
		Resolve: u.userResolver.GetAllDBUsersPagination,
	}
	return &field
}

func (u *userResolver) GetAllDBUsersById() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	}

	field := graphql.Field{
		Type: userDtoType,
		Description: "Get all DB users by id",
		Args: args,
		Resolve: u.userResolver.GetAllDBUsersById,
	}

	return &field
} 
