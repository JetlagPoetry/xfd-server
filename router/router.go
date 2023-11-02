package router

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"xfd-backend/handler"
	"xfd-backend/pkg/middleware"
	"xfd-backend/pkg/response"
	"xfd-backend/pkg/xerr"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(middleware.Logger())
	r.Use(middleware.UserAuthMiddleware("/api/v1/test/hello", "/api/v1/user/login", "/api/v1/mall/categories", "/api/v1/area")) // 登录校验, 参数为跳过登录的路由
	testGroup := r.Group("/api/v1/test")
	{
		testGroup.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, response.RespSuccess(c, "hello world"))
		})
		testGroup.GET("/error", func(c *gin.Context) {
			c.JSON(http.StatusOK, response.RespError(c, xerr.WithCode(xerr.InvalidParams, errors.New("encounter error"))))
		})
	}
	userGroup := r.Group("/api/v1/user")
	{
		userGroup.POST("/login", handler.User.Login)           // 登录
		userGroup.POST("/submitRole", handler.User.SubmitRole) // 选择身份，提交认证信息
		userGroup.POST("/refreshToken", handler.User.RefreshToken)
		userGroup.GET("/info", handler.User.GetUserInfo)
		userGroup.POST("/modifyInfo", handler.User.ModifyInfo)
		userGroup.POST("/assignAdmin", handler.User.AssignAdmin)
	}
	purchaseGroup := r.Group("/api/v1/purchase")
	{
		purchaseGroup.GET("/getOrders", handler.Purchase.GetOrders)
		purchaseGroup.POST("/submitOrder", handler.Purchase.SubmitOrder)
		//purchaseGroup.POST("/modifyOrder", handler.Purchase.ModifyOrder)
		purchaseGroup.POST("/modifyOrderStatus", handler.Purchase.ModifyOrderStatus)
		purchaseGroup.GET("/getQuotes", handler.Purchase.GetQuotes)
		purchaseGroup.POST("/submitQuote", handler.Purchase.SubmitQuote)
		purchaseGroup.GET("/getStatistics", handler.Purchase.SubmitQuote)
	}
	supplyGroup := r.Group("/api/v1/supply")
	{
		supplyGroup.GET("/getPurchases", handler.Supply.GetPurchases)
		supplyGroup.GET("/getQuotes", handler.Supply.GetQuotes)
	}
	orgGroup := r.Group("/api/v1/org")
	{
		orgGroup.POST("/applyPoint", handler.Org.ApplyPoint)            // 申请积分
		orgGroup.POST("/verifyPoint", handler.Org.VerifyPoint)          // 提交积分审核
		orgGroup.GET("/getApplyToVerify", handler.Org.GetApplyToVerify) // 获取待审核积分申请
		orgGroup.GET("/getApplys", handler.Org.GetApplys)               // 获取积分申请记录
		orgGroup.GET("/getApplyDetail", handler.Org.GetApplys)          // 查看本批积分明细

		orgGroup.POST("/verifyAccount", handler.Org.VerifyAccount)              // 提交用户审核
		orgGroup.GET("/getAccountToVerify", handler.Org.GetAccountToVerify)     // 获取待审核用户申请
		orgGroup.GET("/getAccountVerifyList", handler.Org.GetAccountVerifyList) // 获取用户申请记录

		orgGroup.GET("/getOrgMembers", handler.Org.GetOrgMembers)
		orgGroup.GET("/getPointRecordsByUser", handler.Org.GetPointRecordsByUser)
	}
	mallGroup := r.Group("/api/v1/mall")
	{
		mallGroup.GET("/categories", handler.Mall.GetCategories) //获取商城分类信息
	}
	v1Group := r.Group("/api/v1")
	{
		v1Group.GET("/area", handler.Mall.GetArea) //获取区域地址代码
	}

	return r
}
