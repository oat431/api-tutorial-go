package dao

import (
	"github.com/oat431/api-tutorial-go/models"
	"github.com/oat431/api-tutorial-go/repository"
)

type UserV2Dao interface {
	GetAllUsers() []models.User
	GetUserById(id string) models.User
}

type userV2repository struct {
	repository repository.UserRepository
}

func CreateUserV2Repository(repo repository.UserRepository) *userV2repository {
	return &userV2repository{repo}
}

func (userRepository *userV2repository) GetAllUsers() []models.User {
	return userRepository.repository.GetAllUsers()
}

func (userRepository *userV2repository) GetUserById(id string) models.User {
	return userRepository.repository.FindByID(id)
}
