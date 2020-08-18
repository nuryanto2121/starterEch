package reposarole

import (
	"context"
	"fmt"
	isarolemenu "property/framework/interface/sa/sa_role_menu"
	sa_models "property/framework/models/sa"
	"property/framework/pkg/logging"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type repoSaRoleMenu struct {
	Conn *gorm.DB
}

// NewRepoSaRoleMenu :
func NewRepoSaRoleMenu(Conn *gorm.DB) isarolemenu.Repository {
	return &repoSaRoleMenu{Conn}
}

func (db *repoSaRoleMenu) GetMenuRole(ctx context.Context, RoleID uuid.UUID) (result []*sa_models.MenuRole, err error) {
	var (
		logger = logging.Logger{}
	)

	query := db.Conn.Table("sa_menu a").Select(`a.menu_id,a.title,a.menu_url,a.parent_menu_id,a.icon_class,a.order_seq,a.level,
	(row_number()OVER(PARTITION BY a.parent_menu_id order by a.order_seq)::varchar||a.level::"varchar"||a.order_seq::varchar)::integer as ipath,
	b.is_read ,b.is_write`).Joins(`inner join sa_role_menu b on a.menu_id = b.menu_id`).Where(`b.role_id =?`, RoleID).Order(`ipath,a.parent_menu_id`).Find(&result)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error
	if err != nil {
		return result, err
	}

	return result, nil

}
func (db *repoSaRoleMenu) CreateSaRoleMenu(ctx context.Context, roleData *sa_models.SaRoleMenu) error {
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

func (db *repoSaRoleMenu) UpdateSaRoleMenu(ctx context.Context, roleData *sa_models.SaRoleMenu) error {
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

func (db *repoSaRoleMenu) DeleteSaRoleMenu(ctx context.Context, roleID uuid.UUID) error {
	var (
		logger = logging.Logger{}
		err    error
	)

	query := db.Conn.Exec("Delete From sa_role_menu WHERE role_id = ?", roleID)
	logger.Query(fmt.Sprintf("%v", query.QueryExpr())) //cath to log query string
	err = query.Error

	if err != nil {
		return err
	}
	return nil
}
