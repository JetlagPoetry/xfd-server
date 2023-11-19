package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
	"xfd-backend/database/db"
	"xfd-backend/database/db/dao"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/common"
	"xfd-backend/pkg/consts"
	"xfd-backend/pkg/types"
	"xfd-backend/pkg/utils"
	"xfd-backend/pkg/xerr"
)

type OrgService struct {
	userService            *UserService
	userDao                *dao.UserDao
	userVerifyDao          *dao.UserVerifyDao
	orgDao                 *dao.OrganizationDao
	PointApplicationDao    *dao.PointApplicationDao
	PointApplicationTmpDao *dao.PointApplicationTmpDao
	PointRemainDao         *dao.PointRemainDao
	PointRecordDao         *dao.PointRecordDao
}

func NewOrgService() *OrgService {
	return &OrgService{
		userService:            NewUserService(),
		userDao:                dao.NewUserDao(),
		userVerifyDao:          dao.NewUserVerifyDao(),
		orgDao:                 dao.NewOrganizationDao(),
		PointApplicationDao:    dao.NewPointApplicationDao(),
		PointApplicationTmpDao: dao.NewPointApplicationTmpDao(),
		PointRemainDao:         dao.NewPointRemainDao(),
		PointRecordDao:         dao.NewPointRecordDao(),
	}
}

type OrgMember struct {
	Phone string
	Name  string
	Point decimal.Decimal
}

func (s *OrgService) ApplyPoint(ctx context.Context, req types.OrgApplyPointReq) (*types.OrgApplyPointResp, xerr.XErr) {
	// 校验用户
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorUploadFile, err)
	}
	if user.UserRole != model.UserRoleBuyer {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not buyer"))
	}
	org, err := s.orgDao.GetByID(ctx, user.OrganizationID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorUploadFile, err)
	}

	fileBytes, err := ioutil.ReadAll(req.File)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return nil, xerr.WithCode(xerr.ErrorUploadFile, err)
	}
	fileReader := bytes.NewReader(fileBytes)
	// 解析csv文件，校验格式
	members, totalPoint, xErr := s.parseCSV(ctx, fileReader, req.FileHeader)
	if xErr != nil {
		return nil, xErr
	}

	// 校验csv文件内容
	_, xErr = s.verifyCsv(ctx, members)
	if xErr != nil {
		return nil, xErr
	}

	// 上传csv文件
	fileReaderCopy := bytes.NewReader(fileBytes)
	link, err := s.uploadCSV(ctx, fileReaderCopy, req.FileHeader)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorUploadFile, err)
	}

	// 新增积分单申请记录
	tx := db.Get().Begin()
	xErr = s.applyPoint(tx, org, userID, totalPoint, link, req)
	if xErr != nil {
		tx.Rollback()
		return nil, xerr.WithCode(xerr.ErrorDatabase, xErr)
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.OrgApplyPointResp{}, nil
}

