package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
	"xfd-backend/database/db"
	"xfd-backend/database/db/enum"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/types"
)

type OrderDao struct {
}

func NewOrderDao() *OrderDao {
	return &OrderDao{}
}

/*create*/

// AddShoppingCart 添加购物车
func (d *OrderDao) AddShoppingCart(ctx *gin.Context, shoppingCart *model.ShoppingCart) (id int32, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.ShoppingCart{}).Create(shoppingCart).Error
	if err != nil {
		return 0, err
	}
	return shoppingCart.ID, nil
}

func (d *OrderDao) CreateOrderRefund(ctx context.Context, orderRefund *model.OrderRefund) (id int32, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.OrderRefund{}).Create(orderRefund).Error
	if err != nil {
		return 0, err
	}
	return orderRefund.ID, nil
}

func (d *OrderDao) GetMyShoppingCartList(c *gin.Context, req types.ShoppingCartListReq) (shoppingCartList []*model.ShoppingCart, count int64, err error) {
	result := db.Get().Debug().Model(&model.ShoppingCart{}).Where("user_id = ?", req.UserID)
	result = result.Count(&count)
	result = result.Order("updated_at desc,goods_id,created_at desc, id desc").
		Offset(req.Offset()).
		Limit(req.Limit()).
		Find(&shoppingCartList)

	// 错误处理
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return shoppingCartList, count, nil
}

func (d *OrderDao) CustomerGetQueryOrderList(ctx *gin.Context, req types.OrderListReq) (queryOrderList []*types.QueryOrder, count int64, err error) {
	queryOrder := db.Get().Debug().Model(&model.OrderInfo{}).
		Select("id,order_sn,status,name,quantity,unit_price,post_price,total_price,image,product_attr, shipment_company,shipment_sn,estimated_delivery_time,goods_supplier_user_id,created_at,payed_at,delivery_time,signer_name,singer_mobile,signer_address").
		Where("user_id = ? and status in (2,3,4,5,6,8)", req.UserID)
	if req.Status != 0 {
		queryOrder = queryOrder.Where("status = ?", req.Status)
	}
	// 获取订单列表总数
	queryOrder.Count(&count)

	// 添加排序、分页等操作
	queryOrder = queryOrder.Order("created_at desc,goods_id").
		Offset(req.Offset()).
		Limit(req.Limit())

	// 执行查询
	if err = queryOrder.Find(&queryOrderList).Error; err != nil {
		return nil, 0, err
	}
	return queryOrderList, count, nil
}

func (d *OrderDao) SupplierGetQueryOrderList(ctx *gin.Context, req types.OrderListReq) (queryOrderList []*types.QueryOrder, count int64, err error) {
	queryOrder := db.Get().Debug().Model(&model.OrderInfo{}).
		Select("id,order_sn,status,name,quantity,unit_price,post_price,total_price,image, product_attr, shipment_company,shipment_sn,estimated_delivery_time,created_at,payed_at,delivery_time,signer_name,singer_mobile,signer_address").
		Where("goods_supplier_user_id= ? and status in (3,4,5,6,8)", req.UserID)
	queryOrder = queryOrder.Where(&model.OrderInfo{
		Status:  req.Status,
		OrderSn: req.OrderSn,
	})
	// 获取订单列表总数
	queryOrder.Count(&count)

	// 添加排序、分页等操作
	queryOrder = queryOrder.Order("created_at desc,goods_id").
		Offset(req.Offset()).
		Limit(req.Limit())

	// 执行查询
	if err = queryOrder.Find(&queryOrderList).Error; err != nil {
		return nil, 0, err
	}
	return queryOrderList, count, nil
}

