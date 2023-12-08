package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/martian/log"
	"github.com/shopspring/decimal"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	log2 "log"
	"sync"
	"time"
	"xfd-backend/database/db"
	"xfd-backend/database/db/dao"
	"xfd-backend/database/db/enum"
	"xfd-backend/database/db/model"
	"xfd-backend/database/redis"
	"xfd-backend/database/repo"
	"xfd-backend/pkg/common"
	"xfd-backend/pkg/consts"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/utils"
	"xfd-backend/pkg/wechatpay"
	"xfd-backend/pkg/xerr"
)

type OrderService struct {
	goods          *dao.GoodsDao
	userDao        *dao.UserDao
	orderDao       *dao.OrderDao
	userAddressDao *dao.UserAddressDao
	orgDao         *dao.OrganizationDao
	pointRemainDao *dao.PointRemainDao
	pointRecordDao *dao.PointRecordDao
	userVerifyDao  *dao.UserVerifyDao
	inventoryLock  sync.Mutex
}

func NewOrderService() *OrderService {
	return &OrderService{
		goods:          dao.NewGoodsDao(),
		userDao:        dao.NewUserDao(),
		orderDao:       dao.NewOrderDao(),
		orgDao:         dao.NewOrganizationDao(),
		pointRemainDao: dao.NewPointRemainDao(),
		pointRecordDao: dao.NewPointRecordDao(),
		userAddressDao: dao.NewUserAddressDao(),
		userVerifyDao:  dao.NewUserVerifyDao(),
	}
}

func (s *OrderService) AddShoppingCart(ctx *gin.Context, req types.ShoppingCartAddReq) xerr.XErr {
	user, xrr := s.CheckUserRole(ctx, model.UserRoleCustomer)
	if xrr != nil {
		return xrr
	}
	//3.校验productValid
	productVariant, err := s.goods.GetProductVariantByProductVariantID(ctx, req.ProductVariantID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if productVariant == nil {
		return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("product %d not found", req.ProductVariantID))
	}
	//4.校验库存
	if *productVariant.Stock < req.Quantity {
		return xerr.WithCode(xerr.ErrorProductMoreCart, fmt.Errorf("product %d stock is %d want quantity %d not enough", req.ProductVariantID, *productVariant.Stock, req.Quantity))
	}
	goods, err := s.goods.GetGoodsByGoodsID(ctx, productVariant.GoodsID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if goods == nil || goods.Status == enum.GoodsStatusOffSale || goods.RetailStatus == enum.GoodsRetailSoldOut {
		return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("goods %d not found", productVariant.GoodsID))
	}
	//5.校验是否已经添加过购物车
	shoppingCart, err := s.orderDao.GetShoppingCartByUserIDAndProductVariantID(ctx, user.UserID, req.ProductVariantID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if shoppingCart != nil {
		//6.已经添加过购物车，数量增加
		updateValue := &model.ShoppingCart{
			Quantity: shoppingCart.Quantity + req.Quantity,
		}
		_, err = s.orderDao.UpdateShoppingCartByID(ctx, shoppingCart.ID, updateValue)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
		return nil
	}
	//7.新增购物车
	newShoppingCart := &model.ShoppingCart{
		UserID:              user.UserID,
		ProductVariantID:    req.ProductVariantID,
		GoodsID:             productVariant.GoodsID,
		GoodsSupplierUserID: goods.UserID,
		SKUCode:             productVariant.SKUCode,
		Quantity:            req.Quantity,
	}
	_, rr := s.orderDao.AddShoppingCart(ctx, newShoppingCart)
	if rr != nil {
		return xerr.WithCode(xerr.ErrorDatabase, rr)
	}
	return nil
}

func (s *OrderService) DeleteShoppingCart(ctx *gin.Context, req types.ShoppingCartDeleteReq) xerr.XErr {
	user, xrr := s.CheckUserRole(ctx, model.UserRoleCustomer)
	if xrr != nil {
		return xrr
	}
	err := s.orderDao.DeleteShoppingCartByIDsAndUserID(ctx, req.ShoppingCartIDs, user.UserID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil
}

func (s *OrderService) CheckUserRole(ctx *gin.Context, role model.UserRole) (*model.User, xerr.XErr) {
	//1.获取用户ID
	userID := common.GetUserID(ctx)
	//2.获取用户角色
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if user == nil {
		return nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("user %s not found", userID))
	}
	userRole := user.UserRole
	if role != 0 && userRole != role {
		return nil, xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, fmt.Errorf("user %s user role is %d, can not support", userID, userRole))
	}
	return user, nil
}

func (s *OrderService) ModifyShoppingCart(ctx *gin.Context, req types.ShoppingCartModifyReq) xerr.XErr {
	//1.校验用户是否是消费者
	user, xrr := s.CheckUserRole(ctx, model.UserRoleCustomer)
	if xrr != nil {
		return xrr
	}
	//2.校验购物车是否存在
	shoppingCart, err := s.orderDao.GetShoppingCartByID(ctx, req.ShoppingCartID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if shoppingCart == nil {
		return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("shopping cart %d not found", req.ShoppingCartID))
	}
	if shoppingCart.UserID != user.UserID {
		return xerr.WithCode(xerr.ErrorOperationForbidden, fmt.Errorf("shopping cart %d not belong to user %s", req.ShoppingCartID, user.UserID))
	}
	//3.校验购物车数量
	if req.ModifyQuantityType == enum.ModifyQuantityReduce && shoppingCart.Quantity <= req.Quantity {
		return xerr.WithCode(xerr.ErrorCartReduceMoreQuantity, fmt.Errorf("shopping cart %d quantity %d less than %d", req.ShoppingCartID, shoppingCart.Quantity, req.Quantity))
	}
	//4.校验产品库存
	productVariant, err := s.goods.GetProductVariantByProductVariantID(ctx, shoppingCart.ProductVariantID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if productVariant == nil {
		return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("product %d not found", shoppingCart.ProductVariantID))
	}

	if req.ModifyQuantityType == enum.ModifyQuantityAdd && *productVariant.Stock < req.Quantity+shoppingCart.Quantity {
		return xerr.WithCode(xerr.ErrorProductMoreCart, fmt.Errorf("product %d stock is %d want quantity %d not enough", shoppingCart.ProductVariantID, *productVariant.Stock, req.Quantity))
	}
	var updateQuantity int
	if req.ModifyQuantityType == enum.ModifyQuantityReduce {
		updateQuantity = shoppingCart.Quantity - req.Quantity
	} else {
		updateQuantity = shoppingCart.Quantity + req.Quantity
	}
	//5.更新购物车数量
	updateValue := &model.ShoppingCart{Quantity: updateQuantity}
	_, err = s.orderDao.UpdateShoppingCartByID(ctx, req.ShoppingCartID, updateValue)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil
}

