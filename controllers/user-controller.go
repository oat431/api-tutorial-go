package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oat431/api-tutorial-go/services"
)

func GetAllUsers(c *gin.Context) {
	users := services.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
