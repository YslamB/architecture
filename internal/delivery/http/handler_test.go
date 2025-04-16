package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"

	"github.com/TurkmenistanRailways/payment/internal/config"
	postgresrepo "github.com/TurkmenistanRailways/payment/internal/repository/postgres_repository"
	"github.com/TurkmenistanRailways/payment/internal/service"
	"github.com/TurkmenistanRailways/payment/internal/storage/postgres"
)

func TestGetUserByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	conf := config.Init()
	db := postgres.Init()

	router := gin.Default()
	userRepository := postgresrepo.NewPostgresUserRepository(db, conf)
	userService := service.NewUserService(userRepository, conf)
	userHandler := NewUserHandler(userService, conf)
	router.GET("/users/:id", userHandler.GetUserByID)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
