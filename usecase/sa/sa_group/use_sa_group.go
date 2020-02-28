package usesagroup

import (
	"context"
	"math"
	isagroup "property/framework/interface/sa/sa_group"
	"property/framework/models"
	util "property/framework/pkg/utils"
	"reflect"
	"time"
)

type useSaGroup struct {
	repoSaGroup   isagroup.Repository
	contexTimeOut time.Duration
}

// NewUseSaGroup :
func NewUseSaGroup(a isagroup.Repository, timeout time.Duration) isagroup.UseCase {
	return &useSaGroup{
		repoSaGroup:   a,
		contexTimeOut: timeout,
	}
}

func (u *useSaGroup) GetBySaGroup(ctx context.Context, groupID int16) (models.SaGroup, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()

	var (
		result = models.SaGroup{}
		err    error
	)

	result, err = u.repoSaGroup.GetBySaGroup(ctx, groupID)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (u *useSaGroup) GetList(ctx context.Context, queryparam models.ParamList) (models.ResponseModelList, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()
	var (
		result = models.ResponseModelList{}
		tgroup = models.SaGroup{}
		err    error
	)

	/*membuat Where like dari struct*/
	if queryparam.Search != "" {
		value := reflect.ValueOf(tgroup)
		types := reflect.TypeOf(&tgroup)
		queryparam.Search = util.GetWhereLikeStruct(value, types, queryparam.Search, "")
	}
	result.Data, err = u.repoSaGroup.GetList(ctx, queryparam)
	if err != nil {
		return result, err
	}

	result.Total, err = u.repoSaGroup.CountGroupList(ctx, queryparam)
	if err != nil {
		return result, err
	}

	// d := float64(result.Total) / float64(queryparam.PerPage)
	result.LastPage = int(math.Ceil(float64(result.Total) / float64(queryparam.PerPage)))
	result.Page = queryparam.Page

	return result, nil
}

func (u *useSaGroup) CreateSaGroup(ctx context.Context, groupData *models.SaGroup) error {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()
	var (
		err error
	)

	groupData.UpdatedBy = groupData.CreatedBy
	groupData.CreatedAt = util.GetTimeNow()
	groupData.UpdatedAt = util.GetTimeNow()
	err = u.repoSaGroup.CreateSaGroup(ctx, groupData)
	if err != nil {
		return err
	}
	return nil
}

func (u *useSaGroup) UpdateSaGroup(ctx context.Context, groupData *models.SaGroup) error {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()
	var (
		err error
	)
	groupData.UpdatedAt = util.GetTimeNow()
	err = u.repoSaGroup.UpdateSaGroup(ctx, groupData)
	if err != nil {
		return err
	}
	return nil

}

func (u *useSaGroup) DeleteSaGroup(ctx context.Context, groupID int16) error {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()

	err := u.repoSaGroup.DeleteSaGroup(ctx, groupID)
	if err != nil {
		return err
	}
	return nil
}
