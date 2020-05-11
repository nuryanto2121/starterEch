package usesaclient

import (
	"context"
	isabranch "property/framework/interface/sa/sa_branch"
	isaclient "property/framework/interface/sa/sa_client"
	isacompany "property/framework/interface/sa/sa_company"
	isauser "property/framework/interface/sa/sa_user"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	util "property/framework/pkg/utils"
	"time"

	"github.com/mitchellh/mapstructure"
)

type useClient struct {
	repoClient     isaclient.Repository
	repoCompany    isacompany.Repository
	repoUser       isauser.Repository
	repoBranch     isabranch.Repository
	contextTimeOut time.Duration
}

// NewUseClient :
func NewUseClient(cl isaclient.Repository, co isacompany.Repository, us isauser.Repository, br isabranch.Repository, timeout time.Duration) isaclient.Usecase {
	return &useClient{
		repoClient:     cl,
		repoCompany:    co,
		repoUser:       us,
		repoBranch:     br,
		contextTimeOut: timeout,
	}
}

func (u *useClient) RegisterClient(ctx context.Context, clientData *sa_models.SaClient) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	// Create Client
	clientData.UpdatedBy = clientData.CreatedBy
	// clientData.CreatedAt = util.GetTimeNow()
	// clientData.UpdatedAt = util.GetTimeNow()
	err = u.repoClient.CreateSaClient(ctx, clientData)
	if err != nil {
		return err
	}

	// Create COmpany
	cntCompany, err := u.repoCompany.CountCompanyList(ctx, models.ParamList{})
	if err != nil {
		return err
	}
	companyData, err := regisCompany(clientData, cntCompany)
	if err != nil {
		return err
	}
	err = u.repoCompany.CreateSaCompany(ctx, &companyData)
	if err != nil {
		return err
	}

	// create Branch
	cntBranch, err := u.repoBranch.CountBranchList(ctx, models.ParamList{})
	if err != nil {
		return err
	}
	branchData, err := regisBranch(&companyData, cntBranch)
	if err != nil {
		return err
	}
	err = u.repoBranch.CreateSaBranch(ctx, &branchData)
	if err != nil {
		return err
	}

	// Create User

	return nil
}

// regisCompany :
func regisCompany(clientData *sa_models.SaClient, cntCompany int) (sa_models.SaCompany, error) {
	companyData := sa_models.SaCompany{}
	err := mapstructure.Decode(clientData, &companyData)
	if err != nil {
		return sa_models.SaCompany{}, err
	}
	companyData.CompanyName = "Company_" + util.StrTo(cntCompany).String() + "_" + clientData.ClientName
	companyData.StartDate = util.GetTimeNow()
	companyData.FinYear = int16(util.GetTimeNow().Year())
	companyData.FinPeriod = int16(util.GetTimeNow().Month())

	return companyData, nil
}

// regisBranch :
func regisBranch(companyData *sa_models.SaCompany, cntBranch int) (sa_models.SaBranch, error) {
	branchData := sa_models.SaBranch{}
	err := mapstructure.Decode(companyData, &branchData)
	if err != nil {
		return sa_models.SaBranch{}, err
	}
	branchData.BranchName = "Branch_" + util.StrTo(cntBranch).String() + "_" + util.StrTo(companyData.CompanyID).String()
	branchData.StartDate = util.GetTimeNow()

	return branchData, nil
}
