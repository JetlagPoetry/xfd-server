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
}

type VerifyAccountResp struct {
}

type GetAccountToVerifyReq struct {
}

type GetAccountToVerifyResp struct {
}

type GetAccountsReq struct {
}

type GetAccountsResp struct {
}
