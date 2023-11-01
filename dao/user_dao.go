package dao

import (
	"github.com/oat431/api-tutorial-go/configs"
	"github.com/oat431/api-tutorial-go/models"
)

func GetAllUsers() []models.User {
	db := configs.ConnectDB()
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		panic(result.Error)
	}
	return users
}

func GetUserById(id string) models.User {
	db := configs.ConnectDB()
	var user models.User
	result := db.First(&user, id)
	if result.Error != nil {
		panic(result.Error)
	}
	return user
}
