package isamenu

import (
	"context"
	"property/framework/models"
	sa_models "property/framework/models/sa"
)

// Repository :
type Repository interface {
	GetBySaMenu(ctx context.Context, ID int) (sa_models.SaMenu, error)
	// GetList(ctx context.Context, queryparam models.ParamList) ([]*sa_models.SaMenu, error)
	GetList(ctx context.Context, LevelNo int, ParentMenuID int) ([]*sa_models.SaMenu, error)
	CreateSaMenu(ctx context.Context, menuData *sa_models.SaMenu) error
	UpdateSaMenu(ctx context.Context, menuID int, menuData interface{}) error
	DeleteSaMenu(ctx context.Context, ID int) error
	CountMenuList(ctx context.Context, queryparam models.ParamList) (int, error)
}

// UseCase :
type UseCase interface {
	GetBySaMenu(ctx context.Context, ID int) (sa_models.SaMenu, error)
	GetList(ctx context.Context, LevelNo int, ParentMenuID int) ([]*sa_models.SaMenu, error)
	CreateSaMenu(ctx context.Context, menuData *sa_models.SaMenu) error
	UpdateSaMenu(ctx context.Context, menuID int, menuData interface{}) error
	DeleteSaMenu(ctx context.Context, ID int) error
}
