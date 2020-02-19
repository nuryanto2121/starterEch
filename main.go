package main

import (
	"github.com/labstack/echo"
	"net/http"
	"property/framework/pkg/connection"
	"property/framework/pkg/setting"
)

func init() {
	setting.Setup()
	connection.Setup()
}

func main() {
	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		dd := setting.FileConfigSetting
		return ctx.JSON(http.StatusOK, dd)
	})

	r.Start(":9000")
}
