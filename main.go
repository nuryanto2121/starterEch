package main

import (
	"fmt"
	"log"
	"reflect"

	_midd "property/framework/middleware"
	"property/framework/models"
	"property/framework/pkg/connection"
	"property/framework/pkg/logging"
	"property/framework/pkg/setting"
	"property/framework/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	setting.Setup()
	logging.Setup()
	connection.Setup()
}

// @title Starter
// @version 1.0
// @description Backend REST API for golang starter

// @contact.name Nuryanto
// @contact.url https://www.linkedin.com/in/nuryanto-1b2721156/
// @contact.email nuryantofattih@gmail.com

//// @securityDefinitions.apikey ApiKeyAuth
//// @in header
//// @name Authorization

func main() {
	e := echo.New()
	middL := _midd.InitMiddleware()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "runtime")
	e.Use(middL.CORS)

	app := routes.Echo{E: e}

	app.InitialRouter()

	sPort := fmt.Sprintf(":%d", setting.FileConfigSetting.Server.HTTPPort)
	ee := models.SaUser{}

	dd := reflect.ValueOf(ee)
	ff := reflect.TypeOf(&ee)
	// tes1(ee)
	tes2(dd, ff)
	// maxHeaderBytes := 1 << 20
	// s := &http.Server{
	// 	Addr:           sPort,
	// 	ReadTimeout:    setting.FileConfigSetting.Server.ReadTimeout,
	// 	WriteTimeout:   setting.FileConfigSetting.Server.WriteTimeout,
	// 	MaxHeaderBytes: maxHeaderBytes,
	// }
	// // e.Logger.Fatal(e.StartServer(s))
	// s.ListenAndServe()

	// log.Fatal(e.StartServer(s))
	log.Fatal(e.Start(sPort))
	// log.Fatal(e.Start(":" + string(setting.FileConfigSetting.Server.HTTPPort)))
}

func tes2(v reflect.Value, x reflect.Type) {
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

// func getStructTag(f reflect.StructField, types string) string {
// 	switch types {
// 	case "int16":
// 		return int16(fmt.Sprintf("%v", f.Tag))
// 	case models.ErrNotFound:
// 		return http.StatusNotFound
// 	case models.ErrConflict:
// 		return http.StatusConflict
// 	default:
// 		return http.StatusInternalServerError
// 	}
// 	return string(f.Tag)
// }

func tes1(data interface{}) {

	e := reflect.ValueOf(data)

	for i := 0; i < e.NumField(); i++ {
		dd := e.Field(i)
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		fmt.Printf("%v %v %v\n", varName, varType, varValue)
		// field, ok := reflect.TypeOf(data).Elem().FieldByName(varName) // not json:name
		fmt.Printf("%v\n", dd)

		// if !ok {
		// 	panic("Field not found")
		// }
	}
	// fmt.Printf("%v\n", data)
	// v := reflect.ValueOf(data)
	// fmt.Printf("%v\n", v)
	// values := make([]interface{}, v.NumField())
	// // fmt.Printf("%v\n", values)
	// // for i := 0; i < v.NumField(); i++ {
	// // 	fmt.Printf("%v\n", v.Field(i))
	// // 	values[i] = v.Field(i).Interface()
	// // 	fmt.Printf("%v\n", v.Field(i).Interface())
	// // }

	// fmt.Println(values)

	// t := v.Type()
	// for i := 0; i < t.NumField(); i++ {
	// 	aa := t.Field(i) // key
	// 	fmt.Printf("%v\n", aa) // key
	// 	fmt.Printf("%v\n", t.Field(i))
	// 	fmt.Printf("%v\n", v.Field(i)) // value
	// 	// if strings.Split(t.Field(i).Tag.Get("json"), ",")[0] == name {
	// 	// }
	// }
}
