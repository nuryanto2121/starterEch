package usersausercompany

import (
	"context"
	"math"
	isausercompany "property/framework/interface/sa/sa_user_company"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	util "property/framework/pkg/utils"
	"reflect"
	"time"

	uuid "github.com/satori/go.uuid"
)

type useSaUserCompany struct {
	repoSaUserCompany isausercompany.Repository
	contextTimeOut    time.Duration
}

// NewUseSaUserCompany :
func NewUseSaUserCompany(a isausercompany.Repository, timeout time.Duration) isausercompany.Usecase {
	return &useSaUserCompany{
		repoSaUserCompany: a,
		contextTimeOut:    timeout,
	}
}

func (u *useSaUserCompany) GetBySaUserCompany(ctx context.Context, userID uuid.UUID) (result sa_models.SaUserCompany, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	a := sa_models.SaUserCompany{}
	result, err = u.repoSaUserCompany.GetBySaUserCompany(ctx, userID)
	if err != nil {
		return a, err
	}
	return result, nil
}

func (u *useSaUserCompany) GetList(ctx context.Context, queryparam models.ParamList) (result models.ResponseModelList, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	/*membuat Where like dari struct*/
	tuser := sa_models.SaUserCompany{}
	if queryparam.Search != "" {
		value := reflect.ValueOf(tuser)
		types := reflect.TypeOf(&tuser)
		queryparam.Search = util.GetWhereLikeStruct(value, types, queryparam.Search, "") // fmt.Sprintf("user_name LIKE '%s' OR email_addr LIKE '%s' OR handphone_no LIKE '%s'", search, search, search)
	}
	result.Data, err = u.repoSaUserCompany.GetList(ctx, queryparam)
	if err != nil {
		return result, err
	}

	result.Total, err = u.repoSaUserCompany.CountUserCompanyList(ctx, queryparam)
	if err != nil {
		return result, err
	}
	d := float64(result.Total) / float64(queryparam.PerPage)
	result.LastPage = int(math.Ceil(d))
	result.Page = queryparam.Page

	return result, nil
}

func (u *useSaUserCompany) CreateSaUserCompany(ctx context.Context, userData *sa_models.SaUserCompany) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	userData.UpdatedBy = userData.CreatedBy
	userData.CreatedAt = util.GetTimeNow()
	userData.UpdatedAt = util.GetTimeNow()
	err = u.repoSaUserCompany.CreateSaUserCompany(ctx, userData)
	if err != nil {
		return err
	}
	return nil
}

func (u *useSaUserCompany) UpdateSaUserCompany(ctx context.Context, userData *sa_models.SaUserCompany) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	userData.UpdatedAt = util.GetTimeNow()
	err = u.repoSaUserCompany.UpdateSaUserCompany(ctx, userData)
	if err != nil {
		return err
	}

	return nil
}

func (u *useSaUserCompany) DeleteSaUserCompany(ctx context.Context, userID uuid.UUID) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	err = u.repoSaUserCompany.DeleteSaUserCompany(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
