package isauser

import (
	"context"
	models "property/framework/models"
)

// Repository :
type Repository interface {
	GetBySaUser(ctx context.Context, userID int64) (result models.SaUser, err error)
	GetAllSaUser(ctx context.Context) (result []models.SaUser, err error)
	CreateSaUser(ctx context.Context, userData *models.SaUser) (err error)
	UpdateSaUser(ctx context.Context, userData *models.SaUser) (err error)
	DeleteSaUser(ctx context.Context, userID int64) (err error)
}

// Usercase :
type Usercase interface {
	GetBySaUser(ctx context.Context, userID int64) (result models.SaUser, err error)
	GetAllSaUser(ctx context.Context) (result []models.SaUser, err error)
	CreateSaUser(ctx context.Context, userData *models.SaUser) (err error)
	UpdateSaUser(ctx context.Context, userData *models.SaUser) (err error)
	DeleteSaUser(ctx context.Context, userID int64) (err error)
}
