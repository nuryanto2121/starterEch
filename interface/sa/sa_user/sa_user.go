package isauser

import (
	"context"
	"property/framework/models"
	sa_models "property/framework/models/sa"
)

// Repository :
type Repository interface {
	GetBySaUser(ctx context.Context, userID int16) (result sa_models.SaUser, err error)
	GetList(ctx context.Context, queryparam models.ParamList) (result []*sa_models.SaUser, err error)
	CreateSaUser(ctx context.Context, userData *sa_models.SaUser) (err error)
	UpdateSaUser(ctx context.Context, userData *sa_models.SaUser) (err error)
	DeleteSaUser(ctx context.Context, userID int16) (err error)
	CountUserList(ctx context.Context, queryparam models.ParamList) (result int, err error)
}

// Usercase :
type Usercase interface {
	GetBySaUser(ctx context.Context, userID int16) (result sa_models.SaUser, err error)
	GetList(ctx context.Context, queryparam models.ParamList) (result models.ResponseModelList, err error)
	CreateSaUser(ctx context.Context, userData *sa_models.SaUser) (err error)
	UpdateSaUser(ctx context.Context, userData *sa_models.SaUser) (err error)
	DeleteSaUser(ctx context.Context, userID int16) (err error)
}
