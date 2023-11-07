package routes

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/oat431/api-tutorial-go/configs"
	"github.com/oat431/api-tutorial-go/controllers"
	"github.com/oat431/api-tutorial-go/dao"
	"github.com/oat431/api-tutorial-go/repository"
	"github.com/oat431/api-tutorial-go/services"
	"github.com/oat431/api-tutorial-go/resolver"
	"github.com/oat431/api-tutorial-go/schema"
)

var (
	DB             = configs.ConnectDB()
	userRepository = repository.CreateUserRepository(DB)
	userDao        = dao.CreateUserDao(userRepository)
	userService    = services.CreateUserService(userDao)
	userController = controllers.CreateUserController(userService)

	userResolver   = resolver.CreateUserResolver(userService)
	userSchema     = schema.CreateUserSchema(userResolver)
	rootSchema     = schema.CreateRootSchema(userSchema)
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(configs.SetupCORS()))

	gh := configs.SetupGql(rootSchema)

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
		userV2.GET("/pages", userController.GetAllDBUsersPagination)
		userV2.GET("/:id", userController.GetAllDBUsersById)
		userV2.POST("/", userController.CreateUser)
		userV2.PUT("/:id", userController.UpdateUser)
		userV2.DELETE("/:id", userController.DeleteUser)
	}

	graph := r.Group("/graphql")
	{
		graph.GET("/", gh)
		graph.POST("/", gh)
	}

	log.Println("Database connected")

	return r
}
