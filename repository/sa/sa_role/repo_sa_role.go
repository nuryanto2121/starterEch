package reposarole

import (
	"context"
	"fmt"
	isarole "property/framework/interface/sa/sa_role"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	"property/framework/pkg/logging"
	"property/framework/pkg/setting"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type repoSaRole struct {
	Conn *gorm.DB
}

// NewRepoSaRole :
func NewRepoSaRole(Conn *gorm.DB) isarole.Repository {
	return &repoSaRole{Conn}
}

func (db *repoSaRole) GetBySaRole(ctx context.Context, roleID uuid.UUID) (sa_models.SaRole, error) {
	var (
		dataRole = sa_models.SaRole{}
		logger   = logging.Logger{}
		err      error
	)
	query := db.Conn.Where("role_id = ?", roleID).First(&dataRole)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr()))
	err = query.Error

	if err != nil {
		//
		if err == gorm.ErrRecordNotFound {
			return dataRole, models.ErrNotFound
		}
		return dataRole, err
	}

	return dataRole, err
}
func (db *repoSaRole) GetJsonMenuAccess(ctx context.Context, roleID uuid.UUID) (result string, err error) {
	var (
		logger = logging.Logger{}
	)
	// type Result struct {
	// 	get_permission_json_company_branch string
	// }
	// var dd Result

	// Scan
	type Result struct {
		Name string
	}

	var _result Result
	// var values ...interface{}
	query := db.Conn.Raw("SELECT get_menu_access_json as name From public.get_menu_access_json( ?)", roleID).Scan(&_result)

	// query := db.Conn.Raw().Scan(&dd)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr()))
	err = query.Error

	if err != nil {
		//
		if err == gorm.ErrRecordNotFound {
			return result, models.ErrNotFound
		}
		return result, err
	}
	result = _result.Name //dd.get_permission_json_company_branch
	return result, err
}
func (db *repoSaRole) GetList(ctx context.Context, queryparam models.ParamList) ([]*sa_models.SaRole, error) {
	var (
		pageNum  = 0
		pageSize = setting.FileConfigSetting.App.PageSize
		sWhere   = ""
		logger   = logging.Logger{}
		orderBy  = "created_at desc"
		err      error
		result   = []*sa_models.SaRole{}
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

func (db *repoSaRole) CreateSaRole(ctx context.Context, roleData *sa_models.SaRole) error {
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

func (db *repoSaRole) UpdateSaRole(ctx context.Context, roleID uuid.UUID, dataRole interface{}) error {
	var (
		logger = logging.Logger{}
		err    error
	)

	query := db.Conn.Model(sa_models.SaRole{}).Where("role_id = ?", roleID).Updates(dataRole)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	// err = db.Conn.Save(userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaRole) DeleteSaRole(ctx context.Context, roleID uuid.UUID) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	userData := &sa_models.SaRole{}
	userData.RoleID = roleID

	query := db.Conn.Delete(&userData)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error

	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaRole) CountRoleList(ctx context.Context, queryparam models.ParamList) (int, error) {
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

	query := db.Conn.Model(&sa_models.SaRole{}).Where(sWhere).Count(&result)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error

	return result, err
}
