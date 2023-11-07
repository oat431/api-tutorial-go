package resolver

import (
	"strconv"
	"github.com/graphql-go/graphql"
	"github.com/oat431/api-tutorial-go/services"
)

type UserResolver interface {
	GetAllDBUsers(params graphql.ResolveParams) (interface{},error)
	GetAllDBUsersPagination(params graphql.ResolveParams) (interface{},error)
	GetAllDBUsersById(params graphql.ResolveParams) (interface{},error)
	// CreateUser(params graphql.ResolveParams) (interface{}, error)
	// UpdateUser(params graphql.ResolveParams) (interface{}, error)
	// DeleteUser(params graphql.ResolveParams) (interface{}, error)
}

type userService struct {
	service services.UserV2Service
}

func CreateUserResolver(service services.UserV2Service) *userService {
	return &userService{service}
}

func (s *userService) GetAllDBUsers(params graphql.ResolveParams) (interface{},error) {
	users := s.service.GetAllDBUsers()
	return users, nil
}

func (s *userService) GetAllDBUsersPagination(params graphql.ResolveParams) (interface{},error) {
	// pages := s.service.GetAllDBUsersPagination()
	return nil,nil;
	// return pages
}

func (s *userService) GetAllDBUsersById(params graphql.ResolveParams) (interface{},error) {
	id := params.Args["id"].(int)
	user := s.service.GetAllDBUsersById(strconv.Itoa(id))
	return user,nil
}