package app

import (
	"fmt"

	"github.com/astaxie/beego/validation"
)

// MarkErrors :
func MarkErrors(errors []*validation.Error) string {
	res := ""
	for _, err := range errors {
		res = fmt.Sprintf("%s %s", err.Key, err.Message)
	}

	return res
}
