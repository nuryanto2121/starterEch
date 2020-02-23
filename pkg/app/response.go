package app

import "github.com/labstack/echo/v4"

// Res :
type Res struct {
	R echo.Context
}

// Status  int         `json:"status"`
// 	Message string      `json:"message"`
// 	Error   bool        `json:"error"`
// 	Data    interface{} `json:"data"`
// Response :
type ResponseModel struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response :
func (e Res) Response(httpCode int, errMsg string, data interface{}) error {
	response := ResponseModel{
		// Status:  httpCode,
		// Message: errMsg,
		// Error:   err,
		// Data:    data,
		Code: httpCode,
		Msg:  errMsg,
		Data: data,
	}
	return e.R.JSON(httpCode, response)
	// return string(util.Stringify(response))
}