func (s *OrderService) GetShoppingCartList(ctx *gin.Context, req types.ShoppingCartListReq) (*types.ShoppingCartListResp, xerr.XErr) {
	//1.校验用户是否是消费者
	user, xrr := s.CheckUserRole(ctx, model.UserRoleCustomer)
	if xrr != nil {
		return nil, xrr
	}
	req.UserID = user.UserID
	//todo: 优化不用删除，不满足条件的不查询出来，比如SKU下架，Goods库存不足
	shoppingCartList, total, err := s.orderDao.GetMyShoppingCartList(ctx, req)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	result := types.ShoppingCartListResp{PageResult: types.PageResult{PageNum: req.PageNum, PageSize: req.PageSize, TotalNum: total}}
	if len(shoppingCartList) == 0 {
		return &result, nil
	}
	shoppingCartListResp := make([]*types.ShoppingCartDetail, 0)
	for i, shoppingCart := range shoppingCartList {
		shoppingCartDetail := &types.ShoppingCartDetail{ShoppingCart: shoppingCartList[i]}
		productVariant, rr := s.goods.GetProductVariantByProductVariantID(ctx, shoppingCart.ProductVariantID)
		if rr != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, rr)
		}
		if productVariant == nil || *productVariant.Stock == 0 {
			err = s.orderDao.DeleteShoppingCartByProductVariantID(ctx, shoppingCart.ProductVariantID)
			if err != nil {
				log.Errorf(" delete invalid product shopping cart %d failed, err=%v", shoppingCart.ID, err)
				continue
			}
		}
		goods, _rr := s.goods.GetGoodsByGoodsID(ctx, shoppingCart.GoodsID)
		if _rr != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, _rr)
		}
		if goods == nil || goods.Status == enum.GoodsStatusOffSale || goods.RetailStatus == enum.GoodsRetailSoldOut {
			log.Errorf("goods %d not found", productVariant.GoodsID)
			err = s.orderDao.DeleteShoppingCartByGoodsID(ctx, shoppingCart.GoodsID)
			if err != nil {
				log.Errorf(" delete invalid product shopping cart %d failed, err=%v", shoppingCart.ID, err)
				continue
			}
		}
		if productVariant != nil {
			shoppingCartDetail.Stock = productVariant.Stock
			shoppingCartDetail.Price = productVariant.Price.Round(2).String()
			shoppingCartDetail.Name = goods.Name
			shoppingCartDetail.CoverURL = goods.GoodsFrontImage
			shoppingCartDetail.ProductAttr = productVariant.ProductAttr
		}
		shoppingCartListResp = append(shoppingCartListResp, shoppingCartDetail)
	}
	result.List = shoppingCartListResp
	return &result, nil
}

func (s *OrderService) CreateOrder(ctx *gin.Context, req types.CreateOrderReq) (result *types.CreateOrderResp, xErr xerr.XErr) {
	user, userAddress, shoppingCart, goods, productVariant, goodUser, xrr := s.checkOrderNormalInfo(ctx, req)
	if xrr != nil {
		return nil, xrr
	}
	transactionHandler := repo.NewTransactionHandler(db.Get())
	if xErr = transactionHandler.WithTransaction(ctx, func(ctx context.Context) xerr.XErr {
		if result, xErr = s.createOrderWithTX(ctx, req, user, userAddress, shoppingCart, goods, productVariant, goodUser); xErr != nil {
			return xErr
		}
		return nil
	}); xErr != nil {
		return nil, xErr
	}
	return result, nil
}

