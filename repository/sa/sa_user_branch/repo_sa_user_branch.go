package reposauserbranch

import (
	"context"
	"fmt"
	isauserbranch "property/framework/interface/sa/sa_user_branch"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	"property/framework/pkg/logging"
	"property/framework/pkg/setting"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type repoSaUserBranch struct {
	Conn *gorm.DB
}

// NewRepoSaUserBranch :
func NewRepoSaUserBranch(Conn *gorm.DB) isauserbranch.Repository {
	return &repoSaUserBranch{Conn}
}

func (db *repoSaUserBranch) GetBySaUserBranch(ctx context.Context, userID uuid.UUID) (result sa_models.SaUserBranch, err error) {
	var (
		a      = sa_models.SaUserBranch{}
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

func (db *repoSaUserBranch) GetList(ctx context.Context, queryparam models.ParamList) (result []*sa_models.SaUserBranch, err error) {
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

func (db *repoSaUserBranch) CreateSaUserBranch(ctx context.Context, userbranchData *sa_models.SaUserBranch) (err error) {
	var (
		logger = logging.Logger{}
	)
	query := db.Conn.Create(userbranchData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Create(userbranchData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaUserBranch) UpdateSaUserBranch(ctx context.Context, userbranchData *sa_models.SaUserBranch) (err error) {
	var (
		logger = logging.Logger{}
	)
	query := db.Conn.Save(userbranchData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Save(userbranchData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaUserBranch) DeleteSaUserBranch(ctx context.Context, userID uuid.UUID) (err error) {
	var (
		logger = logging.Logger{}
	)
	userbranchData := &sa_models.SaUserBranch{}
	userbranchData.UserID = userID

	// query := db.Conn.Delete(&userbranchData
	query := db.Conn.Exec("Delete From sa_user_branch WHERE user_id = ?", userID)
	// query := db.Conn.Delete(&userbranchData, "user_id = ?", userID)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Where("userbranch_id = ?", userID).Delete(&userbranchData).Error
	// err = db.Conn.Delete(&userbranchData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaUserBranch) CountUserBranchList(ctx context.Context, queryparam models.ParamList) (result int, err error) {
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

	query := db.Conn.Model(&sa_models.SaUserBranch{}).Where(sWhere).Count(&result)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error

	return result, err
}
