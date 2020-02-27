package models

import "time"

// SaUser :
type SaUser struct {
	UserID      int16     `json:"user_id" gorm:"PRIMARY_KEY"`
	Passwd      string    `json:"passwd" gorm:"type:varchar(60);not null"`
	GroupID     int16     `json:"group_id" gorm:"type:integer;not null"`
	LevelNo     int16     `json:"level_no" gorm:"type:integer;default:0;not null"`
	UserName    string    `json:"user_name" gorm:"type:varchar(60);not null"`
	EmailAddr   string    `json:"email_addr" gorm:"type:varchar(100);unique_index"`
	HandphoneNo string    `json:"handphone_no" gorm:"type:varchar(20)"`
	CompanyID   int16     `json:"company_id" gorm:"type:integer;not null"`
	ProjectID   int16     `json:"project_id" gorm:"type:integer;not null"`
	PictureURL  string    `json:"picture_url" gorm:"type:varchar(100)"`
	UserStatus  int16     `json:"user_status" gorm:"type:integer;default:1"`
	CreatedBy   string    `json:"created_by" gorm:"type:varchar(20);not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamp(0) without time zone;default:now()"`
	UpdatedBy   string    `json:"updated_by" gorm:"type:varchar(20);not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:timestamp(0) without time zone;default:now()"`
}
