package contsagroup

import (
	"context"
	"fmt"
	"net/http"
	isagroup "property/framework/interface/sa/sa_group"
	"property/framework/models"
	"property/framework/pkg/app"
	"property/framework/pkg/logging"
	util "property/framework/pkg/utils"

	"github.com/astaxie/beego/validation"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

// ContSaGroup :
type ContSaGroup struct {
	useSaGroup isagroup.UseCase
}

// NewContSaGroup :
func NewContSaGroup(e *echo.Echo, a isagroup.UseCase) {
	controller := &ContSaGroup{
		useSaGroup: a,
	}

	e.GET("/api/group/:id", controller.GetBySaGroup)
	e.GET("/api/group", controller.GetList)
	e.POST("/api/group", controller.CreateSaGroup)
	e.PUT("/api/group/:id", controller.UpdateSaGroup)
	e.DELETE("/api/group/:id", controller.DeleteSaGroup)
}

// GetBySaGroup :
// @Summary GetById SaGroup
// @Tags Group
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.ResponseModel
// @Router /api/group/{id} [get]
func (u *ContSaGroup) GetBySaGroup(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{}
		appE   = app.Res{R: e}                       // wajib
		id     = util.StrTo(e.Param("id")).MustInt() //kalo bukan int => 0
		valid  validation.Validation                 // wajib
	)
	valid.Min(id, 1, "id").Message("ID must be greater than 0")
	logger.Info(id)
	if valid.HasErrors() {
		return appE.ResponseError(http.StatusBadRequest, app.MarkErrors(valid.Errors), nil)
	}

	dataGroup, err := u.useSaGroup.GetBySaGroup(ctx, int16(id))
	if err != nil {
		return appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusOK, "Ok", dataGroup)
}

// GetList :
// @Summary GetList SaGroup
// @Tags Group
// @Produce  json
// @Param page query int true "Page"
// @Param perpage query int true "PerPage"
// @Param search query string false "Search"
// @Param initsearch query string false "InitSearch"
// @Param sortfield query string false "SortField"
// @Success 200 {object} models.ResponseModelList
// @Router /api/group [get]
func (u *ContSaGroup) GetList(e echo.Context) error {
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

	responseList, err = u.useSaGroup.GetList(ctx, paramquery)
	if err != nil {
		// return e.JSON(http.StatusBadRequest, err.Error())
		return appE.ResponseErrorList(util.GetStatusCode(err), fmt.Sprintf("%v", err), responseList)
	}

	// return e.JSON(http.StatusOK, ListDataUser)
	return appE.ResponseList(http.StatusOK, "", responseList)
}

// AddGroupForm :
type AddGroupForm struct {
	Descs     string  `json:"descs" valid:"MaxSize(60)"`
	Num       float32 `json:"num"`
	CreatedBy string  `json:"created_by" valid:"Required"`
}

// CreateSaGroup :
// @Summary Add Group
// @Tags Group
// @Produce json
// @Param req body contsagroup.AddGroupForm true "req param #changes are possible to adjust the form of the registration form from forntend"
// @Success 200 {object} app.ResponseModel
// @Router /api/group [post]
func (u *ContSaGroup) CreateSaGroup(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{} // wajib
		appE   = app.Res{R: e}    // wajib
		group  models.SaGroup
		form   AddGroupForm
	)
	// validasi and bind to struct
	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	// mapping to struct model saGroup
	err := mapstructure.Decode(form, &group)
	if err != nil {
		return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)

	}
	err = u.useSaGroup.CreateSaGroup(ctx, &group)
	if err != nil {
		return appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusCreated, "Ok", group)
}

// EditGroupForm :
type EditGroupForm struct {
	Descs     string `json:"descs" valid:"MaxSize(5)"`
	UpdatedBy string `json:"Updated_by" valid:"Required"`
}

// UpdateSaGroup :
// @Summary Update Group
// @Tags Group
// @Produce json
// @Param id path int true "ID"
// @Param req body contsagroup.EditGroupForm true "req param #changes are possible to adjust the form of the registration form from forntend"
// @Success 200 {object} app.ResponseModel
// @Router /api/group/{id} [put]
func (u *ContSaGroup) UpdateSaGroup(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{} // wajib
		appE   = app.Res{R: e}    // wajib
		group  models.SaGroup
		valid  validation.Validation                 // wajib
		id     = util.StrTo(e.Param("id")).MustInt() //kalo bukan int => 0
		form   = EditGroupForm{}
	)

	valid.Min(id, 1, "id").Message("ID must be greater than 0")
	logger.Info(id)
	if valid.HasErrors() {
		return appE.ResponseError(http.StatusBadRequest, app.MarkErrors(valid.Errors), nil)
	}

	// validasi and bind to struct
	httpCode, errMsg := app.BindAndValid(e, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	// mapping to struct model saSuser
	err := mapstructure.Decode(form, &group)
	if err != nil {
		return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)

	}
	group.GroupID = int16(id)
	err = u.useSaGroup.UpdateSaGroup(ctx, &group)
	if err != nil {
		return appE.ResponseError(util.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}
	return appE.ResponseError(http.StatusCreated, fmt.Sprintf("%v", group), nil)
}

// DeleteSaGroup :
// @Summary Delete group
// @Tags Group
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.ResponseModel
// @Router /api/group/{id} [delete]
func (u *ContSaGroup) DeleteSaGroup(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		logger = logging.Logger{}
		appE   = app.Res{R: e}                       // wajib
		id     = util.StrTo(e.Param("id")).MustInt() //kalo bukan int => 0
		valid  validation.Validation                 // wajib
	)

	valid.Min(id, 1, "id").Message("ID must be greater than 0")
	logger.Info(id)
	if valid.HasErrors() {
		return appE.ResponseError(http.StatusBadRequest, app.MarkErrors(valid.Errors), nil)
	}
	err := u.useSaGroup.DeleteSaGroup(ctx, int16(id))
	if err != nil {
		return appE.ResponseError(util.GetStatusCode(err), err.Error(), nil)
	}
	return appE.Response(http.StatusNoContent, "", nil)
}
