package service

import (
	"context"
	"xfd-backend/database/db/dao"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/xerr"
)

type CommonService struct {
	userDao *dao.UserDao
}

func NewCommonService() *CommonService {
	return &CommonService{
		userDao: dao.NewUserDao(),
	}
}

func (s *CommonService) UploadToOSS(ctx context.Context, req *types.CommonUploadReq) (*types.CommonUploadResp, xerr.XErr) {
	return nil, nil
}
