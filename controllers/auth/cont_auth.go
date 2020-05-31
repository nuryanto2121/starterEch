package contauth

import (
	"context"
	"fmt"
	"net/http"
	isaclient "property/framework/interface/sa/sa_client"
	isauser "property/framework/interface/sa/sa_user"
	sa_models "property/framework/models/sa"
	"property/framework/pkg/app"
	"property/framework/pkg/logging"
	util "property/framework/pkg/utils"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

// ContAuth :
type ContAuth struct {
	useSaClient isaclient.Usecase
	useSaUser   isauser.Usercase
}

func NewContAuth(e *echo.Echo, useSaClient isaclient.Usecase, useSaUser isauser.Usercase) {
	cont := &ContAuth{
		useSaClient: useSaClient,
		useSaUser:   useSaUser,
	}

	e.POST("/api/auth/register", cont.Register)
	e.POST("/api/auth/login", cont.Login)
}

// RegisterForm :
type RegisterForm struct {
	ClientName       string    `json:"client_name" valid:"Required"`
	Address          string    `json:"address,omitempty"`
	PostCd           string    `json:"post_cd,omitempty"`
	TelephoneNo      string    `json:"telephone_no,omitempty"`
	EmailAddr        string    `json:"email_addr,omitempty"`
	ContactPerson    string    `json:"contact_person,omitempty"`
	ClientType       string    `json:"client_type,omitempty"`
	JoiningDate      time.Time `json:"joining_date,omitempty"`
	StartBillingDate time.Time `json:"start_billing_date,omitempty"`
	ExpiryDate       time.Time `json:"expiry_date,omitempty"`
	CreatedBy        string    `json:"created_by" valid:"Required"`
}

// Register :
// @Summary Add Client
// @Tags Auth
// @Produce json
// @Param req body contauth.RegisterForm true "req param #changes are possible to adjust the form of the registration form from frontend"
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

		form = RegisterForm{}
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

//LoginForm :
type LoginForm struct {
	UserName string `json:"u" valid:"Required"`
	Password string `json:"p" valid:"Required"`
}

// Register :
// @Summary Login
// @Tags Auth
// @Produce json
// @Param req body contauth.LoginForm true "req param #changes are possible to adjust the form of the registration form from frontend"
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

		form = LoginForm{}
	)

	// validasi and bind to struct
	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	if form.UserName == "" {
		return appE.ResponseError(http.StatusUnauthorized, "Email Or UserName can't be blank.", nil)
	}

	if form.Password == "" {
		return appE.ResponseError(http.StatusUnauthorized, "Password can't be blank.", nil)
	}

	DataUser, err := u.useSaUser.GetByEmailSaUser(ctx, form.UserName)
	if err != nil {
		// return appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
		return appE.ResponseError(http.StatusUnauthorized, "Invalid User Or Email.", nil)
	}
	// form.Password, _ = util.Hash(form.Password)

	// if ok, _ := util.Compare(form.Password, DataUser.Passwd); !ok {
	// 	return appE.ResponseError(http.StatusUnauthorized, "Invalid Password.", nil)
	// }
	if !util.ComparePassword(DataUser.Passwd, util.GetPassword(form.Password)) {
		return appE.ResponseError(http.StatusUnauthorized, "Invalid Password.", nil)
	}

	token, err := util.GenerateToken(DataUser.UserID.String(), DataUser.RoleID.String(), DataUser.CompanyID)
	if err != nil {
		return appE.ResponseError(http.StatusInternalServerError, "Status Internal Server Error", nil)
	}

	restUser := map[string]interface{}{
		"user_id":      DataUser.UserID,
		"clinet_id":    DataUser.ClientID,
		"user_name":    DataUser.UserName,
		"level_no":     DataUser.LevelNo,
		"role_id":      DataUser.RoleID,
		"email_addr":   DataUser.EmailAddr,
		"handphone_no": DataUser.HandphoneNo,
		"company_id":   DataUser.CompanyID,
		"picture_url":  DataUser.PictureURL,
	}
	response := map[string]interface{}{
		"token":     token,
		"data_user": restUser,
	}

	return appE.Response(http.StatusOK, "Ok", response)
}
