package models

import "time"

type SaMenu struct {
	MenuID       int       `json:"menu_id" gorm:"primary_key"`
	Title        string    `json:"title" gorm:"type:varchar(60);unique_index;not null"`
	MenuUrl      string    `json:"menu_url" gorm:"type:varchar(60);not null"`
	ParentMenuID int       `json:"parent_menu_id" gorm:"type:integer"`
	IconClass    string    `json:"icon_class" gorm:"type:varchar(60)"`
	OrderSeq     int       `json:"order_seq" gorm:"type:integer;not null"`
	Level        int       `json:"level" gorm:"type:integer;not null"`
	CreatedBy    string    `json:"created_by" gorm:"type:varchar(20);not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp(0) without time zone;default:now()"`
	UpdatedBy    string    `json:"updated_by" gorm:"type:varchar(20);not null"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp(0) without time zone;default:now()"`
}
