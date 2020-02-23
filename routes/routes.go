package routes

import (
	"property/framework/pkg/connection"
	"property/framework/pkg/setting"
	//format nama table + package
	_sausercont "property/framework/controllers/sa/sa_user"
	_sauserrepo "property/framework/repository/sa/sa_user"
	_sauseruse "property/framework/usecase/sa/sa_user"
	"time"

	"github.com/labstack/echo/v4"
)

type Echo struct {
	E *echo.Echo
}

// InitialRouter :
func (e *Echo) InitialRouter() {
	timeoutContext := time.Duration(setting.FileConfigSetting.Server.ReadTimeout) * time.Second

	/*sa user*/
	repoSaUser := _sauserrepo.NewRepoSaUser(connection.Conn)
	useSaUser := _sauseruse.NewUseSaUser(repoSaUser, timeoutContext)
	_sausercont.NewContSaUser(e.E, useSaUser)
}
