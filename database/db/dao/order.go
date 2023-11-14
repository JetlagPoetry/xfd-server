package dao

import (
	"context"
	"github.com/gin-gonic/gin"
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

/*get*/

// GetShoppingCartByUserIDAndProductVariantID 根据用户ID和产品ID获取购物车
func (d *OrderDao) GetShoppingCartByUserIDAndProductVariantID(ctx *gin.Context, userID string, productVariantID int32) (shoppingCart *model.ShoppingCart, err error) {
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
func (d *OrderDao) DeleteShoppingCartByID(ctx *gin.Context, id int32) error {
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
