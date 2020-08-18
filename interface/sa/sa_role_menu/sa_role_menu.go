package isarolemenu

import (
	"context"
	sa_models "property/framework/models/sa"

	uuid "github.com/satori/go.uuid"
)

// Repository :
type Repository interface {
	GetMenuRole(ctx context.Context, RoleID uuid.UUID) (result []*sa_models.MenuRole, err error)
	CreateSaRoleMenu(ctx context.Context, clientData *sa_models.SaRoleMenu) (err error)
	UpdateSaRoleMenu(ctx context.Context, clientData *sa_models.SaRoleMenu) (err error)
	DeleteSaRoleMenu(ctx context.Context, userID uuid.UUID) (err error)
}
