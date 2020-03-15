package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type SaCompany struct {
	ClientID      uuid.UUID `json:"client_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CompanyID     int16     `json:"company_id" gorm:"PRIMARY_KEY"`
	CompanyName   string    `json:"company_name" gorm:"type:varchar(60);not null;unique"`
	Address       string    `json:"address" gorm:"type:varchar(255)"`
	PostCd        string    `json:"post_cd" gorm:"type:varchar(10)"`
	TelephoneNo   string    `json:"telephone_no" gorm:"type:varchar(15)"`
	EmailAddr     string    `json:"email_addr" gorm:"type:varchar(20)"`
	ContactPerson string    `json:"contact_person" gorm:"type:varchar(20)"`
	FinYear       int16     `json:"fin_year" gorm:"type:integer"`
	FinPeriod     int16     `json:"fin_period" gorm:"type:integer"`
	StartDate     time.Time `json:"start_date" gorm:"type:timestamp(0) without time zone"`
	ExpiryDate    time.Time `json:"expiry_date" gorm:"type:timestamp(0) without time zone"`
	CreatedBy     string    `json:"created_by" gorm:"type:varchar(20);not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"type:timestamp(0) without time zone;default:now()"`
	UpdatedBy     string    `json:"updated_by" gorm:"type:varchar(20);not null"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"type:timestamp(0) without time zone;default:now()"`
}
