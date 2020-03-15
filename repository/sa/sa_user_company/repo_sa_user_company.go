package reposausercompany

import (
	"context"
	"fmt"
	isausercompany "property/framework/interface/sa/sa_user_company"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	"property/framework/pkg/logging"
	"property/framework/pkg/setting"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type repoSaUserCompany struct {
	Conn *gorm.DB
}

// NewRepoSaUserCompany :
func NewRepoSaUserCompany(Conn *gorm.DB) isausercompany.Repository {
	return &repoSaUserCompany{Conn}
}

func (db *repoSaUserCompany) GetBySaUserCompany(ctx context.Context, userID uuid.UUID) (result sa_models.SaUserCompany, err error) {
	var (
		a      = sa_models.SaUserCompany{}
		logger = logging.Logger{}
	)
	query := db.Conn.Where("user_id = ?", userID).First(&a)
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

func (db *repoSaUserCompany) GetList(ctx context.Context, queryparam models.ParamList) (result []*sa_models.SaUserCompany, err error) {
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

func (db *repoSaUserCompany) CreateSaUserCompany(ctx context.Context, usercompanyData *sa_models.SaUserCompany) (err error) {
	var (
		logger = logging.Logger{}
	)
	query := db.Conn.Create(usercompanyData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Create(usercompanyData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaUserCompany) UpdateSaUserCompany(ctx context.Context, usercompanyData *sa_models.SaUserCompany) (err error) {
	var (
		logger = logging.Logger{}
	)
	query := db.Conn.Save(usercompanyData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Save(usercompanyData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaUserCompany) DeleteSaUserCompany(ctx context.Context, userID uuid.UUID) (err error) {
	var (
		logger = logging.Logger{}
	)
	usercompanyData := &sa_models.SaUserCompany{}
	usercompanyData.UserID = userID

	query := db.Conn.Delete(&usercompanyData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Where("usercompany_id = ?", userID).Delete(&usercompanyData).Error
	// err = db.Conn.Delete(&usercompanyData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaUserCompany) CountUserCompanyList(ctx context.Context, queryparam models.ParamList) (result int, err error) {
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

	query := db.Conn.Model(&sa_models.SaUserCompany{}).Where(sWhere).Count(&result)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error

	return result, err
}
