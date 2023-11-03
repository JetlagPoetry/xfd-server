package types

type PageRequest struct {
	PageNum  int `form:"pageNum" binding:"required,gte=1"`
	PageSize int `form:"pageSize" binding:"required,gte=1"`
}

type PageResult struct {
	//PageNum  int   `json:"pageNum"`
	//PageSize int   `json:"pageSize"`
	TotalNum int64 `json:"totalNum"`
}
