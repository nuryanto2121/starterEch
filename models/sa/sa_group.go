package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// SaGroup :
type SaGroup struct {
	GroupID   uuid.UUID `json:"group_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Descs     string    `json:"descs" gorm:"type:varchar(60)"`
	CreatedBy string    `json:"created_by" gorm:"type:varchar(20);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp(0) without time zone;default:now()"`
	UpdatedBy string    `json:"updated_by" gorm:"type:varchar(20);not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp(0) without time zone;default:now()"`
}
