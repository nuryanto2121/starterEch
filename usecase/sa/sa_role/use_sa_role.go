package usesarole

import (
	"context"
	"math"
	isarole "property/framework/interface/sa/sa_role"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	util "property/framework/pkg/utils"
	"reflect"
	"time"

	uuid "github.com/satori/go.uuid"
)

type useSaRole struct {
	repoSaRole    isarole.Repository
	contexTimeOut time.Duration
}

// NewUseSaRole :
func NewUseSaRole(a isarole.Repository, timeout time.Duration) isarole.UseCase {
	return &useSaRole{
		repoSaRole:    a,
		contexTimeOut: timeout,
	}
}

func (u *useSaRole) GetBySaRole(ctx context.Context, roleID uuid.UUID) (sa_models.SaRole, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()

	var (
		result = sa_models.SaRole{}
		err    error
	)

	result, err = u.repoSaRole.GetBySaRole(ctx, roleID)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (u *useSaRole) GetList(ctx context.Context, queryparam models.ParamList) (models.ResponseModelList, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()
	var (
		result = models.ResponseModelList{}
		trole  = sa_models.SaRole{}
		err    error
	)

	/*membuat Where like dari struct*/
	if queryparam.Search != "" {
		value := reflect.ValueOf(trole)
		types := reflect.TypeOf(&trole)
		queryparam.Search = util.GetWhereLikeStruct(value, types, queryparam.Search, "")
	}
	result.Data, err = u.repoSaRole.GetList(ctx, queryparam)
	if err != nil {
		return result, err
	}

	result.Total, err = u.repoSaRole.CountRoleList(ctx, queryparam)
	if err != nil {
		return result, err
	}

	// d := float64(result.Total) / float64(queryparam.PerPage)
	result.LastPage = int(math.Ceil(float64(result.Total) / float64(queryparam.PerPage)))
	result.Page = queryparam.Page

	return result, nil
}

func (u *useSaRole) CreateSaRole(ctx context.Context, roleData *sa_models.SaRole) error {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()
	var (
		err error
	)

	roleData.UpdatedBy = roleData.CreatedBy
	roleData.CreatedAt = util.GetTimeNow()
	roleData.UpdatedAt = util.GetTimeNow()
	err = u.repoSaRole.CreateSaRole(ctx, roleData)
	if err != nil {
		return err
	}
	return nil
}

func (u *useSaRole) UpdateSaRole(ctx context.Context, roleData *sa_models.SaRole) error {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()
	var (
		err error
	)
	roleData.UpdatedAt = util.GetTimeNow()
	err = u.repoSaRole.UpdateSaRole(ctx, roleData)
	if err != nil {
		return err
	}
	return nil

}

func (u *useSaRole) DeleteSaRole(ctx context.Context, roleID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()

	err := u.repoSaRole.DeleteSaRole(ctx, roleID)
	if err != nil {
		return err
	}
	return nil
}
