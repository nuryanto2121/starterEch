package reposauser

import (
	"context"
	"github.com/jinzhu/gorm"
	isauser "property/framework/interface/sa/sa_user"
	models "property/framework/models/sa"
)

type repoSaUser struct {
	Conn *gorm.DB
}

func NewRepoSaUser(Conn *gorm.DB) isauser.Repository {
	return &repoSaUser{Conn}
}

func (db *repoSaUser) GetBySaUser(ctx context.Context, userID int16) (result *models.SaUser, err error) {

	err = db.Conn.Where("user_id = ?", userID).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (db *repoSaUser) GetAllSaUser(ctx context.Context) (result []*models.SaUser, err error) {

	err = db.Conn.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (db *repoSaUser) CreateSaUser(ctx context.Context, userData *models.SaUser) (err error) {

	err = db.Conn.Create(userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaUser) UpdateSaUser(ctx context.Context, userData *models.SaUser) (err error) {
	err = db.Conn.Save(userData).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoSaUser) DeleteSaUser(ctx context.Context, userID int16) (err error) {
	userData := &models.SaUser{}
	userData.UserID = userID

	err = db.Conn.Where("user_id = ?", userID).Delete(&userData).Error
	if err != nil {
		return err
	}
	return nil
}
