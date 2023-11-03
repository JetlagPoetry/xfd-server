package service

import (
	"context"
	"gorm.io/gorm"
	"time"
	"xfd-backend/database/db"
	"xfd-backend/database/db/dao"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/common"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/xerr"
)

type OrgService struct {
	userDao             *dao.UserDao
	userVerifyDao       *dao.UserVerifyDao
	orgDao              *dao.OrganizationDao
	orgPointApplication *dao.OrgPointApplicationDao
}

func NewOrgService() *OrgService {
	return &OrgService{
		userDao:             dao.NewUserDao(),
		userVerifyDao:       dao.NewUserVerifyDao(),
		orgDao:              dao.NewOrganizationDao(),
		orgPointApplication: dao.NewOrgPointApplicationDao(),
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

func (s *OrgService) ProcessPointExpired(ctx context.Context) xerr.XErr {
	// 定时任务处理过期积分

	return nil
}

func (s *OrgService) GetApplyToVerify(ctx context.Context, req *types.OrgGetApplyToVerifyReq) (*types.OrgGetApplyToVerifyResp, xerr.XErr) {
	// 查找下一个未修改状态的审核单并返回
	apply, err := s.orgPointApplication.GetByStatus(ctx, model.OrgPointApplicationStatusUnknown)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	organization, err := s.orgDao.GetByID(ctx, apply.OrganizationID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	user, err := s.userDao.GetByUserID(ctx, apply.VerifyUserID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	verify, err := s.userVerifyDao.GetByUserID(ctx, apply.VerifyUserID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.OrgGetApplyToVerifyResp{
		OrganizationName:  organization.Name,
		OrganizationCode:  organization.Code,
		Comment:           apply.Comment,
		UserID:            user.UserID,
		Username:          user.Username,
		UserCertificateNo: verify.CertificateNo,
		UserPosition:      verify.Position,
		UserPhone:         user.Phone,
		SubmitTime:        apply.CreatedAt.Unix(),
		ApplyURL:          apply.FileURL,
	}, nil
}

func (s *OrgService) GetApplys(ctx context.Context, req *types.OrgGetApplysReq) (*types.OrgGetApplysResp, xerr.XErr) {
	// 获取待审核数量
	needVerify, err := s.orgPointApplication.CountByStatus(ctx, model.OrgPointApplicationStatusUnknown)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 获取积分申请列表
	applyList, count, err := s.orgPointApplication.Lists(ctx, req.PageRequest, req.OrgID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	list := make([]*types.PointOrder, 0)
	for _, apply := range applyList {
		org, err := s.orgDao.GetByID(ctx, apply.OrganizationID)
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		list = append(list, &types.PointOrder{
			OrganizationName: org.Name,
			OrganizationCode: org.Code,
			Comment:          apply.Comment,
			SubmitTime:       apply.CreatedAt.Unix(),
			VerifyTime:       apply.VerifyTime.Unix(),
			VerifyComment:    apply.VerifyComment,
			VerifyUserID:     apply.VerifyUserID,
			VerifyUsername:   apply.VerifyUsername,
			PointOrderStatus: apply.Status,
			ApplyURL:         apply.FileURL,
		})
	}

	return &types.OrgGetApplysResp{
		List:       list,
		NeedVerify: int(needVerify),
		TotalNum:   int(count),
	}, nil
}
func (s *OrgService) GetOrganizations(ctx context.Context, req *types.GetOrganizationsReq) (*types.GetOrganizationsResp, xerr.XErr) {
	orgList, count, err := s.orgDao.Lists(ctx, req.PageRequest, req.Name)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	list := make([]*types.Organization, 0)
	for _, org := range orgList {
		count, err := s.userDao.CountByOrganization(ctx, int(org.ID))
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}

		list = append(list, &types.Organization{
			Name:        org.Name,
			Code:        org.Code,
			TotalMember: int(count),
			PointMember: 0, // todo  what's the point?
			TotalPoint:  0, // todo use redis
		})
	}

	return &types.GetOrganizationsResp{
		List:     list,
		TotalNum: int(count),
	}, nil
}
func (s *OrgService) GetOrgMembers(ctx context.Context, req *types.GetOrgMembersReq) (*types.GetOrgMembersResp, xerr.XErr) {
	
	return nil, nil
}
func (s *OrgService) GetPointRecordsByUser(ctx context.Context, req *types.GetPointRecordsByUserReq) (*types.GetPointRecordsByUserResp, xerr.XErr) {

	return nil, nil
}

func (s *OrgService) VerifyAccount(ctx context.Context, req *types.VerifyAccountReq) (*types.VerifyAccountResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	userVerify, err := s.userVerifyDao.GetByID(ctx, req.ID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	newUserVerify := &model.UserVerify{
		Status:         req.Status,
		Comment:        req.Comment,
		VerifyTime:     time.Now(),
		VerifyUsername: user.Username,
	}
	// 拒绝用户，直接修改状态
	if req.Status == model.UserVerifyStatusRejected {
		err = s.userVerifyDao.UpdateByID(ctx, req.ID, newUserVerify)
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}

	// 通过用户审核，查看公司是否存在，如果不存在则创建公司
	tx := db.Get().Begin()
	err = s.verifyAccount(tx, req, userVerify, newUserVerify)
	if err != nil {
		tx.Rollback()
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return nil, nil
}

func (s *OrgService) verifyAccount(tx *gorm.DB, req *types.VerifyAccountReq, userVerify, updateValue *model.UserVerify) error {
	err := s.userVerifyDao.UpdateByIDInTx(tx, req.ID, updateValue)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	org, err := s.orgDao.GetByCodeInTx(tx, userVerify.OrganizationCode)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 创建公司
	if org == nil {
		org = &model.Organization{
			Name:  userVerify.Organization,
			Code:  userVerify.OrganizationCode,
			Point: 0,
		}
		err = s.orgDao.CreateInTx(tx, org)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}

	// 绑定公司
	err = s.userDao.UpdateByUserIDInTx(tx, userVerify.UserID, &model.User{OrganizationID: int(org.ID), OrganizationName: org.Name})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return nil
}

func (s *OrgService) GetAccountToVerify(ctx context.Context, req *types.GetAccountToVerifyReq) (*types.GetAccountToVerifyResp, xerr.XErr) {
	// 获取下一个未审核的用户认证申请
	userVerify, err := s.userVerifyDao.GetByStatus(ctx, model.UserVerifyStatusSubmitted)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if userVerify == nil {
		return nil, xerr.WithCode(xerr.ErrorVerifyEmpty, err)
	}

	return &types.GetAccountToVerifyResp{
		ID:               int(userVerify.ID),
		Role:             userVerify.UserRole,
		Organization:     userVerify.Organization,
		OrganizationCode: userVerify.OrganizationCode,
		OrganizationURL:  userVerify.OrganizationURL,
		IdentityURLA:     userVerify.IdentityURLA,
		IdentityURLB:     userVerify.IdentityURLB,
		RealName:         userVerify.RealName,
		CertificateNo:    userVerify.CertificateNo,
		Position:         userVerify.Position,
		Phone:            userVerify.Phone,
	}, nil
}

func (s *OrgService) GetAccountVerifyList(ctx context.Context, req *types.GetAccountVerifyListReq) (*types.GetAccountVerifyListResp, xerr.XErr) {
	// 获取待审核总数
	toVerify, err := s.userVerifyDao.CountByStatus(ctx, model.UserVerifyStatusSubmitted)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 获取用户认证列表
	verifyList, count, err := s.userVerifyDao.List(ctx, req.PageRequest)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	list := make([]*types.AccountVerifyRecord, 0)
	for _, verify := range verifyList {
		list = append(list, &types.AccountVerifyRecord{
			ID:               int(verify.ID),
			Role:             verify.UserRole,
			Organization:     verify.Organization,
			OrganizationCode: verify.OrganizationCode,
			OrganizationURL:  verify.OrganizationURL,
			IdentityURLA:     verify.IdentityURLA,
			IdentityURLB:     verify.IdentityURLB,
			RealName:         verify.RealName,
			CertificateNo:    verify.CertificateNo,
			Position:         verify.Position,
			Phone:            verify.Phone,
			Status:           verify.Status,
			Comment:          verify.Comment,
			VerifyTime:       verify.VerifyTime.Unix(),
			CreateTime:       verify.CreatedAt.Unix(),
		})
	}

	return &types.GetAccountVerifyListResp{
		ToVerify: toVerify,
		List:     list,
		TotalNum: count,
	}, nil
}