func (s *OrderService) createOrderWithTX(ctx context.Context, req types.CreateOrderReq, user *model.User, addr *model.UserAddress, shoppingCart *model.ShoppingCart, goods *model.Goods, productVariant *model.ProductVariant, goodUser *model.User) (*types.CreateOrderResp, xerr.XErr) {
	orderSn := fmt.Sprintf("DD%s%s", utils.GenerateOrder(), utils.GenerateRandomNumber(3))
	totalPrice := decimal.Zero
	postPrice := decimal.Zero
	if user.Point.Equal(decimal.Zero) {
		return nil, xerr.WithCode(xerr.ErrorOrderCreate, fmt.Errorf("user %s point is %s, please add point", user.UserID, user.Point.String()))
	}

	//扣除库存
	leftStock := *productVariant.Stock - shoppingCart.Quantity
	rowsAffected, err := s.goods.UpdateProductVariantByID(ctx, productVariant.ID, &model.ProductVariant{Stock: &leftStock})
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if rowsAffected == 0 {
		return nil, xerr.WithCode(xerr.ErrorOrderCreate, fmt.Errorf("reduce stock failed,please check shopping cart %d,name is %s,product variant id is %d", shoppingCart.ID, goods.Name, productVariant.ID))
	}
	//删除购物车
	err = s.orderDao.DeleteShoppingCartByIDCTX(ctx, shoppingCart.ID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	//计算价格
	totalPrice = totalPrice.Add(productVariant.Price.Mul(decimal.NewFromInt(int64(shoppingCart.Quantity)))).Add(goods.RetailShippingFee)
	postPrice = postPrice.Add(goods.RetailShippingFee)

	orderInfo := &model.OrderInfo{
		UserID:                        user.UserID,
		UserPhone:                     user.Phone,
		UserOrganizationName:          user.OrganizationName,
		ShoppingCartID:                req.ShoppingCartID,
		UnitPrice:                     productVariant.Price,
		TotalPrice:                    totalPrice,
		PostPrice:                     postPrice,
		Status:                        enum.OrderInfoCreated,
		OrderSn:                       orderSn,
		ExpiredAt:                     time.Now().Add(time.Minute * 15),
		SignerAddress:                 addr.Province + addr.City + addr.Region + addr.Address,
		SignerName:                    addr.Name,
		SingerMobile:                  addr.Phone,
		EstimatedDeliveryTime:         s.GetEstimatedDeliveryTime(goods.RetailShippingTime),
		Quantity:                      shoppingCart.Quantity,
		Name:                          goods.Name,
		Image:                         goods.GoodsFrontImage,
		ProductAttr:                   productVariant.ProductAttr,
		ProductVariantID:              productVariant.ID,
		SKUCode:                       productVariant.SKUCode,
		GoodsID:                       goods.ID,
		GoodsSupplierUserID:           goods.UserID,
		GoodsSupplierOrganizationName: goodUser.OrganizationName,
	}
	if req.Remark == "xfd-success" {
		orderInfo.Status = enum.OderInfoPaidSuccess
		payTime := time.Now().Add(time.Minute * 2)
		orderInfo.PayedAt = &payTime
		orderInfo.OtherMessage = "mock测试支付成功"
	}

	orderID, err := s.orderDao.CreateOrderInfo(ctx, orderInfo)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	orderInfo.ID = orderID
	//扣除积分
	payData, payWx, xrr := s.payForOrder(ctx, req.Code, user, orderInfo)
	if xrr != nil {
		return nil, xrr
	}
	if !payWx {
		// 只用积分支付
		orderInfo.Status = enum.OderInfoPaidSuccess
	} else {
		// 积分+微信支付
		orderInfo.Status = enum.OrderInfoPaidWaiting
	}
	return &types.CreateOrderResp{OrderID: orderInfo.ID, OrderSn: orderInfo.OrderSn, OrderStatus: orderInfo.Status, PayWx: payWx, PayData: payData}, nil
}

func (s *OrderService) PayOrder(ctx *gin.Context, req types.PayOrderReq) (*types.PayOrderResp, xerr.XErr) {
	//userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByPhone(ctx, "13300000000")
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	transactionHandler := repo.NewTransactionHandler(db.Get())
	if xErr := transactionHandler.WithTransaction(ctx, func(ctx context.Context) xerr.XErr {
		if _, _, xErr := s.payForOrder(ctx, "", user, &model.OrderInfo{ID: 123, TotalPrice: decimal.NewFromFloat(9.95)}); xErr != nil {
			return xErr
		}
		return nil
	}); xErr != nil {
		return nil, xErr
	}

	return &types.PayOrderResp{}, nil
}

// org->point_application->user->point_remain->point_record
func (s *OrderService) payForOrder(ctx context.Context, code string, user *model.User, order *model.OrderInfo) (*jsapi.PrepayWithRequestPaymentResponse, bool, xerr.XErr) {
	if user.OrganizationID == 0 {
		return nil, false, xerr.WithCode(xerr.ErrorUserOrgNotFound, errors.New("user not belong to org"))
	}

	ok := redis.Lock(fmt.Sprintf("user-point:user_id:%s", user.UserID), time.Minute*5)
	if !ok {
		return nil, false, xerr.WithCode(xerr.ErrorRedisLock, errors.New("get lock failed"))
	}
	defer redis.Unlock(fmt.Sprintf("user-point:user_id:%s", user.UserID))

	//var (
	//	org *model.Organization
	//	err error
	//)
	//org, err = s.orgDao.GetByIDForUpdateCTX(ctx, user.OrganizationID)
	//if err != nil {
	//	return nil, false, xerr.WithCode(xerr.ErrorDatabase, err)
	//}
	//user, err = s.userDao.GetByUserIDForUpdateCTX(ctx, user.UserID)
	//if err != nil {
	//	return nil, false, xerr.WithCode(xerr.ErrorDatabase, err)
	//}
	//if user == nil {
	//	return nil, false, xerr.WithCode(xerr.ErrorUserNotFound, errors.New("user not found"))
	//}
	//if user.Point.Equals(decimal.Zero) {
	//	return nil, false, xerr.WithCode(xerr.ErrorUserPointEmpty, errors.New("user do not have point"))
	//}
	//pointPrice, wxPrice := decimal.Zero, decimal.Zero
	//totalPrice := order.TotalPrice
	//payWx := false
	//var payData *jsapi.PrepayWithRequestPaymentResponse
	//payedAt := time.Now()
	//if user.Point.GreaterThanOrEqual(totalPrice) {
	//	// 只用积分支付
	//	pointPrice = totalPrice
	//	wxPrice = decimal.Zero
	//	xErr := s.payWithPoint(ctx, order, user, org, pointPrice, payWx)
	//	if xErr != nil {
	//		return nil, false, xErr
	//	}
	//	updateValue := &model.OrderInfo{
	//		PointPrice: totalPrice,
	//		PayedAt:    &payedAt,
	//		Status:     enum.OderInfoPaidSuccess,
	//	}
	//	xrr := s.UpdateOrderStatus(ctx, order.ID, updateValue)
	//	if xrr != nil {
	//		return nil, false, xerr.WithCode(xerr.ErrorDatabase, xrr)
	//	}
	//} else {
	//	// 积分+微信支付
	//	payWx = true
	//	pointPrice = user.Point
	//	wxPrice = totalPrice.Sub(user.Point)
	//
	//	xErr := s.payWithPoint(ctx, order, user, org, pointPrice, payWx)
	//	if xErr != nil {
	//		return nil, false, xErr
	//	}
	//	payData, xErr = s.payWithWx(ctx, code, order, user, wxPrice)
	//	if xErr != nil {
	//		return nil, false, xErr
	//	}
	//	updateValue := &model.OrderInfo{
	//		PointPrice: pointPrice,
	//		WxPrice:    wxPrice,
	//		Status:     enum.OrderInfoPaidWaiting,
	//		TradeNo:    *payData.PrepayId,
	//	}
	//	xrr := s.UpdateOrderStatus(ctx, order.ID, updateValue)
	//	if xrr != nil {
	//		return nil, false, xerr.WithCode(xerr.ErrorDatabase, xrr)
	//	}
	//}
	pointPrice, wxPrice := decimal.Zero, decimal.Zero
	totalPrice := order.TotalPrice
	payWx := false
	var payData *jsapi.PrepayWithRequestPaymentResponse
	// 积分+微信支付
	payWx = true
	pointPrice = user.Point
	wxPrice = totalPrice.Sub(user.Point)

	payData, xErr := s.payWithWx(ctx, code, order, user, wxPrice)
	if xErr != nil {
		return nil, false, xErr
	}
	updateValue := &model.OrderInfo{
		PointPrice: pointPrice,
		WxPrice:    wxPrice,
		Status:     enum.OrderInfoPaidWaiting,
		TradeNo:    *payData.PrepayId,
	}
	xrr := s.UpdateOrderStatus(ctx, order.ID, updateValue)
	if xrr != nil {
		return nil, false, xerr.WithCode(xerr.ErrorDatabase, xrr)
	}
	return payData, payWx, nil
}

func (s *OrderService) payWithPoint(ctx context.Context, order *model.OrderInfo, user *model.User, org *model.Organization, point decimal.Decimal, payWx bool) xerr.XErr {
	userLeft := user.Point.Sub(point)
	var (
		remainList []*model.PointRemain
		err        error
	)
	err = s.userDao.UpdateByUserIDInTxCTX(ctx, user.UserID, &model.User{Point: userLeft})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	err = s.orgDao.UpdateByIDInTxCTX(ctx, user.OrganizationID, &model.Organization{Point: org.Point.Sub(point)})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	remainList, err = s.pointRemainDao.ListValidByUserIDCTX(ctx, user.UserID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	pointNeed := point.Copy()
	remainIDs := make([]int, 0)
	for _, remain := range remainList {
		if pointNeed.Equal(decimal.Zero) {
			break
		}
		remainIDs = append(remainIDs, int(remain.ID))
		var spend decimal.Decimal
		if remain.PointRemain.LessThan(pointNeed) {
			// 将pointRemain花费完
			spend = remain.PointRemain
			pointNeed = pointNeed.Sub(remain.PointRemain)
		} else {
			// 将pointLeft花费完
			spend = pointNeed
			pointNeed = pointNeed.Sub(pointNeed)
		}
		record := &model.PointRecord{
			UserID:             user.UserID,
			OrganizationID:     user.OrganizationID,
			ChangePoint:        spend.Mul(decimal.NewFromInt(-1)),
			PointApplicationID: remain.PointApplicationID,
			PointID:            int(remain.ID),
			OrderID:            int(order.ID),
			Type:               model.PointRecordTypeSpend,
			Status: func() model.PointRecordStatus {
				if payWx {
					return model.PointRecordStatusInit
				} else {
					return model.PointRecordStatusConfirmed
				}
			}(),
			Comment: consts.PointCommentSpend,
		}
		err = s.pointRemainDao.UpdateByIDInTxCTX(ctx, int(remain.ID), &model.PointRemain{PointRemain: remain.PointRemain.Sub(spend)})
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
		err = s.pointRecordDao.CreateInTxCTX(ctx, record)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}
	if pointNeed.GreaterThan(decimal.Zero) {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil
}

func (s *OrderService) payWithWx(ctx context.Context, code string, order *model.OrderInfo, user *model.User, wxPrice decimal.Decimal) (*jsapi.PrepayWithRequestPaymentResponse, xerr.XErr) {
	// 获取openid
	openResp, xErr := wechatpay.GetWxOpenID(ctx, code)
	if xErr != nil {
		return nil, xErr
	}

	// 请求wx预付单，保存
	orderResp, xErr := wechatpay.CreateOrder(ctx, order.OrderSn, "艺图小程序商城", openResp.OpenID,
		wxPrice.Mul(decimal.NewFromInt(100)).Floor().IntPart())
	if xErr != nil {
		return nil, xErr
	}
	// 更新订单信息
	err := s.orderDao.UpdateOrderInfoByID(ctx, order.ID, &model.OrderInfo{Status: enum.OrderInfoPaidWaiting})
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return orderResp, nil
}

func (s *OrderService) ApplyExchange(ctx *gin.Context, req types.ApplyExchangeReq) xerr.XErr {
	_, xrr := s.CheckUserRole(ctx, model.UserRoleAdmin)
	if xrr != nil {
		return xrr
	}
	orderInfo, err := s.orderDao.GetOrderInfoByID(ctx, req.QueryOrderID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if orderInfo == nil {
		return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("query order id%d not found", req.QueryOrderID))
	}
	if orderInfo.Status != enum.OderInfoAfterSale && orderInfo.Status != enum.OderInfoReceived && orderInfo.Status != enum.OderInfoShipped {
		return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("query order id%d status is %d, can not apply exchange", req.QueryOrderID, orderInfo.Status))
	}
	refundList, err := s.orderDao.GetOrderRefundByOrderID(ctx, req.QueryOrderID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if len(refundList) > 0 {
		for _, refund := range refundList {
			if refund.ManuallyClosed == enum.ManuallyClosedYes || refund.AfterSaleType == enum.AfterSaleTypeReturnAndRefund {
				return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("query order id%d has refund or closed id%d, can not apply exchange", req.QueryOrderID, refund.ID))
			}
		}
	}
	createNewRefund := &model.OrderRefund{
		OrderID:       req.QueryOrderID,
		OderSn:        orderInfo.OrderSn,
		AfterSaleType: enum.AfterSaleTypeExchange,
		Reason:        req.Reason,
	}
	transactionHandler := repo.NewTransactionHandler(db.Get())
	if xErr := transactionHandler.WithTransaction(ctx, func(ctx context.Context) xerr.XErr {
		if _err := s.applyRefundRecord(ctx, orderInfo, createNewRefund); _err != nil {
			return _err
		}
		return nil
	}); xErr != nil {
		return xErr
	}
	return nil
}

func (s *OrderService) refundPoint(ctx context.Context, user, operator *model.User, point decimal.Decimal, orderID int) xerr.XErr {
	// 加公司积分
	org, err := s.orgDao.GetByIDForUpdateCTX(ctx, user.OrganizationID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	err = s.orgDao.UpdateByIDInTxCTX(ctx, user.OrganizationID, &model.Organization{Point: org.Point.Add(point)})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 加员工积分
	user, err = s.userDao.GetByUserIDForUpdateCTX(ctx, user.UserID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	err = s.userDao.UpdateByUserIDInTxCTX(ctx, user.UserID, &model.User{Point: user.Point.Add(point)})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	records, err := s.pointRecordDao.ListByOrderIDInTxCTX(ctx, user.UserID, int(orderID))
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	for _, record := range records {
		remain, err := s.pointRemainDao.GetByIDForUpdateCTX(ctx, record.PointID)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}

		err = s.pointRemainDao.UpdateByIDInTxCTX(ctx, record.PointID, &model.PointRemain{PointRemain: remain.PointRemain.Sub(record.ChangePoint)})
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
		newRecord := &model.PointRecord{
			UserID:             user.UserID,
			OrganizationID:     user.OrganizationID,
			ChangePoint:        record.ChangePoint.Mul(decimal.NewFromInt(-1)),
			PointApplicationID: record.PointApplicationID,
			PointID:            record.PointID,
			OrderID:            record.OrderID,
			Type:               model.PointRecordTypeRefund,
			Status:             model.PointRecordStatusConfirmed,
			Comment:            consts.PointCommentRefund,
		}
		if operator != nil {
			newRecord.OperateUserID = operator.UserID
			newRecord.OperateUsername = operator.Username
		}
		err = s.pointRecordDao.CreateInTxCTX(ctx, newRecord)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}

	return nil
}

func (s *OrderService) revertPoint(ctx context.Context, user *model.User, point decimal.Decimal, orderID int) xerr.XErr {
	// 加公司积分
	org, err := s.orgDao.GetByIDForUpdateCTX(ctx, user.OrganizationID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	err = s.orgDao.UpdateByIDInTxCTX(ctx, user.OrganizationID, &model.Organization{Point: org.Point.Add(point)})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 加员工积分
	user, err = s.userDao.GetByUserIDForUpdateCTX(ctx, user.UserID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	err = s.userDao.UpdateByUserIDInTxCTX(ctx, user.UserID, &model.User{Point: user.Point.Add(point)})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	records, err := s.pointRecordDao.ListByOrderIDInTxCTX(ctx, user.UserID, int(orderID))
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	for _, record := range records {
		remain, err := s.pointRemainDao.GetByIDForUpdateCTX(ctx, record.PointID)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}

		err = s.pointRemainDao.UpdateByIDInTxCTX(ctx, record.PointID, &model.PointRemain{PointRemain: remain.PointRemain.Sub(record.ChangePoint)})
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}

	}

	err = s.pointRecordDao.UpdateByOrderID(ctx, orderID, &model.PointRecord{Status: model.PointRecordStatusCancelled})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return nil
}

func (s *OrderService) inventoryLockProcess(ctx context.Context, shoppingCart *model.ShoppingCart, goods *model.Goods, productVariant *model.ProductVariant) xerr.XErr {
	inventoryLock, err := s.goods.GetInventoryByProductVariantID(ctx, shoppingCart.ProductVariantID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if inventoryLock == nil {
		//生成锁定库存
		inventoryLock = &model.Inventory{
			ProductVariantID: shoppingCart.ProductVariantID,
			SKUCode:          shoppingCart.SKUCode,
			GoodsID:          shoppingCart.GoodsID,
			LockStock:        shoppingCart.Quantity,
		}
		err = s.goods.CreateInventory(ctx, inventoryLock)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}
	if inventoryLock != nil {
		if *productVariant.Stock < shoppingCart.Quantity+inventoryLock.LockStock {
			return xerr.WithCode(xerr.ErrorSomeoneNotPaid, fmt.Errorf(" you can not create order,product %d stock not enough,but someone not paid,please check shopping cart %d,name is %s", shoppingCart.ProductVariantID, shoppingCart.ID, goods.Name))
		}
		updateInventory := &model.Inventory{LockStock: shoppingCart.Quantity + inventoryLock.LockStock}
		_, err = s.goods.UpdateInventoryByID(ctx, productVariant.ID, updateInventory)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}
	return nil
}

func (s *OrderService) GetEstimatedDeliveryTime(retailDeliveryTime enum.RetailDeliveryTime) time.Time {
	switch retailDeliveryTime {
	case enum.RetailDeliveryTimeWithin24Hours:
		return time.Now().Add(time.Hour * 24)
	case enum.RetailDeliveryTimeWithin48Hours:
		return time.Now().Add(time.Hour * 48)
	case enum.RetailDeliveryTimeWithin7Days:
		return time.Now().Add(time.Hour * 24 * 7)
	default:
		return time.Now()
	}
}

// CreatePreOrder 用户获取预订单信息
func (s *OrderService) CreatePreOrder(ctx *gin.Context, req types.CreateOrderReq) (*types.CreatePreOrderResp, xerr.XErr) {
	//1.校验用户是否是消费者
	user, userAddress, shoppingCart, goods, productVariant, _, xrr := s.checkOrderNormalInfo(ctx, req)
	if xrr != nil {
		return nil, xrr
	}

	totalPrice := decimal.Zero
	orderAddress := &types.PreOrderAddress{
		Name:    userAddress.Name,
		Phone:   userAddress.Phone,
		Address: userAddress.Province + userAddress.City + userAddress.Region + userAddress.Address,
	}

	//计算价格
	totalPrice = totalPrice.Add(productVariant.Price.Mul(decimal.NewFromInt(int64(shoppingCart.Quantity)))).Add(goods.RetailShippingFee)
	pointPrice, wxPrice := decimal.Zero, decimal.Zero
	if user.Point.GreaterThanOrEqual(totalPrice) {
		// 只用积分支付
		pointPrice = totalPrice
	} else {
		// 积分+微信支付
		pointPrice = user.Point
		wxPrice = totalPrice.Sub(user.Point)
	}
	result := &types.CreatePreOrderResp{
		PreOrderAddress: orderAddress,
		TotalPrice:      totalPrice.Round(2).String(),
		PointPrice:      pointPrice.Round(2).String(),
		WxPrice:         wxPrice.Round(2).String(),
		UserPoint:       user.Point.Round(2).String(),
		ShoppingCartID:  shoppingCart.ID,
		SKUCode:         shoppingCart.SKUCode,
		Price:           productVariant.Price.Round(2).String(),
		Quantity:        shoppingCart.Quantity,
		Name:            goods.Name,
		CoverURL:        goods.GoodsFrontImage,
		ProductAttr:     productVariant.ProductAttr,
		PostPrice:       goods.RetailShippingFee.Round(2).String(),
	}

	return result, nil
}

func (s *OrderService) UpdateOrderStatus(ctx context.Context, orderID int32, value *model.OrderInfo) xerr.XErr {
	rowsAffected, err := s.orderDao.UpdateOrderInfoByIDCTX(ctx, orderID, value)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if rowsAffected == 0 {
		return xerr.WithCode(xerr.ErrorOrderCreate, fmt.Errorf("update order status failed,order id is %d", orderID))
	}
	return nil
}

func (s *OrderService) GetOrderList(ctx *gin.Context, req types.OrderListReq) (*types.OrderListResp, xerr.XErr) {
	user, rr := s.CheckUserRole(ctx, 0)
	if rr != nil {
		return nil, rr
	}
	if user.UserRole != model.UserRoleAdmin {
		if req.CheckMiniAppParams() != nil {
			return nil, xerr.WithCode(xerr.InvalidParams, req.CheckMiniAppParams())
		}
	}
	switch user.UserRole {
	case model.UserRoleCustomer:
		req.UserID = user.UserID
		orderList, total, err := s.orderDao.CustomerGetQueryOrderList(ctx, req)
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		return &types.OrderListResp{PageResult: types.PageResult{TotalNum: total, PageNum: req.PageNum, PageSize: req.PageSize}, List: orderList}, nil
	case model.UserRoleSupplier:
		req.UserID = user.UserID
		orderList, total, err := s.orderDao.SupplierGetQueryOrderList(ctx, req)
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		return &types.OrderListResp{PageResult: types.PageResult{TotalNum: total, PageNum: req.PageNum, PageSize: req.PageSize}, List: orderList}, nil
	case model.UserRoleAdmin:
		orderList, total, err := s.orderDao.AdminGetQueryOrderList(ctx, req)
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		return &types.OrderListResp{PageResult: types.PageResult{TotalNum: total, PageNum: req.PageNum, PageSize: req.PageSize}, List: orderList}, nil
	}
	return nil, nil
}

func (s *OrderService) FillShipmentInfo(ctx *gin.Context, req types.FillShipmentInfoReq) xerr.XErr {
	//1.校验用户是否是供应商
	user, rr := s.CheckUserRole(ctx, model.UserRoleSupplier)
	if rr != nil {
		return rr
	}
	//2.校验子订单有消息
	orderInfo, err := s.orderDao.GetOrderInfoByID(ctx, req.QueryOrderID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if orderInfo == nil {
		return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("query order id%d not found", req.QueryOrderID))
	}
	if orderInfo.GoodsSupplierUserID != user.UserID {
		return xerr.WithCode(xerr.ErrorOperationForbidden, fmt.Errorf("user %s is not goods supplier,cannot fill shipment info", user.UserID))
	}
	currentTime := time.Now()
	updateValue := &model.OrderInfo{
		Status:          enum.OderInfoShipped,
		ShipmentCompany: req.ShipmentCompany,
		ShipmentSn:      req.ShipmentSn,
		DeliveryTime:    &currentTime,
	}

	if orderInfo.ShipmentSn != "" {
		updateValue.OtherMessage = orderInfo.OtherMessage + "\n" + "历史物流信息:快递公司-" + orderInfo.ShipmentCompany + "单号-" + orderInfo.ShipmentSn + "时间-" + orderInfo.DeliveryTime.Format("2006-01-02 15:04:05")
	}
	err = s.orderDao.UpdateOrderInfoByID(ctx, req.QueryOrderID, updateValue)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil
}

func (s *OrderService) checkOrderNormalInfo(ctx *gin.Context, req types.CreateOrderReq) (*model.User, *model.UserAddress, *model.ShoppingCart, *model.Goods, *model.ProductVariant, *model.User, xerr.XErr) {
	//1.校验用户是否是消费者
	user, xrr := s.CheckUserRole(ctx, model.UserRoleCustomer)
	if xrr != nil {
		return nil, nil, nil, nil, nil, nil, xrr
	}
	if user.Point.IsZero() {
		return nil, nil, nil, nil, nil, nil, xerr.WithCode(xerr.ErrorOrderCreate, fmt.Errorf("user point is zero"))
	}
	//2.校验用户地址是否存在
	userAddress, err := s.userAddressDao.GetByID(ctx, req.UserAddressID)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if userAddress == nil {
		return nil, nil, nil, nil, nil, nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("user address %d not found", req.UserAddressID))
	}
	if userAddress.UserID != user.UserID {
		return nil, nil, nil, nil, nil, nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("user address %d not belong to user %s", req.UserAddressID, user.UserID))
	}
	//3.校验购物车商品是否存在
	shoppingCart, rr := s.orderDao.GetShoppingCartByUserIDAndShoppingCartID(ctx, user.UserID, req.ShoppingCartID)
	if rr != nil {
		return nil, nil, nil, nil, nil, nil, xerr.WithCode(xerr.ErrorDatabase, rr)
	}
	if shoppingCart == nil {
		return nil, nil, nil, nil, nil, nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("shopping cart %d not found", req.ShoppingCartID))
	}
	//校验商品信息
	goods, rr := s.goods.GetGoodsByGoodsID(ctx, shoppingCart.GoodsID)
	if rr != nil {
		return nil, nil, nil, nil, nil, nil, xerr.WithCode(xerr.ErrorDatabase, rr)
	}
	if goods == nil || goods.Status == enum.GoodsStatusOffSale || goods.RetailStatus == enum.GoodsRetailSoldOut {
		return nil, nil, nil, nil, nil, nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("goods %d not found,please check shopping cart %d,name is %s", shoppingCart.GoodsID, shoppingCart.ID, goods.Name))
	}
	goodUser, err := s.userDao.GetByUserID(ctx, goods.UserID)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if goodUser == nil {
		return nil, nil, nil, nil, nil, nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("goods user  %s not found", goods.UserID))
	}
	//校验产品库存
	productVariant, rr := s.goods.GetProductVariantByProductVariantID(ctx, shoppingCart.ProductVariantID)
	if rr != nil {
		return nil, nil, nil, nil, nil, nil, xerr.WithCode(xerr.ErrorDatabase, rr)
	}
	if productVariant == nil {
		return nil, nil, nil, nil, nil, nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("product %d not found", shoppingCart.ProductVariantID))
	}
	if *productVariant.Stock < shoppingCart.Quantity {
		return nil, nil, nil, nil, nil, nil, xerr.WithCode(xerr.ErrorStockNotEnough, fmt.Errorf("product %d stock not enough,please check shopping cart %d,name is %s", shoppingCart.ProductVariantID, shoppingCart.ID, goods.Name))
	}
	return user, userAddress, shoppingCart, goods, productVariant, goodUser, nil
}

