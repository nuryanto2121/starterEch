package models

import "time"

// SaGroup :
type SaGroup struct {
	GroupID   int16     `json:"group_id" gorm:"PRIMARY_KEY"`
	Descs     string    `json:"descs" gorm:"type:varchar(60)"`
	Num       int16     `json:"num" gorm:"type:integer"`
	CreatedBy string    `json:"created_by" gorm:"type:varchar(20);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp(0) without time zone;default:now()"`
	UpdatedBy string    `json:"updated_by" gorm:"type:varchar(20);not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp(0) without time zone;default:now()"`
}
