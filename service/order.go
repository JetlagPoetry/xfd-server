package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/martian/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"os"
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
		return xerr.WithCode(xerr.ErrorProductExistCart, fmt.Errorf("product %d already in shopping cart", req.ProductVariantID))
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
	xrr := s.payForOrder(ctx, req.Code, user, orderInfo)
	if xrr != nil {
		return nil, xrr
	}
	if user.Point.GreaterThanOrEqual(totalPrice) {
		// 只用积分支付
		orderInfo.Status = enum.OderInfoPaidSuccess
	} else {
		// 积分+微信支付
		orderInfo.Status = enum.OrderInfoPaidWaiting
	}
	return &types.CreateOrderResp{OrderID: orderInfo.ID, OrderSn: orderInfo.OrderSn, OrderStatus: orderInfo.Status}, nil
}

func (s *OrderService) PayOrder(ctx *gin.Context, req types.PayOrderReq) (*types.PayOrderResp, xerr.XErr) {
	//userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByPhone(ctx, "13300000000")
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	transactionHandler := repo.NewTransactionHandler(db.Get())
	if xErr := transactionHandler.WithTransaction(ctx, func(ctx context.Context) xerr.XErr {
		if xErr := s.payForOrder(ctx, "", user, &model.OrderInfo{ID: 123, TotalPrice: decimal.NewFromFloat(9.95)}); xErr != nil {
			return xErr
		}
		return nil
	}); xErr != nil {
		return nil, xErr
	}

	return &types.PayOrderResp{}, nil
}

