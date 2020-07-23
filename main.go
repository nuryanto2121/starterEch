package main

import (
	"fmt"
	"log"

	"property/framework/pkg/connection"
	"property/framework/pkg/logging"
	sqlxposgresdb "property/framework/pkg/postgresqlxdb"
	"property/framework/pkg/setting"
	"property/framework/redisdb"

	"property/framework/routes"

	// _midd "property/framework/middleware"
	//format nama table + package

	// _sauserusecompany "property/framework/usecase/sa/sa_user_company"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	setting.Setup()
	logging.Setup()
	connection.Setup()
	sqlxposgresdb.Setup()
	redisdb.Setup()
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
	// middL := _midd.InitMiddleware()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Static("/static", "wwwroot")
	// e.Use(middL.CORS)
	e.Use(middleware.CORS())
	app := routes.Echo{E: e}

	app.InitialRouter()

	// middL := _midd.InitMiddleware()
	// e.Use(middL.CORS)

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
