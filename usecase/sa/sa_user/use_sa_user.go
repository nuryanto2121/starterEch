package usesauser

import (
	"context"
	isauser "property/framework/interface/sa/sa_user"
	models "property/framework/models"
	util "property/framework/pkg/utils"
	"time"
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

func (u *useSaUser) GetBySaUser(ctx context.Context, userID int16) (result *models.SaUser, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	result, err = u.repoSaUser.GetBySaUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *useSaUser) GetAllSaUser(ctx context.Context) (result []*models.SaUser, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	result, err = u.repoSaUser.GetAllSaUser(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *useSaUser) CreateSaUser(ctx context.Context, userData *models.SaUser) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	userData.CreatedAt = util.GetTimeNow()
	userData.UpdatedAt = util.GetTimeNow()
	err = u.repoSaUser.CreateSaUser(ctx, userData)
	if err != nil {
		return err
	}
	return nil
}

func (u *useSaUser) UpdateSaUser(ctx context.Context, userData *models.SaUser) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	userData.UpdatedAt = util.GetTimeNow()
	err = u.repoSaUser.UpdateSaUser(ctx, userData)
	if err != nil {
		return err
	}

	return nil
}

func (u *useSaUser) DeleteSaUser(ctx context.Context, userID int16) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	err = u.repoSaUser.DeleteSaUser(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
