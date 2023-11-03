package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oat431/api-tutorial-go/payload/request"
	"github.com/oat431/api-tutorial-go/services"
)

type UserV2Controller interface {
	GetAllDBUsers(c *gin.Context)
	GetAllDBUsersById(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userV2Service struct {
	service services.UserV2Service
}

func CreateUserController(service services.UserV2Service) *userV2Service {
	return &userV2Service{service}
}

func (userV2Service *userV2Service) GetAllDBUsers(c *gin.Context) {
	users := userV2Service.service.GetAllDBUsers()
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (userV2Service *userV2Service) GetAllDBUsersById(c *gin.Context) {
	id := c.Param("id")
	user := userV2Service.service.GetAllDBUsersById(id)
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (userV2Service *userV2Service) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := userV2Service.service.CreateUser(userRequest)
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (userV2Service *userV2Service) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := userV2Service.service.UpdateUser(id, userRequest)
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (userV2Service *userV2Service) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userV2Service.service.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