// csv格式：手机号、姓名、积分数
func (s *OrgService) parseCSV(ctx context.Context, file io.Reader, header *multipart.FileHeader) ([]*OrgMember, *decimal.Decimal, xerr.XErr) {
	//if filepath.Ext(header.Filename) != "xls" && filepath.Ext(header.Filename) != "xlsx" {
	//	return nil, 0, xerr.WithCode(xerr.ErrorInvalidFileExt, errors.New("无效的文件扩展名"))
	//}

	xlFile, err := excelize.OpenReader(file)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return nil, nil, xerr.WithCode(xerr.ErrorInvalidCsvFormat, errors.New("文件格式错误"))
	}
	rows, err := xlFile.GetRows("Sheet1")
	if err != nil {
		return nil, nil, xerr.WithCode(xerr.ErrorInvalidCsvFormat, errors.New("文件格式错误"))
	}
	if len(rows) > 1001 {
		return nil, nil, xerr.WithCode(xerr.ErrorInvalidCsvFormat, errors.New("超过1000条"))
	}
	list := make([]*OrgMember, 0)
	totalPoint := decimal.Zero
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if len(row) < 3 {
			return nil, nil, xerr.WithCode(xerr.ErrorInvalidCsvFormat, errors.New("文件格式错误"))
		}
		point, err := strconv.Atoi(row[2])
		if err != nil {
			return nil, nil, xerr.WithCode(xerr.ErrorInvalidCsvFormat, errors.New("文件格式错误"))
		}
		phone := row[0]
		if !utils.Mobile(phone) {
			return nil, nil, xerr.WithCode(xerr.ErrorInvalidCsvFormat, errors.New("包含错误的手机号"))
		}
		list = append(list, &OrgMember{
			Phone: row[0],
			Name:  row[1],
			Point: decimal.NewFromInt(int64(point)),
		})
		totalPoint = totalPoint.Add(decimal.NewFromInt(int64(point)))
	}

	return list, &totalPoint, nil
}
func (s *OrgService) verifyCsv(ctx context.Context, members []*OrgMember) (map[string]*model.User, xerr.XErr) {
	phoneMap := make(map[string]bool)
	phoneList := make([]string, 0)
	for _, member := range members {
		phoneMap[member.Phone] = true
		phoneList = append(phoneList, member.Phone)
	}

	if len(phoneMap) < len(members) {
		return nil, xerr.WithCode(xerr.ErrorInvalidFileExt, errors.New("重复手机号"))
	}
	list, err := s.userDao.ListByPhoneList(ctx, phoneList)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	userMap := make(map[string]*model.User)
	for _, user := range list {
		if user.UserRole == model.UserRoleBuyer || user.UserRole == model.UserRoleSupplier {
			return nil, xerr.WithCode(xerr.ErrorInvalidCsvFormat, errors.New("已认证过身份"))
		}
		userMap[user.Phone] = user
	}

	return userMap, nil
}

func (s *OrgService) uploadCSV(ctx context.Context, file io.Reader, header *multipart.FileHeader) (string, error) {
	fileSize := header.Size
	maxSize := int64(50 * 1024 * 1024) // 50MB
	if fileSize > maxSize {
		return "", errors.New("file size exceed")
	}
	link, err := utils.Upload(ctx, "xfd-t", "point"+"/"+utils.GenerateFileName()+filepath.Ext(header.Filename), file)
	if err != nil {
		return "", err
	}
	return link, nil
}

func (s *OrgService) downloadXLS(ctx context.Context, url string) ([]*OrgMember, error) {
	// 发送HTTP请求，获取文件内容
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("请求文件失败:", err)
		return nil, err
	}
	defer response.Body.Close()

	xlFile, err := excelize.OpenReader(response.Body)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return nil, xerr.WithCode(xerr.ErrorInvalidCsvFormat, errors.New("文件格式错误"))
	}
	rows, err := xlFile.GetRows("Sheet1")
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorInvalidCsvFormat, errors.New("文件格式错误"))
	}
	if len(rows) > 1001 {
		return nil, xerr.WithCode(xerr.ErrorInvalidCsvFormat, errors.New("超过1000条"))
	}
	list := make([]*OrgMember, 0)
	for i, row := range rows {
		if i == 0 {
			continue
		}
		point, err := strconv.Atoi(row[2])
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorInvalidCsvFormat, errors.New("文件格式错误"))
		}
		phone := row[0]
		if !utils.Mobile(phone) {
			return nil, xerr.WithCode(xerr.ErrorInvalidCsvFormat, errors.New("包含错误的手机号"))
		}
		list = append(list, &OrgMember{
			Phone: row[0],
			Name:  row[1],
			Point: decimal.NewFromInt(int64(point)),
		})
	}

	return list, nil

}

func (s *OrgService) applyPoint(tx *gorm.DB, org *model.Organization, userID string, totalPoint *decimal.Decimal, fileURL string, req types.OrgApplyPointReq) xerr.XErr {
	app := &model.PointApplication{
		OrganizationID: int(org.ID),
		TotalPoint:     *totalPoint,
		FileURL:        fileURL,
		Status:         model.PointApplicationStatusUnknown,
		ApplyUserID:    userID,
		Comment:        req.Comment,
		StartTime:      req.StartTime,
		EndTime:        req.EndTime,
		VerifyTime:     consts.TimeZeroValue,
	}
	err := s.PointApplicationDao.CreateInTx(tx, app)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil
}

