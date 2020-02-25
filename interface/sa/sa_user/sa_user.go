package isauser

import (
	"context"
	models "property/framework/models"
)

// Repository :
type Repository interface {
	GetBySaUser(ctx context.Context, userID int16) (result models.SaUser, err error)
	GetList(ctx context.Context, queryparam models.ParamList) (result []*models.SaUser, err error)
	CreateSaUser(ctx context.Context, userData *models.SaUser) (err error)
	UpdateSaUser(ctx context.Context, userData *models.SaUser) (err error)
	DeleteSaUser(ctx context.Context, userID int16) (err error)
	CountUserList(ctx context.Context, queryparam models.ParamList) (result int, err error)
}

// Usercase :
type Usercase interface {
	GetBySaUser(ctx context.Context, userID int16) (result models.SaUser, err error)
	GetList(ctx context.Context, queryparam models.ParamList) (result models.ResponseModelList, err error)
	CreateSaUser(ctx context.Context, userData *models.SaUser) (err error)
	UpdateSaUser(ctx context.Context, userData *models.SaUser) (err error)
	DeleteSaUser(ctx context.Context, userID int16) (err error)
}
