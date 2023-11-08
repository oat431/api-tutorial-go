package resolver

import (
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/oat431/api-tutorial-go/payload/response"
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
	page := params.Args["page"].(int)
	size := params.Args["limit"].(int)
	
	if page <= 0 {
		page = 1
	} else {
		page = page - 1
	}

	url := "https://try-gin-deploy.onrender.com/api/v2/users/pages?page=" + strconv.Itoa(page) + "&size=" +strconv.Itoa(size)
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	var data response.PageUserDto
	json.Unmarshal(body, &data)
	return data,nil;
}

func (s *userService) GetAllDBUsersById(params graphql.ResolveParams) (interface{},error) {
	id := params.Args["id"].(int)
	user := s.service.GetAllDBUsersById(strconv.Itoa(id))
	return user,nil
}