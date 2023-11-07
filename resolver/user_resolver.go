package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/oat431/api-tutorial-go/services"
)

type UserResolver interface {
	GetAllDBUsers(params graphql.ResolveParams) (interface{}, error)
	GetAllDBUsersPagination(params graphql.ResolveParams) (interface{}, error)
	GetAllDBUsersById(params graphql.ResolveParams) (interface{}, error)
	CreateUser(params graphql.ResolveParams) (interface{}, error)
	UpdateUser(params graphql.ResolveParams) (interface{}, error)
	DeleteUser(params graphql.ResolveParams) (interface{}, error)
}

type userService struct {
	service services.UserV2Service
}

func CreateUserResolver(service services.UserV2Service) *userService {
	return &userService{service}
}