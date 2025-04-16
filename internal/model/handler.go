package model

import "github.com/gin-gonic/gin"

type UserHandler interface {
	CreateUser(c *gin.Context)
	GetUserByID(c *gin.Context)
}
