package usesauser

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	isafileupload "property/framework/interface/sa/sa_file_upload"
	isauser "property/framework/interface/sa/sa_user"
	isauserbranch "property/framework/interface/sa/sa_user_branch"
	isausercompany "property/framework/interface/sa/sa_user_company"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	util "property/framework/pkg/utils"
	"reflect"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type useSaUser struct {
	repoSaUser        isauser.Repository
	repoSaUserCompany isausercompany.Repository
	repoSaUserBranch  isauserbranch.Repository
	repoSaFileUpload  isafileupload.UseCase
	contextTimeOut    time.Duration
}

// NewUseSaUser :
func NewUseSaUser(a isauser.Repository, b isausercompany.Repository, c isauserbranch.Repository, d isafileupload.UseCase, timeout time.Duration) isauser.Usecase {
	return &useSaUser{
		repoSaUser:        a,
		repoSaUserCompany: b,
		repoSaUserBranch:  c,
		repoSaFileUpload:  d,
		contextTimeOut:    timeout,
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

	dataFIle, _ := u.repoSaFileUpload.GetBySaFileUpload(ctx, result.FileID)
	result.DataFile.FileID = dataFIle.FileID
	result.DataFile.FileName = dataFIle.FileName
	result.DataFile.FilePath = dataFIle.FilePath
	result.DataFile.FileType = dataFIle.FileType
	result.Passwd = ""
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

func (u *useSaUser) GetJsonPermission(ctx context.Context, userID uuid.UUID, clientID uuid.UUID) (result []map[string]interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	_result, err := u.repoSaUser.GetJsonPermission(ctx, userID, clientID)
	if err != nil {
		return result, err
	}
	// var data []map[string]interface{}

	json.Unmarshal([]byte(_result), &result)
	log.Printf("Unmarshaled: %v", result)

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
		queryparam.Search = util.GetWhereLikeStruct(value, types, queryparam.Search, "user_name,name,email_addr,handphone_no") // fmt.Sprintf("user_name LIKE '%s' OR email_addr LIKE '%s' OR handphone_no LIKE '%s'", search, search, search)
	}

	if queryparam.InitSearch != "" {
		queryparam.InitSearch = strings.ReplaceAll(queryparam.InitSearch, "=", " iLIKE ")
	}
	dataList, err := u.repoSaUser.GetList(ctx, queryparam)
	if err != nil {
		return result, err
	}

	for _, data := range dataList {

		dt, _ := u.repoSaFileUpload.GetBySaFileUpload(ctx, data.FileID)
		data.DataFile.FileID = dt.FileID
		data.DataFile.FileName = dt.FileName
		data.DataFile.FilePath = dt.FilePath
		data.DataFile.FileType = dt.FileType

	}
	result.Data = dataList
	result.Total, err = u.repoSaUser.CountUserList(ctx, queryparam)
	if err != nil {
		return result, err
	}
	d := float64(result.Total) / float64(queryparam.PerPage)
	result.LastPage = int(math.Ceil(d))
	result.Page = queryparam.Page

	return result, nil
}

func (u *useSaUser) CreateSaUser(ctx context.Context, userData *sa_models.SaUser, dataPermission *[]models.Permission) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	userData.Passwd = "" //util.HashAndSalt(util.GetPassword(userData.Passwd))
	userData.UpdatedBy = userData.CreatedBy
	userData.CreatedAt = util.GetTimeNow()
	userData.UpdatedAt = util.GetTimeNow()
	err = u.repoSaUser.CreateSaUser(ctx, userData)
	if err != nil {
		return err
	}

	// insert sa user company
	for _, dataUserCompany := range *dataPermission {
		var userCompany = sa_models.SaUserCompany{}
		userCompany.CompanyID = dataUserCompany.CompanyID
		userCompany.UserID = userData.UserID
		userCompany.CreatedAt = util.GetTimeNow()
		userCompany.CreatedBy = userData.CreatedBy
		userCompany.UpdatedAt = util.GetTimeNow()
		userCompany.UpdatedBy = userData.CreatedBy
		err = u.repoSaUserCompany.CreateSaUserCompany(ctx, &userCompany)
		if err != nil {
			return err
		}

		// insert sa user branch
		for _, dataBranch := range dataUserCompany.DataBranch {
			var datauserBranch = sa_models.SaUserBranch{}
			datauserBranch.BranchID = dataBranch.BranchID
			datauserBranch.CompanyID = dataUserCompany.CompanyID
			datauserBranch.UserID = userData.UserID
			datauserBranch.CreatedAt = util.GetTimeNow()
			datauserBranch.CreatedBy = userData.CreatedBy
			datauserBranch.UpdatedAt = util.GetTimeNow()
			datauserBranch.UpdatedBy = userData.CreatedBy
			err = u.repoSaUserBranch.CreateSaUserBranch(ctx, &datauserBranch)
			if err != nil {
				return err
			}
		}

	}
	//Send email verity
	// Gen Token Email
	// TokenEmail := util.GetEmailToken(userData.EmailAddr)

	// urlButton := setting.FileConfigSetting.App.UrlVerityUser + "/" + TokenEmail

	// fmt.Printf(urlButton)
	// mailService := usemail.Verify{
	// 	Email:      userData.EmailAddr,
	// 	Name:       userData.Name,
	// 	ButtonLink: urlButton,
	// }

	// err = mailService.SendVerify()
	// if err != nil {
	// 	return err //util.GoutputErr(err)
	// }

	// u.repoSaUserCompany.CreateSaUserCompany(ctx,)

	return nil
}

