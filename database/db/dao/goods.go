package dao

import (
	"context"
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

// CreateGoods 创建商品SPU
func (d *GoodsDao) CreateGoods(ctx context.Context, goods *model.Goods) (id int32, err error) {
	err = db.Get().Model(&model.Goods{}).Create(goods).Error
	if err != nil {
		return 0, err
	}
	return goods.ID, nil
}

// CreateSpecification 创建商品规格
func (d *GoodsDao) CreateSpecification(ctx context.Context, specification *model.Specification) (id int32, err error) {
	err = db.Get().Model(&model.Specification{}).Create(specification).Error
	if err != nil {
		return 0, err
	}

	return specification.ID, nil
}

// CreateSpecificationValue 创建商品规格属性
func (d *GoodsDao) CreateSpecificationValue(ctx context.Context, specificationValue *model.SpecificationValue) (id int32, err error) {
	err = db.Get().Model(&model.SpecificationValue{}).Create(specificationValue).Error
	if err != nil {
		return 0, err
	}
	return specificationValue.ID, nil
}

// CreateProductVariant 创建SKU
func (d *GoodsDao) CreateProductVariant(ctx context.Context, productVariant *model.ProductVariant) (id int32, err error) {
	err = db.Get().Model(&model.ProductVariant{}).Create(productVariant).Error
	if err != nil {
		return 0, err
	}
	return productVariant.ID, nil
}

//分页查询商品信息

func (d *GoodsDao) GetGoodsList(ctx context.Context, req types.GoodsListReq) (goodsList []*types.GoodsList, count int64, err error) {
	result := db.Get().Debug().Model(&model.Goods{}).Where("deleted = 0 and status = 1")
	if req.IsRetail == 1 {
		result = result.Where("is_retail = 1 and retail_status = 1")
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

// GetProductVariantListByGoodsID 获取SKU信息
func (d *GoodsDao) GetProductVariantListByGoodsID(ctx context.Context, goodsID int32, productVariantType enum.ProductVariantType) (productVariants []*model.ProductVariant, err error) {
	err = db.Get().Debug().Model(&model.ProductVariant{}).
		Where("deleted = 0 and status = 1").
		Where(&model.ProductVariant{
			GoodsID: goodsID,
			Type:    productVariantType,
		}).Find(&productVariants).Error
	if err != nil {
		return nil, err
	}
	return productVariants, nil
}

func (d *GoodsDao) GetMinPriceList(ctx context.Context, req types.GoodsListReq) (minPriceResult []*types.MinPriceResult, count int64, err error) {
	dbQuery := db.Get().Debug().Table("product_variant as pv").
		Select("pv.goods_id, MIN(pv.price) AS min_price, g.id, g.name, g.goods_front_image, g.images").
		Joins("INNER JOIN goods g on pv.goods_id = g.id").
		Where("pv.deleted = 0 and g.deleted = 0 and g.status = 1")
	if req.IsRetail == 1 {
		dbQuery = dbQuery.Where("is_retail = 1 and retail_status = 1")
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

func (d *GoodsDao) GetGoodsListByIDs(ctx context.Context, goodsIDs []int32, req types.GoodsListReq) (goodsList []*types.GoodsList, count int64, err error) {
	result := db.Get().Debug().Model(&model.Goods{}).Where("deleted = 0 and status = 1")
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
