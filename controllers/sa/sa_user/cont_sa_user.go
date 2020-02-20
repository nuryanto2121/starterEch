package contsauser

import (
	"context"
	"net/http"
	isauser "property/framework/interface/sa/sa_user"
	"strconv"

	"github.com/labstack/echo"
)

type ContSaUser struct {
	useSaUser isauser.Usercase
}

// NewContSaUser :
func NewContSaUser(e *echo.Echo, useSaUser isauser.Usercase) {
	controller := &ContSaUser{
		useSaUser: useSaUser,
	}
	e.GET("user/:id", controller.GetBySaUser)
	e.GET("user", controller.GetBySaUser)
}
func (u *ContSaUser) GetBySaUser(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	userID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	DataUser, err := u.useSaUser.GetBySaUser(ctx, int16(userID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, DataUser)
}

func (u *ContSaUser) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "success")
}
