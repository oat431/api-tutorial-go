package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/oat431/api-tutorial-go/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello-world", controllers.HelloWorld)
		v1.GET("/users", controllers.GetAllUsers)
	}
	return r
}
