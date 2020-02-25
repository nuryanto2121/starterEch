package models

// ResponseModelList :
type ResponseModelList struct {
	Page     int         `json:"page"`
	Total    int         `json:"total"`
	LastPage int         `json:"last_page"`
	Data     interface{} `json:"data"`
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
}
