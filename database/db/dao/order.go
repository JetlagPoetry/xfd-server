package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"xfd-backend/database/db"
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

// CreateOrderProductVariantDetails 创建订单产品详情
func (d *OrderDao) CreateOrderProductVariantDetails(ctx context.Context, orderProductVariantDetails []*model.OrderProductVariantDetail) ([]int32, error) {
	ids := make([]int32, len(orderProductVariantDetails))
	batchSize := 100 // 每批次插入的最大数据量

	for i := 0; i < len(orderProductVariantDetails); i += batchSize {
		end := i + batchSize
		if end > len(orderProductVariantDetails) {
			end = len(orderProductVariantDetails)
		}

		batch := orderProductVariantDetails[i:end]
		result := db.GetRepo().GetDB(ctx).Model(&model.OrderProductVariantDetail{}).CreateInBatches(batch, len(batch))
		if result.Error != nil {
			return nil, result.Error
		}

		for j, createdOrderProductVariantDetail := range batch {
			ids[i+j] = createdOrderProductVariantDetail.ID
		}
	}
	return ids, nil

}

func (d *OrderDao) GetMyShoppingCartList(c *gin.Context, req types.ShoppingCartListReq) (shoppingCartList []*model.ShoppingCart, count int64, err error) {
	result := db.Get().Debug().Model(&model.ShoppingCart{}).Where("user_id = ?", req.UserID)
	result = result.Count(&count)
	result = result.Order("goods_id,created_at desc, id desc").
		Offset(req.Offset()).
		Limit(req.Limit()).
		Find(&shoppingCartList)

	// 错误处理
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return shoppingCartList, count, nil
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
	err = db.GetRepo().GetDB(ctx).Model(&model.ShoppingCart{}).
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
	err = db.GetRepo().GetDB(ctx).Model(&model.ShoppingCart{}).
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

func (d *OrderDao) UpdateOrderInfoByID(ctx context.Context, id int, update *model.OrderInfo) error {
	updateResult := db.GetRepo().GetDB(ctx).Model(&model.OrderInfo{}).Where("id = ?", id).Updates(update)
	return updateResult.Error
}

func (d *OrderDao) UpdateOrderProductVariantDetailByOrderSnTX(tx *gorm.DB, orderSn string, m *model.OrderProductVariantDetail) error {
	updateResult := tx.Model(&model.OrderProductVariantDetail{}).Where("order_sn = ?", orderSn).Updates(m)
	return updateResult.Error
}

func (d *OrderDao) UpdateOrderByOrderSn(ctx context.Context, orderSn string, value *model.OrderInfo) error {
	updateResult := db.GetRepo().GetDB(ctx).Model(&model.OrderInfo{}).Where("order_sn = ?", orderSn).Updates(value)
	return updateResult.Error
}

func (d *OrderDao) UpdateOrderProductVariantDetailByOrderSn(ctx context.Context, orderSn string, value *model.OrderProductVariantDetail) error {
	updateResult := db.GetRepo().GetDB(ctx).Model(&model.OrderProductVariantDetail{}).Where("order_sn = ?", orderSn).Updates(value)
	return updateResult.Error
}
