package models

type Output struct {
	Code int         `json:"code"`
	Err  error       `json:"is_error"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}
