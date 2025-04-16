package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/TurkmenistanRailways/payment/internal/config"
	"github.com/TurkmenistanRailways/payment/internal/model"
)

type UserHandler struct {
	UserService model.UserService
}

func NewUserHandler(service model.UserService, conf *config.Config) model.UserHandler {

	if conf.MOCK_HANDLER == "true" {
		return &MockUserHandler{UserService: service}
	}
	return &UserHandler{UserService: service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UserService.CreateUser(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.UserService.GetUserByID(c.Request.Context(), idInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
