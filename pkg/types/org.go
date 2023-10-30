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
}

type OrgGetApplysReq struct {
}

type OrgGetApplysResp struct {
	List       []*PointOrder `json:"list"`
	NeedVerify int           `json:"needVerify"`
}

type PointOrder struct {
}

type VerifyAccountReq struct {
	ID      int                    `json:"id"`
	Status  model.UserVerifyStatus `json:"status"`
	Comment string                 `json:"comment"`
}

type VerifyAccountResp struct {
}

type GetOrgMembersReq struct {
}

type GetOrgMembersResp struct {
}
type GetPointRecordsByUserReq struct {
}

type GetPointRecordsByUserResp struct {
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
	TotalNum int                    `json:"totalNum"`
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
