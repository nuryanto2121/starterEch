package isagroup

import (
	"context"
	"property/framework/models"
	sa_models "property/framework/models/sa"

	uuid "github.com/satori/go.uuid"
)

// Repository :
type Repository interface {
	GetBySaGroup(ctx context.Context, groupID uuid.UUID) (sa_models.SaGroup, error)
	GetList(ctx context.Context, queryparam models.ParamList) ([]*sa_models.SaGroup, error)
	CreateSaGroup(ctx context.Context, groupData *sa_models.SaGroup) error
	UpdateSaGroup(ctx context.Context, groupData *sa_models.SaGroup) error
	DeleteSaGroup(ctx context.Context, groupID uuid.UUID) error
	CountGroupList(ctx context.Context, queryparam models.ParamList) (int, error)
}

// UseCase :
type UseCase interface {
	GetBySaGroup(ctx context.Context, groupID uuid.UUID) (sa_models.SaGroup, error)
	GetList(ctx context.Context, queryparam models.ParamList) (models.ResponseModelList, error)
	CreateSaGroup(ctx context.Context, groupData *sa_models.SaGroup) error
	UpdateSaGroup(ctx context.Context, groupData *sa_models.SaGroup) error
	DeleteSaGroup(ctx context.Context, groupID uuid.UUID) error
}
