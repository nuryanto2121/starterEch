package contsamenu

import (
	"context"
	"fmt"
	"net/http"
	isamenu "property/framework/interface/sa/sa_menu"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	"property/framework/pkg/app"
	"property/framework/pkg/logging"
	"property/framework/pkg/setting"
	util "property/framework/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mitchellh/mapstructure"
)

// ContSaMenu :
type ContSaMenu struct {
	useSaMenu isamenu.UseCase
}

// NewContSaMenu :
func NewContSaMenu(e *echo.Echo, a isamenu.UseCase) {
	controller := &ContSaMenu{
		useSaMenu: a,
	}

	r := e.Group("/api/menu")
	// Configure middleware with custom claims
	var screet = setting.FileConfigSetting.App.JwtSecret
	config := middleware.JWTConfig{
		Claims:     &util.Claims{},
		SigningKey: []byte(screet),
	}
	r.Use(middleware.JWTWithConfig(config))

	r.GET("/:id", controller.GetBySaMenu)
	r.GET("", controller.GetList)
	r.POST("", controller.CreateSaMenu)
	r.PUT("/:id", controller.UpdateSaMenu)
	r.DELETE("/:id", controller.DeleteSaMenu)
}

// GetBySaMenu :
// @Summary GetById SaMenu
// @Security ApiKeyAuth
// @Tags Menu
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} app.ResponseModel
// @Router /api/menu/{id} [get]
func (u *ContSaMenu) GetBySaMenu(e echo.Context) error {
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
	MenuID := util.StrTo(id).MustInt()
	logger.Info(id)

	dataMenu, err := u.useSaMenu.GetBySaMenu(ctx, MenuID)
	if err != nil {
		return appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusOK, "Ok", dataMenu)
}

// GetList :
// @Summary GetList SaMenu
// @Security ApiKeyAuth
// @Tags Menu
// @Produce  json
// @Param page query int true "Page"
// @Param perpage query int true "PerPage"
// @Param search query string false "Search"
// @Param initsearch query string false "InitSearch"
// @Param sortfield query string false "SortField"
// @Success 200 {object} models.ResponseModelList
// @Router /api/menu [get]
func (u *ContSaMenu) GetList(e echo.Context) error {
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

	responseList, err = u.useSaMenu.GetList(ctx, paramquery)
	if err != nil {
		// return e.JSON(http.StatusBadRequest, err.Error())
		return appE.ResponseErrorList(util.GetStatusCode(err), fmt.Sprintf("%v", err), responseList)
	}

	// return e.JSON(http.StatusOK, ListDataUser)
	return appE.ResponseList(http.StatusOK, "", responseList)
}

// CreateSaMenu :
// @Summary Add Menu
// @Security ApiKeyAuth
// @Tags Menu
// @Produce json
// @Param req body models.AddMenuForm true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} app.ResponseModel
// @Router /api/menu [post]
func (u *ContSaMenu) CreateSaMenu(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{} // wajib
		appE   = app.Res{R: e}    // wajib
		menu   sa_models.SaMenu
		form   sa_models.AddMenuForm
	)
	// validasi and bind to struct
	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	// mapping to struct model saMenu
	err := mapstructure.Decode(form, &menu)
	if err != nil {
		return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)

	}
	err = u.useSaMenu.CreateSaMenu(ctx, &menu)
	if err != nil {
		return appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusCreated, "Ok", menu)
}

// UpdateSaMenu :
// @Summary Update Menu
// @Security ApiKeyAuth
// @Tags Menu
// @Produce json
// @Param id path string true "ID"
// @Param req body models.EditMenuForm true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} app.ResponseModel
// @Router /api/menu/{id} [put]
func (u *ContSaMenu) UpdateSaMenu(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{} // wajib
		appE   = app.Res{R: e}    // wajib
		menu   sa_models.SaMenu
		err    error
		// valid  validation.Validation                 // wajib
		id   = e.Param("id") //kalo bukan int => 0
		form = sa_models.EditMenuForm{}
	)

	MenuID := util.StrTo(id).MustInt()
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

	// mapping to struct model saSuser
	err = mapstructure.Decode(form, &menu)
	if err != nil {
		return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)

	}
	menu.MenuID = MenuID
	err = u.useSaMenu.UpdateSaMenu(ctx, &menu)
	if err != nil {
		return appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}
	return appE.ResponseError(http.StatusCreated, fmt.Sprintf("%v", menu), nil)
}

// DeleteSaMenu :
// @Summary Delete menu
// @Security ApiKeyAuth
// @Tags Menu
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} app.ResponseModel
// @Router /api/menu/{id} [delete]
func (u *ContSaMenu) DeleteSaMenu(e echo.Context) error {
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

	MenuID := util.StrTo(id).MustInt()
	logger.Info(id)
	if err != nil {
		return appE.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	err = u.useSaMenu.DeleteSaMenu(ctx, MenuID)
	if err != nil {
		return appE.ResponseError(util.GetStatusCode(err), err.Error(), nil)
	}
	return appE.Response(http.StatusNoContent, "", nil)
}
