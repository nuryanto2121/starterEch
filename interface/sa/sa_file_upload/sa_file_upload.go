package isafileupload

import (
	"context"
	"property/framework/models"
	sa_models "property/framework/models/sa"

	uuid "github.com/satori/go.uuid"
)

// Repository :
type Repository interface {
	GetBySaFileUpload(ctx context.Context, fileID uuid.UUID) (sa_models.SaFileUpload, error)
	GetList(ctx context.Context, queryparam models.ParamList) ([]*sa_models.SaFileUpload, error)
	CreateSaFileUpload(ctx context.Context, fileData *sa_models.SaFileUpload) error
	UpdateSaFileUpload(ctx context.Context, fileData *sa_models.SaFileUpload) error
	DeleteSaFileUpload(ctx context.Context, fileID uuid.UUID) error
	CountFileUploadList(ctx context.Context, queryparam models.ParamList) (int, error)
}

// UseCase :
type UseCase interface {
	GetBySaFileUpload(ctx context.Context, fileID uuid.UUID) (sa_models.SaFileUpload, error)
	GetList(ctx context.Context, queryparam models.ParamList) (models.ResponseModelList, error)
	CreateSaFileUpload(ctx context.Context, fileData *sa_models.SaFileUpload) error
	UpdateSaFileUpload(ctx context.Context, fileData *sa_models.SaFileUpload) error
	DeleteSaFileUpload(ctx context.Context, fileID uuid.UUID) error
}
