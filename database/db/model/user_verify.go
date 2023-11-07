package model

import (
	"gorm.io/gorm"
	"time"
)

type UserVerify struct {
	gorm.Model
	UserID           string           `gorm:"column:user_id;not null" json:"user_id"`
	UserRole         UserRole         `gorm:"column:user_role;not null" json:"user_role"`
	Organization     string           `gorm:"column:organization;not null" json:"organization"`
	OrganizationCode string           `gorm:"column:organization_code;not null" json:"organization_code"`
	OrganizationURL  string           `gorm:"column:organization_url;not null" json:"organization_url"`
	IdentityURLA     string           `gorm:"column:identity_url_a;not null" json:"identity_url_a"`
	IdentityURLB     string           `gorm:"column:identity_url_b;not null" json:"identity_url_b"`
	RealName         string           `gorm:"column:real_name;not null" json:"real_name"`
	CertificateNo    string           `gorm:"column:certificate_no;not null" json:"certificate_no"`
	Position         string           `gorm:"column:position;not null" json:"position"`
	Phone            string           `gorm:"column:phone;not null" json:"phone"`
	Status           UserVerifyStatus `gorm:"column:status;not null" json:"status"`
	Comment          string           `gorm:"column:comment;not null" json:"comment"`
	VerifyTime       time.Time        `gorm:"column:verify_time;not null" json:"verify_time"`
	VerifyUsername   string           `gorm:"column:verify_username;not null" json:"verify_username"`
}

type UserVerifyStatus int32

const (
	UserVerifyStatusUnknown   UserVerifyStatus = 0
	UserVerifyStatusSubmitted UserVerifyStatus = 1
	UserVerifyStatusRejected  UserVerifyStatus = 2
	UserVerifyStatusSuccess   UserVerifyStatus = 3
)

func (UserVerify) TableName() string {
	return "user_verify"
}
