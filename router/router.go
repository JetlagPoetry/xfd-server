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
	r.Use(middleware.UserAuthMiddleware("/api/v1/test/hello", "/api/v1/user/login", "/api/v1/mall", "/api/v1/common", "/api/v1/goods", "/api/v1/backstage")) // 登录校验, 参数为跳过登录的路由
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
		purchaseGroup.GET("/getPurchases", handler.Purchase.GetPurchases) // 查看本人采购单
		purchaseGroup.POST("/submitOrder", handler.Purchase.SubmitOrder)  // 提交采购单
		//purchaseGroup.POST("/modifyOrder", handler.Purchase.ModifyOrder)
		purchaseGroup.POST("/modifyOrderStatus", handler.Purchase.ModifyOrderStatus) // 删除或结束采购单
		purchaseGroup.GET("/getQuotes", handler.Purchase.GetQuotes)                  // 查看采购单的报价列表
		purchaseGroup.GET("/getStatistics", handler.Purchase.GetStatistics)          // 查看采购商统计数据
		purchaseGroup.POST("/answerQuote", handler.Purchase.AnswerQuote)             // 采购商回复报价
	}
	supplyGroup := r.Group("/api/v1/supply")
	{
		supplyGroup.GET("/getPurchases", handler.Supply.GetPurchases)             // 查看所有采购单
		supplyGroup.GET("/getQuotes", handler.Supply.GetQuotes)                   // 查看采购单对应报价
		supplyGroup.POST("/submitQuote", handler.Supply.SubmitQuote)              // 提交报价
		supplyGroup.GET("/getQuotedPurchases", handler.Supply.GetQuotedPurchases) // 查看报价过的所有采购单
		supplyGroup.GET("/getStatistics", handler.Supply.GetStatistics)           // 查看采购商统计数据
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

		orgGroup.GET("/getOrganizations", handler.Org.GetOrganizations)
		orgGroup.GET("/getOrgMembers", handler.Org.GetOrgMembers)
		orgGroup.GET("/getPointRecordsByUser", handler.Org.GetPointRecordsByUser)

	}
	//todo:/api/v1/area  500 而不是404
	mallGroup := r.Group("/api/v1/mall")
	{
		mallGroup.GET("/categories", handler.Mall.GetCategories) //获取商城分类信息
	}
	goodsGroup := r.Group("/api/v1/goods")
	{
		goodsGroup.POST("/addGoods", handler.Goods.AddGoods)        //添加商品
		goodsGroup.GET("/getGoodsList", handler.Goods.GetGoodsList) //获取商品列表
	}
	backstageGroup := r.Group("/api/v1/backstage")
	{
		backstageGroup.GET("/getGoodsList", handler.Goods.BackstageGetGoodsList) //后台管理获取商品列表
	}
	commonGroup := r.Group("/api/v1/common")
	{
		commonGroup.GET("/area", handler.Common.GetArea)                   //获取区域地址代码
		commonGroup.POST("/uploadFile", handler.Common.UploadFile)         //上传图片
		commonGroup.DELETE("/uploadFile", handler.Common.DeleteUploadFile) //删除图片
	}

	return r
}
