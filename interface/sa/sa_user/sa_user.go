package isauser

import (
	"context"
	"property/framework/models"
	sa_models "property/framework/models/sa"

	uuid "github.com/satori/go.uuid"
)

// Repository :
type Repository interface {
	GetBySaUser(ctx context.Context, userID uuid.UUID) (result sa_models.SaUser, err error)
	GetByEmailSaUser(ctx context.Context, email string) (result sa_models.SaUser, err error)
	GetList(ctx context.Context, queryparam models.ParamList) (result []*sa_models.SaUser, err error)
	GetJsonPermission(ctx context.Context, userID uuid.UUID, clientID uuid.UUID) (result string, err error)
	CreateSaUser(ctx context.Context, userData *sa_models.SaUser) (err error)
	UpdateSaUser(ctx context.Context, userData *sa_models.SaUser) (err error)
	DeleteSaUser(ctx context.Context, userID uuid.UUID) (err error)
	CountUserList(ctx context.Context, queryparam models.ParamList) (result int, err error)
}

// Usecase :
type Usecase interface {
	GetBySaUser(ctx context.Context, userID uuid.UUID) (result sa_models.SaUser, err error)
	GetByEmailSaUser(ctx context.Context, email string) (result sa_models.SaUser, err error)
	GetJsonPermission(ctx context.Context, userID uuid.UUID, clientID uuid.UUID) (result []map[string]interface{}, err error)
	GetList(ctx context.Context, queryparam models.ParamList) (result models.ResponseModelList, err error)
	CreateSaUser(ctx context.Context, userData *sa_models.SaUser, dataPermission *[]models.Permission) (err error)
	UpdateSaUser(ctx context.Context, userData *sa_models.SaUser, dataPermission *[]models.Permission) (err error)
	DeleteSaUser(ctx context.Context, userID uuid.UUID) (err error)
}
