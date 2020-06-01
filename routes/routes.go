package routes

import (
	"property/framework/pkg/connection"
	"property/framework/pkg/setting"

	//format nama table + package
	_sausercont "property/framework/controllers/sa/sa_user"
	_sauserrepo "property/framework/repository/sa/sa_user"
	_sauseruse "property/framework/usecase/sa/sa_user"

	_sausercompany "property/framework/repository/sa/sa_user_company"
	// _sauserusecompany "property/framework/usecase/sa/sa_user_company"

	_sauserbranch "property/framework/repository/sa/sa_user_branch"

	_sarolecont "property/framework/controllers/sa/sa_role"
	_sarolerepo "property/framework/repository/sa/sa_role"
	_saroleuse "property/framework/usecase/sa/sa_role"

	_saauthcont "property/framework/controllers/auth"

	_sacompanyrepo "property/framework/repository/sa/sa_company"

	_sabranchrepo "property/framework/repository/sa/sa_branch"

	_saclientrepo "property/framework/repository/sa/sa_client"
	_saclientuse "property/framework/usecase/sa/sa_client"

	_safilecont "property/framework/controllers/fileupload"
	_safilerepo "property/framework/repository/sa/sa_file_upload"
	_safilieuse "property/framework/usecase/sa/sa_file_upload"

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

	/*sa user branch*/
	repoSaUserBranch := _sauserbranch.NewRepoSaUserBranch(connection.Conn)
	/*sa user company*/
	repoSaUserCompany := _sausercompany.NewRepoSaUserCompany(connection.Conn)
	// useSaUserCompany := _sauserusecompany.NewUseSaUserCompany(repoSaUserCompany, timeoutContext)

	/*sa user*/
	repoSaUser := _sauserrepo.NewRepoSaUser(connection.Conn)
	useSaUser := _sauseruse.NewUseSaUser(repoSaUser, repoSaUserCompany, repoSaUserBranch, timeoutContext)
	_sausercont.NewContSaUser(e.E, useSaUser)

	/*sa Role*/
	repoSaRole := _sarolerepo.NewRepoSaRole(connection.Conn)
	useSaRole := _saroleuse.NewUseSaRole(repoSaRole, timeoutContext)
	_sarolecont.NewContSaRole(e.E, useSaRole)

	/*sa Company*/
	repoSaCompany := _sacompanyrepo.NewRepoSaCompany(connection.Conn)

	/*sa Branch*/
	repoSaBranch := _sabranchrepo.NewRepoSaBranch(connection.Conn)

	/*sa Client*/
	repoSaClient := _saclientrepo.NewRepoSaClient(connection.Conn)
	useSaClient := _saclientuse.NewUseClient(repoSaClient, repoSaCompany, repoSaUser, repoSaBranch, timeoutContext)

	//file upload
	repoSaFileUpload := _safilerepo.NewRepoSaFileUpload(connection.Conn)
	useSaFileUpload := _safilieuse.NewUseSaFileUpload(repoSaFileUpload, timeoutContext)
	_safilecont.NewContFileUpload(e.E, useSaFileUpload)

	//_saauthcont
	_saauthcont.NewContAuth(e.E, useSaClient, useSaUser)
}
