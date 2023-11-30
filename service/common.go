package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	goredis "github.com/go-redis/redis"
	"time"
	"xfd-backend/database/db/dao"
	"xfd-backend/database/redis"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/xerr"
)

type CommonService struct {
	userDao   *dao.UserDao
	configDao *dao.ConfigDao
}

func NewCommonService() *CommonService {
	return &CommonService{
		userDao: dao.NewUserDao(),
	}
}

func (s *CommonService) GetConfig(ctx *gin.Context, req types.CommonGetConfigReq) (*types.CommonGetConfigResp, xerr.XErr) {
	// 查看redis缓存，如果存在，直接返回
	cache, err := redis.RedisClient.Get(fmt.Sprintf(redis.CONFIG_KEY, req.Key)).Result()
	if err != goredis.Nil && err != nil {
		return nil, xerr.WithCode(xerr.ErrorRedis, err)
	}
	if len(cache) > 0 {
		return &types.CommonGetConfigResp{Value: cache}, nil
	}

	// 如果不存在，从db中读取
	config, err := s.configDao.GetByKey(ctx, req.Key)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 更新缓存
	_ = redis.RedisClient.Set(fmt.Sprintf(redis.CONFIG_KEY, req.Key), config.Value, 5*time.Minute)
	return &types.CommonGetConfigResp{Value: config.Value}, nil
}
