package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type SaRoleMenu struct {
	RoleID    uuid.UUID `json:"role_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	MenuLevel string    `json:"menu_level" gorm:"type:varchar(5)"`
	MenuID    int       `json:"menu_id" gorm:"type:varchar(255)"`
	CreatedBy string    `json:"created_by" gorm:"type:varchar(20);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp(0) without time zone;default:now()"`
	UpdatedBy string    `json:"updated_by" gorm:"type:varchar(20);not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp(0) without time zone;default:now()"`
}
