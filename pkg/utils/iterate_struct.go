package util

import (
	"fmt"
	"reflect"
	"strings"
)

// Tes2 :
func Tes2(v reflect.Value, x reflect.Type) {
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {

		// fmt.Printf("%v\n", v.Type().Field(i))
		varName := v.Type().Field(i).Name
		varType := v.Type().Field(i).Type
		varValue := v.Field(i).Interface()
		fmt.Printf("%v %v %v\n", varName, varType, varValue)
		fmt.Printf("%v\n", t.Field(i))
		field, _ := x.Elem().FieldByName(fmt.Sprintf("%v", varName))
		fmt.Printf("%v\n", field.Tag)
		// fmt.Println(getStructTag(field, fmt.Sprintf("%v", varType)))
		// if strings.Split(t.Field(i).Tag.Get("json"), ",")[0] == name {
		// 	fmt.Printf("the value is %q\n", v.Field(i).Interface().(string))
		// }
	}
}

func getWhereLikeStruct(v reflect.Value, t reflect.Type, searchParam string) string {
	result := ""
	vt := v.Type()
	for i := 0; i < vt.NumField(); i++ {
		varName := fmt.Sprintf("%v", v.Type().Field(i).Name) //field Name
		varType := v.Type().Field(i).Type //fmt.Sprintf("%v", v.Type().Field(i).Type) // field type data
		// varValue := fmt.Sprintf("%v", v.Field(i).Interface()) //
		field, _ := t.Elem().FieldByName(fmt.Sprintf("%v", varName)) // getTag json
		varTagJson := fmt.Sprintf("%v", field.Tag)                   //get value json

		i1 := strings.Index(varTagJson, `"`)
		str1 := varTagJson[i1+1:]

		i2 := strings.Index(str1, `"`)
		str2 := str1[:i2]
		varFieldtable := fmt.Sprintln(str2)
		switch varType {
			case int16
		}
	}

	return result
}
