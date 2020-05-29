package reposauser

import (
	"context"
	"fmt"
	isauser "property/framework/interface/sa/sa_user"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	"property/framework/pkg/logging"
	"property/framework/pkg/setting"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type repoSaUser struct {
	Conn *gorm.DB
}

// NewRepoSaUser :
func NewRepoSaUser(Conn *gorm.DB) isauser.Repository {
	return &repoSaUser{Conn}
}

func (db *repoSaUser) GetBySaUser(ctx context.Context, userID uuid.UUID) (result sa_models.SaUser, err error) {
	var (
		a      = sa_models.SaUser{}
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

func (db *repoSaUser) GetByEmailSaUser(ctx context.Context, email string) (result sa_models.SaUser, err error) {
	var (
		a      = sa_models.SaUser{}
		logger = logging.Logger{}
	)
	query := db.Conn.Where("email_addr = ?", email).Or("user_name = ?", email).First(&a)
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

func (db *repoSaUser) GetList(ctx context.Context, queryparam models.ParamList) (result []*sa_models.SaUser, err error) {
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

func (db *repoSaUser) CreateSaUser(ctx context.Context, userData *sa_models.SaUser) (err error) {
	var (
		logger = logging.Logger{}
	)
	query := db.Conn.Create(userData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Create(userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaUser) UpdateSaUser(ctx context.Context, userData *sa_models.SaUser) (err error) {
	var (
		logger = logging.Logger{}
	)
	query := db.Conn.Save(userData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Save(userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaUser) DeleteSaUser(ctx context.Context, userID uuid.UUID) (err error) {
	var (
		logger = logging.Logger{}
	)
	userData := &sa_models.SaUser{}
	userData.UserID = userID

	query := db.Conn.Delete(&userData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
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
		} else {
			sWhere += queryparam.Search
		}
	}
	// end where

	query := db.Conn.Model(&sa_models.SaUser{}).Where(sWhere).Count(&result)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error

	return result, err
}
