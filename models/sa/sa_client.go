package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

// SaClient :
type SaClient struct {
	ClientID         uuid.UUID `json:"client_id" gorm:"type:uuid;primary_key;"`
	ClientName       string    `json:"client_name" gorm:"type:varchar(60);not null"`
	Address          string    `json:"address" gorm:"type:varchar(150)"`
	PostCd           string    `json:"post_cd" gorm:"type:varchar(60)"`
	TelephoneNo      string    `json:"telephone_no" gorm:"type:varchar(15)"`
	EmailAddr        string    `json:"email_addr" gorm:"type:varchar(60)"`
	ContactPerson    string    `json:"contact_person" gorm:"type:varchar(60)"`
	ClientType       string    `json:"client_type" gorm:"type:varchar(60)"`
	JoiningDate      time.Time `json:"joining_date" gorm:"type:timestamp(0) without time zone;default:now()"`
	StartBillingDate time.Time `json:"start_billing_date" gorm:"type:timestamp(0) without time zone"`
	ExpiryDate       time.Time `json:"expiry_date" gorm:"type:timestamp(0) without time zone"`
	CreatedBy        string    `json:"created_by" gorm:"type:varchar(20);not null"`
	CreatedAt        time.Time `json:"created_at" gorm:"type:timestamp(0) without time zone;default:now()"`
	UpdatedBy        string    `json:"updated_by" gorm:"type:varchar(20);not null"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"type:timestamp(0) without time zone;default:now()"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *SaClient) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.New()
	if err != nil {
		return err
	}
	return scope.SetColumn("ClientID", uuid)
}
