package services

import (
	"github.com/oat431/api-tutorial-go/models"
	"github.com/oat431/api-tutorial-go/payload/request"
	"github.com/oat431/api-tutorial-go/payload/response"
	"github.com/oat431/api-tutorial-go/utils"
)

func CreateUser(request request.UserRequest) response.UserDto {
	var user models.User
	var size = len(utils.MockUsers())
	user.ID = size + 1
	user.Firstname = request.Firstname
	user.Lastname = request.Lastname
	user.Email = request.Email
	user.Birthday = request.Birthday

	return utils.MapToUserDto(user)
}

func GetAllUsers() []response.UserDto {
	return utils.MockUsers()
}

func GetUserById(id int) response.UserDto {
	users := utils.MockUsers()
	for _, user := range users {
		if user.ID == id {
			return user
		}
	}
	return response.UserDto{}
}

func UpdateUser(request request.UserRequest, id int) response.UserDto {
	var users = utils.MockRawUsers()
	var updateUser models.User
	for _, user := range users {
		if user.ID == id {
			updateUser = user
			break
		}
	}
	updateUser.Firstname = request.Firstname
	updateUser.Lastname = request.Lastname
	updateUser.Email = request.Email
	updateUser.Birthday = request.Birthday

	return utils.MapToUserDto(updateUser)
}

func DeleteUser(id int) []response.UserDto {
	users := utils.MockUsers()
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
		}
	}
	return users
}
