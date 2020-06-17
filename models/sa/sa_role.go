package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// SaRole :
type SaRole struct {
	RoleID     uuid.UUID `json:"role_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Descs      string    `json:"descs" gorm:"type:varchar(60);not null;unique_index"`
	Remarks    string    `json:"remarks" gorm:"type:varchar(255)"`
	RoleStatus int       `json:"role_status" gorm:"type:integer"`
	CreatedBy  string    `json:"created_by" gorm:"type:varchar(20);not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:timestamp(0) without time zone;default:now()"`
	UpdatedBy  string    `json:"updated_by" gorm:"type:varchar(20);not null"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:timestamp(0) without time zone;default:now()"`
}

// AddRoleForm :
type AddRoleForm struct {
	Descs      string             `json:"descs" valid:"Required;MaxSize(60)"`
	RoleStatus int                `json:"role_status"`
	Remarks    string             `json:"notes"`
	MenuAccess []MenuAccessLevel1 `json:"menu_access"`
}

// EditRoleForm :
type EditRoleForm struct {
	Descs      string             `json:"descs" valid:"Required;MaxSize(60)"`
	RoleStatus int                `json:"role_status"`
	Remarks    string             `json:"notes"`
	MenuAccess []MenuAccessLevel1 `json:"menu_access"`
	UpdatedBy  string             `json:"updated_by"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

// MenuAccess :
type MenuAccess struct {
	MenuID       int    `json:"menu_id,omitempty"`
	Title        string `json:"title,omitempty"`
	ParentMenuID int    `json:"parent_menu_id"`
	OrderSeq     int    `json:"order_seq"`
	Level        int    `json:"level_no"`
	IsRead       bool   `json:"is_read"`
	IsWrite      bool   `json:"is_write"`
}

// MenuAccessLevel2 :
type MenuAccessLevel2 struct {
	MenuAccess
	Level3 []MenuAccess `json:"level_3"`
}

// MenuAccessLevel1 :
type MenuAccessLevel1 struct {
	MenuAccess
	Level2 []MenuAccessLevel2 `json:"level_2"`
}
