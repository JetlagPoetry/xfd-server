package router

import (
	"github.com/gin-gonic/gin"
	"xfd-backend/handler"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	userGroup := r.Group("/api/v1/user")
	{
		userGroup.GET("/info", handler.User.GetUserInfo)
	}
	return r
}
