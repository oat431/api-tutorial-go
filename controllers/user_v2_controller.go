package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oat431/api-tutorial-go/services"
)

type UserV2Controller interface {
	GetAllDBUsers(c *gin.Context)
	GetAllDBUsersById(c *gin.Context)
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
