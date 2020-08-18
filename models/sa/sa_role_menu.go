package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type SaRoleMenu struct {
	RoleID    uuid.UUID `json:"role_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	MenuID    int       `json:"menu_id" gorm:"primary_key;type:integer;not null"`
	IsRead    bool      `json:"is_read" gorm:"type:boolean"`
	IsWrite   bool      `json:"is_write" gorm:"type:boolean"`
	CreatedBy string    `json:"created_by" gorm:"type:varchar(20);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp(0) without time zone;default:now()"`
	UpdatedBy string    `json:"updated_by" gorm:"type:varchar(20);not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp(0) without time zone;default:now()"`
}

type MenuRole struct {
	MenuID       int    `json:"menu_id"`
	Title        string `json:"title" `
	MenuUrl      string `json:"menu_url" `
	ParentMenuID int    `json:"parent_menu_id"`
	IconClass    string `json:"icon_class" `
	OrderSeq     int    `json:"order_seq" `
	Level        int    `json:"level" `
	IsRead       bool   `json:"is_read" `
	IsWrite      bool   `json:"is_write" `
	Ipath        int    `json:"ipath"`
}
