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

var userDtoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserDto",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"age": &graphql.Field{
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
		"page": &graphql.Field{
			Type: graphql.Int,
		},
		"size": &graphql.Field{
			Type: graphql.Int,
		},
		"maxPage": &graphql.Field{
			Type: graphql.Int,
		},
		"totalPages": &graphql.Field{
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

var userRequest = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserRequest",
	Fields: graphql.InputObjectConfigFieldMap{
		"firstname": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"lastname": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"birth_date": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

func (u *userResolver) GetAllDBUsers() *graphql.Field {
	field := graphql.Field{
		Type:        graphql.NewList(userDtoType),
		Description: "Get all DB users",
		Resolve:     u.userResolver.GetAllDBUsers,
	}
	return &field
}

func (u *userResolver) GetAllDBUsersPagination() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"limit": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	}
	field := graphql.Field{
		Type:        pageUserDtoType,
		Description: "Get all DB users as Pagination",
		Args:        args,
		Resolve:     u.userResolver.GetAllDBUsersPagination,
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
		Type:        userDtoType,
		Description: "Get all DB users by id",
		Args:        args,
		Resolve:     u.userResolver.GetAllDBUsersById,
	}

	return &field
}

func (u *userResolver) CreateUser() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"request": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(userRequest),
		},
	}

	field := graphql.Field{
		Type:        userDtoType,
		Description: "Create user",
		Args:        args,
		Resolve:     u.userResolver.CreateUser,
	}

	return &field
}

func (u *userResolver) UpdateUser() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"request": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(userRequest),
		},
	}

	field := graphql.Field{
		Type:        userDtoType,
		Description: "Update user",
		Args:        args,
		Resolve:     u.userResolver.UpdateUser,
	}

	return &field
}

func (u *userResolver) DeleteUser() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	}

	field := graphql.Field{
		Type:        userDtoType,
		Description: "Delete user",
		Args:        args,
		Resolve:     u.userResolver.DeleteUser,
	}

	return &field
}
