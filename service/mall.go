package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"xfd-backend/database/db/dao"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/xerr"
	"xfd-backend/service/cache"
)

type MallService struct {
	category *dao.CategoryDao
}

func NewMallService() *MallService {
	return &MallService{
		category: dao.NewCategoryDao(),
	}
}

func (s *MallService) GetCategories(c context.Context, req types.CategoryListReq) ([]*model.Category, xerr.XErr) {
	//todo:增加缓存
	categoriesList, err := s.category.GetCategoriesList(c, req.Level, req.ParentID)
	if err != nil {
		return nil, nil
	}
	if len(categoriesList) == 0 {
		return nil, nil
	}

	return categoriesList, nil
}

func (s *MallService) SetCategoryCache(ctx context.Context) xerr.XErr {
	categoryList, err := s.category.ListAll(ctx)
	if err != nil {
		return nil
	}

	log.Println("[SetCategoryCache] category get, len=", len(categoryList))

	m := make(map[int32]*cache.Category)
	for _, c := range categoryList {
		m[c.ID] = &cache.Category{
			ID:               c.ID,
			Name:             c.Name,
			ParentCategoryID: c.ParentCategoryID,
			Level:            c.Level,
			Image:            c.Image,
		}
	}
	for _, c := range categoryList {
		if c.ParentCategoryID == 0 || m[c.ParentCategoryID] == nil {
			continue
		}
		if m[c.ParentCategoryID].SubCategoryIDs == nil {
			m[c.ParentCategoryID].SubCategoryIDs = make([]int32, 0)
		}
		m[c.ParentCategoryID].SubCategoryIDs = append(m[c.ParentCategoryID].SubCategoryIDs, c.ID)
	}

	cache.SetCategory(m)

	return nil
}

func (s *MallService) AddCategory(c *gin.Context, req types.CategoryAddReq) xerr.XErr {
	if req.CheckParams() != nil {
		return xerr.WithCode(xerr.InvalidParams, req.CheckParams())
	}
	//校验父分类是否存在
	if req.ParentID != 0 {
		category, err := s.category.GetCategoryByID(c, req.ParentID)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
		if category == nil {
			return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("parent category %d not found", req.ParentID))
		}
		if category.Level != req.Level {
			return xerr.WithCode(xerr.InvalidParams, fmt.Errorf("parent category %d level is %d not equal %d", req.ParentID, category.Level, req.Level))
		}
	}
	category := &model.Category{
		Name:             req.Name,
		Level:            req.Level,
		ParentCategoryID: req.ParentID,
		Image:            req.Image,
	}
	_, err := s.category.CreateCategory(c, category)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil
}
