package isacompany

import (
	"context"
	"property/framework/models"
	sa_models "property/framework/models/sa"
)

// Repository :
type Repository interface {
	GetBySaCompany(ctx context.Context, companyID int) (result sa_models.SaCompany, err error)
	GetList(ctx context.Context, queryparam models.ParamList) (result []*sa_models.SaCompany, err error)
	CreateSaCompany(ctx context.Context, companyData *sa_models.SaCompany) (err error)
	UpdateSaCompany(ctx context.Context, companyData *sa_models.SaCompany) (err error)
	DeleteSaCompany(ctx context.Context, companyID int) (err error)
	CountCompanyList(ctx context.Context, queryparam models.ParamList) (result int, err error)
}

// Companycase :
type Usecase interface {
	GetBySaCompany(ctx context.Context, companyID int) (result sa_models.SaCompany, err error)
	GetList(ctx context.Context, queryparam models.ParamList) (result models.ResponseModelList, err error)
	CreateSaCompany(ctx context.Context, companyData *sa_models.SaCompany) (err error)
	UpdateSaCompany(ctx context.Context, companyData *sa_models.SaCompany) (err error)
	DeleteSaCompany(ctx context.Context, companyID int) (err error)
}
