package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// SaUser :
type SaUser struct {
	UserID      uuid.UUID `json:"user_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Passwd      string    `json:"passwd" gorm:"type:varchar(60);not null"`
	RoleID      uuid.UUID `json:"role_id" gorm:"type:uuid;not null"`
	LevelNo     int       `json:"level_no" gorm:"type:integer;default:0;not null"`
	UserName    string    `json:"user_name" gorm:"type:varchar(60);not null"`
	Name        string    `json:"name" gorm:"type:varchar(60);not null"`
	EmailAddr   string    `json:"email_addr" gorm:"type:varchar(100);unique_index"`
	HandphoneNo string    `json:"handphone_no" gorm:"type:varchar(20)"`
	CompanyID   int       `json:"company_id" gorm:"type:integer;not null"`
	ProjectID   int       `json:"project_id" gorm:"type:integer;not null"`
	PictureURL  string    `json:"picture_url" gorm:"type:varchar(100)"`
	UserStatus  int       `json:"user_status" gorm:"type:integer;default:1"`
	CreatedBy   string    `json:"created_by" gorm:"type:varchar(20);not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamp(0) without time zone;default:now()"`
	UpdatedBy   string    `json:"updated_by" gorm:"type:varchar(20);not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:timestamp(0) without time zone;default:now()"`
}
