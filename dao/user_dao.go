package dao

import (
	"github.com/oat431/api-tutorial-go/models"
	"github.com/oat431/api-tutorial-go/repository"
)

type UserV2Dao interface {
	GetAllUsers() []models.User
	GetUserById(id string) models.User
	CreateUser(user models.User) models.User
	UpdateUser(user models.User) models.User
	DeleteUser(user models.User)
}

type userV2repository struct {
	repository repository.UserRepository
}

func CreateUserDao(repo repository.UserRepository) *userV2repository {
	return &userV2repository{repo}
}

func (userRepository *userV2repository) GetAllUsers() []models.User {
	return userRepository.repository.GetAllUsers()
}

func (userRepository *userV2repository) GetUserById(id string) models.User {
	return userRepository.repository.FindByID(id)
}

func (userRepository *userV2repository) CreateUser(user models.User) models.User {
	return userRepository.repository.Create(user)
}

func (userRepository *userV2repository) UpdateUser(user models.User) models.User {
	return userRepository.repository.Update(user)
}

func (userRepository *userV2repository) DeleteUser(user models.User) {
	userRepository.repository.Delete(user)
}
