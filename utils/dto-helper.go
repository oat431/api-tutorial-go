package utils

import (
	"time"

	"github.com/morkid/paginate"
	"github.com/oat431/api-tutorial-go/models"
	"github.com/oat431/api-tutorial-go/payload/request"
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

func MapToUserModel(request request.UserRequest) models.User {
	return models.User{
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Email:     request.Email,
		Birthday:  request.Birthday,
	}
}

func MapToPageUser(pages paginate.Page) response.PageUserDto {
	var users []response.UserDto
	user_model := pages.Items.(*[]models.User)
	res := &user_model

	for _, user := range **res {
		userDto := MapToUserDto(user)
		users = append(users, userDto)
	}
	var pageUserDto response.PageUserDto
	pageUserDto.Items = users
	pageUserDto.Page = pages.Page
	pageUserDto.Size = pages.Size
	pageUserDto.MaxPage = pages.MaxPage
	pageUserDto.TotalPages = pages.TotalPages
	pageUserDto.Total = pages.Total
	pageUserDto.Last = pages.Last
	pageUserDto.First = pages.First

	return pageUserDto
}
