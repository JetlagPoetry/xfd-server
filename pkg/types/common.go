package types

type CommonDeleteUploadReq struct {
	Link string `form:"link" binding:"required,url"`
}

type CommonUploadResp struct {
	Link string `json:"link"`
}
