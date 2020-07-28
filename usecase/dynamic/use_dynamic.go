package usedynamic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	idynamic "property/framework/interface/dynamic"
	"property/framework/models"
	tool "property/framework/pkg/tools"
	util "property/framework/pkg/utils"
	"strconv"
	"strings"
	"time"
)

type useOptionTemplate struct {
	repoOption     idynamic.Repository
	contextTimeOut time.Duration
	claims         util.Claims
}

func NewUserSysUser(a idynamic.Repository, timeout time.Duration) idynamic.Usecase {
	return &useOptionTemplate{repoOption: a, contextTimeOut: timeout}
}

func (u *useOptionTemplate) Execute(ctx context.Context, claims util.Claims, data map[string]interface{}) (result interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	// parameter wajib
	OptionUrl := fmt.Sprintf("%v", data["option_url"])
	Method := fmt.Sprintf("%v", data["method"])
	LineNo, err := strconv.Atoi(fmt.Sprintf("%v", data["line_no"])) //data["line_no"].(int)
	if err != nil {
		return nil, err
	}

	if _, ok := data["option_url"]; ok {
		delete(data, "option_url")
	}
	if _, ok := data["method"]; ok {
		delete(data, "method")
	}
	if _, ok := data["line_no"]; ok {
		delete(data, "line_no")
	}

	OptionDbList, err := u.repoOption.GetOptionByUrl(ctx, OptionUrl)
	if err != nil {
		return nil, err
	}
	var DataOption = tool.FilterOptionList(OptionDbList, LineNo, Method)[0]
	fmt.Printf("%v", DataOption)

	SpName := DataOption.SP

	DataParameter, err := u.repoOption.GetParamFunction(ctx, SpName)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v", DataParameter)
	DataPostSP, err := tool.SetParameterSP(DataParameter, data, claims)
	if err != nil {
		return nil, err
	}

	sQuery := tool.QueryFunction(SpName, DataParameter)
	fmt.Printf(sQuery)
	resultPost, err := u.repoOption.CRUD(ctx, sQuery, DataPostSP)
	if err != nil {
		return nil, err
	}

	return resultPost, nil
}

func (u *useOptionTemplate) ExecuteMulti(ctx context.Context, claims util.Claims, data models.PostMulti, method string) (result interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	// parameter wajib
	OptionUrl := data.MenuUrl
	Method := method
	LineNo := data.LineNo

	OptionDbList, err := u.repoOption.GetOptionByUrl(ctx, OptionUrl)
	if err != nil {
		return nil, err
	}
	var DataOption = tool.FilterOptionList(OptionDbList, LineNo, Method)[0]
	fmt.Printf("%v", DataOption)

	SpName := DataOption.SP

	DataParameter, err := u.repoOption.GetParamFunction(ctx, SpName)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v", DataParameter)

	dataString, err := json.Marshal(data.InData)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(dataString))

	DataPostSP := make(map[string]interface{})
	DataPostSP["in_data"] = string(dataString)

	sQuery := tool.QueryJson(SpName)
	fmt.Printf(sQuery)
	resultPost, err := u.repoOption.CRUD(ctx, sQuery, DataPostSP)
	if err != nil {
		return nil, err
	}

	return resultPost, nil
}

func (u *useOptionTemplate) Delete(ctx context.Context, claims util.Claims, ParamGet models.ParamGet) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	OptionDbList, err := u.repoOption.GetOptionByUrl(ctx, ParamGet.MenuUrl)
	if err != nil {
		return err
	}
	var DataOption = tool.FilterOptionList(OptionDbList, ParamGet.LineNo, "DELETE")[0]
	fmt.Printf("%v", DataOption)
	SpName := DataOption.SP

	DataParameter, err := u.repoOption.GetParamFunction(ctx, SpName)
	if err != nil {
		return err
	}
	fmt.Printf("%v", DataParameter)
	DataPostSP := make(map[string]interface{}, 0)
	DataPostSP["p_row_id"] = ParamGet.ID
	DataPostSP["p_lastupdatestamp"] = ParamGet.Lastupdatestamp

	sQuery := tool.QueryFunction(SpName, DataParameter)
	fmt.Printf(sQuery)
	resultPost, err := u.repoOption.CRUD(ctx, sQuery, DataPostSP)
	if err != nil {
		return err
	}
	fmt.Printf("%v", resultPost)
	return nil
}
func (u *useOptionTemplate) GetDataBy(ctx context.Context, claims util.Claims, ParamGet models.ParamGet) (result interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	OptionDbList, err := u.repoOption.GetOptionByUrl(ctx, ParamGet.MenuUrl)
	if err != nil {
		return nil, err
	}
	var DataOption = tool.FilterOptionList(OptionDbList, ParamGet.LineNo, "GETBYID")[0]
	fmt.Printf("%v", DataOption)
	SpName := DataOption.SP

	DataParameter, err := u.repoOption.GetParamFunction(ctx, SpName)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v", DataParameter)
	DataPostSP := make(map[string]interface{}, 0)
	DataPostSP["p_row_id"] = ParamGet.ID
	DataPostSP["p_lastupdatestamp"] = ParamGet.Lastupdatestamp

	sQuery := tool.QueryFunctionByID(SpName, DataParameter)
	fmt.Printf(sQuery)
	resultPost, err := u.repoOption.CRUD(ctx, sQuery, DataPostSP)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v", resultPost)

	return resultPost, nil
}

