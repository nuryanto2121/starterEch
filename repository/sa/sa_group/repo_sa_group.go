package reposagroup

import (
	"context"
	"fmt"
	isagroup "property/framework/interface/sa/sa_group"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	"property/framework/pkg/logging"
	"property/framework/pkg/setting"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type repoSaGroup struct {
	Conn *gorm.DB
}

// NewRepoSaGroup :
func NewRepoSaGroup(Conn *gorm.DB) isagroup.Repository {
	return &repoSaGroup{Conn}
}

func (db *repoSaGroup) GetBySaGroup(ctx context.Context, groupID uuid.UUID) (sa_models.SaGroup, error) {
	var (
		dataGroup = sa_models.SaGroup{}
		logger    = logging.Logger{}
		err       error
	)
	query := db.Conn.Where("group_id = ?", groupID).First(&dataGroup)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr()))
	err = query.Error

	if err != nil {
		//
		if err == gorm.ErrRecordNotFound {
			return dataGroup, models.ErrNotFound
		}
		return dataGroup, err
	}

	return dataGroup, err
}

func (db *repoSaGroup) GetList(ctx context.Context, queryparam models.ParamList) ([]*sa_models.SaGroup, error) {
	var (
		pageNum  = 0
		pageSize = setting.FileConfigSetting.App.PageSize
		sWhere   = ""
		logger   = logging.Logger{}
		orderBy  = "created_at desc"
		err      error
		result   = []*sa_models.SaGroup{}
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

func (db *repoSaGroup) CreateSaGroup(ctx context.Context, groupData *sa_models.SaGroup) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Create(groupData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Create(userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaGroup) UpdateSaGroup(ctx context.Context, groupData *sa_models.SaGroup) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Save(groupData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Save(userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaGroup) DeleteSaGroup(ctx context.Context, groupID uuid.UUID) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	userData := &sa_models.SaGroup{}
	userData.GroupID = groupID

	query := db.Conn.Delete(&userData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error

	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaGroup) CountGroupList(ctx context.Context, queryparam models.ParamList) (int, error) {
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
