package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/martian/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"xfd-backend/database/db"
	"xfd-backend/database/db/dao"
	"xfd-backend/database/db/enum"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/common"
	"xfd-backend/pkg/consts"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/xerr"
)

type OrderService struct {
	goods          *dao.GoodsDao
	userDao        *dao.UserDao
	orderDao       *dao.OrderDao
	orgDao         *dao.OrganizationDao
	pointRemainDao *dao.PointRemainDao
	pointRecordDao *dao.PointRecordDao
}

func NewOrderService() *OrderService {
	return &OrderService{
		goods:          dao.NewGoodsDao(),
		userDao:        dao.NewUserDao(),
		orderDao:       dao.NewOrderDao(),
		orgDao:         dao.NewOrganizationDao(),
		pointRemainDao: dao.NewPointRemainDao(),
		pointRecordDao: dao.NewPointRecordDao(),
	}
}

func (s *OrderService) AddShoppingCart(ctx *gin.Context, req types.ShoppingCartAddReq) xerr.XErr {
	user, xrr := s.CheckUserIsCustomer(ctx)
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
		UserID:           user.UserID,
		ProductVariantID: req.ProductVariantID,
		GoodsID:          productVariant.GoodsID,
		SKUCode:          productVariant.SKUCode,
		Name:             goods.Name,
		CoverURL:         goods.GoodsFrontImage,
		Quantity:         req.Quantity,
		ProductAttr:      productVariant.ProductAttr,
	}
	_, rr := s.orderDao.AddShoppingCart(ctx, newShoppingCart)
	if rr != nil {
		return xerr.WithCode(xerr.ErrorDatabase, rr)
	}
	return nil
}

func (s *OrderService) DeleteShoppingCart(ctx *gin.Context, req types.ShoppingCartDeleteReq) xerr.XErr {
	user, xrr := s.CheckUserIsCustomer(ctx)
	if xrr != nil {
		return xrr
	}
	err := s.orderDao.DeleteShoppingCartByIDsAndUserID(ctx, req.ShoppingCartIDs, user.UserID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil
}

func (s *OrderService) CheckUserIsCustomer(ctx *gin.Context) (*model.User, xerr.XErr) {
	//1.获取用户ID
	userID := common.GetUserID(ctx)
	//2.获取用户角色
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	userRole := user.UserRole
	if userRole != model.UserRoleCustomer {
		return nil, xerr.WithCode(xerr.ErrorAuthInsufficientAuthority, fmt.Errorf("user %s user role is %d, can not use shopping cart", userID, userRole))
	}
	return user, nil
}

func (s *OrderService) ModifyShoppingCart(ctx *gin.Context, req types.ShoppingCartModifyReq) xerr.XErr {
	//1.校验用户是否是消费者
	user, xrr := s.CheckUserIsCustomer(ctx)
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
	user, xrr := s.CheckUserIsCustomer(ctx)
	if xrr != nil {
		return nil, xrr
	}
	req.UserID = user.UserID
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
			err = s.orderDao.DeleteShoppingCartByID(ctx, shoppingCart.ID)
			if err != nil {
				log.Errorf(" delete invalid product shopping cart %d failed, err=%v", shoppingCart.ID, err)
				continue
			}
		}
		if productVariant != nil {
			shoppingCartDetail.Stock = productVariant.Stock
		}
		shoppingCartListResp = append(shoppingCartListResp, shoppingCartDetail)
	}
	result.List = shoppingCartListResp
	return &result, nil
}

func (s *OrderService) CreateOrder(ctx *gin.Context, req types.CreateOrderReq) (*types.CreateOrderResp, xerr.XErr) {
	// todo implement

	return &types.CreateOrderResp{}, nil
}

func (s *OrderService) createOrder(tx *gorm.DB, req types.CreateOrderReq) (*types.CreateOrderResp, xerr.XErr) {
	// 写订单

	// 扣购物车

	// 扣库存

	// 算价

	// 扣钱
	xErr := s.payForOrder(tx, nil, nil, decimal.Zero)
	if xErr != nil {
		return nil, xErr
	}

	return &types.CreateOrderResp{}, nil
}

