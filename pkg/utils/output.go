package util

import (
	"errors"
	"fmt"
	models "property/framework/models"
)

func Goutput(data interface{}, statusCode int) (out models.Output) {

	out.Data = data
	out.Err = nil
	out.Code = statusCode //GetStatusCode(err)
	return out
}

func GoutputErr(err error) (out models.Output) {
	out.Err = err
	out.Data = nil
	out.Code = GetStatusCode(err)
	out.Msg = fmt.Sprintf("%v", err)
	return out
}
func GoutputErrCode(code int, msg string) (out models.Output) {
	out.Err = errors.New(msg)
	out.Data = nil
	out.Code = code //GetStatusCode(err)
	out.Msg = msg
	return out
}