func (s *OrgService) VerifyPoint(ctx context.Context, req types.OrgVerifyPointReq) (*types.OrgVerifyPointResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorUploadFile, err)
	}
	if user.UserRole != model.UserRoleRoot && user.UserRole != model.UserRoleAdmin {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not admin"))
	}

	tx := db.Get().Begin()
	xErr := s.verifyPoint(tx, req, user)
	if xErr != nil {
		tx.Rollback()
		return nil, xerr.WithCode(xerr.ErrorDatabase, xErr)
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return nil, nil
}

func (s *OrgService) verifyPoint(tx *gorm.DB, req types.OrgVerifyPointReq, user *model.User) xerr.XErr {
	app, err := s.PointApplicationDao.GetByIDForUpdate(tx, req.ID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorUploadFile, err)
	}

	if app.Status != model.PointApplicationStatusUnknown {
		return xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("application has been approved"))
	}

	// 修改审核记录表
	update := &model.PointApplication{
		Status:         req.VerifyStatus,
		VerifyTime:     time.Now(),
		VerifyComment:  req.VerifyComment,
		VerifyUserID:   user.UserID,
		VerifyUsername: user.Username,
	}
	err = s.PointApplicationDao.UpdateByIDInTx(tx, req.ID, update)
	if err != nil {
		return xerr.WithCode(xerr.ErrorUploadFile, err)
	}
	return nil
}

func (s *OrgService) ProcessPointVerify(ctx context.Context) xerr.XErr {
	list, err := s.PointApplicationDao.ListByStatus(ctx, model.PointApplicationStatusApproved)
	if err != nil {
		return xerr.WithCode(xerr.ErrorUploadFile, err)
	}

	for _, apply := range list {
		members, err := s.downloadXLS(ctx, apply.FileURL)
		if err != nil {
			continue
		}
		userMap, err := s.verifyCsv(ctx, members)
		if err != nil {
			continue
		}

		tx := db.Get().Begin()
		xErr := s.processPointVerify(tx, apply, members, userMap)
		if xErr != nil {
			tx.Rollback()
			return xerr.WithCode(xerr.ErrorDatabase, xErr)
		}

		// 提交事务
		if err = tx.Commit().Error; err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}
	return nil
}

