package dao

import (
	"context"
	"xfd-backend/database/db"
	"xfd-backend/database/db/enum"
	"xfd-backend/database/db/model"
)

type CategoryDao struct {
}

func NewCategoryDao() *CategoryDao {
	return &CategoryDao{}
}

// GetCategoriesList 获取商品分类
func (d *CategoryDao) GetCategoriesList(ctx context.Context, level, parentCategoryID int32) ([]*model.Category, error) {
	var categories []*model.Category
	result := db.Get().Model(&model.Category{}).
		Where(&model.Category{
			Level:            enum.GoodsCategoryLevel(level),
			Status:           1,
			ParentCategoryID: parentCategoryID}).
		Preload("SubCategory").
		Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (d *CategoryDao) ListAll(ctx context.Context) (list []*model.Category, err error) {
	result := db.Get().Model(&model.Category{}).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
