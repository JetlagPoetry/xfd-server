package types

import (
	"mime/multipart"
	"xfd-backend/database/db/model"
)

type OrgApplyPointReq struct {
	File       multipart.File        `json:"file"`
	FileHeader *multipart.FileHeader `json:"fileHeader"`
	Comment    string                `json:"comment"`
}

type OrgApplyPointResp struct {
}

type OrgVerifyPointReq struct {
	ID      string                    `json:"id"`
	Status  model.OrgPointApplication `json:"status"`
	Comment string                    `json:"comment"`
}

type OrgVerifyPointResp struct {
}

type OrgGetApplyToVerifyReq struct {
}

type OrgGetApplyToVerifyResp struct {
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
	OrgID int `json:"orgID"`
}

type OrgGetApplysResp struct {
	List       []*PointOrder `json:"list"`
	NeedVerify int           `json:"needVerify"`
	TotalNum   int           `json:"totalNum"`
}

type PointOrder struct {
	OrganizationName string                          `json:"organizationName"`
	OrganizationCode string                          `json:"organizationCode"`
	Comment          string                          `json:"comment"`
	SubmitTime       int64                           `json:"submitTime"`
	VerifyTime       int64                           `json:"verifyTime"`
	VerifyComment    string                          `json:"verifyComment"`
	VerifyUserID     string                          `json:"verifyUserID"`
	VerifyUsername   string                          `json:"verifyUsername"`
	PointOrderStatus model.OrgPointApplicationStatus `json:"pointOrderStatus"`
	ApplyURL         string                          `json:"applyURL"`
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
	Name string `json:"name"`
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
	Username string `json:"username"`
	Phone    string `json:"phone"`
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
	Point            string `json:"point"`
	CreateTime       int64  `json:"createTime"`
	// todo 积分失效时间？？
}

type GetPointRecordsByUserReq struct {
	UserID int `json:"userID"`
}

type GetPointRecordsByUserResp struct {
}
