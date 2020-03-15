package reposaclient

import (
	"context"
	"fmt"
	isclient "property/framework/interface/sa/sa_client"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	"property/framework/pkg/logging"
	"property/framework/pkg/setting"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// repoSaClient :
type repoSaClient struct {
	Conn *gorm.DB
}

// NewRepoSaClient :
func NewRepoSaClient(Conn *gorm.DB) isclient.Repository {
	return &repoSaClient{Conn}
}

func (db *repoSaClient) GetBySaClient(ctx context.Context, clientID uuid.UUID) (result sa_models.SaClient, err error) {
	var (
		a      = sa_models.SaClient{}
		logger = logging.Logger{}
	)
	query := db.Conn.Where("client_id = ?", clientID).First(&a)
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

func (db *repoSaClient) GetList(ctx context.Context, queryparam models.ParamList) (result []*sa_models.SaClient, err error) {
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

func (db *repoSaClient) CreateSaClient(ctx context.Context, clientData *sa_models.SaClient) (err error) {
	var (
		logger = logging.Logger{}
	)
	query := db.Conn.Create(clientData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Create(clientData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaClient) UpdateSaClient(ctx context.Context, clientData *sa_models.SaClient) (err error) {
	var (
		logger = logging.Logger{}
	)
	query := db.Conn.Save(clientData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Save(clientData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaClient) DeleteSaClient(ctx context.Context, clientID uuid.UUID) (err error) {
	var (
		logger = logging.Logger{}
	)
	clientData := &sa_models.SaClient{}
	clientData.ClientID = clientID

	query := db.Conn.Delete(&clientData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error

	if err != nil {
		return err
	}
	return nil
}

// CountClientList :
func (db *repoSaClient) CountClientList(ctx context.Context, queryparam models.ParamList) (result int, err error) {
	return result, nil
}
