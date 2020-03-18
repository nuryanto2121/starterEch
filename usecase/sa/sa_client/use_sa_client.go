package usesaclient

import (
	"context"
	isabranch "property/framework/interface/sa/sa_branch"
	isaclient "property/framework/interface/sa/sa_client"
	isacompany "property/framework/interface/sa/sa_company"
	isauser "property/framework/interface/sa/sa_user"
	sa_models "property/framework/models/sa"
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
	companyData, err := regisCompany(clientData)
	err = u.repoCompany.CreateSaCompany(ctx, &companyData)

	// create Branch

	// Create User
	return nil
}

// regisCompany :
func regisCompany(clientData *sa_models.SaClient) (sa_models.SaCompany, error) {
	companyData := sa_models.SaCompany{}
	// companyData.ClientID = clientData.ClientID
	// companyData.CompanyName = clientData.ClientName
	// companyData.Address = clientData.Address
	// companyData.EmailAddr = clientData.EmailAddr
	// companyData.ContactPerson = clientData.ContactPerson
	// companyData.CreatedBy = clientData.CreatedBy
	// companyData.UpdatedBy = clientData.UpdatedBy
	err := mapstructure.Decode(clientData, &companyData)
	if err != nil {
		return sa_models.SaCompany{}, err
	}
	companyData.CompanyName = "Company " + clientData.ClientName

	return companyData, nil
}
