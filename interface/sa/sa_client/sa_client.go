package isaclient

import (
	"context"
	"property/framework/models"
	sa_models "property/framework/models/sa"

	uuid "github.com/satori/go.uuid"
)

// Repository :
type Repository interface {
	GetBySaClient(ctx context.Context, clientID uuid.UUID) (result sa_models.SaClient, err error)
	GetList(ctx context.Context, queryparam models.ParamList) (result []*sa_models.SaClient, err error)
	CreateSaClient(ctx context.Context, clientData *sa_models.SaClient) (err error)
	UpdateSaClient(ctx context.Context, clientData *sa_models.SaClient) (err error)
	DeleteSaClient(ctx context.Context, clientID uuid.UUID) (err error)
	CountClientList(ctx context.Context, queryparam models.ParamList) (result int, err error)
}

// Clientcase :
type Usecase interface {
	// GetBySaClient(ctx context.Context, clientID uuid.UUID) (result sa_models.SaClient, err error)
	// GetByEmailSaClient(ctx context.Context, email string) (result sa_models.SaClient, err error)
	// GetList(ctx context.Context, queryparam models.ParamList) (result models.ResponseModelList, err error)
	RegisterClient(ctx context.Context, clientData *sa_models.SaClient) (err error)
	// LoginClient(ctx context.Context, clientData *sa_models.SaClient) (err error)
	// UpdateSaClient(ctx context.Context, clientData *sa_models.SaClient) (err error)
	// DeleteSaClient(ctx context.Context, clientID uuid.UUID) (err error)
}
