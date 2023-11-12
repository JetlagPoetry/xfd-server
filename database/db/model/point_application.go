package model

import (
	"gorm.io/gorm"
	"time"
)

type PointApplication struct {
	gorm.Model
	OrganizationID int                    `gorm:"column:organization_id;not null" json:"organization_id"`
	TotalPoint     float64                `gorm:"column:total_point;not null" json:"total_point"`
	FileURL        string                 `gorm:"column:file_url;not null" json:"file_url"`
	Status         PointApplicationStatus `gorm:"column:status;not null" json:"status"`
	Comment        string                 `gorm:"column:comment;not null" json:"comment"`
	StartTime      time.Time              `gorm:"column:start_time;not null" json:"start_time"`
	EndTime        time.Time              `gorm:"column:end_time;not null" json:"end_time"`
	ApplyUserID    string                 `gorm:"column:apply_user_id;not null" json:"apply_user_id"`
	VerifyTime     time.Time              `gorm:"column:verify_time;not null" json:"verify_time"`
	VerifyComment  string                 `gorm:"column:verify_comment;not null" json:"verify_comment"`
	VerifyUserID   string                 `gorm:"column:verify_user_id;not null" json:"verify_user_id"`
	VerifyUsername string                 `gorm:"column:verify_username;not null" json:"verify_username"`
}

func (u *PointApplication) TableName() string {
	return "point_application"
}

type PointApplicationStatus int

const (
	PointApplicationStatusUnknown  PointApplicationStatus = 0
	PointApplicationStatusApproved PointApplicationStatus = 1 // 审核通过
	PointApplicationStatusRejected PointApplicationStatus = 2 // 审核失败
	PointApplicationStatusFinish   PointApplicationStatus = 3 // 发放完成
	PointApplicationStatusClear    PointApplicationStatus = 4 // 已清零
	PointApplicationStatusExpired  PointApplicationStatus = 5 // 已过期
)
