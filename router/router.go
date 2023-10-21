package router

import (
	"github.com/gin-gonic/gin"
	"xfd-backend/handler"
	"xfd-backend/pkg/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(middleware.Logger())
	//r.Use(middleware.UserAuthMiddleware("/api/v1/user/login")) // 登录校验
	userGroup := r.Group("/api/v1/user")
	{
		//userGroup.GET("/getInfo", handler.User.GetUserInfo)
		//userGroup.GET("/getToken", handler.User.GetOpenID)
		userGroup.POST("/login", handler.User.Login)
		//userGroup.POST("/modifyInfo", handler.User.ModifyInfo)
		//userGroup.POST("/modifyInfo", handler.User.ModifyInfo)
	}
	return r
}
