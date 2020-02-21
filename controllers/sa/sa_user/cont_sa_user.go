package contsauser

import (
	"context"
	"net/http"
	isauser "property/framework/interface/sa/sa_user"
	"property/framework/models"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

// ContSaUser :
type ContSaUser struct {
	useSaUser isauser.Usercase
}

// NewContSaUser :
func NewContSaUser(e *echo.Echo, useSaUser isauser.Usercase) {
	controller := &ContSaUser{
		useSaUser: useSaUser,
	}
	e.GET("/user/:id", controller.GetBySaUser)
	e.GET("/user", controller.GetAllSaUser)
	e.POST("/user", controller.CreateSaUser)
	e.PUT("/user/:id", controller.UpdateSaUser)
	e.DELETE("/user/:id", controller.DeleteSaUser)
}

// GetBySaUser :
func (u *ContSaUser) GetBySaUser(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	userID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}
	DataUser, err := u.useSaUser.GetBySaUser(ctx, int16(userID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, DataUser)
}

// GetAllSaUser :
func (u *ContSaUser) GetAllSaUser(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	ListDataUser, err := u.useSaUser.GetAllSaUser(ctx)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, ListDataUser)
}

// CreateSaUser :
func (u *ContSaUser) CreateSaUser(e echo.Context) error {
	var user models.SaUser
	err := e.Bind(&user)
	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&user); !ok {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = u.useSaUser.CreateSaUser(ctx, &user)
	if err != nil {
		return e.JSON(getStatusCode(err), err.Error())
	}
	return e.JSON(http.StatusCreated, user)
}

// UpdateSaUser :
func (u *ContSaUser) UpdateSaUser(e echo.Context) error {
	var user models.SaUser
	err := e.Bind(&user)
	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	userID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}
	user.UserID = int16(userID)

	if ok, err := isRequestValid(&user); !ok {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = u.useSaUser.UpdateSaUser(ctx, &user)
	if err != nil {
		return e.JSON(getStatusCode(err), err.Error())
	}
	return e.JSON(http.StatusCreated, user)
}

// DeleteSaUser :
func (u *ContSaUser) DeleteSaUser(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	userID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}
	err = u.useSaUser.DeleteSaUser(ctx, int16(userID))
	if err != nil {
		return e.JSON(getStatusCode(err), err.Error())
	}
	return e.NoContent(http.StatusNoContent)
}

// HealthCheck :
func (u *ContSaUser) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "success")
}

// isRequestValid
func isRequestValid(m *models.SaUser) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// getStatusCode
func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
