package models

import "time"

// SaBranch :
type SaBranch struct {
	CompanyID      int16     `json:"company_id" gorm:"primary_key;type:integer;not null"`
	BranchID       int16     `json:"branch_id" gorm:"primary_key;type:integer;not null"`
	BranchName     string    `json:"branch_name" gorm:"type:varchar(60);not null"`
	Address        string    `json:"address" gorm:"type:varchar(150)"`
	PostCd         string    `json:"post_cd" gorm:"type:varchar(60)"`
	TelephoneNo    string    `json:"telephone_no" gorm:"type:varchar(15)"`
	EmailAddr      string    `json:"email_addr" gorm:"type:varchar(60)"`
	ContactPerson  string    `json:"contact_person" gorm:"type:varchar(60)"`
	BaseDivID      int16     `json:"base_div_id	" gorm:"type:varchar(60)"`
	BaseDeptID     int16     `json:"base_dept_id	" gorm:"type:varchar(60)"`
	BaseCurrencyID int16     `json:"base_currency_id	" gorm:"type:varchar(60)"`
	TaxRegNo       string    `json:"tax_reg_no" gorm:"type:varchar(60)"`
	StartDate      time.Time `json:"start_date" gorm:"type:timestamp(0) without time zone"`
	ExpiryDate     time.Time `json:"expiry_date" gorm:"type:timestamp(0) without time zone"`
	LogoUrl        string    `json:"logo_url" gorm:"type:varchar(60)"`
	CreatedBy      string    `json:"created_by" gorm:"type:varchar(20);not null"`
	CreatedAt      time.Time `json:"created_at" gorm:"type:timestamp(0) without time zone;default:now()"`
	UpdatedBy      string    `json:"updated_by" gorm:"type:varchar(20);not null"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"type:timestamp(0) without time zone;default:now()"`
}