func (s *OrderService) ConfirmReceipt(ctx *gin.Context, req types.ConfirmReceiptReq) xerr.XErr {
	//1.校验用户是否是消费者
	user, xrr := s.CheckUserRole(ctx, model.UserRoleCustomer)
	if xrr != nil {
		return xrr
	}
	//2.订单有效性
	orderInfo, err := s.orderDao.GetOrderInfoByID(ctx, req.QueryOrderID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if orderInfo == nil {
		return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("query order id%d not found", req.QueryOrderID))
	}
	if orderInfo.UserID != user.UserID {
		return xerr.WithCode(xerr.ErrorOperationForbidden, fmt.Errorf("user %s is not order owner,cannot confirm receipt", user.UserID))
	}
	if orderInfo.Status != enum.OderInfoShipped && orderInfo.Status != enum.OderInfoPaidSuccess {
		return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("order status is %d,not allow confirm receipt", orderInfo.Status))
	}
	currentTime := time.Now()
	rowsAffected, er := s.orderDao.UpdateOrderInfoByIDCTX(ctx, req.QueryOrderID, &model.OrderInfo{Status: enum.OderInfoReceived, ConfirmTime: &currentTime})
	if er != nil {
		return xerr.WithCode(xerr.ErrorDatabase, er)
	}
	if rowsAffected == 0 {
		return xerr.WithCode(xerr.ErrorOperationForbidden, fmt.Errorf("order status is %d,not allow confirm receipt", orderInfo.Status))
	}
	return nil
}

