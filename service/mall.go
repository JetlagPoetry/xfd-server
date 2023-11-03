package service

import (
	"github.com/gin-gonic/gin"
	"xfd-backend/database/db/dao"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/xerr"
)

type MallService struct {
	category *dao.CategoryDao
}

func NewMallService() *MallService {
	return &MallService{
		category: dao.NewCategoryDao(),
	}
}

func (s *MallService) GetCategories(c *gin.Context, req types.CategoryListReq) ([]*model.Category, xerr.XErr) {
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
