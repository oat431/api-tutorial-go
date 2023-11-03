package services

import (
	"github.com/oat431/api-tutorial-go/dao"
	"github.com/oat431/api-tutorial-go/payload/request"
	"github.com/oat431/api-tutorial-go/payload/response"
	"github.com/oat431/api-tutorial-go/utils"
)

type UserV2Service interface {
	GetAllDBUsers() []response.UserDto
	GetAllDBUsersById(id string) response.UserDto
	CreateUser(user request.UserRequest) response.UserDto
	UpdateUser(id string, user request.UserRequest) response.UserDto
	DeleteUser(id string)
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

func (userDao *userV2Dao) CreateUser(user request.UserRequest) response.UserDto {
	userModel := utils.MapToUserModel(user)
	userModel = userDao.dao.CreateUser(userModel)
	userDto := utils.MapToUserDto(userModel)
	return userDto
}

func (userDao *userV2Dao) UpdateUser(id string, user request.UserRequest) response.UserDto {
	updatedUser := userDao.dao.GetUserById(id)
	updatedUser.Firstname = user.Firstname
	updatedUser.Lastname = user.Lastname
	updatedUser.Email = user.Email
	updatedUser.Birthday = user.Birthday
	updatedUser = userDao.dao.UpdateUser(updatedUser)
	userDto := utils.MapToUserDto(updatedUser)
	return userDto
}

func (userDao *userV2Dao) DeleteUser(id string) {
	user := userDao.dao.GetUserById(id)
	userDao.dao.DeleteUser(user)
}
