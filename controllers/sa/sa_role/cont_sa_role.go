package contsarole

import (
	"context"
	"fmt"
	"net/http"
	isarole "property/framework/interface/sa/sa_role"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	"property/framework/pkg/app"
	"property/framework/pkg/logging"
	"property/framework/pkg/setting"
	util "property/framework/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mitchellh/mapstructure"

	uuid "github.com/satori/go.uuid"
)

// ContSaRole :
type ContSaRole struct {
	useSaRole isarole.UseCase
}

// NewContSaRole :
func NewContSaRole(e *echo.Echo, a isarole.UseCase) {
	controller := &ContSaRole{
		useSaRole: a,
	}

	r := e.Group("/api/role")
	// Configure middleware with custom claims
	var screet = setting.FileConfigSetting.App.JwtSecret
	config := middleware.JWTConfig{
		Claims:     &util.Claims{},
		SigningKey: []byte(screet),
	}
	r.Use(middleware.JWTWithConfig(config))

	r.GET("/:id", controller.GetBySaRole)
	r.GET("/menu_access", controller.GetMenuAccess)
	r.GET("", controller.GetList)
	r.POST("", controller.CreateSaRole)
	r.PUT("/:id", controller.UpdateSaRole)
	r.DELETE("/:id", controller.DeleteSaRole)
}

// GetBySaRole :
// @Summary GetById SaRole
// @Security ApiKeyAuth
// @Tags Role
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} app.ResponseModel
// @Router /api/role/{id} [get]
func (u *ContSaRole) GetBySaRole(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{}
		appE   = app.Res{R: e} // wajib
		id     = e.Param("id") //kalo bukan int => 0
		// valid  validation.Validation                 // wajib
	)
	RoleID, err := uuid.FromString(id)
	logger.Info(id)
	if err != nil {
		return appE.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}

	dataRole, err := u.useSaRole.GetBySaRole(ctx, RoleID)
	if err != nil {
		return appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusOK, "Ok", dataRole)
}

// GetList :
// @Summary Get string Array Menu Access Role
// @Security ApiKeyAuth
// @Tags Role
// @Produce  json
// @Param role_id query string false "RoleID"
// @Success 200 {object} app.ResponseModel
// @Router /api/role/menu_access [get]
func (u *ContSaRole) GetMenuAccess(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{}
		appE   = app.Res{R: e} // wajib
		//valid      validation.Validation // wajib

		err error
	)

	// roleID := e.QueryParam("role_id")
	roleID, _ := uuid.FromString(e.QueryParam("role_id"))
	logger.Info(util.Stringify(roleID))

	dd, err := u.useSaRole.GetJsonMenuAccess(ctx, roleID)
	if err != nil {
		return appE.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	return appE.Response(http.StatusCreated, "Ok", dd)
}

// GetList :
// @Summary GetList SaRole
// @Security ApiKeyAuth
// @Tags Role
// @Produce  json
// @Param page query int true "Page"
// @Param perpage query int true "PerPage"
// @Param search query string false "Search"
// @Param initsearch query string false "InitSearch"
// @Param sortfield query string false "SortField"
// @Success 200 {object} models.ResponseModelList
// @Router /api/role [get]
func (u *ContSaRole) GetList(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{}
		appE   = app.Res{R: e} // wajib
		//valid      validation.Validation // wajib
		paramquery   = models.ParamList{} // ini untuk list
		responseList = models.ResponseModelList{}
		err          error
	)

	httpCode, errMsg := app.BindAndValid(e, &paramquery)
	logger.Info(util.Stringify(paramquery))
	if httpCode != 200 {
		return appE.ResponseErrorList(http.StatusBadRequest, errMsg, responseList)
	}

	responseList, err = u.useSaRole.GetList(ctx, paramquery)
	if err != nil {
		// return e.JSON(http.StatusBadRequest, err.Error())
		return appE.ResponseErrorList(util.GetStatusCode(err), fmt.Sprintf("%v", err), responseList)
	}

	// return e.JSON(http.StatusOK, ListDataUser)
	return appE.ResponseList(http.StatusOK, "", responseList)
}

// CreateSaRole :
// @Summary Add Role
// @Security ApiKeyAuth
// @Tags Role
// @Produce json
// @Param req body models.AddRoleForm true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} app.ResponseModel
// @Router /api/role [post]
func (u *ContSaRole) CreateSaRole(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger     = logging.Logger{} // wajib
		appE       = app.Res{R: e}    // wajib
		role       sa_models.SaRole
		form       sa_models.AddRoleForm
		menuAccess []sa_models.MenuAccessLevel1
	)

	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(*util.Claims)
	// validasi and bind to struct
	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	// mapping to struct model saRole
	err := mapstructure.Decode(form, &role)
	if err != nil {
		return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)

	}

	err = mapstructure.Decode(form.MenuAccess, &menuAccess)
	if err != nil {
		return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)

	}

	role.CreatedBy = claims.UserName
	role.UpdatedBy = claims.UserName
	err = u.useSaRole.CreateSaRole(ctx, &role, &menuAccess)
	if err != nil {
		return appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusCreated, "Ok", role)
}

// UpdateSaRole :
// @Summary Update Role
// @Security ApiKeyAuth
// @Tags Role
// @Produce json
// @Param id path string true "ID"
// @Param req body models.EditRoleForm true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} app.ResponseModel
// @Router /api/role/{id} [put]
func (u *ContSaRole) UpdateSaRole(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{} // wajib
		appE   = app.Res{R: e}    // wajib
		// role   sa_models.SaRole
		err error
		// valid  validation.Validation                 // wajib
		id   = e.Param("id") //kalo bukan int => 0
		form = sa_models.EditRoleForm{}
		// menuAccess []sa_models.MenuAccessLevel1
	)
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(*util.Claims)

	form.UpdatedBy = claims.UserName
	RoleID, err := uuid.FromString(id)
	logger.Info(id)
	if err != nil {
		return appE.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}

	// validasi and bind to struct
	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	err = u.useSaRole.UpdateSaRole(ctx, RoleID, &form)
	if err != nil {
		return appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}
	return appE.Response(http.StatusCreated, "Ok", nil)
}

// DeleteSaRole :
// @Summary Delete role
// @Security ApiKeyAuth
// @Tags Role
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} app.ResponseModel
// @Router /api/role/{id} [delete]
func (u *ContSaRole) DeleteSaRole(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{}
		err    error
		appE   = app.Res{R: e} // wajib
		id     = e.Param("id") //kalo bukan int => 0
		// valid  validation.Validation                 // wajib
	)

	RoleID, err := uuid.FromString(id)
	logger.Info(id)
	if err != nil {
		return appE.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	err = u.useSaRole.DeleteSaRole(ctx, RoleID)
	if err != nil {
		return appE.ResponseError(util.GetStatusCode(err), err.Error(), nil)
	}
	return appE.Response(http.StatusNoContent, "", nil)
}
