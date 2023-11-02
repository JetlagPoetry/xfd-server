package service

import (
	"xfd-backend/database/db/dao"
)

type CommonService struct {
	userDao *dao.UserDao
}

func NewCommonService() *CommonService {
	return &CommonService{
		userDao: dao.NewUserDao(),
	}
}
