package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
	"xfd-backend/database/db"
	"xfd-backend/database/db/enum"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/types"
)

type GoodsDao struct {
}

func NewGoodsDao() *GoodsDao {
	return &GoodsDao{}
}

/*create*/

// CreateGoods 创建商品SPU
func (d *GoodsDao) CreateGoods(ctx context.Context, goods *model.Goods) (id int32, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.Goods{}).Create(goods).Error
	if err != nil {
		return 0, err
	}
	return goods.ID, nil
}

// CreateSpecification 创建商品规格
func (d *GoodsDao) CreateSpecification(ctx context.Context, specification *model.Specification) (id int32, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.Specification{}).Create(specification).Error
	if err != nil {
		return 0, err
	}

	return specification.ID, nil
}

// CreateSpecificationValue 创建商品规格属性
func (d *GoodsDao) CreateSpecificationValue(ctx context.Context, specificationValue *model.SpecificationValue) (id int32, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.SpecificationValue{}).Create(specificationValue).Error
	if err != nil {
		return 0, err
	}
	return specificationValue.ID, nil
}

// CreateProductVariant 创建SKU
func (d *GoodsDao) CreateProductVariant(ctx context.Context, productVariant *model.ProductVariant) (id int32, err error) {
	err = db.GetRepo().GetDB(ctx).Model(&model.ProductVariant{}).Create(productVariant).Error
	if err != nil {
		return 0, err
	}
	return productVariant.ID, nil
}

/*get*/

// GetGoodsList 分页查询商品信息
func (d *GoodsDao) GetGoodsList(ctx context.Context, req types.GoodsListReq) (goodsList []*types.GoodsList, count int64, err error) {
	result := db.Get().Debug().Model(&model.Goods{}).Where("status = 1 and retail_status = 1")
	if req.IsRetail == 1 {
		result = result.Where("is_retail = 1")
	}
	if req.ListType == enum.GoodsListTypeQuery {
		result = result.Where("name like ?", "%"+req.QueryText+"%")
	}
	result = result.Where(&model.Goods{
		CategoryAID: req.CategoryAID,
		CategoryBID: req.CategoryBID,
		CategoryCID: req.CategoryCID,
		UserID:      req.UserID}).
		Select("id, name, goods_front_image, images")
	result = result.Count(&count)
	if req.ListType == enum.GoodsListTypeSale {
		result = result.Order("sold_num desc")
	}
	result = result.Order("created_at desc, id desc").
		Offset(req.Offset()).
		Limit(req.Limit()).
		Find(&goodsList)

	// 错误处理
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return goodsList, count, nil
}

// GetMinPriceList 获取最低价商品列表
func (d *GoodsDao) GetMinPriceList(ctx context.Context, req types.GoodsListReq) (minPriceResult []*types.MinPriceResult, count int64, err error) {
	dbQuery := db.Get().Debug().Table("product_variant as pv").
		Select("pv.goods_id, MIN(pv.price) AS min_price, g.id, g.name, g.goods_front_image, g.images").
		Joins("INNER JOIN goods g on pv.goods_id = g.id").
		Where("g.status = 1 and g.retail_status = 1")
	if req.IsRetail == 1 {
		dbQuery = dbQuery.Where("is_retail = 1")
	}
	dbQuery = dbQuery.Group("pv.goods_id")
	dbQuery = dbQuery.Count(&count)
	dbQuery = dbQuery.Order("min_price asc,pv.goods_id desc").
		Offset(req.Offset()).
		Limit(req.Limit()).
		Scan(&minPriceResult)
	if dbQuery.Error != nil {
		return nil, 0, dbQuery.Error
	}
	return minPriceResult, count, nil
}

// GetGoodsListByIDs 通过商品ID列表获取商品信息
func (d *GoodsDao) GetGoodsListByIDs(ctx context.Context, goodsIDs []int32, req types.GoodsListReq) (goodsList []*types.GoodsList, count int64, err error) {
	result := db.Get().Debug().Model(&model.Goods{}).Where("status = 1")
	if req.IsRetail == 1 {
		result = result.Where("is_retail = 1 and retail_status = 1")
	}
	result = result.Where("id IN ?", goodsIDs).
		Select("id, name, goods_front_image, images")
	// 错误处理
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return goodsList, count, nil
}

// GetGoodsByGoodsID 通过商品ID获取商品信息
func (d *GoodsDao) GetGoodsByGoodsID(ctx context.Context, id int32) (goods *model.Goods, err error) {
	var goodsList []*model.Goods
	err = db.Get().Debug().Model(&model.Goods{}).
		Where("id = ?", id).
		Find(&goodsList).Error
	if err != nil {
		return nil, err
	}
	if len(goodsList) == 0 {
		return nil, nil
	}
	return goodsList[0], nil
}

