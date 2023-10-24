package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oat431/api-tutorial-go/services"
)

func GetAllUsers(c *gin.Context) {
	users := services.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	user_id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id should be a number",
		})
		return
	}
	user := services.GetUserById(user_id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
