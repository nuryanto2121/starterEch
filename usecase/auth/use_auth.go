package useauth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	iauth "property/framework/interface/auth"
	isaclient "property/framework/interface/sa/sa_client"
	isafileupload "property/framework/interface/sa/sa_file_upload"
	isauser "property/framework/interface/sa/sa_user"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	"property/framework/pkg/setting"
	util "property/framework/pkg/utils"
	"property/framework/usecase/usemail"
	"time"

	"github.com/mitchellh/mapstructure"
)

// useAuth :
type useAuth struct {
	repoSaUser isauser.Repository
	// repoSaClient   isaclient.Repository
	useSaClient     isaclient.Usecase
	useSaFileUpload isafileupload.UseCase
	contextTimeOut  time.Duration
}

func NewUserAuth(a isauser.Repository, b isaclient.Usecase, c isafileupload.UseCase, timeout time.Duration) iauth.Usecase {
	return &useAuth{
		repoSaUser:      a,
		useSaClient:     b,
		useSaFileUpload: c,
		contextTimeOut:  timeout,
	}
}

func (u *useAuth) ForgotPassword(ctx context.Context, dataForgot *models.ForgotForm) (output models.Output) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	dataUser, err := u.repoSaUser.GetByEmailSaUser(ctx, dataForgot.EmailAddr)
	if err != nil {
		return util.GoutputErrCode(http.StatusUnauthorized, "Your User not valid.") //appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}
	fmt.Printf("%v", dataUser)

	// Gen Token Email
	TokenEmail := util.GetEmailToken(dataForgot.EmailAddr)

	urlButton := setting.FileConfigSetting.App.UrlForgotPassword + "/" + TokenEmail
	mailService := usemail.Forgot{
		Email:      dataForgot.EmailAddr,
		Name:       dataUser.Name,
		ButtonLink: urlButton,
	}

	err = mailService.SendForgot()
	if err != nil {
		return util.GoutputErr(err)
	}

	// logger.Info(util.Stringify(dataUser))
	// if dataUser.EmailAddr == "" {
	// 	return appE.ResponseError(util.GetStatusCode(err), "Your Email not Valid.", nil)
	// }
	// semd mail
	output.Err = nil
	return output
}

func (u *useAuth) Login(ctx context.Context, dataLogin *models.LoginForm) (output models.Output) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	DataUser, err := u.repoSaUser.GetByEmailSaUser(ctx, dataLogin.UserName)
	if err != nil {
		return util.GoutputErrCode(http.StatusUnauthorized, "Your User/Email not valid.") //appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
		//return util.GoutputErr(err) //return result, err
	}

	if !util.ComparePassword(DataUser.Passwd, util.GetPassword(dataLogin.Password)) {
		return util.GoutputErr(models.ErrInvalidLogin) //return result, models.ErrInvalidLogin
	}

	token, err := util.GenerateToken(DataUser.UserID.String(), DataUser.UserName)
	if err != nil {
		return util.GoutputErr(err) //return result, models.ErrInternalServerError
	}

	dataFile, err := u.useSaFileUpload.GetBySaFileUpload(ctx, DataUser.FileID)

	err = mapstructure.Decode(dataFile, &dataFile)
	if err != nil {
		return util.GoutputErr(err) //return result, models.ErrInternalServerError

	}

	restUser := map[string]interface{}{
		"user_id":      DataUser.UserID,
		"client_id":    DataUser.ClientID,
		"user_name":    DataUser.UserName,
		"level_no":     DataUser.LevelNo,
		"role_id":      DataUser.RoleID,
		"email_addr":   DataUser.EmailAddr,
		"handphone_no": DataUser.HandphoneNo,
		"company_id":   DataUser.CompanyID,
		"picture_url":  dataFile,
	}
	response := map[string]interface{}{
		"token":     token,
		"data_user": restUser,
	}

	return util.Goutput(response, 200)
}

func (u *useAuth) Register(ctx context.Context, dataRegister *sa_models.SaClient) (output models.Output) { //(result map[string]interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	err := u.useSaClient.RegisterClient(ctx, dataRegister)
	if err != nil {
		return util.GoutputErr(err)
	}

	return util.Goutput(nil, 200)
}

func (u *useAuth) ResetPassword(ctx context.Context, dataReset *models.ResetPasswd) (output models.Output) { //(result map[string]interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	if dataReset.Passwd != dataReset.ConfirmPasswd {
		return util.GoutputErr(errors.New("Password and Confirm Password not same."))
	}

	email, err := util.ParseEmailToken(dataReset.TokenEmail)
	if err != nil {
		email = dataReset.TokenEmail
	}

	dataUser, err := u.repoSaUser.GetByEmailSaUser(ctx, email)
	if err != nil {
		return util.GoutputErr(err)
	}

	dataUser.Passwd, _ = util.Hash(dataReset.Passwd)

	err = u.repoSaUser.UpdateSaUser(ctx, &dataUser)
	if err != nil {
		return util.GoutputErr(err)
	}

	return util.Goutput(nil, 200)
}

func (u *useAuth) Verify(ctx context.Context, dataReset *models.ResetPasswd) (output models.Output) { //(result map[string]interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	if dataReset.Passwd != dataReset.ConfirmPasswd {
		return util.GoutputErr(errors.New("Password and Confirm Password not same."))
	}

	email, err := util.ParseEmailToken(dataReset.TokenEmail)
	if err != nil {
		return util.GoutputErr(err)
	}

	dataUser, err := u.repoSaUser.GetByEmailSaUser(ctx, email)
	if err != nil {
		return util.GoutputErr(err)
	}

	dataUser.Passwd, _ = util.Hash(dataReset.Passwd)
	dataUser.UserStatus = 1

	err = u.repoSaUser.UpdateSaUser(ctx, &dataUser)
	if err != nil {
		return util.GoutputErr(err)
	}

	return util.Goutput(nil, 200)
}
