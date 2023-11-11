package resolver

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/oat431/api-tutorial-go/payload/request"
	"github.com/oat431/api-tutorial-go/payload/response"
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

func (s *userService) GetAllDBUsers(params graphql.ResolveParams) (interface{}, error) {
	users := s.service.GetAllDBUsers()
	return users, nil
}

func (s *userService) GetAllDBUsersPagination(params graphql.ResolveParams) (interface{}, error) {
	page := params.Args["page"].(int)
	size := params.Args["limit"].(int)

	if page <= 0 {
		page = 1
	} else {
		page = page - 1
	}

	url := "https://try-gin-deploy.onrender.com/api/v2/users/pages?page=" + strconv.Itoa(page) + "&size=" + strconv.Itoa(size)
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	var data response.PageUserDto
	json.Unmarshal(body, &data)
	return data, nil
}

func (s *userService) GetAllDBUsersById(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	user := s.service.GetAllDBUsersById(strconv.Itoa(id))
	return user, nil
}

func (s *userService) CreateUser(params graphql.ResolveParams) (interface{}, error) {
	userRequest := params.Args["request"].(request.UserRequest)
	userDto := s.service.CreateUser(userRequest)
	return userDto, nil
}

func (s *userService) UpdateUser(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	userRequest := params.Args["request"].(request.UserRequest)
	userDto := s.service.UpdateUser(strconv.Itoa(id), userRequest)
	return userDto, nil
}

func (s *userService) DeleteUser(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	s.service.DeleteUser(strconv.Itoa(id))
	return "User deleted", nil
}
