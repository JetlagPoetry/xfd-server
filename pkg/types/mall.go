package types

import (
	"fmt"
	"xfd-backend/database/db/enum"
)

type CategoryListReq struct {
	Level    int32 `form:"level" binding:"required,gte=1,lte=3"`
	ParentID int32 `form:"parentID" binding:"numeric"`
	ParentId int32 `form:"parentId" binding:"numeric"`
}

func (c CategoryListReq) CheckParams() error {
	if c.Level != 1 {
		if c.ParentId == 0 && c.ParentID == 0 {
			return fmt.Errorf("parentId should be filled")
		}
	}
	return nil
}

type CategoryListResp struct {
	List []*CategoryDetail `json:"category,omitempty"`
}

type CategoryDetail struct {
	ParentID int              `json:"parentID"`
	Name     string           `json:"name"`
	Image    string           `json:"image"`
	Children []CategoryDetail `json:"children"`
}

type AreaReq struct {
	Code int `json:"code" form:"code" binding:"numeric,gte=0"`
}
type AreaList struct {
	Name     string  `json:"name"`
	Code     int     `json:"code"`
	Level    int     `json:"level"`
	Children []*Area `json:"children,omitempty"`
}

type Area struct {
	Name  string `json:"name"`
	Code  int    `json:"code"`
	Level int    `json:"level"`
}

type CategoryAddReq struct {
	ParentID        int32                   `json:"parentID"`
	Level           enum.GoodsCategoryLevel `json:"level" binding:"required,gte=1,lte=3"`
	CategoryDetails []*AddCategoryDetails   `json:"categoryDetails" binding:"required,len=1"`
}

type AddCategoryDetails struct {
	Name  string `json:"name" binding:"required"`
	Image string `json:"image" `
}

func (r *CategoryAddReq) CheckParams() error {
	if r.ParentID == 0 && r.Level != enum.LevelOne {
		return fmt.Errorf("parentID must be filled")
	}
	return nil
}

type CategoryDeleteReq struct {
	IDs []int32 `json:"ids" binding:"required"`
}

func (r *CategoryDeleteReq) CheckParams() error {
	if len(r.IDs) == 0 {
		return fmt.Errorf("ids must be filled")
	}
	for _, id := range r.IDs {
		if id == 0 {
			return fmt.Errorf("ids must be filled")
		}
	}
	return nil
}

type CategoryModifyReq struct {
	ID    int32  `json:"id" binding:"required"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
