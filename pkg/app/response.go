package app

import (
	"property/framework/models"
	"property/framework/pkg/logging"
	util "property/framework/pkg/utils"

	"github.com/labstack/echo/v4"
)

// Res :
type Res struct {
	R echo.Context
	// List ResponseModelList
}

// Status  int         `json:"status"`
// 	Message string      `json:"message"`
// 	Error   bool        `json:"error"`
// 	Data    interface{} `json:"data"`

// ResponseModel :
type ResponseModel struct {
	// Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

// Response :
func (e Res) Response(httpCode int, errMsg string, data interface{}) error {
	var logger = logging.Logger{}
	response := ResponseModel{
		// Status:  httpCode,
		// Message: errMsg,
		// Error:   err,
		// Data:    data,
		// Code: httpCode,
		Msg:  errMsg,
		Data: data,
	}
	logger.Info(string(util.Stringify(response)))
	return e.R.JSON(httpCode, response)
	// return string(util.Stringify(response))
}

// ResponseError :
func (e Res) ResponseError(httpCode int, errMsg string, data interface{}) error {
	var logger = logging.Logger{}
	response := ResponseModel{
		// Status:  httpCode,
		// Message: errMsg,
		// Error:   err,
		// Data:    data,
		// Code: httpCode,
		Msg:  errMsg,
		Data: data,
	}
	logger.Error(string(util.Stringify(response)))
	return e.R.JSON(httpCode, response)
	// return string(util.Stringify(response))
}

// ResponseList :
func (e Res) ResponseList(httpCode int, errMsg string, data models.ResponseModelList) error {
	var logger = logging.Logger{}
	data.Msg = errMsg

	logger.Info(string(util.Stringify(data)))
	return e.R.JSON(httpCode, data)
	// return string(util.Stringify(response))
}

// ResponseErrorList :
func (e Res) ResponseErrorList(httpCode int, errMsg string, data models.ResponseModelList) error {
	var logger = logging.Logger{}
	data.Msg = errMsg

	logger.Error(string(util.Stringify(data)))
	return e.R.JSON(httpCode, data)
	// return string(util.Stringify(response))
}
