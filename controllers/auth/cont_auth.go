package contauth

import (
	"context"
	"net/http"
	iauth "property/framework/interface/auth"

	// isaclient "property/framework/interface/sa/sa_client"
	// isafileupload "property/framework/interface/sa/sa_file_upload"
	// isauser "property/framework/interface/sa/sa_user"
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
	useAuth iauth.Usecase
	// useSaClient     isaclient.Usecase
	// useSaUser       isauser.Usecase
	// useSaFileUpload isafileupload.UseCase
}

func NewContAuth(e *echo.Echo, useAuth iauth.Usecase) {
	cont := &ContAuth{
		useAuth: useAuth,
		// useSaClient:     useSaClient,
		// useSaUser:       useSaUser,
		// useSaFileUpload: useSaFileUpload,
	}

	e.POST("/api/auth/register", cont.Register)
	e.POST("/api/auth/login", cont.Login)
	e.POST("/api/auth/forgot", cont.ForgotPassword)
	e.POST("/api/auth/reset", cont.ResetPasswd)
	e.POST("/api/auth/verify", cont.Verify)
}

// Register :
// @Summary Add Client
// @Tags Auth
// @Produce json
// @Param OS header string true "OS Device"
// @Param Version header string true "OS Device"
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
		return appE.ResponseErr(util.GoutputErr(err)) //return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)
	}

	out := u.useAuth.Register(ctx, &client)
	// err = u.useSaClient.RegisterClient(ctx, &client)
	if out.Err != nil {
		// return appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
		return appE.ResponseErr(out)
	}

	return appE.Response(http.StatusCreated, "Ok", out.Data)
}

// Register :
// @Summary Login
// @Tags Auth
// @Produce json
// @Param OS header string true "OS Device"
// @Param Version header string true "OS Device"
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

		form = models.LoginForm{}
		// dataFiles = sa_models.SaFileOutput{}
	)

	// validasi and bind to struct
	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	out := u.useAuth.Login(ctx, &form) //u.useSaUser.GetByEmailSaUser(ctx, form.UserName)
	if out.Err != nil {
		return appE.ResponseErr(out)
		//return appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
		// return appE.ResponseError(http.StatusUnauthorized, err, nil)
	}

	return appE.Response(http.StatusOK, "Ok", out.Data)
}

// Register :
// @Summary Forgot Password
// @Tags Auth
// @Produce json
// @Param OS header string true "OS Device"
// @Param Version header string true "OS Device"
// @Param req body models.ForgotForm true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} app.ResponseModel
// @Router /api/auth/forgot [post]
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
	out := u.useAuth.ForgotPassword(ctx, &form)
	if out.Err != nil {
		return appE.ResponseErr(out)
	}

	return appE.Response(http.StatusOK, "Ok", "Please Check Your Email")
}

// Register :
// @Summary Reset Password
// @Tags Auth
// @Produce json
// @Param OS header string true "OS Device"
// @Param Version header string true "OS Device"
// @Param req body models.ResetPasswd true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} app.ResponseModel
// @Router /api/auth/reset [post]
func (u *ContAuth) ResetPasswd(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{} // wajib
		appE   = app.Res{R: e}    // wajib
		// client sa_models.SaClient

		form = models.ResetPasswd{}
	)
	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}
	out := u.useAuth.ResetPassword(ctx, &form)
	if out.Err != nil {
		return appE.ResponseErr(out)
	}

	return appE.Response(http.StatusOK, "Ok", "Please Login")
}

// Register :
// @Summary Verify / Aktivasi User
// @Tags Auth
// @Produce json
// @Param OS header string true "OS Device"
// @Param Version header string true "OS Device"
// @Param req body models.ResetPasswd true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} app.ResponseModel
// @Router /api/auth/verify [post]
func (u *ContAuth) Verify(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{} // wajib
		appE   = app.Res{R: e}    // wajib
		// client sa_models.SaClient

		form = models.ResetPasswd{}
	)
	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}
	out := u.useAuth.Verify(ctx, &form)
	if out.Err != nil {
		return appE.ResponseErr(out)
	}

	return appE.Response(http.StatusOK, "Ok", "Account Anda telah aktiv, Silahkan Login")
}
