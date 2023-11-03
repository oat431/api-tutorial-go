package repository

import (
	"github.com/oat431/api-tutorial-go/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() []models.User
	FindByID(id string) models.User
}

type userDB struct {
	connection *gorm.DB
}

func CreateUserRepository(db *gorm.DB) *userDB {
	return &userDB{
		connection: db,
	}
}

func (db *userDB) GetAllUsers() []models.User {
	var users []models.User
	result := db.connection.Find(&users)
	if result.Error != nil {
		panic(result.Error)
	}
	return users
}

func (db *userDB) FindByID(id string) models.User {
	var user models.User
	result := db.connection.First(&user, id)
	if result.Error != nil {
		panic(result.Error)
	}
	return user
}
