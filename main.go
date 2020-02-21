package main

import (
	"fmt"
	"log"
	_contsauser "property/framework/controllers/sa/sa_user"
	"property/framework/middleware"
	"property/framework/pkg/connection"
	"property/framework/pkg/setting"
	_reposauser "property/framework/repository/sa/sa_user"
	_usesauser "property/framework/usecase/sa/sa_user"
	"time"

	"github.com/labstack/echo"
)

func init() {
	setting.Setup()
	connection.Setup()
}

func main() {
	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	timeoutContext := time.Duration(setting.FileConfigSetting.Server.ReadTimeout) * time.Second

	repoSaUser := _reposauser.NewRepoSaUser(connection.Conn)
	useSaUser := _usesauser.NewUseSaUser(repoSaUser, timeoutContext)
	_contsauser.NewContSaUser(e, useSaUser)

	sPort := fmt.Sprintf(":%d", setting.FileConfigSetting.Server.HTTPPort)
	log.Fatal(e.Start(sPort))
	// log.Fatal(e.Start(":" + string(setting.FileConfigSetting.Server.HTTPPort)))
}
