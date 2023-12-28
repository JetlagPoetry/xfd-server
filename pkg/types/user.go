package types

import "xfd-backend/database/db/model"

type Jscode2SessionResponse struct {
	ErrorMsg   string `json:"errmsg"`
	ErrorCode  int32  `json:"errcode"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	OpenID     string `json:"openid"`
}

type UserSendCodeReq struct {
	Phone string `json:"phone"`
}

type UserSendCodeResp struct {
}

type UserLoginReq struct {
	Phone  string `json:"phone"`
	Source Source `json:"source"`
	Code   string `json:"code"`
}

type Source int32

const (
	SourceUnknown Source = 0
	SourceMiniApp Source = 1
	SourceCMS     Source = 2
)

type UserLoginResp struct {
	AccessToken string         `json:"accessToken"` // 访问令牌
	TokenType   string         `json:"tokenType"`   // 令牌类型
	ExpiresAt   int64          `json:"expiresAt"`   // 令牌到期时间戳
	UserRole    model.UserRole `json:"userRole"`    // 用户角色
	UserID      string         `json:"userID"`
	Username    string         `json:"username"`
	AvatarURL   string         `json:"avatarURL"`
	//VerifyStatus  model.UserVerifyStatus `json:"verifyStatus"`  // 认证状态
	//VerifyComment string                 `json:"verifyComment"` // 认证备注
	//NotifyVerify  bool                   `json:"notifyVerify"`  // 是否提示认证成功
}

type UserSubmitRoleReq struct {
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

type UserSubmitRoleResp struct {
}

type UserRefreshTokenResp struct {
	AccessToken string `json:"accessToken"` // 访问令牌
	TokenType   string `json:"tokenType"`   // 令牌类型
	ExpiresAt   int64  `json:"expiresAt"`   // 令牌到期时间戳
}

type GetVerifyInfoResp struct {
	Username      string                 `json:"username"`
	UserRole      model.UserRole         `json:"userRole"`
	VerifyStatus  model.UserVerifyStatus `json:"verifyStatus"`
	VerifyComment string                 `json:"verifyComment"`
	Organization  string                 `json:"organization"`
	NotifyVerify  bool                   `json:"notifyVerify"`  // 是否提示认证成功
	VerifyHistory bool                   `json:"verifyHistory"` // 是否有曾经通过的审核
}

type GetUserInfoResp struct {
	Username  string         `json:"username"`
	AvatarURL string         `json:"avatarUrl"`
	UserRole  model.UserRole `json:"userRole"`
	//VerifyStatus   model.UserVerifyStatus `json:"verifyStatus"`
	//VerifyComment  string                 `json:"verifyComment"`
	Organization   string `json:"organization"`
	OrganizationID int    `json:"organizationID"`
	Point          string `json:"point"`
	//NotifyVerify   bool                   `json:"notifyVerify"` // 是否提示认证成功
}

//type UserVerifyStatus int32
//
//const (
//	UserVerifyStatusUnfinished UserVerifyStatus = 0
//	UserVerifyStatusDone       UserVerifyStatus = 1
//)

type UserModifyInfoReq struct {
	Username  string `json:"username"`
	AvatarURL string `json:"avatarUrl"`
}

type UserModifyInfoResp struct {
}

type UserAssignAdminReq struct {
	Phone    string         `json:"phone"`
	Username string         `json:"username"`
	Role     model.UserRole `json:"role"`
}

type UserAssignAdminResp struct {
}

type UserGetAdminsReq struct {
	PageRequest
}

type UserGetAdminsResp struct {
	List     []*UserAdmin `json:"list"`
	TotalNum int          `json:"totalNum"`
	RootNum  int          `json:"rootNum"`
}

type UserAdmin struct {
	UserID   string         `json:"userID"`
	Username string         `json:"username"`
	Phone    string         `json:"phone"`
	Role     model.UserRole `json:"role"`
	RoleName string         `json:"roleName"`
	Comment  string         `json:"comment"`
}

type UserDeleteUserReq struct {
	UserID string `json:"userID"`
}

type UserDeleteUserResp struct {
}

type UserImSigReq struct {
}

type UserImSigResp struct {
	UserSig   string `json:"userSig"`
	ExpiredAt int64  `json:"expiredAt"`
}

type UserGetAddressListReq struct {
}

type UserGetAddressListResp struct {
	List []*UserAddress `json:"list"`
}

type UserAddress struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Phone     string ` json:"phone"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Region    string `json:"region"`
	Address   string `json:"address"`
	IsDefault bool   `json:"isDefault"`
}

type UserGetDefaultAddressReq struct {
}

type UserGetDefaultAddressResp struct {
	UserAddress
	NoDefault bool `json:"noDefault"`
}

type UserAddAddressReq struct {
	Name      string `json:"name"`
	Phone     string ` json:"phone"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Region    string `json:"region"`
	Address   string `json:"address"`
	IsDefault bool   `json:"isDefault"`
}

type UserAddAddressResp struct {
}

type UserModifyAddressReq struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Region    string `json:"region"`
	Address   string `json:"address"`
	IsDefault bool   `json:"isDefault"`
}

type UserModifyAddressResp struct {
}

type UserDeleteAddressReq struct {
	ID int `json:"id"`
}

type UserDeleteAddressResp struct {
}
