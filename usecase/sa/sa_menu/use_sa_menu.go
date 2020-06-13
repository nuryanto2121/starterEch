package usesamenu

import (
	"context"
	"math"
	isamenu "property/framework/interface/sa/sa_menu"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	util "property/framework/pkg/utils"
	"reflect"
	"time"
)

type useSaMenu struct {
	menuSaMenu    isamenu.Repository
	contexTimeOut time.Duration
}

// NewUseSaMenu :
func NewUseSaMenu(a isamenu.Repository, timeout time.Duration) isamenu.UseCase {
	return &useSaMenu{
		menuSaMenu:    a,
		contexTimeOut: timeout,
	}
}

func (u *useSaMenu) GetBySaMenu(ctx context.Context, menuID int) (sa_models.SaMenu, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()

	var (
		result = sa_models.SaMenu{}
		err    error
	)

	result, err = u.menuSaMenu.GetBySaMenu(ctx, menuID)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (u *useSaMenu) GetList(ctx context.Context, queryparam models.ParamList) (models.ResponseModelList, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()
	var (
		result = models.ResponseModelList{}
		tmenu  = sa_models.SaMenu{}
		err    error
	)

	/*membuat Where like dari struct*/
	if queryparam.Search != "" {
		value := reflect.ValueOf(tmenu)
		types := reflect.TypeOf(&tmenu)
		queryparam.Search = util.GetWhereLikeStruct(value, types, queryparam.Search, "")
	}
	result.Data, err = u.menuSaMenu.GetList(ctx, queryparam)
	if err != nil {
		return result, err
	}

	result.Total, err = u.menuSaMenu.CountMenuList(ctx, queryparam)
	if err != nil {
		return result, err
	}

	// d := float64(result.Total) / float64(queryparam.PerPage)
	result.LastPage = int(math.Ceil(float64(result.Total) / float64(queryparam.PerPage)))
	result.Page = queryparam.Page

	return result, nil
}

func (u *useSaMenu) CreateSaMenu(ctx context.Context, menuData *sa_models.SaMenu) error {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()
	var (
		err error
	)

	menuData.UpdatedBy = menuData.CreatedBy
	menuData.CreatedAt = util.GetTimeNow()
	menuData.UpdatedAt = util.GetTimeNow()
	err = u.menuSaMenu.CreateSaMenu(ctx, menuData)
	if err != nil {
		return err
	}
	return nil
}

func (u *useSaMenu) UpdateSaMenu(ctx context.Context, menuData *sa_models.SaMenu) error {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()
	var (
		err error
	)
	menuData.UpdatedAt = util.GetTimeNow()
	err = u.menuSaMenu.UpdateSaMenu(ctx, menuData)
	if err != nil {
		return err
	}
	return nil

}

func (u *useSaMenu) DeleteSaMenu(ctx context.Context, menuID int) error {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()

	err := u.menuSaMenu.DeleteSaMenu(ctx, menuID)
	if err != nil {
		return err
	}
	return nil
}
