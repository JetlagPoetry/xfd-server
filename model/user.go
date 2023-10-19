package model

type GetUserInfoReq struct {
}

type GetUserInfoResp struct {
	UserName       string
	Email          string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string
	Money          string
}
