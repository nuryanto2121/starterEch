package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type SaFileUpload struct {
	FileID    uuid.UUID `json:"file_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	FileName  string    `json:"file_name" gorm:"type:varchar(100);not null"`
	FilePath  string    `json:"file_path" gorm:"type:varchar(150);not null"`
	FileType  string    `json:"file_type" gorm:"type:varchar(10)"`
	CreatedBy string    `json:"created_by" gorm:"type:varchar(20);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp(0) without time zone;default:now()"`
	UpdatedBy string    `json:"updated_by" gorm:"type:varchar(20);not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp(0) without time zone;default:now()"`
}

type SaFileOutput struct {
	FileName string `json:"file_name" gorm:"type:varchar(100);not null"`
	FilePath string `json:"file_path" gorm:"type:varchar(150);not null"`
	FileType string `json:"file_type" gorm:"type:varchar(10)"`
}
