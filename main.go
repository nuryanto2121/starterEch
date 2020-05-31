package main

import (
	"fmt"
	"log"

	_midd "property/framework/middleware"
	"property/framework/pkg/connection"
	"property/framework/pkg/logging"
	"property/framework/pkg/setting"
	"property/framework/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	setting.Setup()
	logging.Setup()
	connection.Setup()
}

// @title Starter
// @version 1.0
// @description Backend REST API for golang starter

// @contact.name Nuryanto
// @contact.url https://www.linkedin.com/in/nuryanto-1b2721156/
// @contact.email nuryantofattih@gmail.com

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	e := echo.New()
	middL := _midd.InitMiddleware()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "wwwroot")
	e.Use(middL.CORS)

	app := routes.Echo{E: e}

	app.InitialRouter()

	sPort := fmt.Sprintf(":%d", setting.FileConfigSetting.Server.HTTPPort)

	// maxHeaderBytes := 1 << 20
	// s := &http.Server{
	// 	Addr:           sPort,
	// 	ReadTimeout:    setting.FileConfigSetting.Server.ReadTimeout,
	// 	WriteTimeout:   setting.FileConfigSetting.Server.WriteTimeout,
	// 	MaxHeaderBytes: maxHeaderBytes,
	// }
	// // e.Logger.Fatal(e.StartServer(s))
	// s.ListenAndServe()

	// log.Fatal(e.StartServer(s))
	log.Fatal(e.Start(sPort))
	// log.Fatal(e.Start(":" + string(setting.FileConfigSetting.Server.HTTPPort)))
}
