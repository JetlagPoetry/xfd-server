package service

import (
	"context"
	"xfd-backend/database/db/dao"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/xerr"
)

type OrgService struct {
	userDao *dao.UserDao
}

func NewOrgService() *OrgService {
	return &OrgService{
		userDao: dao.NewUserDao(),
	}
}

func (s *OrgService) ApplyPoint(ctx context.Context, req *types.OrgApplyPointReq) (*types.OrgApplyPointResp, xerr.XErr) {
	// 解析csv文件，校验格式

	// 上传csv文件

	// 新增积分单申请记录

	return nil, nil
}

func (s *OrgService) VerifyPoint(ctx context.Context, req *types.OrgVerifyPointReq) (*types.OrgVerifyPointResp, xerr.XErr) {
	// 修改审核记录表

	return nil, nil
}

func (s *OrgService) ProcessPointVerify(ctx context.Context) xerr.XErr {
	// 定时任务处理审核记录

	// 插入发积分流水

	return nil
}

func (s *OrgService) ProcessPointDistribute(ctx context.Context) xerr.XErr {
	// 处理发积分流水

	// for 检查用户是否已经注册

	// 下发积分、修改流水状态

	return nil
}

func (s *OrgService) GetApplyToVerify(ctx context.Context, req *types.OrgGetApplyToVerifyReq) (*types.OrgGetApplyToVerifyResp, xerr.XErr) {
	// 查找下一个未修改状态的审核单并返回

	return nil, nil
}

func (s *OrgService) GetApplys(ctx context.Context, req *types.OrgGetApplysReq) (*types.OrgGetApplysResp, xerr.XErr) {
	// 获取积分申请列表

	return nil, nil
}

func (s *OrgService) VerifyAccount(ctx context.Context, req *types.VerifyAccountReq) (*types.VerifyAccountResp, xerr.XErr) {
	// 审核用户，拒绝直接返回

	// 通过用户审核，查看公司是否存在，如果不存在则创建公司

	return nil, nil
}

func (s *OrgService) GetAccountToVerify(ctx context.Context, req *types.GetAccountToVerifyReq) (*types.GetAccountToVerifyResp, xerr.XErr) {
	// 获取下一个未审核的用户认证申请

	return nil, nil
}

func (s *OrgService) GetAccounts(ctx context.Context, req *types.GetAccountsReq) (*types.GetAccountsResp, xerr.XErr) {
	// 获取用户认证列表

	return nil, nil
}
