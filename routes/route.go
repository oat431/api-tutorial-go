package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/oat431/api-tutorial-go/configs"
	"github.com/oat431/api-tutorial-go/controllers"
	"github.com/oat431/api-tutorial-go/dao"
	"github.com/oat431/api-tutorial-go/repository"
	"github.com/oat431/api-tutorial-go/services"
)

var (
	DB             = configs.ConnectDB()
	userRepository = repository.CreateUserRepository(DB)
	userDao        = dao.CreateUserDao(userRepository)
	userService    = services.CreateUserService(userDao)
	userController = controllers.CreateUserController(userService)
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello-world", controllers.HelloWorld)
	}

	user := r.Group("/api/v1/users")
	{
		user.POST("/", controllers.CreateUser)
		user.GET("/", controllers.GetAllUsers)
		user.GET("/:id", controllers.GetUserById)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}

	userV2 := r.Group("/api/v2/users")
	{
		userV2.GET("/", userController.GetAllDBUsers)
		userV2.GET("/:id", userController.GetAllDBUsersById)
	}
	log.Println("Database connected")

	return r
}
