package reposauser

import (
	"context"
	isauser "property/framework/interface/sa/sa_user"
	models "property/framework/models"

	"github.com/jinzhu/gorm"
)

type repoSaUser struct {
	Conn *gorm.DB
}

// NewRepoSaUser :
func NewRepoSaUser(Conn *gorm.DB) isauser.Repository {
	return &repoSaUser{Conn}
}

func (db *repoSaUser) GetBySaUser(ctx context.Context, userID int16) (result models.SaUser, err error) {

	a := models.SaUser{}
	// var c *models.SaUser
	// var b []*models.SaUser
	err = db.Conn.Where("user_id = ?", userID).First(&a).Error
	// member := models.SaUser{}
	// err = db.Conn.Model(&models.SaUser{}).Where("user_id = ?", userID).First(&result).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return a, err
	}
	// fmt.Printf("%d", len(b))
	// if len(b) > 0 {
	// 	c = b[0]
	// }

	return a, err
	// return a, nil
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
