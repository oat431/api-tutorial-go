package services

import (
	"github.com/oat431/api-tutorial-go/payload/response"
	"github.com/oat431/api-tutorial-go/utils"
)

func GetAllUsers() []response.UserDto {
	return utils.MockUsers()
}
