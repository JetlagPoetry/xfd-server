package types

type WxOpenIDResp struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrorCode  int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type WxOrderRequest struct {
	MchID       string `json:"mchid"`
	OutTradeNo  string `json:"out_trade_no"`
	AppID       string `json:"appid"`
	Description string `json:"description"`
	NotifyURL   string `json:"notify_url"`
	Amount      struct {
		Total    int    `json:"total"`
		Currency string `json:"currency"`
	} `json:"amount"`
	Payer struct {
		OpenID string `json:"openid"`
	} `json:"payer"`
}

type WxOrderResponse struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	PrepayID string `json:"prepay_id"`
}
