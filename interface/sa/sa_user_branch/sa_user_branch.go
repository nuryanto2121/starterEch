package isauserbranch

import (
	"context"
	"property/framework/models"
	sa_models "property/framework/models/sa"

	uuid "github.com/satori/go.uuid"
)

// Repository :
type Repository interface {
	GetBySaUserBranch(ctx context.Context, userID uuid.UUID) (result sa_models.SaUserBranch, err error)
	GetList(ctx context.Context, queryparam models.ParamList) (result []*sa_models.SaUserBranch, err error)
	CreateSaUserBranch(ctx context.Context, clientData *sa_models.SaUserBranch) (err error)
	UpdateSaUserBranch(ctx context.Context, clientData *sa_models.SaUserBranch) (err error)
	DeleteSaUserBranch(ctx context.Context, userID uuid.UUID) (err error)
	CountUserBranchList(ctx context.Context, queryparam models.ParamList) (result int, err error)
}

// UserBranchcase :
type Usecase interface {
	GetBySaUserBranch(ctx context.Context, userID uuid.UUID) (result sa_models.SaUserBranch, err error)
	GetList(ctx context.Context, queryparam models.ParamList) (result models.ResponseModelList, err error)
	CreateSaUserBranch(ctx context.Context, clientData *sa_models.SaUserBranch) (err error)
	UpdateSaUserBranch(ctx context.Context, clientData *sa_models.SaUserBranch) (err error)
	DeleteSaUserBranch(ctx context.Context, userID uuid.UUID) (err error)
}