func (s *OrderService) GetOrderDetail(ctx *gin.Context, req types.ConfirmReceiptReq) (*types.OrderDetailResp, xerr.XErr) {
	isOrderClosed := false
	_, xrr := s.CheckUserRole(ctx, model.UserRoleAdmin)
	if xrr != nil {
		return nil, xrr
	}
	orderInfo, err := s.orderDao.GetOrderInfoByID(ctx, req.QueryOrderID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if orderInfo == nil {
		return nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("query order id%d not found", req.QueryOrderID))
	}
	orderDetailInfo := types.OrderInfo{
		OrderID:    orderInfo.ID,
		OrderSn:    orderInfo.OrderSn,
		Status:     orderInfo.Status,
		CreatedAt:  &orderInfo.CreatedAt,
		PayedAt:    orderInfo.PayedAt,
		TotalPrice: orderInfo.TotalPrice.Round(2).String(),
		WxPrice:    orderInfo.WxPrice.Round(2).String(),
		PointPrice: orderInfo.PointPrice.Round(2).String(),
	}
	goodsInfo := types.GoodsInfo{
		Name:        orderInfo.Name,
		GoodsID:     orderInfo.GoodsID,
		SKUCode:     orderInfo.SKUCode,
		Image:       orderInfo.Image,
		ProductAttr: orderInfo.ProductAttr,
		Quantity:    orderInfo.Quantity,
		UintPrice:   orderInfo.UnitPrice.Round(2).String(),
		PostPrice:   orderInfo.PostPrice.Round(2).String(),
	}
	buyerInfo := types.BuyerInfo{
		UserName:             orderInfo.UserID,
		UserPhone:            orderInfo.UserPhone,
		UserOrganizationName: orderInfo.UserOrganizationName,
		SingerName:           orderInfo.SignerName,
		SingerPhone:          orderInfo.SingerMobile,
		SingerAddr:           orderInfo.SignerAddress,
	}
	var sellerInfo types.SellerInfo

	sellerUser, err := s.userDao.GetByUserID(ctx, orderInfo.GoodsSupplierUserID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if sellerUser != nil {
		sellerInfo = types.SellerInfo{
			Name:             sellerUser.Username,
			Phone:            sellerUser.Phone,
			OrganizationName: sellerUser.OrganizationName,
		}
	}
	verify, err := s.userVerifyDao.GetByUserID(ctx, orderInfo.GoodsSupplierUserID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if verify != nil {
		sellerInfo.Position = verify.Position
	}
	refundList, err := s.orderDao.GetOrderRefundByOrderID(ctx, req.QueryOrderID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	afterSaleRecords := make([]*types.AfterSaleRecord, 0)
	var manuallyCloseOrder *types.ManuallyCloseOrder
	if len(refundList) != 0 {
		for _, refund := range refundList {
			if refund.ManuallyClosed == enum.ManuallyClosedYes {
				manuallyCloseOrder = &types.ManuallyCloseOrder{
					CreatedAt:       &refund.CreatedAt,
					Reason:          refund.Reason,
					ReturnPointType: refund.ReturnPointType,
				}
				isOrderClosed = true
				break
			}
			afterSaleRecord := &types.AfterSaleRecord{
				CreatedAt:       &refund.CreatedAt,
				Reason:          refund.Reason,
				AfterSaleType:   refund.AfterSaleType,
				ReturnPointType: refund.ReturnPointType,
			}
			if refund.AfterSaleType == enum.AfterSaleTypeReturnAndRefund {
				isOrderClosed = true
			}
			afterSaleRecords = append(afterSaleRecords, afterSaleRecord)
		}
	}
	orderRecord := types.OrderRecord{
		CreatedAt:          orderInfo.CreatedAt,
		PayedAt:            orderInfo.PayedAt,
		DeliveryTime:       orderInfo.DeliveryTime,
		ConfirmTime:        orderInfo.ConfirmTime,
		ManuallyCloseOrder: manuallyCloseOrder,
		AfterSaleRecords:   afterSaleRecords,
	}
	result := &types.OrderDetailResp{
		IsOrderClosed: isOrderClosed,
		OrderInfo:     orderDetailInfo,
		GoodsInfo:     goodsInfo,
		BuyerInfo:     buyerInfo,
		SellerInfo:    sellerInfo,
		OrderRecord:   orderRecord,
	}
	return result, nil
}

func (s *OrderService) CloseOrder(ctx *gin.Context, req types.CloseOrderReq) xerr.XErr {
	operator, xrr := s.CheckUserRole(ctx, model.UserRoleAdmin)
	if xrr != nil {
		return xrr
	}
	orderInfo, err := s.orderDao.GetOrderInfoByID(ctx, req.QueryOrderID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if orderInfo == nil {
		return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("query order id%d not found", req.QueryOrderID))
	}
	if orderInfo.Status != enum.OderInfoPaidSuccess {
		return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("order status is %d,not allow close", orderInfo.Status))
	}
	user, err := s.userDao.GetByUserID(ctx, orderInfo.UserID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if user == nil {
		return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("order user not found"))
	}

	transactionHandler := repo.NewTransactionHandler(db.Get())
	if xErr := transactionHandler.WithTransaction(ctx, func(ctx context.Context) xerr.XErr {
		if _err := s.closeOrder(ctx, req, user, operator, orderInfo); _err != nil {
			return _err
		}
		return nil
	}); xErr != nil {
		return xErr
	}
	return nil

}

func (s *OrderService) closeOrder(ctx context.Context, req types.CloseOrderReq, user, operator *model.User, orderInfo *model.OrderInfo) xerr.XErr {
	orderRefund, xErr := s.lockOrderForClose(ctx, req, orderInfo)
	if xErr != nil {
		return xErr
	}
	if orderRefund.ReturnPointStatus == enum.ReturnPointStatusWaitReturn && !orderRefund.NeedRefundPoint.Equal(decimal.Zero) {
		xErr = s.refundPoint(ctx, user, operator, orderRefund.NeedRefundPoint, int(orderRefund.OrderID))
		if xErr != nil {
			return xErr
		}
		//更新orderRefund状态
		_, err := s.orderDao.UpdateOrderRefundByID(ctx, orderRefund.ID, &model.OrderRefund{ReturnPointStatus: enum.ReturnPointStatusReturned})
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}
	return nil
}

func (s *OrderService) lockOrderForClose(ctx context.Context, req types.CloseOrderReq, orderInfo *model.OrderInfo) (*model.OrderRefund, xerr.XErr) {
	ok := redis.Lock(fmt.Sprintf("lock-order:order_sn:%s", orderInfo.OrderSn), time.Minute*5)
	if !ok {
		return nil, xerr.WithCode(xerr.ErrorRedisLock, errors.New("get lock failed"))
	}
	defer redis.Unlock(fmt.Sprintf("lock-order:order_sn:%s", orderInfo.OrderSn))
	rowsAffected, er := s.orderDao.UpdateOrderInfoByIDCTX(ctx, req.QueryOrderID, &model.OrderInfo{Status: enum.OderInfoAfterSale})
	if er != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, er)
	}
	if rowsAffected == 0 {
		return nil, xerr.WithCode(xerr.ErrorDatabase, fmt.Errorf("close order status failed,order id is %d", req.QueryOrderID))
	}
	createRefund := &model.OrderRefund{
		OrderID:         req.QueryOrderID,
		OderSn:          orderInfo.OrderSn,
		ManuallyClosed:  enum.ManuallyClosedYes,
		Reason:          req.Reason,
		ReturnPointType: req.ReturnPointType,
	}
	if req.ReturnPointType == enum.ReturnPointYes {
		createRefund.NeedRefundPoint = orderInfo.PointPrice
		createRefund.ReturnPointStatus = enum.ReturnPointStatusWaitReturn
	}
	id, err := s.orderDao.CreateOrderRefund(ctx, createRefund)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	createRefund.ID = id
	return createRefund, nil
}

