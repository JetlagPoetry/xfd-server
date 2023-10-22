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
	Phone string `json:"phone"`
	Code  string `json:"code"` // 验证码
}

type UserLoginResp struct {
	AccessToken   string                 `json:"accessToken"`   // 访问令牌
	TokenType     string                 `json:"tokenType"`     // 令牌类型
	ExpiresAt     int64                  `json:"expiresAt"`     // 令牌到期时间戳
	UserRole      model.UserRole         `json:"userRole"`      // 用户角色
	VerifyStatus  model.UserVerifyStatus `json:"verifyStatus"`  // 认证状态
	VerifyComment string                 `json:"verifyComment"` // 认证备注
	NotifyVerify  bool                   `json:"notifyVerify"`  // 是否提示认证成功
}

type UserSubmitRoleReq struct {
	Role             model.UserRole `json:"role"`
	Organization     string         `json:"organization"`
	OrganizationCode string         `json:"organization_code"`
	OrganizationURL  string         `json:"organization_url"`
	CorporationURLA  string         `json:"corporation_url_a"`
	CorporationURLB  string         `json:"corporation_url_b"`
	RealName         string         `json:"real_name"`
	CertificateNo    string         `json:"certificate_no"`
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

type GetUserInfoResp struct {
	Username     string           `json:"username"`
	AvatarURL    string           `json:"avatarUrl"`
	UserRole     model.UserRole   `json:"userRole"`
	VerifyStatus UserVerifyStatus `json:"verifyStatus"`
	Organization string           `json:"organization"`
	Point        int              `json:"point"`
}

type UserVerifyStatus int32

const (
	UserVerifyStatusUnfinished UserVerifyStatus = 0
	UserVerifyStatusDone       UserVerifyStatus = 1
)

type UserModifyInfoReq struct {
	Username  string `json:"username"`
	AvatarURL string `json:"avatarURL"`
}

type UserModifyInfoResp struct {
}
