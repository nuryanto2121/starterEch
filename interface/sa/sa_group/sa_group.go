package isagroup

import (
	"context"
	"property/framework/models"
	sa_models "property/framework/models/sa"
)

// Repository :
type Repository interface {
	GetBySaGroup(ctx context.Context, groupID int16) (sa_models.SaGroup, error)
	GetList(ctx context.Context, queryparam models.ParamList) ([]*sa_models.SaGroup, error)
	CreateSaGroup(ctx context.Context, groupData *sa_models.SaGroup) error
	UpdateSaGroup(ctx context.Context, groupData *sa_models.SaGroup) error
	DeleteSaGroup(ctx context.Context, groupID int16) error
	CountGroupList(ctx context.Context, queryparam models.ParamList) (int, error)
}

// UseCase :
type UseCase interface {
	GetBySaGroup(ctx context.Context, groupID int16) (sa_models.SaGroup, error)
	GetList(ctx context.Context, queryparam models.ParamList) (models.ResponseModelList, error)
	CreateSaGroup(ctx context.Context, groupData *sa_models.SaGroup) error
	UpdateSaGroup(ctx context.Context, groupData *sa_models.SaGroup) error
	DeleteSaGroup(ctx context.Context, groupID int16) error
}