func (s *OrderService) applyRefundRecord(ctx context.Context, orderInfo *model.OrderInfo, createNewRefund *model.OrderRefund) xerr.XErr {
	ok := redis.Lock(fmt.Sprintf("lock-order:order_sn:%s", orderInfo.OrderSn), time.Minute*5)
	if !ok {
		return xerr.WithCode(xerr.ErrorRedisLock, errors.New("get lock failed"))
	}
	defer redis.Unlock(fmt.Sprintf("lock-order:order_sn:%s", orderInfo.OrderSn))
	//更新售后记录
	_, err := s.orderDao.CreateOrderRefund(ctx, createNewRefund)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	//更新订单状态
	if orderInfo.Status != enum.OderInfoAfterSale {
		rowsAffected, err := s.orderDao.UpdateOrderInfoByIDCTX(ctx, createNewRefund.OrderID, &model.OrderInfo{Status: enum.OderInfoAfterSale})
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
		if rowsAffected == 0 {
			return xerr.WithCode(xerr.ErrorDatabase, fmt.Errorf("exchange order status failed,order id is %d", createNewRefund.OrderID))
		}
	}
	return nil
}

func (s *OrderService) ApplyRefund(ctx *gin.Context, req types.ApplyRefundReq) xerr.XErr {
	operator, xrr := s.CheckUserRole(ctx, model.UserRoleAdmin)
	if xrr != nil {
		return xrr
	}
	orderInfo, err := s.orderDao.GetOrderInfoByID(ctx, req.QueryOrderID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if orderInfo == nil {
		return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("query order id%d not found", req.QueryOrderID))
	}
	if orderInfo.Status != enum.OderInfoAfterSale && orderInfo.Status != enum.OderInfoReceived && orderInfo.Status != enum.OderInfoShipped {
		return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("query order id%d status is %d, can not apply refund", req.QueryOrderID, orderInfo.Status))
	}
	user, err := s.userDao.GetByUserID(ctx, orderInfo.UserID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if user == nil {
		return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("order user not found"))
	}
	refundList, err := s.orderDao.GetOrderRefundByOrderID(ctx, req.QueryOrderID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if len(refundList) > 0 {
		for _, refund := range refundList {
			if refund.ManuallyClosed == enum.ManuallyClosedYes || refund.AfterSaleType == enum.AfterSaleTypeReturnAndRefund {
				return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("query order id%d has refund or closed id%d, can not apply refund", req.QueryOrderID, refund.ID))
			}
		}
	}
	transactionHandler := repo.NewTransactionHandler(db.Get())
	if xErr := transactionHandler.WithTransaction(ctx, func(ctx context.Context) xerr.XErr {
		if _err := s.applyRefundTX(ctx, req, orderInfo, operator, user); _err != nil {
			return _err
		}
		return nil
	}); xErr != nil {
		return xErr
	}

	return nil
}