func (d *OrderDao) AdminGetQueryOrderList(ctx *gin.Context, req types.OrderListReq) (queryOrderList []*types.QueryOrder, count int64, err error) {
	queryOrder := db.Get().Debug().Model(&model.OrderInfo{}).
		Select("id,order_sn,status,name,goods_supplier_organization_name,total_price,user_phone,user_organization_name,payed_at").
		Where("status in (3,4,5,6,8)")
	queryOrder = queryOrder.Where(&model.OrderInfo{
		Status:    req.Status,
		OrderSn:   req.OrderSn,
		UserPhone: req.UserPhone,
	})
	if req.UserOrganizationName != "" {
		queryOrder = queryOrder.Where("user_organization_name like ?", "%"+req.UserOrganizationName+"%")
	}
	if req.SupplierOrganizationName != "" {
		queryOrder = queryOrder.Where("goods_supplier_organization_name like ?", "%"+req.SupplierOrganizationName+"%")
	}
	if req.GoodName != "" {
		queryOrder = queryOrder.Where("name like ?", "%"+req.GoodName+"%")
	}
	// 获取订单列表总数
	queryOrder.Count(&count)

	// 添加排序、分页等操作
	queryOrder = queryOrder.Order("payed_at desc,goods_id").
		Offset(req.Offset()).
		Limit(req.Limit())

	// 执行查询
	if err = queryOrder.Find(&queryOrderList).Error; err != nil {
		return nil, 0, err
	}
	return queryOrderList, count, nil
}

// CreateOrderInfo 创建订单信息
func (d *OrderDao) CreateOrderInfo(ctx context.Context, orderInfo *model.OrderInfo) (id int32, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.OrderInfo{}).Create(orderInfo).Error
	if err != nil {
		return 0, err
	}
	return orderInfo.ID, nil
}

/*get*/

// GetShoppingCartByUserIDAndProductVariantID 根据用户ID和产品ID获取购物车
func (d *OrderDao) GetShoppingCartByUserIDAndProductVariantID(ctx context.Context, userID string, productVariantID int32) (shoppingCart *model.ShoppingCart, err error) {
	var shoppingCartList []*model.ShoppingCart
	err = db.Get().Model(&model.ShoppingCart{}).
		Where("user_id = ? and product_variant_id = ?", userID, productVariantID).
		Find(&shoppingCartList).Error
	if err != nil {
		return nil, err
	}
	if len(shoppingCartList) == 0 {
		return nil, nil
	}
	return shoppingCartList[0], nil
}

// GetShoppingCartByUserIDAndShoppingCartIDForUpdate 根据用户ID和购物车ID获取购物车
func (d *OrderDao) GetShoppingCartByUserIDAndShoppingCartIDForUpdate(ctx context.Context, userID string, shoppingCartID int32) (shoppingCart *model.ShoppingCart, err error) {
	var shoppingCartList []*model.ShoppingCart
	err = db.GetRepo().GetDB(ctx).Model(&model.ShoppingCart{}).
		Where("user_id = ? and id = ?", userID, shoppingCartID).Clauses(clause.Locking{Strength: "UPDATE"}).
		Find(&shoppingCartList).Error
	if err != nil {
		return nil, err
	}
	if len(shoppingCartList) == 0 {
		return nil, nil
	}
	return shoppingCartList[0], nil
}

// GetShoppingCartByUserIDAndShoppingCartID 根据用户ID和购物车ID获取购物车
func (d *OrderDao) GetShoppingCartByUserIDAndShoppingCartID(ctx context.Context, userID string, shoppingCartID int32) (shoppingCart *model.ShoppingCart, err error) {
	var shoppingCartList []*model.ShoppingCart
	err = db.Get().Model(&model.ShoppingCart{}).
		Where("user_id = ? and id = ?", userID, shoppingCartID).Find(&shoppingCartList).Error
	if err != nil {
		return nil, err
	}
	if len(shoppingCartList) == 0 {
		return nil, nil
	}
	return shoppingCartList[0], nil
}

func (d *OrderDao) GetShoppingCartByID(ctx *gin.Context, id int32) (shoppingCart *model.ShoppingCart, err error) {
	var shoppingCarts []*model.ShoppingCart
	err = db.GetRepo().GetDB(ctx).Model(&model.ShoppingCart{}).
		Where("id = ?", id).Find(&shoppingCarts).Error
	if err != nil {
		return nil, nil
	}
	if len(shoppingCarts) == 0 {
		return nil, nil
	}
	return shoppingCarts[0], nil
}

