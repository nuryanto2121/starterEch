package contauth

import (
	"context"
	"fmt"
	"net/http"
	isaclient "property/framework/interface/sa/sa_client"
	isafileupload "property/framework/interface/sa/sa_file_upload"
	isauser "property/framework/interface/sa/sa_user"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	"property/framework/pkg/app"
	"property/framework/pkg/logging"
	util "property/framework/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

// ContAuth :
type ContAuth struct {
	useSaClient     isaclient.Usecase
	useSaUser       isauser.Usercase
	useSaFileUpload isafileupload.UseCase
}

func NewContAuth(e *echo.Echo, useSaClient isaclient.Usecase, useSaUser isauser.Usercase, useSaFileUpload isafileupload.UseCase) {
	cont := &ContAuth{
		useSaClient:     useSaClient,
		useSaUser:       useSaUser,
		useSaFileUpload: useSaFileUpload,
	}

	e.POST("/api/auth/register", cont.Register)
	e.POST("/api/auth/login", cont.Login)
}

// Register :
// @Summary Add Client
// @Tags Auth
// @Produce json
// @Param req body models.RegisterForm true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} app.ResponseModel
// @Router /api/auth/register [post]
func (u *ContAuth) Register(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{} // wajib
		appE   = app.Res{R: e}    // wajib
		client sa_models.SaClient

		form = models.RegisterForm{}
	)
	// validasi and bind to struct
	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}
	// mapping to struct model saClient
	err := mapstructure.Decode(form, &client)
	if err != nil {
		return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)
	}

	err = u.useSaClient.RegisterClient(ctx, &client)
	if err != nil {
		return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusCreated, "Ok", client)
}

// Register :
// @Summary Login
// @Tags Auth
// @Produce json
// @Param req body models.LoginForm true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} app.ResponseModel
// @Router /api/auth/login [post]
func (u *ContAuth) Login(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{} // wajib
		appE   = app.Res{R: e}    // wajib
		// client sa_models.SaClient

		form      = models.LoginForm{}
		dataFiles = sa_models.SaFileOutput{}
	)

	// validasi and bind to struct
	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	// if form.UserName == "" {
	// 	return appE.ResponseError(http.StatusUnauthorized, "Email Or UserName can't be blank.", nil)
	// }

	// if form.Password == "" {
	// 	return appE.ResponseError(http.StatusUnauthorized, "Password can't be blank.", nil)
	// }

	DataUser, err := u.useSaUser.GetByEmailSaUser(ctx, form.UserName)
	if err != nil {
		// return appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
		return appE.ResponseError(http.StatusUnauthorized, "Invalid User Or Email.", nil)
	}

	if !util.ComparePassword(DataUser.Passwd, util.GetPassword(form.Password)) {
		return appE.ResponseError(http.StatusUnauthorized, "Invalid Password.", nil)
	}

	token, err := util.GenerateToken(DataUser.UserID.String(), DataUser.UserName)
	if err != nil {
		return appE.ResponseError(http.StatusInternalServerError, "Status Internal Server Error", nil)
	}

	dataFile, err := u.useSaFileUpload.GetBySaFileUpload(ctx, DataUser.FileID)

	err = mapstructure.Decode(dataFile, &dataFiles)
	if err != nil {
		return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)

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
		"picture_url":  dataFiles,
	}
	response := map[string]interface{}{
		"token":     token,
		"data_user": restUser,
	}

	return appE.Response(http.StatusOK, "Ok", response)
}

// Register :
// @Summary Forgot Password
// @Tags Auth
// @Produce json
// @Param req body models.ForgotForm true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} app.ResponseModel
// @Router /api/auth/login [post]
func (u *ContAuth) ForgotPassword(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{} // wajib
		appE   = app.Res{R: e}    // wajib
		// client sa_models.SaClient

		form = models.ForgotForm{}
	)

	// validasi and bind to struct
	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	dataUser, err := u.useSaUser.GetByEmailSaUser(ctx, form.EmailAddr)
	if err != nil {
		return appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}
	logger.Info(util.Stringify(dataUser))
	return appE.Response(http.StatusOK, "Ok", "Please Check Your Email")
}