func (s *OrderService) applyRefundTX(ctx context.Context, req types.ApplyRefundReq, orderInfo *model.OrderInfo, operator *model.User, user *model.User) xerr.XErr {
	createNewRefund := &model.OrderRefund{
		OrderID:         req.QueryOrderID,
		OderSn:          orderInfo.OrderSn,
		AfterSaleType:   enum.AfterSaleTypeReturnAndRefund,
		ReturnPointType: req.ReturnPointType,
		Reason:          req.Reason,
	}
	if req.ReturnPointType == enum.ReturnPointNo {
		xrr := s.applyRefundRecord(ctx, orderInfo, createNewRefund)
		if xrr != nil {
			return xrr
		}
		return nil
	}
	if req.ReturnPointType == enum.ReturnPointYes {
		createNewRefund.NeedRefundPoint = orderInfo.PointPrice
		createNewRefund.ReturnPointStatus = enum.ReturnPointStatusWaitReturn
		//创建售后记录
		xrr := s.applyRefundRecord(ctx, orderInfo, createNewRefund)
		if xrr != nil {
			return xrr
		}
		if !createNewRefund.NeedRefundPoint.Equal(decimal.Zero) {
			xErr := s.refundPoint(ctx, user, operator, createNewRefund.NeedRefundPoint, int(createNewRefund.OrderID))
			if xErr != nil {
				return xErr
			}
			//更新orderRefund状态
			_, err := s.orderDao.UpdateOrderRefundByID(ctx, createNewRefund.ID, &model.OrderRefund{ReturnPointStatus: enum.ReturnPointStatusReturned})
			if err != nil {
				return xerr.WithCode(xerr.ErrorDatabase, err)
			}
		}

	}
	return nil
}

