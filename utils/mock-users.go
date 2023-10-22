package utils

import (
	"github.com/oat431/api-tutorial-go/models"
	"github.com/oat431/api-tutorial-go/payload/response"
)

func MockUsers() []response.UserDto {
	user1 := models.User{
		ID:        1,
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "oat431@email1.com",
		Birthday:  "1999-01-01",
	}

	user2 := models.User{
		ID:        2,
		Firstname: "Jane",
		Lastname:  "Doe",
		Email:     "oat432@email2.com",
		Birthday:  "1998-01-02",
	}

	user3 := models.User{
		ID:        3,
		Firstname: "James",
		Lastname:  "Doe",
		Email:     "oat433@email3.com",
		Birthday:  "1997-01-03",
	}

	userDto1 := MapToUserDto(user1)
	userDto2 := MapToUserDto(user2)
	userDto3 := MapToUserDto(user3)

	var usersDto []response.UserDto
	usersDto = append(usersDto, userDto1)
	usersDto = append(usersDto, userDto2)
	usersDto = append(usersDto, userDto3)

	return usersDto

}
