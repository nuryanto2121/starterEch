package app

import (
	"fmt"
	"net/http"
	util "property/framework/pkg/utils"

	"github.com/astaxie/beego/validation"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

// BindAndValid :
func BindAndValid(c echo.Context, form interface{}) (int, string) {
	err := c.Bind(form)

	if err != nil {
		return http.StatusBadRequest, fmt.Sprintf("invalid request param: %v", err)
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, fmt.Sprintf("internal server error: %v", err)
	}
	if !check {
		return http.StatusBadRequest, MarkErrors(valid.Errors)
	}
	return http.StatusOK, "ok"
}

// GetClaims :
func GetClaims(c echo.Context) (util.Claims, error) {
	var clm util.Claims
	claims := c.Get("claims")
	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(*util.Claims)

	err := mapstructure.Decode(claims, &clm)
	if err != nil {
		return clm, err
	}
	return clm, nil
}
