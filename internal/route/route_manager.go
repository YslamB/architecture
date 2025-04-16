package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/TurkmenistanRailways/payment/internal/config"
	"github.com/TurkmenistanRailways/payment/internal/delivery/http"
	postgresrepo "github.com/TurkmenistanRailways/payment/internal/repository/postgres_repository"
	"github.com/TurkmenistanRailways/payment/internal/service"
)

func Init(r *gin.Engine, db *pgxpool.Pool, conf *config.Config) {
	userRoute := r.Group("/users")
	SetupUserRoutes(userRoute, db, conf)
}

func SetupUserRoutes(r *gin.RouterGroup, db *pgxpool.Pool, conf *config.Config) {
	userRepository := postgresrepo.NewPostgresUserRepository(db, conf)
	userService := service.NewUserService(userRepository, conf)
	userHandler := http.NewUserHandler(userService, conf)
	{
		r.POST("/", userHandler.CreateUser)
		r.GET("/users/:id", userHandler.GetUserByID)
	}
}
