package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type SaUserCompany struct {
	UserID    uuid.UUID `json:"user_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CompanyID int       `json:"company_id" gorm:"primary_key;type:integer;not null"`
	CreatedBy string    `json:"created_by" gorm:"type:varchar(20);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp(0) without time zone;default:now()"`
	UpdatedBy string    `json:"updated_by" gorm:"type:varchar(20);not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp(0) without time zone;default:now()"`
}
