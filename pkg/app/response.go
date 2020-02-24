package app

import (
	util "property/framework/pkg/utils"

	"github.com/labstack/echo/v4"
)

// Res :
type Res struct {
	R echo.Context
}

// Status  int         `json:"status"`
// 	Message string      `json:"message"`
// 	Error   bool        `json:"error"`
// 	Data    interface{} `json:"data"`

// ResponseModel :
type ResponseModel struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response :
func (e Res) Response(httpCode int, errMsg string, data interface{}) (interface{}, error) {
	response := ResponseModel{
		// Status:  httpCode,
		// Message: errMsg,
		// Error:   err,
		// Data:    data,
		Code: httpCode,
		Msg:  errMsg,
		Data: data,
	}
	return string(util.Stringify(response)), e.R.JSON(httpCode, response)
	// return string(util.Stringify(response))
}

// ResponseList :
func (e Res) ResponseList(httpCode int, errMsg string, data interface{}) (interface{}, error) {
	response := ResponseModel{
		// Status:  httpCode,
		// Message: errMsg,
		// Error:   err,
		// Data:    data,
		Code: httpCode,
		Msg:  errMsg,
		Data: data,
	}
	return string(util.Stringify(response)), e.R.JSON(httpCode, response)
	// return string(util.Stringify(response))
}
