package isarole

import (
	"context"
	"property/framework/models"
	sa_models "property/framework/models/sa"

	uuid "github.com/satori/go.uuid"
)

// Repository :
type Repository interface {
	GetBySaRole(ctx context.Context, roleID uuid.UUID) (sa_models.SaRole, error)
	GetList(ctx context.Context, queryparam models.ParamList) ([]*sa_models.SaRole, error)
	CreateSaRole(ctx context.Context, roleData *sa_models.SaRole) error
	UpdateSaRole(ctx context.Context, roleData *sa_models.SaRole) error
	DeleteSaRole(ctx context.Context, roleID uuid.UUID) error
	CountRoleList(ctx context.Context, queryparam models.ParamList) (int, error)
	GetJsonMenuAccess(ctx context.Context, roleID uuid.UUID) (result string, err error)
}

// UseCase :
type UseCase interface {
	GetJsonMenuAccess(ctx context.Context, roleID uuid.UUID) (result []map[string]interface{}, err error)
	GetBySaRole(ctx context.Context, roleID uuid.UUID) (sa_models.SaRole, error)
	GetList(ctx context.Context, queryparam models.ParamList) (models.ResponseModelList, error)
	CreateSaRole(ctx context.Context, roleData *sa_models.SaRole) error
	UpdateSaRole(ctx context.Context, roleData *sa_models.SaRole) error
	DeleteSaRole(ctx context.Context, roleID uuid.UUID) error
}
