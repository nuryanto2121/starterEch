package main

import (
	"net/http"
	"property/framework/pkg/setting"

	"github.com/labstack/echo"
)

func init() {
	setting.Setup()
}

func main() {
	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		dd := setting.FileConfigSetting
		return ctx.JSON(http.StatusOK, dd)
	})

	r.Start(":9000")
}