// GetProductVariantListByGoodsID 通过商品ID获取产品信息
func (d *GoodsDao) GetProductVariantListByGoodsID(ctx context.Context, goodsID int32, productVariantType enum.ProductVariantType, status enum.ProductVariantStatus) (productVariants []*model.ProductVariant, err error) {
	err = db.Get().Debug().Model(&model.ProductVariant{}).
		Where(&model.ProductVariant{
			GoodsID: goodsID,
			Type:    productVariantType,
			Status:  status,
		}).Find(&productVariants).Error
	if err != nil {
		return nil, err
	}
	return productVariants, nil
}

// GetProductVariantByProductVariantID 通过产品ID获取产品信息
func (d *GoodsDao) GetProductVariantByProductVariantID(ctx context.Context, id int32) (productVariant *model.ProductVariant, err error) {
	var productVariants []*model.ProductVariant
	err = db.Get().Debug().Model(&model.ProductVariant{}).
		Where(&model.ProductVariant{
			ID: id,
		}).Find(&productVariants).Error
	if err != nil {
		return nil, err
	}
	if len(productVariants) == 0 {
		return nil, nil
	}
	return productVariants[0], nil
}

// GetMyGoodsList 获取我的商品列表
func (d *GoodsDao) GetMyGoodsList(ctx context.Context, req types.MyGoodsListReq) (goodsList []*types.GoodsList, count int64, err error) {
	result := db.Get().Debug().Model(&model.Goods{}).Where("user_id = ?", req.UserID)

	result = result.Where(&model.Goods{
		UserID: req.UserID}).
		Select("id, name, goods_front_image, images, status, retail_status,sold_num,spu_code,created_at, updated_at")
	switch req.QueryGoodsListStatus {
	case enum.QueryGoodsListStatusOnSale:
		result = result.Where("status = 1")
	case enum.QueryGoodsListStatusOffSale:
		result = result.Where("status = 2")
	case enum.QueryGoodsListStatusSoldOut:
		result = result.Where("retail_status = 2")
	}
	result = result.Count(&count)

	result = result.Order("created_at desc, id desc").
		Offset(req.Offset()).
		Limit(req.Limit()).
		Find(&goodsList)

	// 错误处理
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return goodsList, count, nil
}

// GetSpecificationByGoodsID 通过商品ID获取规格信息
func (d *GoodsDao) GetSpecificationByGoodsID(ctx context.Context, goodsID int32, productVariantType enum.ProductVariantType) (specifications []*model.Specification, err error) {
	err = db.Get().Debug().Model(&model.Specification{}).
		Where(&model.Specification{
			GoodsID: goodsID,
			Type:    productVariantType,
		}).
		Find(&specifications).Error
	if err != nil {
		return nil, err
	}
	return specifications, nil
}

// GetSpecificationByID 通过规格ID获取规格信息
func (d *GoodsDao) GetSpecificationByID(ctx context.Context, id int32) (specification *model.Specification, err error) {
	var specifications []*model.Specification
	err = db.Get().Debug().Model(&model.Specification{}).
		Where(&model.Specification{
			ID: id,
		}).
		Find(&specifications).Error
	if err != nil {
		return nil, err
	}
	return specifications[0], nil
}

// GetSpecificationValueBySpecID 通过规格ID获取规格属性信息
func (d *GoodsDao) GetSpecificationValueBySpecID(ctx context.Context, specID int32) (specificationValues []*model.SpecificationValue, err error) {
	err = db.Get().Model(&model.SpecificationValue{}).
		Where(&model.SpecificationValue{
			SpecificationID: specID,
		}).
		Find(&specificationValues).Error
	if err != nil {
		return nil, err
	}
	return specificationValues, nil
}

// GetSpecificationValueByGoodID 通过商品ID获取规格属性信息
func (d *GoodsDao) GetSpecificationValueByGoodID(ctx context.Context, goodsID int32, productVariantType enum.ProductVariantType) (specificationValues []*model.SpecificationValue, err error) {
	err = db.Get().Model(&model.SpecificationValue{}).
		Where(&model.SpecificationValue{
			GoodsID: goodsID,
			Type:    productVariantType,
		}).
		Find(&specificationValues).Error
	if err != nil {
		return nil, err
	}
	return specificationValues, nil
}

/*update*/

// UpdateGoodsByID 通过商品ID更新商品信息
func (d *GoodsDao) UpdateGoodsByID(ctx context.Context, id int32, updateValue *model.Goods) (int64, error) {
	updateResult := db.GetRepo().GetDB(ctx).Debug().
		Model(&model.Goods{}).
		Where("id =?", id).
		Updates(updateValue)
	return updateResult.RowsAffected, updateResult.Error
}

