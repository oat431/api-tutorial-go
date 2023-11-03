package services

import (
	"github.com/oat431/api-tutorial-go/dao"
	"github.com/oat431/api-tutorial-go/payload/response"
	"github.com/oat431/api-tutorial-go/utils"
)

type UserV2Service interface {
	GetAllDBUsers() []response.UserDto
	GetAllDBUsersById(id string) response.UserDto
}

type userV2Dao struct {
	dao dao.UserV2Dao
}

func CreateUserService(dao dao.UserV2Dao) *userV2Dao {
	return &userV2Dao{dao}
}

func (userDao *userV2Dao) GetAllDBUsers() []response.UserDto {
	users := userDao.dao.GetAllUsers()
	var usersDto []response.UserDto
	for _, user := range users {
		userDto := utils.MapToUserDto(user)
		usersDto = append(usersDto, userDto)
	}
	return usersDto
}

func (userDao *userV2Dao) GetAllDBUsersById(id string) response.UserDto {
	user := userDao.dao.GetUserById(id)
	userDto := utils.MapToUserDto(user)
	return userDto
}
