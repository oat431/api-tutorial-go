package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oat431/api-tutorial-go/services"
)

func GetAllDBUsers(c *gin.Context) {
	users := services.GetAllDBUsers()
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetAllDBUsersById(c *gin.Context) {
	id := c.Param("id")
	user := services.GetAllDBUsersById(id)
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
