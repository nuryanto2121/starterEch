package models

import (
	m "property/framework/models"
	"time"

	uuid "github.com/satori/go.uuid"
)

// SaUser :
type SaUser struct {
	UserID      uuid.UUID    `json:"user_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	ClientID    uuid.UUID    `json:"client_id" gorm:"type:uuid; not null"`
	RoleID      uuid.UUID    `json:"role_id" gorm:"type:uuid;not null"`
	LevelNo     int          `json:"level_no" gorm:"type:integer;default:0;not null"`
	UserName    string       `json:"user_name" gorm:"type:varchar(60);not null"`
	Passwd      string       `json:"passwd" gorm:"type:varchar(60);not null"`
	Name        string       `json:"name" gorm:"type:varchar(60);not null"`
	EmailAddr   string       `json:"email_addr" gorm:"type:varchar(100);unique_index"`
	HandphoneNo string       `json:"handphone_no" gorm:"type:varchar(20)"`
	CompanyID   int          `json:"company_id" gorm:"type:integer;not null"`
	FileID      uuid.UUID    `json:"file_id" gorm:"type:uuid"`
	UserStatus  int          `json:"user_status" gorm:"type:integer;default:1"`
	CreatedBy   string       `json:"created_by" gorm:"type:varchar(20);not null"`
	CreatedAt   time.Time    `json:"created_at" gorm:"type:timestamp(0) without time zone;default:now()"`
	UpdatedBy   string       `json:"updated_by" gorm:"type:varchar(20);not null"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:"type:timestamp(0) without time zone;default:now()"`
	DataFile    SaFileOutput `json:"data_file"`
}

// AddUserForm : param from frond end
type AddUserForm struct {
	ClientID       uuid.UUID      `json:"client_id" valid:"Required"`
	UserName       string         `json:"user_name" valid:"Required"`
	Name           string         `json:"name" valid:"Required"`
	EmailAddr      string         `json:"email_addr" valid:"Required"`
	LevelNo        int            `json:"level_no"`
	UserStatus     int            `json:"user_status"`
	RoleID         uuid.UUID      `json:"role_id" valid:"Required"`
	HandphoneNo    string         `json:"handphone_no"`
	CompanyID      int            `json:"company_id" valid:"Required"`
	FileID         uuid.UUID      `json:"file_id"`
	CreatedBy      string         `json:"created_by" valid:"Required"`
	DataPermission []m.Permission `json:"data_permission"`
	// Passwd         string         `json:"passwd" valid:"Required"`
	// ConfimPasswd   string         `json:"confirm_passwd" valid:"Required"`
}

// EditUserForm : param from frond end
type EditUserForm struct {
	ClientID       uuid.UUID      `json:"client_id" valid:"Required"`
	UserName       string         `json:"user_name" valid:"Required"`
	Name           string         `json:"name" valid:"Required"`
	EmailAddr      string         `json:"email_addr" valid:"Required"`
	LevelNo        int            `json:"level_no"`
	UserStatus     int            `json:"user_status"`
	RoleID         uuid.UUID      `json:"role_id" valid:"Required"`
	HandphoneNo    string         `json:"handphone_no"`
	CompanyID      int            `json:"company_id" valid:"Required"`
	FileID         uuid.UUID      `json:"file_id"`
	UpdatedBy      string         `json:"updated_by" valid:"Required"`
	DataPermission []m.Permission `json:"data_permission"`
}

type GetPermissionForm struct {
	ClinetID uuid.UUID `json:"client_id" valid:"Required"`
	UserID   uuid.UUID `json:"user_id"`
}
