package types

type CategoryListReq struct {
	Level    int32 `form:"level" binding:"required,gte=1,lte=3"`
	ParentID int32 `form:"parentID" binding:"numeric"`
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
