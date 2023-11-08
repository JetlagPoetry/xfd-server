package types

import (
	"errors"
)

type PageRequest struct {
	PageNum  int `form:"pageNum" binding:"required,gte=1"`
	PageSize int `form:"pageSize" binding:"required,gte=1"`
}

type PageResult struct {
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
	TotalNum int64 `json:"totalNum"`
}

func (r *PageRequest) Offset() int {
	return (r.PageNum - 1) * r.PageSize
}

func (r *PageRequest) Limit() int {
	return r.PageSize
}

func (r *PageRequest) CheckParams() error {
	if r.PageNum < 0 {
		return errors.New("page num is negative")
	}
	if r.PageNum == 0 {
		r.PageNum = 1
	}
	if r.PageSize < 0 {
		return errors.New("page size is negative")
	}
	if r.PageSize == 0 {
		r.PageSize = 6
	}
	return nil
}
