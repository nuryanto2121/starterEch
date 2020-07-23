package tool

import (
	"errors"
	"fmt"
	"property/framework/models"
	util "property/framework/pkg/utils"
	"strings"

	"reflect"
)

func SetParameterSP(Parameter []models.ParamFunction, DataPost map[string]interface{}, claims util.Claims) (map[string]interface{}, error) {
	result := make(map[string]interface{}, 0)

	v := reflect.ValueOf(DataPost)
	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			strct := v.MapIndex(key)
			if key.Interface() == "option_url" || key.Interface() == "line_no" {
				continue
			}

			sKey := "p_" + fmt.Sprintf("%v", key.Interface())
			sValue := strct.Interface()
			ParamFunction := FilterParamterList(Parameter, sKey)

			if ParamFunction.ParameterName == "" {
				return nil, errors.New("Post parameter function not valid.")
			}
			result[ParamFunction.ParameterName] = sValue
			fmt.Println(key.Interface(), sValue)
		}
	}

	for _, Key := range Parameter {
		if _, ok := result[Key.ParameterName]; !ok {
			if strings.Contains(Key.ParameterName, "user") || strings.Contains(Key.ParameterName, "created_by") || strings.Contains(Key.ParameterName, "updated_by") {
				result[Key.ParameterName] = claims.UserName
			}
		}
	}

	return result, nil
}

func SetWhereLikeList(FieldWhere string, ParamSearch string) string {
	fields := strings.Split(FieldWhere, ",")
	var result string
	for i := 0; i < len(fields); i++ {
		sField := strings.Split(fields[i], ":")
		if strings.ToLower(sField[0]) == "no" {
			continue
		}
		sField[0] = strings.TrimSpace(sField[0])
		if sField[1] == "T" {
			result += fmt.Sprintf("OR lower(TO_CHAR(%s,'DD/MM/YYYY HH24:MI')) LIKE '%%%s%%' ", sField[0], strings.ToLower(ParamSearch))
		} else {
			result += fmt.Sprintf("OR lower(%s::varchar) LIKE '%%%s%%' ", sField[0], strings.ToLower(ParamSearch))
		}
	}
	i1 := strings.Index(result, `OR`)
	result = result[i1+2:]
	return result
}
