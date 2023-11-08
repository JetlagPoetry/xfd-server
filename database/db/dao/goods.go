package dao

import (
	"context"
	"xfd-backend/database/db"
	"xfd-backend/database/db/enum"
	"xfd-backend/database/db/model"
	"xfd-backend/database/repo"
	"xfd-backend/pkg/types"
)

type GoodsDao struct {
	repo repo.IRepo
}

func NewGoodsDao() *GoodsDao {
	return &GoodsDao{repo: repo.NewRepo(db.Get())}
}

// CreateGoods 创建商品SPU
func (d *GoodsDao) CreateGoods(ctx context.Context, goods *model.Goods) (id int32, err error) {
	err = d.repo.GetDB(ctx).Model(&model.Goods{}).Create(goods).Error
	if err != nil {
		return 0, err
	}
	return goods.ID, nil
}

// CreateSpecification 创建商品规格
func (d *GoodsDao) CreateSpecification(ctx context.Context, specification *model.Specification) (id int32, err error) {
	err = d.repo.GetDB(ctx).Model(&model.Specification{}).Create(specification).Error
	if err != nil {
		return 0, err
	}

	return specification.ID, nil
}

// CreateSpecificationValue 创建商品规格属性
func (d *GoodsDao) CreateSpecificationValue(ctx context.Context, specificationValue *model.SpecificationValue) (id int32, err error) {
	err = d.repo.GetDB(ctx).Model(&model.SpecificationValue{}).Create(specificationValue).Error
	if err != nil {
		return 0, err
	}
	return specificationValue.ID, nil
}

// CreateProductVariant 创建SKU
func (d *GoodsDao) CreateProductVariant(ctx context.Context, productVariant *model.ProductVariant) (id int32, err error) {
	err = d.repo.GetDB(ctx).Model(&model.ProductVariant{}).Create(productVariant).Error
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
		Where("status = 1").
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
		Where("g.deleted = 0 and g.status = 1")
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

func (d *GoodsDao) UpdateGoodsByID(ctx context.Context, id int32, updateValue *model.Goods) (err error) {
	updateResult := db.Get().Debug().Model(&model.Goods{}).Where("id =?", id).Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}
func (d *GoodsDao) DeleteGoodsByID(ctx context.Context, id int32) (err error) {
	updateResult := db.Get().Debug().Where("id =?", id).Delete(&model.Goods{})
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

func (d *GoodsDao) GetMyGoodsList(ctx context.Context, req types.MyGoodsListReq) (goodsList []*types.GoodsList, count int64, err error) {
	result := db.Get().Debug().Model(&model.Goods{}).Where("user_id = ?", req.UserID)

	result = result.Where(&model.Goods{
		UserID: req.UserID}).
		Select("id, name, goods_front_image, images, status, sold_num,spu_code,created_at, updated_at")
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

//GetSpecificationValueBySpecID

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

func (d *GoodsDao) ModifyGoodsByID(ctx context.Context, id int32, updateValue *model.Goods) (err error) {

	updateResult := db.Get().Debug().
		Model(&model.Goods{}).
		Where("id =?", id).
		Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

// updateSpecificationByID
func (d *GoodsDao) UpdateSpecificationByID(ctx context.Context, id int32, updateValue *model.Specification) (err error) {
	updateResult := db.Get().Debug().
		Model(&model.Specification{}).
		Where("id =?", id).
		Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

// updateSpecificationValueByID
func (d *GoodsDao) UpdateSpecificationValueByID(ctx context.Context, id int32, updateValue *model.SpecificationValue) (err error) {
	updateResult := db.Get().Debug().
		Model(&model.SpecificationValue{}).
		Where("id =?", id).
		Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}

// updateProductVariantByID
func (d *GoodsDao) UpdateProductVariantByID(ctx context.Context, id int32, updateValue *model.ProductVariant) (err error) {
	updateResult := db.Get().Debug().
		Model(&model.ProductVariant{}).
		Where("id =?", id).
		Updates(updateValue)
	if err = updateResult.Error; err != nil {
		return err
	}
	return nil
}
