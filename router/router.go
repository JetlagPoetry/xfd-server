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
	//r.Use(middleware.UserAuthMiddleware("/api/v1/user/login")) // 登录校验, 参数为跳过登录的路由
	userGroup := r.Group("/api/v1/user")
	{
		userGroup.POST("/sendCode", handler.User.SendCode)
		userGroup.POST("/login", handler.User.Login)
		userGroup.POST("/submitRole", handler.User.SubmitRole)
		userGroup.POST("/refreshToken", handler.User.RefreshToken)
		userGroup.GET("/info", handler.User.GetUserInfo)
		userGroup.POST("/modifyInfo", handler.User.ModifyInfo)
		//userGroup.GET("/getInfo", handler.User.GetUserInfo)
		//userGroup.GET("/getToken", handler.User.GetOpenID)
		//userGroup.POST("/modifyInfo", handler.User.ModifyInfo)
		//userGroup.POST("/modifyInfo", handler.User.ModifyInfo)
	}
	return r
}
