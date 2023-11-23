package types

type WxOpenIDResp struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrorCode  int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}
