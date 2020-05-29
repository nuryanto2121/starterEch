package isabranch

import (
	"context"
	"property/framework/models"
	sa_models "property/framework/models/sa"
)

// Repository :
type Repository interface {
	GetBySaBranch(ctx context.Context, branchID int) (result sa_models.SaBranch, err error)
	GetList(ctx context.Context, queryparam models.ParamList) (result []*sa_models.SaBranch, err error)
	CreateSaBranch(ctx context.Context, branchData *sa_models.SaBranch) (err error)
	UpdateSaBranch(ctx context.Context, branchData *sa_models.SaBranch) (err error)
	DeleteSaBranch(ctx context.Context, branchID int) (err error)
	CountBranchList(ctx context.Context, queryparam models.ParamList) (result int, err error)
}

// Branchcase :
type Usecase interface {
	GetBySaBranch(ctx context.Context, branchID int) (result sa_models.SaBranch, err error)
	GetList(ctx context.Context, queryparam models.ParamList) (result models.ResponseModelList, err error)
	CreateSaBranch(ctx context.Context, branchData *sa_models.SaBranch) (err error)
	UpdateSaBranch(ctx context.Context, branchData *sa_models.SaBranch) (err error)
	DeleteSaBranch(ctx context.Context, branchID int) (err error)
}
