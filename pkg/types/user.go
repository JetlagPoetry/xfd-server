package types

import "xfd-backend/database/db/model"

type Jscode2SessionResponse struct {
	ErrorMsg   string `json:"errmsg"`
	ErrorCode  int32  `json:"errcode"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	OpenID     string `json:"openid"`
}

type UserLoginReq struct {
	Code     string         `json:"code"`
	UserRole model.UserRole `json:"userRole"`
}

type UserLoginResp struct {
	AccessToken string `json:"accessToken"` // 访问令牌
	TokenType   string `json:"tokenType"`   // 令牌类型
	ExpiresAt   int64  `json:"expiresAt"`   // 令牌到期时间戳
}

type UserRefreshTokenReq struct {
}

type UserRefreshTokenResp struct {
	AccessToken string `json:"accessToken"` // 访问令牌
	TokenType   string `json:"tokenType"`   // 令牌类型
	ExpiresAt   int64  `json:"expiresAt"`   // 令牌到期时间戳
}
