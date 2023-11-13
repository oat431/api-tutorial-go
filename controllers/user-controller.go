package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oat431/api-tutorial-go/payload/request"
	"github.com/oat431/api-tutorial-go/services"
)

func CreateUser(c *gin.Context) {
	var request request.UserRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := services.CreateUser(request)
	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

// @Summary Get all mock user
// @Schemes
// @Description do ping
// @Tags Users V1
// @Accept json
// @Produce json
// @Success 200
// @Router /v1/users/ [get]
func GetAllUsers(c *gin.Context) {
	users := services.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// @summary Get mock user by id
// @Schemes
// @description  Get mock user by given id
// @tags Users V1
// @id GetUserById
// @param id path string true "id of mock user"
// @Produce json
// @Accept json
// @Success 200
// @Router /v1/users/{id} [get]
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
			"error": fmt.Sprintf("user id: %d not found", user_id),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func UpdateUser(c *gin.Context) {
	var request request.UserRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := c.Params.ByName("id")
	user_id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id should be a number",
		})
		return
	}

	chk := services.GetUserById(user_id)
	if chk.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("user id: %d not found", user_id),
		})
		return
	}

	updateUser := services.UpdateUser(request, user_id)
	c.JSON(http.StatusOK, gin.H{
		"user": updateUser,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	user_id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id should be a number",
		})
		return
	}
	chk := services.GetUserById(user_id)
	if chk.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("user id: %d not found", user_id),
		})
		return
	}
	users := services.DeleteUser(user_id)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