// GetOrderInfoByID 根据订单ID获取订单信息
func (d *OrderDao) GetOrderInfoByID(ctx context.Context, id int32) (orderInfo *model.OrderInfo, err error) {
	var orderInfos []*model.OrderInfo
	err = db.GetRepo().GetDB(ctx).Model(&model.OrderInfo{}).
		Where("id = ?", id).Find(&orderInfos).Error
	if err != nil {
		return nil, nil
	}
	if len(orderInfos) == 0 {
		return nil, nil
	}
	return orderInfos[0], nil
}

// GetOrderInfoByOrderSn 根据订单编号获取订单信息
func (d *OrderDao) GetOrderInfoByOrderSn(ctx context.Context, orderSn string) (orderInfo *model.OrderInfo, err error) {
	var orderInfos []*model.OrderInfo
	err = db.GetRepo().GetDB(ctx).Model(&model.OrderInfo{}).
		Where("order_sn = ?", orderSn).Find(&orderInfos).Error
	if err != nil {
		return nil, nil
	}
	if len(orderInfos) == 0 {
		return nil, nil
	}
	return orderInfos[0], nil
}

// GetOrderInfoByIDTX 根据订单ID获取订单信息
func (d *OrderDao) GetOrderInfoByIDTX(tx *gorm.DB, id int) (orderInfo *model.OrderInfo, err error) {
	var orderInfos []*model.OrderInfo
	err = tx.Model(&model.OrderInfo{}).Where("id = ?", id).Find(&orderInfos).Error
	if err != nil {
		return nil, err
	}
	if len(orderInfos) == 0 {
		return nil, nil
	}
	return orderInfos[0], nil
}

// GetOrderRefundByOrderID 根据订单ID获取退款信息
func (d *OrderDao) GetOrderRefundByOrderID(ctx context.Context, orderID int32) (refund []*model.OrderRefund, err error) {
	var refundList []*model.OrderRefund
	err = db.GetRepo().GetDB(ctx).Model(&model.OrderRefund{}).
		Where("order_id = ?", orderID).Find(&refundList).Error
	if err != nil {
		return nil, err
	}
	if len(refundList) == 0 {
		return nil, nil
	}
	return refundList, nil
}

func (d *OrderDao) ListByStatus(ctx context.Context, status enum.OrderInfoStatus) (list []*model.OrderInfo, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.OrderInfo{}).Where("status = ?", status).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

/*update*/

// UpdateShoppingCartByID 通过购物车ID更新购物车信息
func (d *OrderDao) UpdateShoppingCartByID(ctx context.Context, id int32, updateValue *model.ShoppingCart) (int64, error) {
	updateResult := db.GetRepo().GetDB(ctx).Debug().
		Model(&model.ShoppingCart{}).
		Where("id =?", id).
		Updates(updateValue)
	return updateResult.RowsAffected, updateResult.Error
}

/*delete*/

// DeleteShoppingCartByID 根据购物车ID删除购物车
func (d *OrderDao) DeleteShoppingCartByID(ctx context.Context, id int32) error {
	err := db.Get().Model(&model.ShoppingCart{}).
		Where("id = ?", id).
		Delete(&model.ShoppingCart{}).Error
	return err
}

// DeleteShoppingCartByIDCTX 根据购物车ID删除购物车
func (d *OrderDao) DeleteShoppingCartByIDCTX(ctx context.Context, id int32) error {
	err := db.GetRepo().GetDB(ctx).Model(&model.ShoppingCart{}).
		Where("id = ?", id).
		Delete(&model.ShoppingCart{}).Error
	return err
}