func (u *useOptionTemplate) GetList(ctx context.Context, claims util.Claims, queryparam models.ParamDynamicList) (result models.ResponseModelList, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	var (
		iStart         int
		iPerpage       int
		isViewFunction bool
		ViewName       string
		FieldList      []models.ParamFunction
		DefineColumns  string
		// DefineColumnFormat string
	)
	OptionDbList, err := u.repoOption.GetOptionByUrl(ctx, queryparam.MenuUrl)
	if err != nil {
		return result, err
	}

	iStart = queryparam.Page
	iPerpage = queryparam.PerPage
	isViewFunction = queryparam.ParamView != ""
	MenuUrl := queryparam.MenuUrl
	LineNo := queryparam.LineNo
	ParamWhere := queryparam.Search
	InitialWhere := queryparam.InitSearch
	sSortField := queryparam.SortField
	if sSortField == "" {
		sSortField = "ORDER BY updated_at desc"
	} else {
		sSortField = "ORDER BY " + sSortField
	}

	var DataOption = tool.FilterOptionList(OptionDbList, LineNo, "LIST")[0]
	fmt.Printf("%v", DataOption)
	ViewName = DataOption.SP

	DefineColumn, err := u.repoOption.GetDefineColumn(ctx, MenuUrl, LineNo)
	if err != nil {
		return result, err
	}

	FieldList, err = u.repoOption.GetFieldType(ctx, ViewName, isViewFunction)
	if err != nil {
		return result, err
	}
	AllColumnQuery, _, DefineSize, FieldWhere := tool.SetFieldList(FieldList, DefineColumn.ColumnField, 20, true)

	_, AllColumn, _, _ := tool.SetFieldList(FieldList, DefineColumn.ColumnField, 0, true)

	if DefineColumn.ColumnField != "" {
		DefineColumns = DefineColumn.ColumnField
		// DefineColumnFormat =
	} else {
		SpName := "fss_define_column_i"
		DefineColumns = "no," + AllColumn
		DataParameter, err := u.repoOption.GetParamFunction(ctx, SpName)
		if err != nil {
			return result, err
		}
		fmt.Printf("%v", DataParameter)
		DataPostSP := make(map[string]interface{}, 0)
		DataPostSP["p_option_url"] = MenuUrl
		DataPostSP["p_line_no"] = LineNo
		DataPostSP["p_column_field"] = DefineColumns
		DataPostSP["p_created_by"] = claims.UserName

		sQuery := tool.QueryFunctionByID(SpName, DataParameter)
		fmt.Printf(sQuery)
		_, err = u.repoOption.CRUD(ctx, sQuery, DataPostSP)
		if err != nil {
			return result, err
		}
	}

	if InitialWhere != "" {
		InitialWhere = "WHERE " + InitialWhere
	}
	sWhere := strings.Replace(InitialWhere, "claims.user_id", claims.UserID, -1)
	// sWhereLike := ""
	// if ParamWhere != "" {
	// 	sWhereLike = tool.SetWhereLikeList(FieldWhere, ParamWhere)
	// }

	if ParamWhere != "" {
		sWhereLike := tool.SetWhereLikeList(FieldWhere, ParamWhere)
		if sWhere != "" {
			sWhere += " AND " + sWhereLike
		} else {
			sWhere += " WHERE " + sWhereLike
		}
	}

	if queryparam.ParamView != "" {
		ViewName = fmt.Sprintf("%s(%s)", ViewName, queryparam.ParamView)
	}
	iOffset := (iStart * iPerpage) - iPerpage
	// DataList := make(map[string]interface{}, 0)
	// DataList["Limit"] = iPerpage
	// DataList["Offset"] = iOffset

	sQuery := tool.QueryFunctionList(ViewName, sSortField, AllColumnQuery, sWhere)
	result.Data, err = u.repoOption.GetDataList(ctx, sQuery, iPerpage, iOffset)
	if err != nil {
		return result, err
	}

	result.Total, err = u.repoOption.CountList(ctx, ViewName, sWhere)
	if err != nil {
		return result, err
	}
	result.LastPage = int(math.Ceil(float64(result.Total) / float64(queryparam.PerPage)))
	result.Page = queryparam.Page
	result.DefineSize = DefineSize
	result.DefineColumn = DefineColumns
	result.AllColumn = AllColumn
	return result, err
}

