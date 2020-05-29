package usesauser

import (
	"context"
	"math"
	isauser "property/framework/interface/sa/sa_user"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	util "property/framework/pkg/utils"
	"reflect"
	"time"

	uuid "github.com/satori/go.uuid"
)

type useSaUser struct {
	repoSaUser     isauser.Repository
	contextTimeOut time.Duration
}

// NewUseSaUser :
func NewUseSaUser(a isauser.Repository, timeout time.Duration) isauser.Usercase {
	return &useSaUser{
		repoSaUser:     a,
		contextTimeOut: timeout,
	}
}

func (u *useSaUser) GetBySaUser(ctx context.Context, userID uuid.UUID) (result sa_models.SaUser, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	a := sa_models.SaUser{}
	result, err = u.repoSaUser.GetBySaUser(ctx, userID)
	if err != nil {
		return a, err
	}
	return result, nil
}

func (u *useSaUser) GetByEmailSaUser(ctx context.Context, email string) (result sa_models.SaUser, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	a := sa_models.SaUser{}
	result, err = u.repoSaUser.GetByEmailSaUser(ctx, email)
	if err != nil {
		return a, err
	}
	return result, nil
}

func (u *useSaUser) GetList(ctx context.Context, queryparam models.ParamList) (result models.ResponseModelList, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	/*membuat Where like dari struct*/
	tuser := sa_models.SaUser{}
	if queryparam.Search != "" {
		value := reflect.ValueOf(tuser)
		types := reflect.TypeOf(&tuser)
		queryparam.Search = util.GetWhereLikeStruct(value, types, queryparam.Search, "") // fmt.Sprintf("user_name LIKE '%s' OR email_addr LIKE '%s' OR handphone_no LIKE '%s'", search, search, search)
	}
	result.Data, err = u.repoSaUser.GetList(ctx, queryparam)
	if err != nil {
		return result, err
	}

	result.Total, err = u.repoSaUser.CountUserList(ctx, queryparam)
	if err != nil {
		return result, err
	}
	d := float64(result.Total) / float64(queryparam.PerPage)
	result.LastPage = int(math.Ceil(d))
	result.Page = queryparam.Page

	return result, nil
}

func (u *useSaUser) CreateSaUser(ctx context.Context, userData *sa_models.SaUser) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	userData.Passwd, _ = util.HashAndSalt(util.GetPassword(userData.Passwd))
	userData.UpdatedBy = userData.CreatedBy
	userData.CreatedAt = util.GetTimeNow()
	userData.UpdatedAt = util.GetTimeNow()
	err = u.repoSaUser.CreateSaUser(ctx, userData)
	if err != nil {
		return err
	}
	return nil
}

func (u *useSaUser) UpdateSaUser(ctx context.Context, userData *sa_models.SaUser) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	userData.UpdatedAt = util.GetTimeNow()
	err = u.repoSaUser.UpdateSaUser(ctx, userData)
	if err != nil {
		return err
	}

	return nil
}

func (u *useSaUser) DeleteSaUser(ctx context.Context, userID uuid.UUID) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	err = u.repoSaUser.DeleteSaUser(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
