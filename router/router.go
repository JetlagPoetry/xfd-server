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
		userGroup.POST("/login", handler.User.Login)           // 登录
		userGroup.POST("/submitRole", handler.User.SubmitRole) // 选择身份，提交认证信息
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
	orgGroup := r.Group("/api/v1/org")
	{
		orgGroup.POST("/applyPoint", handler.Org.ApplyPoint)            // 申请积分
		orgGroup.POST("/verifyPoint", handler.Org.VerifyPoint)          // 提交积分审核
		orgGroup.GET("/getApplyToVerify", handler.Org.GetApplyToVerify) // 获取待审核积分申请
		orgGroup.GET("/getApplys", handler.Org.GetApplys)               // 获取积分申请记录
		orgGroup.GET("/getApplyDetail", handler.Org.GetApplys)          // 查看本批积分明细

		orgGroup.GET("/getOrgMembers", handler.Org.GetOrgMembers)
		orgGroup.GET("/getPointRecordsByUser", handler.Org.GetPointRecordsByUser)

		orgGroup.POST("/verifyAccount", handler.Org.VerifyAccount)              // 提交用户审核
		orgGroup.GET("/getAccountToVerify", handler.Org.GetAccountToVerify)     // 获取待审核用户申请
		orgGroup.GET("/getAccountVerifyList", handler.Org.GetAccountVerifyList) // 获取用户申请记录
	}
	return r
}
