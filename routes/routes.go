package routes

import (
	"property/framework/pkg/setting"

	sqlxposgresdb "property/framework/pkg/postgresqlxdb"

	_contDynamic "property/framework/controllers/dynamic"
	_repoDynamic "property/framework/repository/dynamic"
	sarouter "property/framework/routes/sa"
	_useDynamic "property/framework/usecase/dynamic"

	_contLookUp "property/framework/controllers/look_up"

	"time"

	"github.com/labstack/echo/v4"

	_midd "property/framework/middleware"
)

// Echo :
type Echo struct {
	E *echo.Echo
}

// InitialRouter :
func (e *Echo) InitialRouter() {
	middL := _midd.InitMiddleware()
	e.E.Use(middL.CORS)
	timeoutContext := time.Duration(setting.FileConfigSetting.Server.ReadTimeout) * time.Second

	repoDynamic := _repoDynamic.NewRepoOptionDB(sqlxposgresdb.DbCon)
	useDynamic := _useDynamic.NewUserSysUser(repoDynamic, timeoutContext)
	_contDynamic.NewContDynamic(e.E, useDynamic)

	_contLookUp.NewContLookUp(e.E, useDynamic)

	sarouter.Router(e.E, timeoutContext)

}
