package types

type PageRequest struct {
	PageNum  int `form:"pageNum"`
	PageSize int `form:"pageSize"`
}

type PageResult struct {
	//PageNum  int   `json:"pageNum"`
	//PageSize int   `json:"pageSize"`
	TotalNum int64 `json:"totalNum"`
}
