package idynamic

import (
	"context"
	"property/framework/models"
	util "property/framework/pkg/utils"
)

type Usecase interface {
	Execute(ctx context.Context, claims util.Claims, data map[string]interface{}) (result interface{}, err error)
	ExecuteMulti(ctx context.Context, claims util.Claims, data models.PostMulti, method string) (result interface{}, err error)
	Delete(ctx context.Context, claims util.Claims, ParamGet models.ParamGet) error
	GetDataBy(ctx context.Context, claims util.Claims, ParamGet models.ParamGet) (result interface{}, err error)
	GetList(ctx context.Context, claims util.Claims, queryparam models.ParamDynamicList) (result models.ResponseModelList, err error)
	GetDataLookUp(ctx context.Context, claims util.Claims, ParamGet models.ParamLookup) (result interface{}, err error)
	GetDataLookUpList(ctx context.Context, claims util.Claims, ParamGet models.ParamLookUpList) (result models.ResponseModelListLookUp, err error)
}
type Repository interface {
	GetOptionByUrl(ctx context.Context, Url string) (result []models.OptionDB, err error)
	GetOptionLookupBy(ctx context.Context, LookUpCd string, ColumnDB string) (result models.OptionLookup, err error)
	GetParamFunction(ctx context.Context, SpName string) (result []models.ParamFunction, err error)
	CRUD(ctx context.Context, sQuery string, data interface{}) (result interface{}, err error)
	GetDataList(ctx context.Context, sQuery string, Limit int, Offset int) (result interface{}, err error)
	GetDataQuery(ctx context.Context, sQuery string) (result interface{}, err error)
	GetDefineColumn(ctx context.Context, MenuUrl string, LineNo int) (result models.DefineColumn, err error)
	GetFieldType(ctx context.Context, SourceFrom string, isViewFunction bool) (result []models.ParamFunction, err error)
	CountList(ctx context.Context, ViewName string, sWhere string) (int, error)
}
