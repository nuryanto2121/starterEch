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
