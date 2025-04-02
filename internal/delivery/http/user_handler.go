package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"architecture/internal/model"
	"architecture/internal/service"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{UserService: service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.UserService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
