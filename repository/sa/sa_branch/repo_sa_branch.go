package reposabranch

import (
	"context"
	"fmt"
	isabranch "property/framework/interface/sa/sa_branch"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	"property/framework/pkg/logging"
	"property/framework/pkg/setting"

	"github.com/jinzhu/gorm"
)

type repoSaBranch struct {
	Conn *gorm.DB
}

// NewRepoSaBranch :
func NewRepoSaBranch(Conn *gorm.DB) isabranch.Repository {
	return &repoSaBranch{Conn}
}

func (db *repoSaBranch) GetBySaBranch(ctx context.Context, branchID int16) (result sa_models.SaBranch, err error) {
	var (
		a      = sa_models.SaBranch{}
		logger = logging.Logger{}
	)
	query := db.Conn.Where("branch_id = ?", branchID).First(&a)
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

func (db *repoSaBranch) GetList(ctx context.Context, queryparam models.ParamList) (result []*sa_models.SaBranch, err error) {
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

func (db *repoSaBranch) CreateSaBranch(ctx context.Context, branchData *sa_models.SaBranch) (err error) {
	var (
		logger = logging.Logger{}
	)
	query := db.Conn.Create(branchData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Create(branchData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaBranch) UpdateSaBranch(ctx context.Context, branchData *sa_models.SaBranch) (err error) {
	var (
		logger = logging.Logger{}
	)
	query := db.Conn.Save(branchData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Save(branchData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaBranch) DeleteSaBranch(ctx context.Context, branchID int16) (err error) {
	var (
		logger = logging.Logger{}
	)
	branchData := &sa_models.SaBranch{}
	branchData.BranchID = branchID

	query := db.Conn.Delete(&branchData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Where("branch_id = ?", branchID).Delete(&branchData).Error
	// err = db.Conn.Delete(&branchData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaBranch) CountBranchList(ctx context.Context, queryparam models.ParamList) (result int, err error) {
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

	query := db.Conn.Model(&sa_models.SaBranch{}).Where(sWhere).Count(&result)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error

	return result, err
}
