package services

import (
	"github.com/oat431/api-tutorial-go/payload/response"
	"github.com/oat431/api-tutorial-go/utils"
)

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
