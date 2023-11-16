package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/martian/log"
	"gorm.io/gorm"
	"xfd-backend/database/db/dao"
	"xfd-backend/database/db/enum"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/common"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/xerr"
)

type OrderService struct {
	goods    *dao.GoodsDao
	userDao  *dao.UserDao
	orderDao *dao.OrderDao
}

func NewOrderService() *OrderService {
	return &OrderService{
		goods:    dao.NewGoodsDao(),
		userDao:  dao.NewUserDao(),
		orderDao: dao.NewOrderDao(),
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

	return &types.CreateOrderResp{}, nil
}