// org->point_application->user->point_remain->point_record
func (s *OrderService) payForOrder(ctx context.Context, code string, user *model.User, order *model.OrderInfo) xerr.XErr {

	if user.OrganizationID == 0 {
		return xerr.WithCode(xerr.ErrorUserOrgNotFound, errors.New("user not belong to org"))
	}

	ok := redis.Lock(fmt.Sprintf("user-point:user_id:%s", user.UserID), time.Minute*5)
	if !ok {
		return xerr.WithCode(xerr.ErrorRedisLock, errors.New("get lock failed"))
	}
	defer redis.Unlock(fmt.Sprintf("user-point:user_id:%s", user.UserID))

	var (
		org *model.Organization
		err error
	)
	org, err = s.orgDao.GetByIDForUpdateCTX(ctx, user.OrganizationID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	user, err = s.userDao.GetByUserIDForUpdateCTX(ctx, user.UserID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if user == nil {
		return xerr.WithCode(xerr.ErrorUserNotFound, errors.New("user not found"))
	}
	if user.Point.Equals(decimal.Zero) {
		return xerr.WithCode(xerr.ErrorUserPointEmpty, errors.New("user do not have point"))
	}
	pointPrice, wxPrice := decimal.Zero, decimal.Zero
	totalPrice := order.TotalPrice
	payWx := false
	//var wechatPay *types.WechatPay
	payedAt := time.Now()
	if user.Point.GreaterThanOrEqual(totalPrice) {
		// 只用积分支付
		pointPrice = totalPrice
		wxPrice = decimal.Zero
		xErr := s.payWithPoint(ctx, order, user, org, pointPrice, payWx)
		if xErr != nil {
			return xErr
		}
		updateValue := &model.OrderInfo{
			PointPrice: totalPrice,
			PayedAt:    &payedAt,
			Status:     enum.OderInfoPaidSuccess,
		}
		xrr := s.UpdateOrderStatus(ctx, order.ID, updateValue)
		if xrr != nil {
			return xerr.WithCode(xerr.ErrorDatabase, xrr)
		}
	} else {
		// 积分+微信支付
		payWx = true
		pointPrice = user.Point
		wxPrice = totalPrice.Sub(user.Point)

		xErr := s.payWithPoint(ctx, order, user, org, pointPrice, payWx)
		if xErr != nil {
			return xErr
		}
		_, xErr = s.payWithWx(ctx, code, order, user, wxPrice)
		if xErr != nil {
			return xErr
		}
		updateValue := &model.OrderInfo{
			PointPrice: pointPrice,
			WxPrice:    wxPrice,
			Status:     enum.OrderInfoPaidWaiting,
			TradeNo:    "wexin",
		}
		xrr := s.UpdateOrderStatus(ctx, order.ID, updateValue)
		if xrr != nil {
			return xerr.WithCode(xerr.ErrorDatabase, xrr)
		}
	}
	//此时总订单初始状态 enum.OrderInfoPaidWaiting（待付款） 子订单初始状态 enum.OrderProductVariantDetailPending（待付款）
	// todo 修改订单，支付金额、微信支付字段、订单状态、修改子订单、订单状态、支付时间
	//todo 返回订单状态（积分支付成功/积分支付失败/微信等待支付）
	//看情况是否需要区分 订单待付款/支付待付款 两种状态 解决方案加一个 订单创建的初始态 enum.OrderInfoCreated enum.OrderProductVariantDetailCreated

	//if !payWx {
	//	err = s.orderDao.UpdateOrderInfoByID(ctx, int(order.ID), &model.OrderInfo{Status: enum.OderInfoPaidSuccess})
	//	if err != nil {
	//		return xerr.WithCode(xerr.ErrorDatabase, err)
	//	}
	//} else {
	//fmt.Sprintf(wechatPay)
	//}

	return nil
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
	remainList, err = s.pointRemainDao.ListByUserIDCTX(ctx, user.UserID)
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

func (s *OrderService) payWithWx(ctx context.Context, code string, order *model.OrderInfo, user *model.User, wxPrice decimal.Decimal) (*types.WechatPay, xerr.XErr) {
	// 请求wx预付单，保存
	openResp, xErr := wechatpay.GetWxOpenID(ctx, code)
	if xErr != nil {
		return nil, xErr
	}

	orderResp, xErr := wechatpay.CreateOrder(ctx, order.OrderSn, "新发地小程序商城", openResp.OpenID, wxPrice.Mul(decimal.NewFromInt(100)).Floor().IntPart())
	if xErr != nil {
		return nil, xErr
	}

	// 更新订单信息
	err := s.orderDao.UpdateOrderInfoByID(ctx, order.ID, &model.OrderInfo{Status: enum.OderInfoPaidSuccess})
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	appID := os.Getenv("WC_ID")
	timestamp := time.Now().Unix()
	nonce := utils.RandStr(32)
	pack := fmt.Sprintf("prepay_id=%s", orderResp.PrepayID)
	data := fmt.Sprintf("%s\n%d\n%s\n%s\n", appID, timestamp, nonce, pack)
	sign, err := wechatpay.WechatPayClient.Sign(ctx, data)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, err)
	}

	return &types.WechatPay{
		AppID:     appID,
		Timestamp: timestamp,
		NonceStr:  nonce,
		Package:   pack,
		SignType:  "RSA",
		PaySign:   sign.Signature,
	}, nil
}

func (s *OrderService) ApplyRefund(ctx *gin.Context, req types.ApplyRefundReq) (*types.ApplyRefundResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if user.UserRole != model.UserRoleAdmin && user.UserRole != model.UserRoleRoot {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not admin"))
	}

	tx := db.Get().Begin()
	xErr := s.applyRefund(tx, req, user)
	if xErr != nil {
		tx.Rollback()
		return nil, xErr
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil, nil
}

func (s *OrderService) applyRefund(tx *gorm.DB, req types.ApplyRefundReq, operator *model.User) xerr.XErr {
	ok := redis.Lock(fmt.Sprintf("lock-order:order_id:%d", req.OrderID), time.Minute*5)
	if !ok {
		return xerr.WithCode(xerr.ErrorRedisLock, errors.New("redis lock"))
	}
	defer redis.Unlock(fmt.Sprintf("lock-order:order_id:%d", req.OrderID))

	// 获取order信息
	orderInfo, err := s.orderDao.GetOrderInfoByIDTX(tx, req.OrderID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if orderInfo == nil {
		return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("order %d not found", req.OrderID))
	}

	userID := ""
	user, err := s.userDao.GetByUserIDInTx(tx, userID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	if req.RefundPoint {
		xErr := s.refundPoint(tx, user, operator, nil)
		if xErr != nil {
			return xErr
		}
	}

	return nil
}

func (s *OrderService) ApplyExchange(ctx *gin.Context, req types.ApplyExchangeReq) (*types.ApplyExchangeResp, xerr.XErr) {
	return nil, nil
}

func (s *OrderService) refundPoint(tx *gorm.DB, user, operator *model.User, order *model.OrderInfo) xerr.XErr {
	point := order.PointPrice
	// 加公司积分
	org, err := s.orgDao.GetByIDForUpdate(tx, user.OrganizationID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	err = s.orgDao.UpdateByIDInTx(tx, user.OrganizationID, &model.Organization{Point: org.Point.Add(point)})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 加员工积分
	user, err = s.userDao.GetByUserIDForUpdate(tx, user.UserID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	err = s.userDao.UpdateByUserIDInTx(tx, user.UserID, &model.User{Point: user.Point.Add(point)})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	records, err := s.pointRecordDao.ListByOrderIDInTx(tx, user.UserID, int(order.ID))
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	for _, record := range records {
		remain, err := s.pointRemainDao.GetByIDForUpdate(tx, record.PointID)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}

		err = s.pointRemainDao.UpdateByIDInTx(tx, record.PointID, &model.PointRemain{PointRemain: remain.PointRemain.Sub(record.ChangePoint)})
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}

		newRecord := &model.PointRecord{
			UserID:             user.UserID,
			OrganizationID:     0,
			ChangePoint:        decimal.Decimal{},
			PointApplicationID: record.PointApplicationID,
			PointID:            record.PointID,
			OrderID:            record.OrderID,
			Type:               model.PointRecordTypeRefund,
			Status:             model.PointRecordStatusConfirmed,
			Comment:            consts.PointCommentRefund,
			OperateUserID:      operator.UserID,
			OperateUsername:    operator.Username,
		}
		err = s.pointRecordDao.CreateInTx(tx, newRecord)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
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
	rowsAffected, er := s.orderDao.UpdateOrderInfoByIDCTX(ctx, req.QueryOrderID, &model.OrderInfo{Status: enum.OderInfoReceived})
	if er != nil {
		return xerr.WithCode(xerr.ErrorDatabase, er)
	}
	if rowsAffected == 0 {
		return xerr.WithCode(xerr.ErrorOperationForbidden, fmt.Errorf("order status is %d,not allow confirm receipt", orderInfo.Status))
	}
	return nil
}
