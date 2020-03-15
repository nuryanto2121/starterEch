package isausercompany

import (
	"context"
	"property/framework/models"
	sa_models "property/framework/models/sa"

	uuid "github.com/satori/go.uuid"
)

// Repository :
type Repository interface {
	GetBySaUserCompany(ctx context.Context, UserID uuid.UUID) (result sa_models.SaUserCompany, err error)
	GetList(ctx context.Context, queryparam models.ParamList) (result []*sa_models.SaUserCompany, err error)
	CreateSaUserCompany(ctx context.Context, userCompanyData *sa_models.SaUserCompany) (err error)
	UpdateSaUserCompany(ctx context.Context, userCompanyData *sa_models.SaUserCompany) (err error)
	DeleteSaUserCompany(ctx context.Context, UserID uuid.UUID) (err error)
	CountUserCompanyList(ctx context.Context, queryparam models.ParamList) (result int, err error)
}

// UserCompanycase :
type Usecase interface {
	GetBySaUserCompany(ctx context.Context, UserID uuid.UUID) (result sa_models.SaUserCompany, err error)
	GetList(ctx context.Context, queryparam models.ParamList) (result models.ResponseModelList, err error)
	CreateSaUserCompany(ctx context.Context, userCompanyData *sa_models.SaUserCompany) (err error)
	UpdateSaUserCompany(ctx context.Context, userCompanyData *sa_models.SaUserCompany) (err error)
	DeleteSaUserCompany(ctx context.Context, UserID uuid.UUID) (err error)
}
