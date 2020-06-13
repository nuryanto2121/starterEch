package reposamenu

import (
	"context"
	"fmt"
	isamenu "property/framework/interface/sa/sa_menu"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	"property/framework/pkg/logging"
	"property/framework/pkg/setting"

	"github.com/jinzhu/gorm"
)

type repoSaMenu struct {
	Conn *gorm.DB
}

func NewRepoSaMenu(Conn *gorm.DB) isamenu.Repository {
	return &repoSaMenu{Conn}
}

func (db *repoSaMenu) GetBySaMenu(ctx context.Context, ID int) (sa_models.SaMenu, error) {
	var (
		dataMenu = sa_models.SaMenu{}
		logger   = logging.Logger{}
		err      error
	)
	query := db.Conn.Where("role_id = ?", ID).First(&dataMenu)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr()))
	err = query.Error

	if err != nil {
		//
		if err == gorm.ErrRecordNotFound {
			return dataMenu, models.ErrNotFound
		}
		return dataMenu, err
	}

	return dataMenu, err
}

func (db *repoSaMenu) GetList(ctx context.Context, queryparam models.ParamList) ([]*sa_models.SaMenu, error) {
	var (
		pageNum  = 0
		pageSize = setting.FileConfigSetting.App.PageSize
		sWhere   = ""
		logger   = logging.Logger{}
		orderBy  = "created_at desc"
		err      error
		result   = []*sa_models.SaMenu{}
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

func (db *repoSaMenu) CreateSaMenu(ctx context.Context, roleData *sa_models.SaMenu) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Create(roleData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Create(userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaMenu) UpdateSaMenu(ctx context.Context, roleData *sa_models.SaMenu) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Save(roleData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Save(userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaMenu) DeleteSaMenu(ctx context.Context, ID int) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	userData := &sa_models.SaMenu{}
	userData.MenuID = ID

	query := db.Conn.Delete(&userData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error

	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaMenu) CountMenuList(ctx context.Context, queryparam models.ParamList) (int, error) {
	var (
		logger = logging.Logger{}
		sWhere = ""
		err    error
		result = 0
	)

	// WHERE
	if queryparam.InitSearch != "" {
		sWhere = queryparam.InitSearch
	}

	if queryparam.Search != "" {
		if sWhere != "" {
			sWhere += " and " + queryparam.Search
		}
	}
	// end where

	query := db.Conn.Model(&sa_models.SaUser{}).Where(sWhere).Count(&result)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error

	return result, err
}