func (s *OrderService) PayOrder(ctx *gin.Context, req types.PayOrderReq) (*types.PayOrderResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	tx := db.Get().Begin()
	xErr := s.payForOrder(tx, user, nil, decimal.NewFromFloat(9.95))
	if xErr != nil {
		tx.Rollback()
		return nil, xErr
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.PayOrderResp{}, nil
}

// org->point_application->user->point_remain->point_record
func (s *OrderService) payForOrder(tx *gorm.DB, user *model.User, order *model.OrderGoods, totalPrice decimal.Decimal) xerr.XErr {
	if user.OrganizationID == 0 {
		return xerr.WithCode(xerr.ErrorUserOrgNotFound, errors.New("user not belong to org"))
	}
	// todo  redis锁用户积分变动

	org, err := s.orgDao.GetByIDForUpdate(tx, user.OrganizationID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	user, err = s.userDao.GetByUserIDForUpdate(tx, user.UserID)
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
	if user.Point.GreaterThanOrEqual(totalPrice) {
		// 只用积分支付
		pointPrice = totalPrice
		wxPrice = decimal.Zero
		xErr := s.payWithPoint(tx, order, user, org, pointPrice, false)
		if xErr != nil {
			return xErr
		}
	} else {
		// 积分+微信支付
		pointPrice = user.Point
		wxPrice = totalPrice.Sub(user.Point)

		xErr := s.payWithPoint(tx, order, user, org, pointPrice, true)
		if xErr != nil {
			return xErr
		}
		xErr = s.payWithWx(tx, order, user, wxPrice)
		if xErr != nil {
			return xErr
		}
	}

	// todo 修改订单，支付金额、微信支付字段

	return nil
}

func (s *OrderService) payWithPoint(tx *gorm.DB, order *model.OrderGoods, user *model.User, org *model.Organization, point decimal.Decimal, payWx bool) xerr.XErr {
	userLeft := user.Point.Sub(point)
	err := s.userDao.UpdateByUserIDInTx(tx, user.UserID, &model.User{Point: userLeft})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	err = s.orgDao.UpdateByIDInTx(tx, user.OrganizationID, &model.Organization{Point: org.Point.Sub(point)})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	remainList, err := s.pointRemainDao.ListByUserID(tx, user.UserID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	pointLeft := point.Copy()
	remainIDs := make([]int, 0)
	for _, remain := range remainList {
		if pointLeft == decimal.Zero {
			break
		}
		remainIDs = append(remainIDs, int(remain.ID))
		var spend decimal.Decimal
		if remain.PointRemain.LessThan(pointLeft) {
			// 将pointRemain花费完
			spend = remain.PointRemain
			pointLeft.Sub(remain.PointRemain)
		} else {
			// 将pointLeft花费完
			spend = pointLeft
			pointLeft.Sub(pointLeft)
		}

		err = s.pointRemainDao.UpdateByIDInTx(tx, int(remain.ID), &model.PointRemain{PointRemain: spend})
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
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
		err = s.pointRecordDao.CreateInTx(tx, record)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}

	if pointLeft.GreaterThan(decimal.Zero) {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return nil
}

func (s *OrderService) payWithWx(tx *gorm.DB, order *model.OrderGoods, user *model.User, wxPrice decimal.Decimal) xerr.XErr {

	return nil
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
	// todo 获取order

	// todo 修改order状态

	// todo 获取用户

	xErr := s.refundPoint(tx, nil, operator, nil)
	if xErr != nil {
		return xErr
	}
	return nil
}

func (s *OrderService) ApplyExchange(ctx *gin.Context, req types.ApplyExchangeReq) (*types.ApplyExchangeResp, xerr.XErr) {
	return nil, nil
}

func (s *OrderService) refundPoint(tx *gorm.DB, user, operator *model.User, order *model.OrderGoods) xerr.XErr {
	point := order.PointPrice
	// todo redis分布式锁

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
