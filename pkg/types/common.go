package types

type CommonDeleteUploadReq struct {
	Link string `form:"link" binding:"required,url"`
}

type CommonUploadResp struct {
	Link string `json:"link"`
}

type CommonGetConfigReq struct {
	Key string `form:"key"`
}

type CommonGetConfigResp struct {
	Value string `json:"value"`
}