func (s *OrderService) GetCustomerService(ctx *gin.Context, req types.GetCustomerServiceReq) (*types.GetCustomerServiceResp, xerr.XErr) {
	_, xrr := s.CheckUserRole(ctx, model.UserRoleCustomer)
	if xrr != nil {
		return nil, xrr
	}
	supplierUser, err := s.userDao.GetByUserID(ctx, req.SupplierUserID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if supplierUser == nil {
		return nil, xerr.WithCode(xerr.InvalidParams, fmt.Errorf("supplier user %s not found", req.SupplierUserID))
	}
	return &types.GetCustomerServiceResp{Phone: supplierUser.Phone}, nil
}

func (s *OrderService) BatchPaymentLookup(ctx context.Context) xerr.XErr {
	list, err := s.orderDao.ListByStatus(ctx, enum.OrderInfoPaidWaiting)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	log2.Println("[OrderService] BatchPaymentLookup called, waiting2pay orders count=", len(list))

	for _, order := range list {
		transactionHandler := repo.NewTransactionHandler(db.Get())
		if xErr := transactionHandler.WithTransaction(ctx, func(ctx context.Context) xerr.XErr {
			_, xErr := s.paymentLookup(ctx, order, false)
			if xErr != nil {
				return xErr
			}
			return nil
		}); xErr != nil {
			return xErr
		}
	}
	return nil
}

// paymentLookup 查看订单是否已支付成功，如果支付成功，则执行paymentConfirm。调用cancel=true时，主动关闭订单
func (s *OrderService) paymentLookup(ctx context.Context, order *model.OrderInfo, cancel bool) (enum.OrderInfoStatus, xerr.XErr) {
	if order.Status != enum.OrderInfoPaidWaiting {
		return order.Status, nil
	}

	lock := redis.Lock(fmt.Sprintf("payment-look-up:order_id:%d", order.ID), time.Minute*5)
	if !lock {
		return 0, xerr.WithCode(xerr.ErrorRedisLock, errors.New("get lock failed"))
	}
	defer redis.Unlock(fmt.Sprintf("payment-look-up:order_id:%d", order.ID))

	resp, cErr := wechatpay.LookupOrder(ctx, order.OrderSn)
	if cErr != nil {
		return 0, cErr
	}

	if *resp.TradeState == consts.WECHAT_PAY_TRADE_STATE {
		status, cErr := s.paymentConfirm(ctx, resp)
		if cErr != nil {
			return 0, cErr
		}
		return status, nil
	}
	if cancel || (order.Status == enum.OrderInfoPaidWaiting && order.CreatedAt.Add(time.Minute*consts.WECHAT_PAY_EXPIRE_MINUTE).Before(time.Now())) {
		cErr = s.paymentCancel(ctx, order)
		if cErr != nil {
			return 0, cErr
		}
	}
	return order.Status, nil
}

func (s *OrderService) paymentConfirm(ctx context.Context, req *payments.Transaction) (enum.OrderInfoStatus, xerr.XErr) {
	order, err := s.orderDao.GetOrderInfoByOrderSn(ctx, *req.OutTradeNo)
	if err != nil {
		return 0, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	// 修改订单状态
	if order.Status != enum.OrderInfoPaidWaiting {
		return order.Status, nil
	}

	successTime, err := time.Parse("2006-01-02T15:04:05+MST", *req.SuccessTime)
	if err != nil {
		return 0, xerr.WithCode(xerr.InvalidParams, err)
	}
	_, err = s.orderDao.UpdateOrderInfoByIDCTX(ctx, order.ID, &model.OrderInfo{
		PayedAt: &successTime,
		Status:  enum.OrderInfoPaidPointConfirmWaiting,
	})
	if err != nil {
		return 0, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return enum.OderInfoPaidSuccess, nil
}

func (s *OrderService) paymentCancel(ctx context.Context, order *model.OrderInfo) xerr.XErr {
	if order.Status != enum.OrderInfoPaidWaiting {
		return xerr.WithCode(xerr.ErrorDatabase, errors.New("order has been processed"))
	}

	cErr := wechatpay.CancelOrder(ctx, order.OrderSn)
	if cErr != nil {
		return cErr
	}

	// 恢复库存
	variant, err := s.goods.GetProductVariantByProductVariantIDForUpdate(ctx, order.ProductVariantID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	leftStock := *variant.Stock + order.Quantity
	rowsAffected, err := s.goods.UpdateProductVariantByID(ctx, order.ProductVariantID, &model.ProductVariant{Stock: &leftStock})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if rowsAffected == 0 {
		return xerr.WithCode(xerr.ErrorOrderCreate, errors.New("reduce stock failed"))
	}

	user, err := s.userDao.GetByUserID(ctx, order.UserID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	err = s.revertPoint(ctx, user, order.PointPrice, int(order.ID))
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return nil
}

func (s *OrderService) PaymentConfirm(ctx context.Context, req *payments.Transaction) xerr.XErr {
	switch *req.TradeState {
	case consts.WECHAT_PAY_TRADE_STATE:
		transactionHandler := repo.NewTransactionHandler(db.Get())
		if xErr := transactionHandler.WithTransaction(ctx, func(ctx context.Context) xerr.XErr {
			if _, err := s.paymentConfirm(ctx, req); err != nil {
				return err
			}
			return nil
		}); xErr != nil {
			return xErr
		}
	}

	return nil
}

func (s *OrderService) BatchPointConfirm(ctx context.Context) xerr.XErr {
	list, err := s.orderDao.ListByStatus(ctx, enum.OrderInfoPaidPointConfirmWaiting)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	log2.Println("[OrderService] BatchPointConfirm called, waiting2pay orders count=", len(list))

	for _, order := range list {
		transactionHandler := repo.NewTransactionHandler(db.Get())
		if xErr := transactionHandler.WithTransaction(ctx, func(ctx context.Context) xerr.XErr {
			xErr := s.pointConfirm(ctx, order)
			if xErr != nil {
				return xErr
			}
			return nil
		}); xErr != nil {
			return xErr
		}
	}
	return nil
}

func (s *OrderService) pointConfirm(ctx context.Context, order *model.OrderInfo) xerr.XErr {
	err := s.pointRecordDao.UpdateByOrderID(ctx, int(order.ID), &model.PointRecord{Status: model.PointRecordStatusConfirmed})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	_, err = s.orderDao.UpdateOrderInfoByIDCTX(ctx, order.ID, &model.OrderInfo{Status: enum.OderInfoPaidSuccess})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil
}
