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
	r.Use(middleware.UserAuthMiddleware("/api/v1/test/hello", "/api/v1/user/login", "/api/v1/user/sendCode")) // 登录校验, 参数为跳过登录的路由

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
		userGroup.POST("/sendCode", handler.User.SendCode)     // 发送验证码
		userGroup.POST("/login", handler.User.Login)           // 登录
		userGroup.POST("/submitRole", handler.User.SubmitRole) // 选择身份，提交认证信息
		userGroup.POST("/refreshToken", handler.User.RefreshToken)
		userGroup.GET("/info", handler.User.GetUserInfo)
		userGroup.POST("/modifyInfo", handler.User.ModifyInfo)
		userGroup.POST("/assignAdmin", handler.User.AssignAdmin)
		userGroup.GET("/getAdmins", handler.User.GetAdmins)
		userGroup.POST("/deleteUser", handler.User.DeleteUser)

		userGroup.GET("/getAddressList", handler.User.GetAddressList)
		userGroup.POST("/addAddress", handler.User.AddAddress)
		userGroup.POST("/modifyAddress", handler.User.ModifyAddress)
		userGroup.POST("/deleteAddress", handler.User.DeleteAddress)
	}
	purchaseGroup := r.Group("/api/v1/purchase")
	{
		purchaseGroup.GET("/getPurchases", handler.Purchase.GetPurchases)            // 查看本人采购单
		purchaseGroup.POST("/submitOrder", handler.Purchase.SubmitOrder)             // 提交采购单
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
		orgGroup.POST("/clearPoint", handler.Org.ClearPoint)            // 积分清零

		orgGroup.POST("/verifyAccount", handler.Org.VerifyAccount)              // 提交用户审核
		orgGroup.GET("/getAccountToVerify", handler.Org.GetAccountToVerify)     // 获取待审核用户申请
		orgGroup.GET("/getAccountVerifyList", handler.Org.GetAccountVerifyList) // 获取用户申请记录

		orgGroup.GET("/getOrganizations", handler.Org.GetOrganizations)             // 获取公司列表，模糊搜索
		orgGroup.GET("/getOrgMembers", handler.Org.GetOrgMembers)                   // 获取本公司员工列表
		orgGroup.GET("/getPointRecordsByApply", handler.Org.GetPointRecordsByApply) // 查看本批积分明细
		orgGroup.GET("/getPointRecordsByUser", handler.Org.GetPointRecordsByUser)   // 查看员工积分明细
		orgGroup.GET("/getPointRecords", handler.Org.GetPointRecords)               // 查看公司/本公司积分明细
	}
	//todo:/api/v1/area  500 而不是404
	mallGroup := r.Group("/api/v1/mall")
	{
		mallGroup.GET("/categories", handler.Mall.GetCategories) //获取商城分类信息
	}
	goodsGroup := r.Group("/api/v1/goods")
	{
		goodsGroup.POST("/addGoods", handler.Goods.AddGoods)                       //添加商品
		goodsGroup.GET("/getGoodsList", handler.Goods.GetGoodsList)                //获取商品列表
		goodsGroup.GET("/getGoodsDetail", handler.Goods.GetGoodsDetail)            //获取商品详情
		goodsGroup.GET("/getMyGoodsList", handler.Goods.GetMyGoodsList)            //后台获取商品列表
		goodsGroup.POST("/modifyMyGoods", handler.Goods.ModifyMyGoods)             //修改商品信息
		goodsGroup.DELETE("/deleteMyGoods", handler.Goods.DeleteMyGoods)           //删除商品
		goodsGroup.GET("/getMyGoodsDetail", handler.Goods.GetMyGoodsDetail)        //获取商品详情
		goodsGroup.POST("/modifyMyGoodsStatus", handler.Goods.ModifyMyGoodsStatus) //修改商品状态
	}
	orderGroup := r.Group("/api/v1/order")
	{
		orderGroup.POST("/addShoppingCart", handler.Order.AddShoppingCart)         //加入购物车
		orderGroup.GET("/getShoppingCart", handler.Order.GetShoppingCartList)      //获取购物车列表
		orderGroup.DELETE("/deleteShoppingCart", handler.Order.DeleteShoppingCart) //删除购物车商品
		orderGroup.POST("/modifyShoppingCart", handler.Order.ModifyShoppingCart)   //修改购物车商品数量
	}
	commonGroup := r.Group("/api/v1/common")
	{
		commonGroup.GET("/area", handler.Common.GetArea)                   //获取区域地址代码
		commonGroup.POST("/uploadFile", handler.Common.UploadFile)         //上传图片
		commonGroup.DELETE("/uploadFile", handler.Common.DeleteUploadFile) //删除图片
	}

	return r
}
