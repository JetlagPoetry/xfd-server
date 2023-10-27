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
	}
	//messageGroup := r.Group("/api/v1/message")
	//{
	//	messageGroup.GET("/getConversations", handler.Message.GetConversations)
	//	messageGroup.GET("/getMessages", handler.Message.GetMessages)
	//	//messageGroup.POST("/send", handler.Message.Send)
	//}
	purchaseGroup := r.Group("/api/v1/purchase")
	{
		purchaseGroup.GET("/getOrders", handler.Purchase.GetOrders)
		purchaseGroup.POST("/submitOrder", handler.Purchase.SubmitOrder)
		purchaseGroup.POST("/modifyOrder", handler.Purchase.ModifyOrder)
		purchaseGroup.POST("/modifyOrderStatus", handler.Purchase.ModifyOrderStatus)
		purchaseGroup.GET("/getQuotes", handler.Purchase.GetQuotes)
		purchaseGroup.POST("/submitQuote", handler.Purchase.SubmitQuote)
	}
	return r
}
