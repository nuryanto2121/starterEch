package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// SaRole :
type SaRole struct {
	RoleID     uuid.UUID `json:"role_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Descs      string    `json:"descs" gorm:"type:varchar(60)"`
	Remarks    string    `json:"remarks" gorm:"type:varchar(255)"`
	RoleStatus int       `json:"role_status" gorm:"type:integer"`
	CreatedBy  string    `json:"created_by" gorm:"type:varchar(20);not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:timestamp(0) without time zone;default:now()"`
	UpdatedBy  string    `json:"updated_by" gorm:"type:varchar(20);not null"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:timestamp(0) without time zone;default:now()"`
}

// AddRoleForm :
type AddRoleForm struct {
	Descs      string `json:"descs" valid:"MaxSize(60)"`
	RoleStatus int    `json:"role_status"`
	Remarks    string `json:"notes"`
}

// EditRoleForm :
type EditRoleForm struct {
	Descs      string `json:"descs" valid:"MaxSize(60)"`
	RoleStatus int    `json:"role_status"`
	Remarks    string `json:"notes"`
}

// GetMenuAccess :
type GetMenuAccess struct {
	RoleID uuid.UUID `json:"role_id" `
}
