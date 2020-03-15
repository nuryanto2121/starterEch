package contauth

import (
	"context"
	"fmt"
	"net/http"
	isaclient "property/framework/interface/sa/sa_client"
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
}

func NewContAuth(e *echo.Echo, useSaClient isaclient.Usecase) {
	cont := &ContAuth{
		useSaClient: useSaClient,
	}

	e.POST("/api/auth/register", cont.Register)
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
// @Param req body contauth.RegisterForm true "req param #changes are possible to adjust the form of the registration form from forntend"
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
