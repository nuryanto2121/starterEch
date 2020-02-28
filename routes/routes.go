package routes

import (
	"property/framework/pkg/connection"
	"property/framework/pkg/setting"

	//format nama table + package
	_sausercont "property/framework/controllers/sa/sa_user"
	_sauserrepo "property/framework/repository/sa/sa_user"
	_sauseruse "property/framework/usecase/sa/sa_user"

	_sagroupcont "property/framework/controllers/sa/sa_group"
	_sagrouprepo "property/framework/repository/sa/sa_group"
	_sagroupuse "property/framework/usecase/sa/sa_group"

	"time"

	"github.com/labstack/echo/v4"
)

// Echo :
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

	/*sa Group*/
	repoSaGroup := _sagrouprepo.NewRepoSaGroup(connection.Conn)
	useSaGroup := _sagroupuse.NewUseSaGroup(repoSaGroup, timeoutContext)
	_sagroupcont.NewContSaGroup(e.E, useSaGroup)
}
