package utils

import (
	"time"

	"github.com/oat431/api-tutorial-go/models"
	"github.com/oat431/api-tutorial-go/payload/response"
)

func GetAge(user models.User) int {
	currentYear := time.Now().Year()
	birthdayYear, err := time.Parse("2006-01-02", user.Birthday)
	if err != nil {
		return 0
	}
	return currentYear - birthdayYear.Year()
}

func MapToUserDto(user models.User) response.UserDto {
	return response.UserDto{
		ID:    user.ID,
		Name:  user.Firstname + " " + user.Lastname,
		Email: user.Email,
		Age:   GetAge(user),
	}
}