func (s *OrgService) processPointVerify(tx *gorm.DB, apply *model.PointApplication, members []*OrgMember, userMap map[string]*model.User) xerr.XErr {
	org, err := s.orgDao.GetByIDForUpdate(tx, apply.OrganizationID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	updateOrg := &model.Organization{
		Point: org.Point.Add(apply.TotalPoint),
	}
	err = s.orgDao.UpdateByIDInTx(tx, apply.OrganizationID, updateOrg)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 插入发积分列表
	list := make([]*model.PointApplicationTmp, 0)
	for _, mem := range members {
		list = append(list, &model.PointApplicationTmp{
			OrganizationID: apply.OrganizationID,
			ApplicationID:  int(apply.ID),
			Username:       mem.Name,
			Phone:          mem.Phone,
			Point:          mem.Point,
		})
	}
	err = s.PointApplicationTmpDao.BatchCreateInTx(tx, list)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	update := &model.PointApplication{
		Status: model.PointApplicationStatusFinish,
	}
	err = s.PointApplicationDao.UpdateByIDInTx(tx, int(apply.ID), update)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil
}

func (s *OrgService) ProcessPointDistribute(ctx context.Context) xerr.XErr {
	// 处理发积分流水
	list, err := s.PointApplicationTmpDao.ListByStatus(ctx, model.PointApplicationTmpInit)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	orgMap := make(map[int]*model.Organization)
	applyMap := make(map[int]*model.PointApplication)
	for _, mem := range list {
		org, ok := orgMap[mem.OrganizationID]
		if !ok || org == nil {
			org, err = s.orgDao.GetByID(ctx, mem.OrganizationID)
			if err != nil {
				return xerr.WithCode(xerr.ErrorDatabase, err)
			}
			orgMap[mem.OrganizationID] = org
		}

		apply, ok := applyMap[mem.ApplicationID]
		if !ok || apply == nil {
			apply, err = s.PointApplicationDao.GetByID(ctx, mem.ApplicationID)
			if err != nil {
				return xerr.WithCode(xerr.ErrorDatabase, err)
			}
			applyMap[mem.ApplicationID] = apply
		}

		tx := db.Get().Begin()
		xErr := s.processPointDistribute(tx, apply, org, mem)
		if xErr != nil {
			tx.Rollback()
			return xerr.WithCode(xerr.ErrorDatabase, xErr)
		}

		// 提交事务
		if err = tx.Commit().Error; err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}

	return nil
}

// 加锁顺序：org->point_application->user->point_remain->point_record
func (s *OrgService) processPointDistribute(tx *gorm.DB, apply *model.PointApplication, org *model.Organization, member *model.PointApplicationTmp) xerr.XErr {
	err := s.PointApplicationTmpDao.UpdateByAppIDInTx(tx, int(apply.ID), &model.PointApplicationTmp{Status: model.PointApplicationTmpFinish})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 检查用户是否已经注册
	user, err := s.userDao.GetByPhoneForUpdate(tx, member.Phone)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	// 未注册的用户注册一下
	if user == nil {
		user = &model.User{
			UserID:           utils.GenUUID(),
			Phone:            member.Phone,
			UserRole:         model.UserRoleCustomer,
			Username:         member.Username,
			OrganizationID:   int(org.ID),
			OrganizationName: org.Name,
			Point:            decimal.Zero,
		}
		if err = s.userDao.CreateInTx(tx, user); err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}

	// 已注册但公司对不上的，触发员工离职，并重新绑定
	if user.OrganizationID != 0 && user.OrganizationID != int(org.ID) && user.Point.GreaterThan(decimal.Zero) {
		err = s.processUserQuit(tx, user.UserID, user.OrganizationID)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}

	update := &model.User{
		Username:         member.Username,
		OrganizationID:   int(org.ID),
		OrganizationName: org.Name,
		Point:            user.Point.Add(member.Point),
		PointStatus:      model.UserPointStatusOwn,
	}
	if err = s.userDao.UpdateByUserIDInTx(tx, user.UserID, update); err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 下发积分
	remain := &model.PointRemain{
		UserID:             user.UserID,
		OrganizationID:     user.OrganizationID,
		PointApplicationID: int(apply.ID),
		Point:              member.Point,
		PointRemain:        member.Point,
	}
	err = s.PointRemainDao.CreateInTx(tx, remain)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 设置积分流水
	record := &model.PointRecord{
		UserID:             user.UserID,
		OrganizationID:     user.OrganizationID,
		ChangePoint:        member.Point,
		PointApplicationID: int(apply.ID),
		PointID:            int(remain.ID),
		Type:               model.PointRecordTypeApplication,
		Status:             model.PointRecordStatusConfirmed,
		Comment:            consts.PointCommentApplication,
		OperateUserID:      apply.VerifyUserID,
		OperateUsername:    apply.VerifyUsername,
	}
	err = s.PointRecordDao.CreateInTx(tx, record)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return nil
}

func (s *OrgService) ProcessPointExpired(ctx context.Context) xerr.XErr {
	// 定时任务处理过期积分
	expires, err := s.PointApplicationDao.ListExpired(ctx)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	for _, expire := range expires {
		tx := db.Get().Begin()
		xErr := s.processPointExpired(tx, expire)
		if xErr != nil {
			tx.Rollback()
			continue
		}

		// 提交事务
		if err = tx.Commit().Error; err != nil {
			continue
		}
	}

	return nil
}

// 加锁顺序：org->point_application->user->point_remain->point_record
func (s *OrgService) processPointExpired(tx *gorm.DB, apply *model.PointApplication) xerr.XErr {
	// 定时任务处理过期积分
	org, err := s.orgDao.GetByIDForUpdate(tx, apply.OrganizationID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 修改apply状态
	err = s.PointApplicationDao.UpdateByIDInTx(tx, int(apply.ID), &model.PointApplication{Status: model.PointApplicationStatusExpired})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	remainList, err := s.PointRemainDao.ListByAppIDInTx(tx, int(apply.ID))
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	totalChange := decimal.Zero
	recordList := make([]*model.PointRecord, 0)
	for _, remain := range remainList {
		user, err := s.userDao.GetByUserIDForUpdate(tx, remain.UserID)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}

		// 修改员工个人积分
		err = s.userDao.UpdateByUserIDInTx(tx, remain.UserID, &model.User{Point: user.Point.Sub(remain.PointRemain)})
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
		totalChange = totalChange.Add(remain.PointRemain)

		recordList = append(recordList, &model.PointRecord{
			UserID:             remain.UserID,
			OrganizationID:     int(org.ID),
			ChangePoint:        remain.PointRemain.Mul(decimal.NewFromInt(-1)),
			PointApplicationID: remain.PointApplicationID,
			PointID:            int(remain.ID),
			Type:               model.PointRecordTypeExpired,
			Status:             model.PointRecordStatusConfirmed,
			Comment:            consts.PointCommentExpire,
		})

	}

	// 修改该公司积分
	err = s.orgDao.UpdateByIDInTx(tx, apply.OrganizationID, &model.Organization{Point: org.Point.Sub(totalChange)})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 清空remain
	err = s.PointRemainDao.UpdateByAppIDInTx(tx, int(apply.ID), &model.PointRemain{PointRemain: decimal.Zero})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 添加积分流水
	err = s.PointRecordDao.BatchCreateInTx(tx, recordList)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil
}

func (s *OrgService) GetApplyToVerify(ctx context.Context, req types.OrgGetApplyToVerifyReq) (*types.OrgGetApplyToVerifyResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorUploadFile, err)
	}
	if user.UserRole != model.UserRoleRoot && user.UserRole != model.UserRoleAdmin {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not admin"))
	}

	// 查找下一个未修改状态的审核单并返回
	apply, err := s.PointApplicationDao.GetByStatus(ctx, model.PointApplicationStatusUnknown)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	count, err := s.PointApplicationDao.CountByStatus(ctx, model.PointApplicationStatusUnknown)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	organization, err := s.orgDao.GetByID(ctx, apply.OrganizationID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	applyUser, err := s.userDao.GetByUserID(ctx, apply.ApplyUserID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	verify, err := s.userVerifyDao.GetByUserID(ctx, apply.ApplyUserID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.OrgGetApplyToVerifyResp{
		ID:                int(apply.ID),
		OrganizationName:  organization.Name,
		OrganizationCode:  organization.Code,
		Comment:           apply.Comment,
		UserID:            applyUser.UserID,
		Username:          applyUser.Username,
		UserCertificateNo: verify.CertificateNo,
		UserPosition:      verify.Position,
		UserPhone:         applyUser.Phone,
		SubmitTime:        apply.CreatedAt.Unix(),
		ApplyURL:          apply.FileURL,
		HasNext:           count > 1,
	}, nil
}

func (s *OrgService) GetApplys(ctx context.Context, req types.OrgGetApplysReq) (*types.OrgGetApplysResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorUploadFile, err)
	}
	if user.UserRole != model.UserRoleRoot && user.UserRole != model.UserRoleAdmin {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not admin"))
	}

	// 获取待审核数量
	needVerify, err := s.PointApplicationDao.CountByStatus(ctx, model.PointApplicationStatusUnknown)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 获取积分申请列表
	applyList, count, err := s.PointApplicationDao.Lists(ctx, req.PageRequest, req.OrgID)
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

func (s *OrgService) ClearPoint(ctx context.Context, req types.OrgClearPointReq) (*types.OrgClearPointResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorUploadFile, err)
	}
	if user.UserRole != model.UserRoleRoot && user.UserRole != model.UserRoleAdmin {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("user is not admin"))
	}

	tx := db.Get().Begin()
	xErr := s.clearPoint(tx, req, user)
	if xErr != nil {
		tx.Rollback()
		return nil, xerr.WithCode(xerr.ErrorDatabase, xErr)
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.OrgClearPointResp{}, nil
}

