package reposacompany

import (
	"context"
	"fmt"
	isacompany "property/framework/interface/sa/sa_company"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	"property/framework/pkg/logging"
	"property/framework/pkg/setting"

	"github.com/jinzhu/gorm"
)

type repoSaCompany struct {
	Conn *gorm.DB
}

// NewRepoSaCompany :
func NewRepoSaCompany(Conn *gorm.DB) isacompany.Repository {
	return &repoSaCompany{Conn}
}

func (db *repoSaCompany) GetBySaCompany(ctx context.Context, companyID int) (result sa_models.SaCompany, err error) {
	var (
		a      = sa_models.SaCompany{}
		logger = logging.Logger{}
	)
	query := db.Conn.Where("company_id = ?", companyID).First(&a)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr()))
	err = query.Error

	if err != nil {
		//
		if err == gorm.ErrRecordNotFound {
			return a, models.ErrNotFound
		}
		return a, err
	}

	return a, err
}

func (db *repoSaCompany) GetList(ctx context.Context, queryparam models.ParamList) (result []*sa_models.SaCompany, err error) {
	var (
		pageNum  = 0
		pageSize = setting.FileConfigSetting.App.PageSize
		sWhere   = ""
		logger   = logging.Logger{}
		orderBy  = "created_at desc"
	)
	// pagination
	if queryparam.Page > 0 {
		pageNum = (queryparam.Page - 1) * queryparam.PerPage
	}
	if queryparam.PerPage > 0 {
		pageSize = queryparam.PerPage
	}
	//end pagination

	// Order
	if queryparam.SortField != "" {
		orderBy = queryparam.SortField
	}
	//end Order by

	// WHERE
	if queryparam.InitSearch != "" {
		sWhere = queryparam.InitSearch
	}

	if queryparam.Search != "" {
		if sWhere != "" {
			sWhere += " and " + queryparam.Search
		} else {
			sWhere += queryparam.Search
		}
	}
	// end where
	if pageNum >= 0 && pageSize > 0 {
		query := db.Conn.Where(sWhere).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)
		logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
		err = query.Error
	} else {
		query := db.Conn.Where(sWhere).Order(orderBy).Find(&result)
		logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
		err = query.Error
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return result, nil
}

func (db *repoSaCompany) CreateSaCompany(ctx context.Context, companyData *sa_models.SaCompany) (err error) {
	var (
		logger = logging.Logger{}
	)
	query := db.Conn.Create(companyData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Create(companyData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaCompany) UpdateSaCompany(ctx context.Context, companyData *sa_models.SaCompany) (err error) {
	var (
		logger = logging.Logger{}
	)
	query := db.Conn.Save(companyData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Save(companyData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaCompany) DeleteSaCompany(ctx context.Context, companyID int) (err error) {
	var (
		logger = logging.Logger{}
	)
	companyData := &sa_models.SaCompany{}
	companyData.CompanyID = companyID

	query := db.Conn.Delete(&companyData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Where("company_id = ?", companyID).Delete(&companyData).Error
	// err = db.Conn.Delete(&companyData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaCompany) CountCompanyList(ctx context.Context, queryparam models.ParamList) (result int, err error) {
	var (
		logger = logging.Logger{}
		sWhere = ""
	)
	result = 0
	// WHERE
	if queryparam.InitSearch != "" {
		sWhere = queryparam.InitSearch
	}

	if queryparam.Search != "" {
		if sWhere != "" {
			sWhere += " and " + queryparam.Search
		} else {
			sWhere += queryparam.Search
		}
	}
	// end where

	query := db.Conn.Model(&sa_models.SaCompany{}).Where(sWhere).Count(&result)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error

	return result, err
}
