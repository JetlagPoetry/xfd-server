package types

import (
	"mime/multipart"
	"xfd-backend/database/db/model"
)

type OrgApplyPointReq struct {
	File       multipart.File        `json:"file"`
	FileHeader *multipart.FileHeader `json:"fileHeader"`
	StartTime  int64                 `json:"startTime"`
	EndTime    int64                 `json:"endTime"`
	Comment    string                `json:"comment"`
}

type OrgApplyPointResp struct {
}

type OrgVerifyPointReq struct {
	ID      int                          `json:"id"`
	Status  model.PointApplicationStatus `json:"status"`
	Comment string                       `json:"comment"`
}

type OrgVerifyPointResp struct {
}

type OrgGetApplyToVerifyReq struct {
}

type OrgGetApplyToVerifyResp struct {
	ID                int    `json:"id"`
	OrganizationName  string `json:"organizationName"`
	OrganizationCode  string `json:"organizationCode"`
	Comment           string `json:"comment"`
	UserID            string `json:"userID"`
	Username          string `json:"username"`
	UserCertificateNo string `json:"userCertificateNo"`
	UserPosition      string `json:"userPosition"`
	UserPhone         string `json:"userPhone"`
	SubmitTime        int64  `json:"submitTime"`
	ApplyURL          string `json:"applyURL"`
}

type OrgGetApplysReq struct {
	PageRequest
	OrgID int `form:"orgID"`
}

type OrgGetApplysResp struct {
	List       []*PointOrder `json:"list"`
	NeedVerify int           `json:"needVerify"`
	TotalNum   int           `json:"totalNum"`
}

type OrgClearPointReq struct {
	OrgID int `json:"orgID"`
}

type OrgClearPointResp struct {
}

type PointOrder struct {
	OrganizationName string                       `json:"organizationName"`
	OrganizationCode string                       `json:"organizationCode"`
	Comment          string                       `json:"comment"`
	SubmitTime       int64                        `json:"submitTime"`
	VerifyTime       int64                        `json:"verifyTime"`
	VerifyComment    string                       `json:"verifyComment"`
	VerifyUserID     string                       `json:"verifyUserID"`
	VerifyUsername   string                       `json:"verifyUsername"`
	PointOrderStatus model.PointApplicationStatus `json:"pointOrderStatus"`
	ApplyURL         string                       `json:"applyURL"`
}

type VerifyAccountReq struct {
	ID      int                    `json:"id"`
	Status  model.UserVerifyStatus `json:"status"`
	Comment string                 `json:"comment"`
}

type VerifyAccountResp struct {
}

type GetAccountToVerifyReq struct {
}

type GetAccountToVerifyResp struct {
	ID               int            `json:"id"`
	Role             model.UserRole `json:"role"`
	Organization     string         `json:"organization"`
	OrganizationCode string         `json:"organizationCode"`
	OrganizationURL  string         `json:"organizationUrl"`
	IdentityURLA     string         `json:"identityUrlA"`
	IdentityURLB     string         `json:"identityUrlB"`
	RealName         string         `json:"realName"`
	CertificateNo    string         `json:"certificateNo"`
	Position         string         `json:"position"`
	Phone            string         `json:"phone"`
}

type GetAccountVerifyListReq struct {
	PageRequest
}

type GetAccountVerifyListResp struct {
	ToVerify int64                  `json:"toVerify"`
	List     []*AccountVerifyRecord `json:"list"`
	TotalNum int64                  `json:"totalNum"`
}

type AccountVerifyRecord struct {
	ID               int                    `json:"id"`
	Role             model.UserRole         `json:"role"`
	Organization     string                 `json:"organization"`
	OrganizationCode string                 `json:"organizationCode"`
	OrganizationURL  string                 `json:"organizationUrl"`
	IdentityURLA     string                 `json:"identityUrlA"`
	IdentityURLB     string                 `json:"identityUrlB"`
	RealName         string                 `json:"realName"`
	CertificateNo    string                 `json:"certificateNo"`
	Position         string                 `json:"position"`
	Phone            string                 `json:"phone"`
	Status           model.UserVerifyStatus `json:"status"`
	Comment          string                 `json:"comment"`
	VerifyTime       int64                  `json:"verifyTime"`
	CreateTime       int64                  `json:"createTime"`
}

type GetOrganizationsReq struct {
	PageRequest
	Name string `form:"name"`
}

type GetOrganizationsResp struct {
	List     []*Organization `json:"list"`
	TotalNum int             `json:"totalNum"`
}

type Organization struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	TotalMember int    `json:"totalMember"`
	PointMember int    `json:"pointMember"`
	TotalPoint  int    `json:"totalPoint"`
}

type GetOrgMembersReq struct {
	PageRequest
	OrgID    int    `form:"orgID"`
	Username string `form:"username"`
	Phone    string `form:"phone"`
}

type GetOrgMembersResp struct {
	List     []*OrgMember `json:"list"`
	TotalNum int          `json:"totalNum"`
}

type OrgMember struct {
	UserID           string `json:"userID"`
	Name             string `json:"name"`
	Phone            string `json:"phone"`
	OrganizationName string `json:"organization_name"`
	Point            int    `json:"point"`
	CreateTime       int64  `json:"createTime"`
}

type GetPointRecordsByApplyReq struct {
	PageRequest
	ApplyID int `form:"applyID"`
}

type GetPointRecordsByApplyResp struct {
	List           []*PointRecord `json:"list"`
	TotalNum       int            `json:"totalNum"`
	PointTotal     int            `json:"pointTotal"`
	PointExpired   int            `json:"pointExpired"`
	PointSpend     int            `json:"pointSpend"`
	PointAvailable int            `json:"pointAvailable"`
}

type GetPointRecordsByUserReq struct {
	PageRequest
	UserID string `form:"userID"`
}

type GetPointRecordsByUserResp struct {
	List     []*PointRecord `json:"list"`
	TotalNum int            `json:"totalNum"`
}

type GetPointRecordsReq struct {
	PageRequest
	OrgID int `form:"orgID"`
}

type GetPointRecordsResp struct {
	List     []*PointRecord `json:"list"`
	TotalNum int            `json:"totalNum"`
}

type PointRecord struct {
	UserID          string                `json:"userID"`
	Username        string                `json:"username"`
	PointTotal      int                   `json:"pointTotal"`
	PointChange     int                   `json:"pointChange"`
	Type            model.PointRecordType `json:"type"`
	Comment         string                `json:"comment"`
	UpdateTime      int64                 `json:"updateTime"`
	OperateUserID   string                `json:"operateUserID"`
	OperateUsername string                `json:"operateUsername"`
}
