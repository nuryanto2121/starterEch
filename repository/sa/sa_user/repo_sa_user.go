package reposauser

import (
	"context"
	isauser "property/framework/interface/sa/sa_user"
	models "property/framework/models"
	"property/framework/pkg/logging"
	"property/framework/pkg/setting"

	"github.com/jinzhu/gorm"
)

type repoSaUser struct {
	Conn *gorm.DB
}

// NewRepoSaUser :
func NewRepoSaUser(Conn *gorm.DB) isauser.Repository {
	return &repoSaUser{Conn}
}

func (db *repoSaUser) GetBySaUser(ctx context.Context, userID int16) (result models.SaUser, err error) {
	var (
		a      = models.SaUser{}
		logger = logging.Logger{}
	)
	query := db.Conn.Where("user_id = ?", userID).First(&a)
	logger.Query(query.QueryExpr())
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

func (db *repoSaUser) GetList(ctx context.Context, queryparam models.ParamList) (result []*models.SaUser, err error) {
	var (
		pageNum  = 0
		pageSize = setting.FileConfigSetting.App.PageSize
		sWhere   = ""
		logger   = logging.Logger{}
	)
	// pagination
	if queryparam.Page > 0 {
		pageNum = (queryparam.Page - 1) * queryparam.PerPage
	}
	if queryparam.PerPage > 0 {
		pageSize = queryparam.PerPage
	}
	//end pagination

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
		query := db.Conn.Where(sWhere).Offset(pageNum).Limit(pageSize).Find(&result)
		logger.Query(query.QueryExpr()) //cath to log query string
		err = query.Error
	} else {
		query := db.Conn.Where(sWhere).Find(&result)
		logger.Query(query.QueryExpr()) //cath to log query string
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

func (db *repoSaUser) CreateSaUser(ctx context.Context, userData *models.SaUser) (err error) {
	var (
		logger = logging.Logger{}
	)
	query := db.Conn.Create(userData)
	logger.Query(query.QueryExpr()) //cath to log query string
	err = query.Error
	// err = db.Conn.Create(userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaUser) UpdateSaUser(ctx context.Context, userData *models.SaUser) (err error) {
	var (
		logger = logging.Logger{}
	)
	query := db.Conn.Save(userData)
	logger.Query(query.QueryExpr()) //cath to log query string
	err = query.Error
	// err = db.Conn.Save(userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaUser) DeleteSaUser(ctx context.Context, userID int16) (err error) {
	var (
		logger = logging.Logger{}
	)
	userData := &models.SaUser{}
	userData.UserID = userID

	query := db.Conn.Delete(&userData)
	logger.Query(query.QueryExpr()) //cath to log query string
	err = query.Error
	// err = db.Conn.Where("user_id = ?", userID).Delete(&userData).Error
	// err = db.Conn.Delete(&userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaUser) CountUserList(ctx context.Context, queryparam models.ParamList) (result int, err error) {
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
		}
	}
	// end where

	query := db.Conn.Model(&models.SaUser{}).Where(sWhere).Count(&result)
	logger.Query(query.QueryExpr()) //cath to log query string
	err = query.Error

	return result, err
}