func (s *OrgService) clearPoint(tx *gorm.DB, req types.OrgClearPointReq, operator *model.User) xerr.XErr {
	org, err := s.orgDao.GetByIDForUpdate(tx, req.OrgID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	appList, err := s.PointApplicationDao.ListByStatusInTx(tx, model.PointApplicationStatusFinish)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 修改该公司所有apply状态
	err = s.PointApplicationDao.UpdateByOrgIDInTx(tx, req.OrgID, &model.PointApplication{Status: model.PointApplicationStatusClear})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 修改该公司积分
	err = s.orgDao.UpdateByIDInTx(tx, req.OrgID, &model.Organization{Point: decimal.Zero})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 修改员工个人积分
	err = s.userDao.UpdateByOrgIDInTx(tx, req.OrgID, &model.User{Point: decimal.Zero})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	appIDs := make([]int, 0)
	for _, app := range appList {
		appIDs = append(appIDs, int(app.ID))
	}
	remainList, err := s.PointRemainDao.ListByAppIDs(tx, appIDs)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 修改员工积分剩余表
	err = s.PointRemainDao.UpdateByOrgIDInTx(tx, req.OrgID, &model.PointRemain{PointRemain: decimal.Zero})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	// 添加积分流水
	recordList := make([]*model.PointRecord, 0)
	for _, remain := range remainList {
		recordList = append(recordList, &model.PointRecord{
			UserID:             remain.UserID,
			OrganizationID:     int(org.ID),
			ChangePoint:        remain.PointRemain.Mul(decimal.NewFromInt(-1)),
			PointApplicationID: remain.PointApplicationID,
			PointID:            int(remain.ID),
			Type:               model.PointRecordTypeCancel,
			Status:             model.PointRecordStatusConfirmed,
			Comment:            consts.PointCommentClear,
			OperateUserID:      operator.UserID,
			OperateUsername:    operator.Username,
		})
	}
	err = s.PointRecordDao.BatchCreateInTx(tx, recordList)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil
}

func (s *OrgService) VerifyAccount(ctx context.Context, req types.VerifyAccountReq) (*types.VerifyAccountResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	userVerify, err := s.userVerifyDao.GetByID(ctx, req.ID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	if userVerify.Status != model.UserVerifyStatusSubmitted {
		return nil, xerr.WithCode(xerr.ErrorOperationForbidden, errors.New("verify finished"))
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
	xErr := s.verifyAccount(tx, req, userVerify, newUserVerify)
	if xErr != nil {
		tx.Rollback()
		return nil, xerr.WithCode(xerr.ErrorDatabase, xErr)
	}
	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return nil, nil
}

func (s *OrgService) verifyAccount(tx *gorm.DB, req types.VerifyAccountReq, userVerify, updateValue *model.UserVerify) xerr.XErr {
	user, err := s.userDao.GetByUserIDInTx(tx, userVerify.UserID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	err = s.userVerifyDao.UpdateByIDInTx(tx, req.ID, updateValue)
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
			Name: userVerify.Organization,
			Code: userVerify.OrganizationCode,
		}
		err = s.orgDao.CreateInTx(tx, org)
		if err != nil {
			return xerr.WithCode(xerr.ErrorDatabase, err)
		}
	}

	// 如果已经有公司，此时认证了新的公司
	if user.OrganizationID > 0 && user.OrganizationID != int(org.ID) {
		// 触发员工在旧公司离职
		err = s.processUserQuit(tx, user.UserID, user.OrganizationID)
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

// org->point_application->user->point_remain->point_record
func (s *OrgService) processUserQuit(tx *gorm.DB, userID string, orgID int) xerr.XErr {
	user, err := s.userDao.GetByUserIDInTx(tx, userID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	org, err := s.orgDao.GetByIDForUpdate(tx, orgID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	err = s.orgDao.UpdateByIDInTx(tx, orgID, &model.Organization{Point: org.Point.Sub(user.Point)})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	user, err = s.userDao.GetByUserIDForUpdate(tx, userID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	err = s.userDao.UpdateByUserIDInTx(tx, userID, &model.User{Point: decimal.Zero})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	remainList, err := s.PointRemainDao.ListByUserID(tx, userID)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	err = s.PointRemainDao.UpdateByUserIDInTx(tx, userID, &model.PointRemain{PointRemain: decimal.Zero})
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}

	recordList := make([]*model.PointRecord, 0)
	for _, remain := range remainList {
		recordList = append(recordList, &model.PointRecord{
			UserID:             userID,
			OrganizationID:     int(org.ID),
			ChangePoint:        remain.PointRemain.Mul(decimal.NewFromInt(-1)),
			PointApplicationID: remain.PointApplicationID,
			PointID:            int(remain.ID),
			Type:               model.PointRecordTypeQuit,
			Status:             model.PointRecordStatusConfirmed,
			Comment:            consts.PointCommentQuit,
		})
	}
	err = s.PointRecordDao.BatchCreateInTx(tx, recordList)
	if err != nil {
		return xerr.WithCode(xerr.ErrorDatabase, err)
	}
	return nil
}

func (s *OrgService) GetAccountToVerify(ctx context.Context, req types.GetAccountToVerifyReq) (*types.GetAccountToVerifyResp, xerr.XErr) {
	// 获取下一个未审核的用户认证申请
	userVerify, err := s.userVerifyDao.GetByStatus(ctx, model.UserVerifyStatusSubmitted)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	if userVerify == nil {
		return nil, xerr.WithCode(xerr.ErrorVerifyEmpty, err)
	}

	count, err := s.userVerifyDao.CountByStatus(ctx, model.UserVerifyStatusSubmitted)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
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
		HasNext:          count > 1,
	}, nil
}

func (s *OrgService) GetAccountVerifyList(ctx context.Context, req types.GetAccountVerifyListReq) (*types.GetAccountVerifyListResp, xerr.XErr) {
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

func (s *OrgService) GetOrganizations(ctx context.Context, req types.GetOrganizationsReq) (*types.GetOrganizationsResp, xerr.XErr) {
	orgList, count, err := s.orgDao.Lists(ctx, req.PageRequest, req.Name)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	list := make([]*types.Organization, 0)
	for _, org := range orgList {
		totalMember, err := s.userDao.CountByOrganization(ctx, int(org.ID))
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}

		pointMember, err := s.userDao.CountByOrganizationAndStatus(ctx, int(org.ID), model.UserPointStatusOwn)
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}

		list = append(list, &types.Organization{
			ID:          int(org.ID),
			Name:        org.Name,
			Code:        org.Code,
			TotalMember: int(totalMember),
			PointMember: int(pointMember),
			TotalPoint:  org.Point.Round(2).String(),
		})
	}

	return &types.GetOrganizationsResp{
		List:     list,
		TotalNum: int(count),
	}, nil
}
func (s *OrgService) GetOrgMembers(ctx context.Context, req types.GetOrgMembersReq) (*types.GetOrgMembersResp, xerr.XErr) {
	userID := common.GetUserID(ctx)

	user, err := s.userDao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	if req.OrgID == 0 {
		req.OrgID = user.OrganizationID
	}
	userList, count, err := s.userDao.ListByOrgID(ctx, req.PageRequest, req.OrgID, req.Username, req.Phone)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	list := make([]*types.OrgMember, 0)
	for _, user := range userList {
		list = append(list, &types.OrgMember{
			UserID:           user.UserID,
			Name:             user.Username,
			Phone:            user.Phone,
			OrganizationName: user.OrganizationName,
			Point:            user.Point.Round(2).String(),
			CreateTime:       user.CreatedAt.Unix(),
		})
	}

	return &types.GetOrgMembersResp{
		List:     list,
		TotalNum: int(count),
	}, nil
}

func (s *OrgService) GetPointRecordsByApply(ctx context.Context, req types.GetPointRecordsByApplyReq) (*types.GetPointRecordsByApplyResp, xerr.XErr) {
	apply, err := s.PointApplicationDao.GetByID(ctx, req.ApplyID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	recordList, count, err := s.PointRecordDao.ListByApplyID(ctx, req.PageRequest, req.ApplyID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	list := make([]*types.PointRecord, 0)
	for _, record := range recordList {
		user, err := s.userDao.GetByUserID(ctx, record.UserID)
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		list = append(list, &types.PointRecord{
			UserID:          record.UserID,
			Username:        user.Username,
			PointTotal:      user.Point.Round(2).String(),
			PointChange:     record.ChangePoint.Round(2).String(),
			Type:            record.Type,
			Comment:         record.Comment,
			UpdateTime:      record.CreatedAt.Unix(),
			OperateUserID:   record.OperateUserID,
			OperateUsername: record.OperateUsername,
		})
	}

	expired, err := s.PointRecordDao.SumByAppIDInTx(ctx, req.ApplyID, model.PointRecordTypeSpend)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	spend, err := s.PointRecordDao.SumByAppIDInTx(ctx, req.ApplyID, model.PointRecordTypeExpired)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	available, err := s.PointRemainDao.SumPointRemainByApplyID(ctx, req.ApplyID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}

	return &types.GetPointRecordsByApplyResp{
		List:           list,
		TotalNum:       int(count),
		PointTotal:     apply.TotalPoint.Round(2).String(),
		PointExpired:   expired.Round(2).String(),
		PointSpend:     spend.Round(2).String(),
		PointAvailable: available.Round(2).String(),
	}, nil
}

func (s *OrgService) GetPointRecordsByUser(ctx context.Context, req types.GetPointRecordsByUserReq) (*types.GetPointRecordsByUserResp, xerr.XErr) {
	user, err := s.userDao.GetByUserID(ctx, req.UserID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	recordList, count, err := s.PointRecordDao.ListByUserID(ctx, req.PageRequest, req.UserID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	list := make([]*types.PointRecord, 0)
	for _, record := range recordList {

		list = append(list, &types.PointRecord{
			UserID:          record.UserID,
			Username:        user.Username,
			PointTotal:      user.Point.Round(2).String(),
			PointChange:     record.ChangePoint.Round(2).String(),
			Type:            record.Type,
			Comment:         record.Comment,
			UpdateTime:      record.CreatedAt.Unix(),
			OperateUserID:   record.OperateUserID,
			OperateUsername: record.OperateUsername,
		})
	}

	return &types.GetPointRecordsByUserResp{
		List:     list,
		TotalNum: int(count),
	}, nil
}

func (s *OrgService) GetPointRecords(ctx context.Context, req types.GetPointRecordsReq) (*types.GetPointRecordsResp, xerr.XErr) {
	userID := common.GetUserID(ctx)
	// 如果orgID为空，则查看本公司积分明细
	if req.OrgID == 0 {
		user, err := s.userDao.GetByUserID(ctx, userID)
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		req.OrgID = user.OrganizationID
	}
	recordList, count, err := s.PointRecordDao.ListByOrgID(ctx, req.PageRequest, req.OrgID)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrorDatabase, err)
	}
	list := make([]*types.PointRecord, 0)
	for _, record := range recordList {
		user, err := s.userDao.GetByUserID(ctx, record.UserID)
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrorDatabase, err)
		}
		list = append(list, &types.PointRecord{
			UserID:          record.UserID,
			Username:        user.Username,
			PointTotal:      user.Point.Round(2).String(),
			PointChange:     record.ChangePoint.Round(2).String(),
			Type:            record.Type,
			Comment:         record.Comment,
			UpdateTime:      record.CreatedAt.Unix(),
			OperateUserID:   record.OperateUserID,
			OperateUsername: record.OperateUsername,
		})
	}

	return &types.GetPointRecordsResp{
		List:     list,
		TotalNum: int(count),
	}, nil
}
