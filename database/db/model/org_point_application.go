package model

import (
	"gorm.io/gorm"
	"time"
)

type OrgPointApplication struct {
	gorm.Model
	OrganizationID int                       `gorm:"column:organization_id;not null" json:"organization_id"`
	FileURL        string                    `gorm:"column:file_url;not null" json:"file_url"`
	Status         OrgPointApplicationStatus `gorm:"column:status;not null" json:"status"`
	Comment        string                    `gorm:"column:comment;not null" json:"comment"`
	VerifyTime     time.Time                 `gorm:"column:verify_time;not null" json:"verify_time"`
	VerifyComment  string                    `gorm:"column:verify_comment;not null" json:"verify_comment"`
	VerifyUserID   string                    `gorm:"column:verify_user_id;not null" json:"verify_user_id"`
	VerifyUsername string                    `gorm:"column:verify_username;not null" json:"verify_username"`
	Deleted        int                       `gorm:"column:deleted" json:"deleted"`
}

func (u *OrgPointApplication) TableName() string {
	return "org_point_application"
}

type OrgPointApplicationStatus int

const (
	OrgPointApplicationStatusUnknown  OrgPointApplicationStatus = 0
	OrgPointApplicationStatusApproved OrgPointApplicationStatus = 1 // 审核通过
	OrgPointApplicationStatusRejected OrgPointApplicationStatus = 2 // 审核失败
)
