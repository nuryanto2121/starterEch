package contlookup

import (
	"context"
	"fmt"
	"net/http"
	idynamic "property/framework/interface/dynamic"
	midd "property/framework/middleware"
	"property/framework/models"
	"property/framework/pkg/app"
	tool "property/framework/pkg/tools"

	"github.com/labstack/echo/v4"
)

type ContLookUp struct {
	useOption idynamic.Usecase
}

func NewContLookUp(e *echo.Echo, a idynamic.Usecase) {
	cont := &ContLookUp{
		useOption: a,
	}

	r := e.Group("api/lookup")

	r.Use(midd.JWT)
	r.POST("", cont.GetData)
}

// GetLookUp :
// @Summary GetLookUp Dynamic
// @Security ApiKeyAuth
// @Tags LookUp
// @Produce  json
// @Param req body models.ParamLookup true "req param #changes are possible to adjust the form of the registration form from frontend"
// @Success 200 {object} app.ResponseModel
// @Router /api/lookup [post]
func (c *ContLookUp) GetData(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	var (
		// logger = logging.Logger{}
		appE = tool.Res{R: e} // wajib
		//valid      validation.Validation // wajib
		paramquery = models.ParamLookup{} // ini untuk list
		Data       interface{}
		err        error
	)

	httpCode, errMsg := app.BindAndValid(e, &paramquery)
	// logger.Info(util.Stringify(paramquery))
	if httpCode != 200 {
		return appE.ResponseError(http.StatusBadRequest, errMsg, nil)
	}

	claims, err := app.GetClaims(e)
	if err != nil {
		return appE.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	Data, err = c.useOption.GetDataLookUp(ctx, claims, paramquery)
	if err != nil {
		return appE.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)
	}

	return appE.Response(http.StatusOK, "", Data)

}