// DeleteShoppingCartByIDsAndUserID 根据购物车IDs和用户ID删除购物车
func (d *OrderDao) DeleteShoppingCartByIDsAndUserID(ctx *gin.Context, ids []int32, userID string) error {
	err := db.Get().Model(&model.ShoppingCart{}).
		Where("id in (?) and user_id = ?", ids, userID).
		Delete(&model.ShoppingCart{}).Error
	return err
}

// DeleteShoppingCartByGoodsID 根据商品ID删除购物车
func (d *OrderDao) DeleteShoppingCartByGoodsID(ctx *gin.Context, goodsID int32) error {
	err := db.Get().Model(&model.ShoppingCart{}).
		Where("goods_id = ?", goodsID).
		Delete(&model.ShoppingCart{}).Error
	return err
}

func (d *OrderDao) DeleteShoppingCartByProductVariantID(ctx context.Context, id int32) error {
	err := db.GetRepo().GetDB(ctx).Model(&model.ShoppingCart{}).
		Where("product_variant_id = ?", id).Delete(&model.ShoppingCart{}).Error
	return err
}

func (d *OrderDao) UpdateOrderInfoByIDTX(tx *gorm.DB, id int, update *model.OrderInfo) error {
	updateResult := tx.Model(&model.OrderInfo{}).Where("id = ?", id).Updates(update)
	return updateResult.Error
}

func (d *OrderDao) UpdateOrderInfoByID(ctx context.Context, id int32, update *model.OrderInfo) error {
	updateResult := db.GetRepo().GetDB(ctx).Model(&model.OrderInfo{}).Where("id = ?", id).Updates(update)
	return updateResult.Error
}

func (d *OrderDao) UpdateOrderByOrderSn(ctx context.Context, orderSn string, value *model.OrderInfo) error {
	updateResult := db.GetRepo().GetDB(ctx).Model(&model.OrderInfo{}).Where("order_sn = ?", orderSn).Updates(value)
	return updateResult.Error
}

func (d *OrderDao) UpdateOrderInfoByIDCTX(ctx context.Context, id int32, updateValue *model.OrderInfo) (int64, error) {
	updateResult := db.GetRepo().GetDB(ctx).Model(&model.OrderInfo{}).Where("id = ?", id).Updates(updateValue)
	return updateResult.RowsAffected, updateResult.Error
}

//UpdateOrderRefundByID

func (d *OrderDao) UpdateOrderRefundByID(ctx context.Context, id int32, updateValue *model.OrderRefund) (int64, error) {
	updateResult := db.GetRepo().GetDB(ctx).Model(&model.OrderRefund{}).Where("id = ?", id).Updates(updateValue)
	return updateResult.RowsAffected, updateResult.Error
}

func (d *OrderDao) UpdateOrderRefundByIDTX(tx *gorm.DB, id int32, updateValue *model.OrderRefund) (int64, error) {
	updateResult := tx.Model(&model.OrderRefund{}).Where("id = ?", id).Updates(updateValue)
	return updateResult.RowsAffected, updateResult.Error
}

func (d *OrderDao) CountWaitingForDeliveryOrderBySupplyUserID(ctx context.Context, userID string) (count int64, err error) {
	query := db.GetRepo().GetDB(ctx).Model(&model.OrderInfo{}).
		Where("goods_supplier_user_id = ? and status = 3", userID).Count(&count)
	if query.Error != nil {
		return 0, query.Error
	}
	return count, nil
}

func (d *OrderDao) AdminGetQueryOrderListByIDs(ctx *gin.Context, req types.ExportOrderReq) (queryOrderList []*types.QueryOrder, err error) {
	queryOrder := db.Get().Debug().Model(&model.OrderInfo{}).
		Select("id,order_sn,status,name,goods_supplier_organization_name,total_price,user_phone,user_organization_name,payed_at").
		Where("status in (3,4,5,6,8) and id in ?", strings.Split(req.QueryOrderIDs, ","))

	// 执行查询
	if err = queryOrder.Find(&queryOrderList).Error; err != nil {
		return nil, err
	}
	return queryOrderList, nil
}
