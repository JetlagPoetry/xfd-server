package dao

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
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

func (d *CategoryDao) GetCategoryByID(ctx context.Context, id int32) (category *model.Category, err error) {
	var categories []*model.Category
	result := db.Get().Model(&model.Category{}).Where(&model.Category{ID: id}).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(categories) == 0 {
		return nil, nil
	}
	return categories[0], nil
}

/*create*/

func (d *CategoryDao) CreateCategory(ctx context.Context, category *model.Category) (id int32, err error) {
	result := db.Get().Model(&model.Category{}).Create(category)
	if result.Error != nil {
		return 0, result.Error
	}
	return category.ID, nil
}

/*update*/

func (d *CategoryDao) UpdateCategory(ctx context.Context, category *model.Category) error {
	result := db.Get().Model(&model.Category{}).Where(&model.Category{ID: category.ID}).Updates(category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/*delete*/

func (d *CategoryDao) DeleteCategory(ctx context.Context, id int32) error {
	// 查询要删除的父分类及其所有子分类
	var category model.Category
	result := db.Get().Preload("SubCategory").First(&category, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("category not found")
		}
		return result.Error
	}

	// 删除所有子分类
	for _, subCategory := range category.SubCategory {
		if err := d.DeleteCategory(ctx, subCategory.ID); err != nil {
			return err
		}
	}

	// 删除父分类
	result = db.Get().Delete(&category)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
