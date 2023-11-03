package dao

import (
	"context"
	"xfd-backend/database/db"
	"xfd-backend/database/db/model"
)

type GoodsDao struct {
}

func NewGoodsDao() *GoodsDao {
	return &GoodsDao{}
}

// CreateGoods 创建商品
func (d *GoodsDao) CreateGoods(ctx context.Context, goods *model.Goods) (id int32, err error) {
	err = db.Get().Model(&model.Goods{}).Create(goods).Error
	if err != nil {
		return 0, err
	}
	return goods.ID, nil
}

// CreateProductVariant 创建商品规格

func (d *GoodsDao) CreateSpecification(ctx context.Context, specification *model.Specification) (id int32, err error) {
	err = db.Get().Model(&model.Specification{}).Create(specification).Error
	if err != nil {
		return 0, err
	}

	return specification.ID, nil
}

// CreateProductVariant 创建商品规格值

func (d *GoodsDao) CreateSpecificationValue(ctx context.Context, specificationValue *model.SpecificationValue) (id int32, err error) {
	err = db.Get().Model(&model.SpecificationValue{}).Create(specificationValue).Error
	if err != nil {
		return 0, err
	}
	return specificationValue.ID, nil
}

func (d *GoodsDao) CreateProductVariant(ctx context.Context, productVariant *model.ProductVariant) (id int32, err error) {
	err = db.Get().Model(&model.ProductVariant{}).Create(productVariant).Error
	if err != nil {
		return 0, err
	}
	return productVariant.ID, nil
}
