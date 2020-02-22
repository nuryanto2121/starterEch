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

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
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
