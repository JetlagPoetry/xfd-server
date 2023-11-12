package model

import (
	"gorm.io/gorm"
)

type PointApplicationTmp struct {
	gorm.Model
	OrganizationID int     `gorm:"column:organization_id;not null" json:"organization_id"`
	ApplicationID  int     `gorm:"column:application_id;not null" json:"application_id"`
	Username       string  `gorm:"column:username;not null" json:"username"`
	Phone          string  `gorm:"column:phone;not null" json:"phone"`
	Point          float64 `gorm:"column:point;not null" json:"point"`
}

func (u *PointApplicationTmp) TableName() string {
	return "point_application_tmp"
}

type PointApplicationTmpStatus int

const (
	PointApplicationTmpInit   PointApplicationTmpStatus = 0
	PointApplicationTmpFinish PointApplicationTmpStatus = 1
)
