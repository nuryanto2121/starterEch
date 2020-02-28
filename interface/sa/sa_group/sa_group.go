package isagroup

import (
	"context"
	"property/framework/models"
)

// Repository :
type Repository interface {
	GetBySaGroup(ctx context.Context, groupID int16) (models.SaGroup, error)
	GetList(ctx context.Context, queryparam models.ParamList) ([]*models.SaGroup, error)
	CreateSaGroup(ctx context.Context, groupData *models.SaGroup) error
	UpdateSaGroup(ctx context.Context, groupData *models.SaGroup) error
	DeleteSaGroup(ctx context.Context, groupID int16) error
	CountGroupList(ctx context.Context, queryparam models.ParamList) (int, error)
}

// UseCase :
type UseCase interface {
	GetBySaGroup(ctx context.Context, groupID int16) (models.SaGroup, error)
	GetList(ctx context.Context, queryparam models.ParamList) (models.ResponseModelList, error)
	CreateSaGroup(ctx context.Context, groupData *models.SaGroup) error
	UpdateSaGroup(ctx context.Context, groupData *models.SaGroup) error
	DeleteSaGroup(ctx context.Context, groupID int16) error
}
