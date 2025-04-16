package http

import (
	"net/http"

	"github.com/TurkmenistanRailways/payment/internal/model"
	"github.com/gin-gonic/gin"
)

type MockUserHandler struct {
	UserService model.UserService
}

func (h *MockUserHandler) CreateUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *MockUserHandler) GetUserByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User retrieved successfully"})
}
