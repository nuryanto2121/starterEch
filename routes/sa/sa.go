package saroutes

import (
	"time"

	"github.com/labstack/echo/v4"

	_sausercont "property/framework/controllers/sa/sa_user"
	"property/framework/pkg/connection"
	_sauserrepo "property/framework/repository/sa/sa_user"
	_sauseruse "property/framework/usecase/sa/sa_user"

	_sausercompany "property/framework/repository/sa/sa_user_company"
	// _sauserusecompany "property/framework/usecase/sa/sa_user_company"

	_sauserbranch "property/framework/repository/sa/sa_user_branch"

	_sarolecont "property/framework/controllers/sa/sa_role"
	_sarolerepo "property/framework/repository/sa/sa_role"
	_saroleuse "property/framework/usecase/sa/sa_role"

	_saauthcont "property/framework/controllers/auth"
	_authuse "property/framework/usecase/auth"

	_sacompanyrepo "property/framework/repository/sa/sa_company"

	_sabranchrepo "property/framework/repository/sa/sa_branch"

	_saclientrepo "property/framework/repository/sa/sa_client"
	_saclientuse "property/framework/usecase/sa/sa_client"

	_safilecont "property/framework/controllers/fileupload"
	_safilerepo "property/framework/repository/sa/sa_file_upload"
	_safilieuse "property/framework/usecase/sa/sa_file_upload"

	_samenucont "property/framework/controllers/sa/sa_menu"
	_samenurepo "property/framework/repository/sa/sa_menu"
	_samenuuse "property/framework/usecase/sa/sa_menu"
)

func Router(e *echo.Echo, timeoutContext time.Duration) {
	//file upload
	repoSaFileUpload := _safilerepo.NewRepoSaFileUpload(connection.Conn)
	useSaFileUpload := _safilieuse.NewUseSaFileUpload(repoSaFileUpload, timeoutContext)
	_safilecont.NewContFileUpload(e, useSaFileUpload)

	/*sa user branch*/
	repoSaUserBranch := _sauserbranch.NewRepoSaUserBranch(connection.Conn)
	/*sa user company*/
	repoSaUserCompany := _sausercompany.NewRepoSaUserCompany(connection.Conn)
	// useSaUserCompany := _sauserusecompany.NewUseSaUserCompany(repoSaUserCompany, timeoutContext)

	/*sa user*/
	repoSaUser := _sauserrepo.NewRepoSaUser(connection.Conn)
	useSaUser := _sauseruse.NewUseSaUser(repoSaUser, repoSaUserCompany, repoSaUserBranch, useSaFileUpload, timeoutContext)
	_sausercont.NewContSaUser(e, useSaUser)

	/*sa Role*/
	repoSaRole := _sarolerepo.NewRepoSaRole(connection.Conn)
	useSaRole := _saroleuse.NewUseSaRole(repoSaRole, timeoutContext)
	_sarolecont.NewContSaRole(e, useSaRole)

	/*sa Menu*/
	repoSaMenu := _samenurepo.NewRepoSaMenu(connection.Conn)
	useSaMenu := _samenuuse.NewUseSaMenu(repoSaMenu, timeoutContext)
	_samenucont.NewContSaMenu(e, useSaMenu)

	/*sa Company*/
	repoSaCompany := _sacompanyrepo.NewRepoSaCompany(connection.Conn)

	/*sa Branch*/
	repoSaBranch := _sabranchrepo.NewRepoSaBranch(connection.Conn)

	/*sa Client*/
	repoSaClient := _saclientrepo.NewRepoSaClient(connection.Conn)
	useSaClient := _saclientuse.NewUseClient(repoSaClient, repoSaCompany, repoSaUser, repoSaBranch, timeoutContext)

	//_saauthcont
	useAuth := _authuse.NewUserAuth(repoSaUser, useSaClient, useSaFileUpload, timeoutContext)
	_saauthcont.NewContAuth(e, useAuth)
}
