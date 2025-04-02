package route

import (
	"architecture/internal/delivery/http"
	"architecture/internal/model"
	postgresrepo "architecture/internal/repository/postgres_repository"
	"architecture/internal/service"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine, db model.DB) {
	userRoute := r.Group("/users")
	SetupUserRoutes(userRoute, db)
}

func SetupUserRoutes(r *gin.RouterGroup, db model.DB) {
	userRepository := postgresrepo.NewPostgresUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := http.NewUserHandler(userService)
	{
		r.POST("/", userHandler.CreateUser)
		// r.GET("/users/:id", userHandler.GetUserByID)
		// r.GET("/users", userHandler.GetAllUsers)
		// r.PUT("/users/:id", userHandler.UpdateUser)
		// r.DELETE("/users/:id", userHandler.DeleteUser)
	}
}
