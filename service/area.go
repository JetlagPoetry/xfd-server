package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"xfd-backend/database/db/dao"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/xerr"
)

type AreaService struct {
	area *dao.AreaDao
}

func NewAreaService() *AreaService {
	return &AreaService{
		area: dao.NewAreaDao(),
	}
}

func (s *AreaService) GetAreaInfo(ctx *gin.Context, req types.AreaReq) ([]*types.AreaList, xerr.XErr) {
	//todo:增加缓存
	info, err := s.area.GetAreaInfo(ctx, req.Code)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if info == nil {
		return nil, nil
	}
	result := make([]*types.AreaList, len(info))
	for i := range info {
		result[i] = &types.AreaList{
			Name:  info[i].Name,
			Code:  info[i].Code,
			Level: info[i].Level,
		}
		if info[i].Level == 2 {
			result[i].Children = s.getChildrenArea(ctx, info[i].Code)
		}
	}
	return result, nil
}

func (s *AreaService) getChildrenArea(ctx context.Context, pcode int) []*types.Area {
	info, err := s.area.GetAreaInfo(ctx, pcode)
	if err != nil {
		return nil
	}
	if info == nil {
		return nil
	}
	childrenInfo := make([]*types.Area, len(info))
	for i := range childrenInfo {
		childrenInfo[i] = &types.Area{
			Name:  info[i].Name,
			Code:  info[i].Code,
			Level: info[i].Level,
		}
	}
	return childrenInfo
}
