package usesarole

import (
	"context"
	"encoding/json"
	"log"
	"math"
	isarole "property/framework/interface/sa/sa_role"

	isarolemenu "property/framework/interface/sa/sa_role_menu"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	util "property/framework/pkg/utils"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
	uuid "github.com/satori/go.uuid"
)

type useSaRole struct {
	repoSaRole     isarole.Repository
	repoSaRoleMenu isarolemenu.Repository
	contexTimeOut  time.Duration
}

// NewUseSaRole :
func NewUseSaRole(a isarole.Repository, b isarolemenu.Repository, timeout time.Duration) isarole.UseCase {
	return &useSaRole{
		repoSaRole:     a,
		repoSaRoleMenu: b,
		contexTimeOut:  timeout,
	}
}

func (u *useSaRole) GetJsonMenuAccess(ctx context.Context, roleID uuid.UUID) (result []map[string]interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()

	_result, err := u.repoSaRole.GetJsonMenuAccess(ctx, roleID)
	if err != nil {
		return result, err
	}
	json.Unmarshal([]byte(_result), &result)
	log.Printf("Unmarshaled: %v", result)
	return result, nil
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

func (u *useSaRole) CreateSaRole(ctx context.Context, roleData *sa_models.SaRole, menuAccess *[]sa_models.MenuAccessLevel1) error {
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

	// insert role_menu level 1
	for _, dataLevel1 := range *menuAccess {
		var roleMenu = sa_models.SaRoleMenu{}
		if dataLevel1.IsRead == true || dataLevel1.IsWrite == true {
			roleMenu.IsRead = dataLevel1.IsRead
			roleMenu.IsWrite = dataLevel1.IsWrite
			roleMenu.MenuID = dataLevel1.MenuID
			roleMenu.RoleID = roleData.RoleID
			roleMenu.CreatedBy = roleData.CreatedBy
			roleMenu.UpdatedBy = roleData.CreatedBy
			err = u.repoSaRoleMenu.CreateSaRoleMenu(ctx, &roleMenu)
			if err != nil {
				return err
			}
		}

		// insert from level 2
		for _, dataLevel2 := range dataLevel1.Level2 {
			var roleMenu = sa_models.SaRoleMenu{}
			if dataLevel2.IsRead == true || dataLevel2.IsWrite == true {
				roleMenu.IsRead = dataLevel2.IsRead
				roleMenu.IsWrite = dataLevel2.IsWrite
				roleMenu.MenuID = dataLevel2.MenuID
				roleMenu.RoleID = roleData.RoleID
				roleMenu.CreatedBy = roleData.CreatedBy
				roleMenu.UpdatedBy = roleData.CreatedBy
				err = u.repoSaRoleMenu.CreateSaRoleMenu(ctx, &roleMenu)
				if err != nil {
					return err
				}
			}
			// insert from level 2
			for _, dataLevel3 := range dataLevel2.Level3 {
				var roleMenu = sa_models.SaRoleMenu{}
				if dataLevel3.IsRead == true || dataLevel3.IsWrite == true {
					roleMenu.IsRead = dataLevel3.IsRead
					roleMenu.IsWrite = dataLevel3.IsWrite
					roleMenu.MenuID = dataLevel3.MenuID
					roleMenu.RoleID = roleData.RoleID
					roleMenu.CreatedBy = roleData.CreatedBy
					roleMenu.UpdatedBy = roleData.CreatedBy
					err = u.repoSaRoleMenu.CreateSaRoleMenu(ctx, &roleMenu)
					if err != nil {
						return err
					}
				}
			}
		}

	}

	return nil
}

func (u *useSaRole) UpdateSaRole(ctx context.Context, roleID uuid.UUID, data interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()
	var (
		err        error
		form       = sa_models.EditRoleForm{}
		menuAccess []sa_models.MenuAccessLevel1
	)

	// mapping to struct model saSuser
	err = mapstructure.Decode(data, &form)
	if err != nil {
		return err
		// return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)

	}
	err = mapstructure.Decode(form.MenuAccess, &menuAccess)
	if err != nil {
		return err
		// return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)

	}

	form.UpdatedAt = util.GetTimeNow()
	err = u.repoSaRole.UpdateSaRole(ctx, roleID, form)
	if err != nil {
		return err
	}

	err = u.repoSaRoleMenu.DeleteSaRoleMenu(ctx, roleID)
	if err != nil {
		return err
	}

	// insert role_menu level 1
	for _, dataLevel1 := range menuAccess {
		var roleMenu = sa_models.SaRoleMenu{}
		if dataLevel1.IsRead == true || dataLevel1.IsWrite == true {
			roleMenu.IsRead = dataLevel1.IsRead
			roleMenu.IsWrite = dataLevel1.IsWrite
			roleMenu.MenuID = dataLevel1.MenuID
			roleMenu.RoleID = roleID
			roleMenu.CreatedBy = form.UpdatedBy
			roleMenu.UpdatedBy = form.UpdatedBy
			err = u.repoSaRoleMenu.CreateSaRoleMenu(ctx, &roleMenu)
			if err != nil {
				return err
			}
		}

		// insert from level 2
		for _, dataLevel2 := range dataLevel1.Level2 {
			var roleMenu = sa_models.SaRoleMenu{}
			if dataLevel2.IsRead == true || dataLevel2.IsWrite == true {
				roleMenu.IsRead = dataLevel2.IsRead
				roleMenu.IsWrite = dataLevel2.IsWrite
				roleMenu.MenuID = dataLevel2.MenuID
				roleMenu.RoleID = roleID
				roleMenu.CreatedBy = form.UpdatedBy
				roleMenu.UpdatedBy = form.UpdatedBy
				err = u.repoSaRoleMenu.CreateSaRoleMenu(ctx, &roleMenu)
				if err != nil {
					return err
				}
			}
			// insert from level 2
			for _, dataLevel3 := range dataLevel2.Level3 {
				var roleMenu = sa_models.SaRoleMenu{}
				if dataLevel3.IsRead == true || dataLevel3.IsWrite == true {
					roleMenu.IsRead = dataLevel3.IsRead
					roleMenu.IsWrite = dataLevel3.IsWrite
					roleMenu.MenuID = dataLevel3.MenuID
					roleMenu.RoleID = roleID
					roleMenu.CreatedBy = form.UpdatedBy
					roleMenu.UpdatedBy = form.UpdatedBy
					err = u.repoSaRoleMenu.CreateSaRoleMenu(ctx, &roleMenu)
					if err != nil {
						return err
					}
				}
			}
		}

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
