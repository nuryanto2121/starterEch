package reposafileupload

import (
	"context"
	"fmt"
	isafileupload "property/framework/interface/sa/sa_file_upload"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	"property/framework/pkg/logging"
	"property/framework/pkg/setting"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type repoSaFileUpload struct {
	Conn *gorm.DB
}

//NewRepoSaFileUpload :
func NewRepoSaFileUpload(Conn *gorm.DB) isafileupload.Repository {
	return &repoSaFileUpload{Conn}
}

func (db *repoSaFileUpload) GetBySaFileUpload(ctx context.Context, fileID uuid.UUID) (sa_models.SaFileUpload, error) {
	var (
		dataFileUpload = sa_models.SaFileUpload{}
		logger         = logging.Logger{}
		err            error
	)
	query := db.Conn.Where("file_id = ?", fileID).First(&dataFileUpload)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr()))
	err = query.Error

	if err != nil {
		//
		if err == gorm.ErrRecordNotFound {
			return dataFileUpload, models.ErrNotFound
		}
		return dataFileUpload, err
	}

	return dataFileUpload, err
}

func (db *repoSaFileUpload) GetList(ctx context.Context, queryparam models.ParamList) ([]*sa_models.SaFileUpload, error) {
	var (
		pageNum  = 0
		pageSize = setting.FileConfigSetting.App.PageSize
		sWhere   = ""
		logger   = logging.Logger{}
		orderBy  = "created_at desc"
		err      error
		result   = []*sa_models.SaFileUpload{}
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

func (db *repoSaFileUpload) CreateSaFileUpload(ctx context.Context, fileData *sa_models.SaFileUpload) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Create(fileData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Create(userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaFileUpload) UpdateSaFileUpload(ctx context.Context, fileData *sa_models.SaFileUpload) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Save(fileData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Save(userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaFileUpload) DeleteSaFileUpload(ctx context.Context, fileID uuid.UUID) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	userData := &sa_models.SaFileUpload{}
	userData.FileID = fileID

	query := db.Conn.Delete(&userData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error

	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaFileUpload) CountFileUploadList(ctx context.Context, queryparam models.ParamList) (int, error) {
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