// UpdateSpecificationByID 通过规格ID更新规格信息
func (d *GoodsDao) UpdateSpecificationByID(ctx context.Context, id int32, updateValue *model.Specification) (int64, error) {
	updateResult := db.GetRepo().GetDB(ctx).Debug().
		Model(&model.Specification{}).
		Where("id =?", id).
		Updates(updateValue)
	return updateResult.RowsAffected, updateResult.Error
}

// UpdateSpecificationValueByID 通过规格属性ID更新规格属性信息
func (d *GoodsDao) UpdateSpecificationValueByID(ctx context.Context, id int32, updateValue *model.SpecificationValue) (int64, error) {
	updateResult := db.GetRepo().GetDB(ctx).Debug().
		Model(&model.SpecificationValue{}).
		Where("id =?", id).
		Updates(updateValue)
	return updateResult.RowsAffected, updateResult.Error
}

// UpdateProductVariantByID 通过产品ID更新产品信息
func (d *GoodsDao) UpdateProductVariantByID(ctx context.Context, id int32, updateValue *model.ProductVariant) (int64, error) {
	updateResult := db.GetRepo().GetDB(ctx).
		Debug().Model(&model.ProductVariant{}).
		Where("id =?", id).
		Updates(updateValue)
	return updateResult.RowsAffected, updateResult.Error
}

/*delete*/

// DeleteGoodsByID 通过商品ID删除商品信息
func (d *GoodsDao) DeleteGoodsByID(ctx context.Context, id int32) (err error) {
	updateResult := db.Get().Debug().Where("id =?", id).Delete(&model.Goods{})
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

// DeleteProductVariantByGoodsID 通过商品ID删除产品信息
func (d *GoodsDao) DeleteProductVariantByGoodsID(ctx *gin.Context, goodsID int32) error {
	err := db.Get().Debug().Where("goods_id =?", goodsID).Delete(&model.ProductVariant{}).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteSpecificationByGoodsID 通过商品ID删除规格信息
func (d *GoodsDao) DeleteSpecificationByGoodsID(ctx context.Context, goodsID int32) error {
	err := db.Get().Debug().Where("goods_id =?", goodsID).Delete(&model.Specification{}).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteSpecificationValueByGoodsID 通过商品ID删除规格属性信息
func (d *GoodsDao) DeleteSpecificationValueByGoodsID(ctx context.Context, goodsID int32) error {
	err := db.Get().Debug().Where("goods_id =?", goodsID).Delete(&model.SpecificationValue{}).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteSpecificationByID 通过规格ID删除规格信息
func (d *GoodsDao) DeleteSpecificationByID(c *gin.Context, id int32) (err error) {
	err = db.GetRepo().GetDB(c).Debug().Where("id =?", id).Delete(&model.Specification{}).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteSpecificationValueBySpecID 通过规格ID删除规格属性信息
func (d *GoodsDao) DeleteSpecificationValueBySpecID(c *gin.Context, id int32) error {
	err := db.GetRepo().GetDB(c).Debug().Where("specification_id =?", id).Delete(&model.SpecificationValue{}).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteSpecificationValueByID 通过规格属性ID删除规格属性信息
func (d *GoodsDao) DeleteSpecificationValueByID(c *gin.Context, vid int32) error {
	err := db.GetRepo().GetDB(c).Debug().Where("id =?", vid).Delete(&model.SpecificationValue{}).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteSpecificationByIDs 通过规格ID列表删除规格信息
func (d *GoodsDao) DeleteSpecificationByIDs(ctx context.Context, ids []int32) (int64, error) {
	deleteResult := db.GetRepo().GetDB(ctx).Debug().Model(&model.Specification{}).
		Where("id IN ?", ids).
		Updates(&model.Specification{DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: true}})

	return deleteResult.RowsAffected, deleteResult.Error
}

// DeleteSpecificationValueByIDs 通过规格属性ID列表删除规格属性信息
func (d *GoodsDao) DeleteSpecificationValueByIDs(ctx context.Context, ids []int32) (int64, error) {
	deleteResult := db.GetRepo().GetDB(ctx).Debug().Model(&model.SpecificationValue{}).
		Where("id IN ?", ids).
		Updates(&model.SpecificationValue{DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: true}})
	return deleteResult.RowsAffected, deleteResult.Error
}

// DeleteProductVariantByIDs 通过产品ID列表删除产品信息
func (d *GoodsDao) DeleteProductVariantByIDs(ctx context.Context, ids []int32) (int64, error) {
	deleteResult := db.GetRepo().GetDB(ctx).Debug().Model(&model.ProductVariant{}).Where("id IN ?", ids).
		Updates(&model.ProductVariant{DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: true}})
	return deleteResult.RowsAffected, deleteResult.Error
}