func (u *useOptionTemplate) GetDataLookUp(ctx context.Context, claims util.Claims, ParamGet models.ParamLookup) (result interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	var (
		OptionLookup = models.OptionLookup{}
		SpName       string
		InitialWHere = ParamGet.InitSearch
		ParamWhere   = ParamGet.Search
	)
	isViewFunction := ParamGet.ParamView != ""
	fmt.Println(isViewFunction)

	OptionLookup, err = u.repoOption.GetOptionLookupBy(ctx, ParamGet.LookUpCd, ParamGet.ColumnDB)
	if err != nil || OptionLookup.OptionLookUpCD == "" {
		if OptionLookup.OptionLookUpCD == "" {
			return nil, errors.New("Please Contact Your Administrator. (table setting Look up null)")
		}
		return nil, err
	}
	source_field := tool.SetFieldListLookup(OptionLookup.SourceField, OptionLookup.DisplayLookup)

	if isViewFunction {
		SpName = fmt.Sprintf("%s(%s)", OptionLookup.ViewName, ParamGet.ParamView)
	} else {
		SpName = OptionLookup.ViewName

	}

	//WHere
	if InitialWHere != "" {
		InitialWHere = "WHERE " + InitialWHere
	}
	sWhere := strings.Replace(InitialWHere, "claims.user_id", claims.UserID, -1)
	// sWhereLike := tool.SetWhereLikeList(OptionLookup.SourceField, ParamWhere)
	if ParamWhere != "" {
		sWhereLike := tool.SetWhereLikeList(OptionLookup.SourceField, ParamWhere)
		if sWhere != "" {
			sWhere += " AND (" + sWhereLike + ") "
		} else {
			sWhere += " WHERE (" + sWhereLike + ") "
		}
	}
	sQuery := tool.QueryFunctionLookUp(SpName, sWhere, source_field, "10000")
	result, err = u.repoOption.GetDataQuery(ctx, sQuery)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (u *useOptionTemplate) GetDataLookUpList(ctx context.Context, claims util.Claims, queryparam models.ParamLookUpList) (result models.ResponseModelListLookUp, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	var (
		iStart         int
		iPerpage       int
		isViewFunction bool
		ViewName       string
		FieldList      []models.ParamFunction
		OptionLookup   = models.OptionLookup{}
		// DefineColumnFormat string
	)

	iStart = queryparam.Page
	iPerpage = queryparam.PerPage
	isViewFunction = queryparam.ParamView != ""
	ParamWhere := queryparam.Search
	InitialWhere := queryparam.InitSearch
	sSortField := queryparam.SortField

	OptionLookup, err = u.repoOption.GetOptionLookupBy(ctx, queryparam.LookUpCd, queryparam.ColumnDB)
	if err != nil || OptionLookup.OptionLookUpCD == "" {
		if OptionLookup.OptionLookUpCD == "" {
			return result, errors.New("Please Contact Your Administrator. (table setting Look up null)")
		}
		return result, err
	}

	if sSortField == "" {
		sSortField = fmt.Sprintf("ORDER BY %s desc", strings.Split(OptionLookup.SourceField, ",")[0])
	} else {
		sSortField = "ORDER BY " + sSortField
	}

	ViewName = OptionLookup.ViewName

	DefineColumn := OptionLookup.SourceField

	FieldList, err = u.repoOption.GetFieldType(ctx, ViewName, isViewFunction)
	if err != nil {
		return result, err
	}
	AllColumnQuery, _, _, FieldWhere := tool.SetFieldList(FieldList, DefineColumn, 20, false)

	AllColumn := DefineColumn

	if InitialWhere != "" {
		InitialWhere = "WHERE " + InitialWhere
	}
	sWhere := strings.Replace(InitialWhere, "claims.user_id", claims.UserID, -1)
	// sWhereLike := ""
	// if ParamWhere != "" {
	// 	sWhereLike = tool.SetWhereLikeList(FieldWhere, ParamWhere)
	// }

	if ParamWhere != "" {
		sWhereLike := tool.SetWhereLikeList(FieldWhere, ParamWhere)
		if sWhere != "" {
			sWhere += " AND " + sWhereLike
		} else {
			sWhere += " WHERE " + sWhereLike
		}
	}

	if queryparam.ParamView != "" {
		ViewName = fmt.Sprintf("%s(%s)", ViewName, queryparam.ParamView)
	}
	iOffset := (iStart * iPerpage) - iPerpage
	// DataList := make(map[string]interface{}, 0)
	// DataList["Limit"] = iPerpage
	// DataList["Offset"] = iOffset

	sQuery := tool.QueryFunctionList(ViewName, sSortField, AllColumnQuery, sWhere)
	result.Data, err = u.repoOption.GetDataList(ctx, sQuery, iPerpage, iOffset)
	if err != nil {
		return result, err
	}

	result.Total, err = u.repoOption.CountList(ctx, ViewName, sWhere)
	if err != nil {
		return result, err
	}
	result.LastPage = int(math.Ceil(float64(result.Total) / float64(queryparam.PerPage)))
	result.Page = queryparam.Page
	result.AllColumn = AllColumn

	return result, nil
}
