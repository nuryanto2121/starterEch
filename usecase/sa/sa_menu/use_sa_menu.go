package usesamenu

import (
	"context"
	isamenu "property/framework/interface/sa/sa_menu"
	sa_models "property/framework/models/sa"
	util "property/framework/pkg/utils"
	"time"

	"github.com/mitchellh/mapstructure"
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
func (u *useSaMenu) GetList(ctx context.Context, LevelNo int, ParentMenuID int) ([]*sa_models.SaMenu, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()

	var (
		result = []*sa_models.SaMenu{}
		err    error
	)

	result, err = u.menuSaMenu.GetList(ctx, LevelNo, ParentMenuID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// func (u *useSaMenu) GetList2(ctx context.Context, queryparam models.ParamList) (models.ResponseModelList, error) {
// 	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
// 	defer cancel()
// 	var (
// 		result = models.ResponseModelList{}
// 		tmenu  = sa_models.SaMenu{}
// 		err    error
// 	)

// 	/*membuat Where like dari struct*/
// 	if queryparam.Search != "" {
// 		value := reflect.ValueOf(tmenu)
// 		types := reflect.TypeOf(&tmenu)
// 		queryparam.Search = util.GetWhereLikeStruct(value, types, queryparam.Search, "")
// 	}
// 	result.Data, err = u.menuSaMenu.GetList(ctx, queryparam)
// 	if err != nil {
// 		return result, err
// 	}

// 	result.Total, err = u.menuSaMenu.CountMenuList(ctx, queryparam)
// 	if err != nil {
// 		return result, err
// 	}

// 	// d := float64(result.Total) / float64(queryparam.PerPage)
// 	result.LastPage = int(math.Ceil(float64(result.Total) / float64(queryparam.PerPage)))
// 	result.Page = queryparam.Page

// 	return result, nil
// }

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

func (u *useSaMenu) UpdateSaMenu(ctx context.Context, menuID int, data interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()
	var (
		err  error
		form = sa_models.EditMenuForm{}
	)

	err = mapstructure.Decode(data, &form)
	if err != nil {
		return err
		// return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)

	}
	form.UpdatedAt = util.GetTimeNow()
	err = u.menuSaMenu.UpdateSaMenu(ctx, menuID, form)
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
