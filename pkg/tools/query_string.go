package tool

import (
	"fmt"
	"property/framework/models"
	queryoption "property/framework/query/option"
	"strings"
)

func QueryFunction(SpName string, ParamList []models.ParamFunction) string {
	var (
		result string
		sparam string
	)

	for _, Key := range ParamList {
		sparam += ":" + Key.ParameterName + ","
	}
	if last := len(sparam) - 1; last >= 0 && sparam[last] == ',' {
		sparam = sparam[:last]
	}

	result = queryoption.QueryExecCUD

	result = strings.Replace(result, "{FunctionName}", SpName, -1)
	result = strings.Replace(result, "{ParameterFunction}", sparam, -1)

	return result
}
func QueryJson(SpName string) string {
	var result string

	result = queryoption.QueryExecCUD
	result = strings.Replace(result, "{FunctionName}", SpName, -1)
	result = strings.Replace(result, "{ParameterFunction}", ":in_data", -1)
	return result
}
func QueryFunctionByID(SpName string, ParamList []models.ParamFunction) string {
	var (
		result string
		sparam string
	)

	for _, Key := range ParamList {
		sparam += ":" + Key.ParameterName + ","
	}
	if last := len(sparam) - 1; last >= 0 && sparam[last] == ',' {
		sparam = sparam[:last]
	}

	result = queryoption.QueryGetByID

	result = strings.Replace(result, "{FunctionName}", SpName, -1)
	result = strings.Replace(result, "{ParameterFunction}", sparam, -1)

	return result
}

func QueryFunctionList(SourceFrom string, sSortField string, sField string, sWhere string) string {
	var (
		result string
		// sparam string
	)
	if sField == "" {
		sField = "*"
	}
	// iOffset := (iStart * iPageSize) - iPageSize

	result = queryoption.QueryList
	result = strings.Replace(result, "{sTable}", SourceFrom, -1)
	result = strings.Replace(result, "{sSortFiled}", sSortField, -1)
	result = strings.Replace(result, "{sField}", sField, -1)
	result = strings.Replace(result, "{sWhere}", sWhere, -1)
	// result = strings.Replace(result, "{iPageSize}", strconv.Itoa(iPageSize), -1)
	// result = strings.Replace(result, "{iOffSet}", strconv.Itoa(iOffset), -1)
	fmt.Printf(result)
	return result
}

func QueryFunctionLookUp(SourceFrom string, Parameter string, Field string, Limit string) string {
	var (
		result string
		// sparam string
	)
	if Field == "" {
		Field = "*"
	}
	result = fmt.Sprintf("SELECT %s FROM %s %s", Field, SourceFrom, Parameter)
	if Limit != "" {
		result += fmt.Sprintf(" LIMIT %s; ", Limit)
	}
	return result
}

// func
