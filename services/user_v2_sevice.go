package services

import (
	"github.com/oat431/api-tutorial-go/dao"
	"github.com/oat431/api-tutorial-go/payload/response"
	"github.com/oat431/api-tutorial-go/utils"
)

func GetAllDBUsers() []response.UserDto {
	users := dao.GetAllUsers()
	var usersDto []response.UserDto
	for _, user := range users {
		userDto := utils.MapToUserDto(user)
		usersDto = append(usersDto, userDto)
	}
	return usersDto
}

func GetAllDBUsersById(id string) response.UserDto {
	user := dao.GetUserById(id)
	userDto := utils.MapToUserDto(user)
	return userDto
}