func (u *useSaUser) UpdateSaUser(ctx context.Context, userData *sa_models.SaUser, dataPermission *[]models.Permission) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	dataUserBefor, ers := u.repoSaUser.GetBySaUser(ctx, userData.UserID)
	fmt.Printf("%v", ers)
	if ers != nil {
		return errors.New("Item is not found : id " + fmt.Sprintf("%s", userData.UserID))
	}
	// dataUserBefor, _ := u.repoSaUser.GetBySaUser(ctx, userData.UserID)
	UID, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")

	if UID != userData.FileID {

		if UID != dataUserBefor.FileID {
			err = u.repoSaFileUpload.DeleteSaFileUpload(ctx, dataUserBefor.FileID)
		}

	}
	userData.UpdatedAt = util.GetTimeNow()
	err = u.repoSaUser.UpdateSaUser(ctx, userData)
	if err != nil {
		return err
	}

	//delete sa user company by user
	err = u.repoSaUserCompany.DeleteSaUserCompany(ctx, userData.UserID)
	if err != nil {
		return err
	}
	// delete sa user branch by user
	err = u.repoSaUserBranch.DeleteSaUserBranch(ctx, userData.UserID)
	if err != nil {
		return err
	}

	// insert sa user company
	for _, dataUserCompany := range *dataPermission {
		var userCompany = sa_models.SaUserCompany{}
		userCompany.CompanyID = dataUserCompany.CompanyID
		userCompany.UserID = userData.UserID
		userCompany.CreatedAt = util.GetTimeNow()
		userCompany.CreatedBy = userData.UpdatedBy
		userCompany.UpdatedAt = util.GetTimeNow()
		userCompany.UpdatedBy = userData.UpdatedBy
		err = u.repoSaUserCompany.CreateSaUserCompany(ctx, &userCompany)
		if err != nil {
			return err
		}

		// insert sa user branch
		for _, dataBranch := range dataUserCompany.DataBranch {
			var datauserBranch = sa_models.SaUserBranch{}
			datauserBranch.BranchID = dataBranch.BranchID
			datauserBranch.CompanyID = dataUserCompany.CompanyID
			datauserBranch.UserID = userData.UserID
			datauserBranch.CreatedAt = util.GetTimeNow()
			datauserBranch.CreatedBy = userData.UpdatedBy
			datauserBranch.UpdatedAt = util.GetTimeNow()
			datauserBranch.UpdatedBy = userData.UpdatedBy
			err = u.repoSaUserBranch.CreateSaUserBranch(ctx, &datauserBranch)
			if err != nil {
				return err
			}
		}

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
