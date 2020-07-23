package middleware

import (
	"fmt"
	"net/http"
	"property/framework/pkg/app"
	sqlxposgresdb "property/framework/pkg/postgresqlxdb"
	"property/framework/pkg/setting"
	util "property/framework/pkg/utils"
	"property/framework/redisdb"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		var (
			code  = http.StatusOK
			msg   = ""
			data  interface{}
			token = e.Request().Header.Get("Authorization")
		)
		data = map[string]string{
			"token": token,
		}

		if token == "" {
			code = http.StatusNetworkAuthenticationRequired
			msg = "Auth Token Required"
		} else {
			existToken := redisdb.GetSession(token)
			if existToken == "" {
				code = http.StatusUnauthorized
				msg = "Token Failed"
			}
			claims, err := util.ParseToken(token)
			if err != nil {
				code = http.StatusUnauthorized
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					msg = "Token Expired"
				default:
					msg = "Token Failed"
				}
			} else {
				var issuer = setting.FileConfigSetting.App.Issuer
				valid := claims.VerifyIssuer(issuer, true)
				if !valid {
					code = http.StatusUnauthorized
					msg = "Issuer is not valid"
				}
				e.Set("claims", claims)
			}
		}

		if code != http.StatusOK {
			resp := app.ResponseModel{
				Msg:  msg,
				Data: data,
			}
			return e.JSON(code, resp)

			// return nil
		}
		return next(e)
	}
}
func Versioning(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		var (
			OS    = e.Request().Header.Get("OS")
			Versi = e.Request().Header.Get("Version")
		)
		Version, err := strconv.Atoi(Versi)

		if Version == 0 {
			resp := app.ResponseModel{
				Msg:  "Please Set Header Version",
				Data: nil,
			}
			return e.JSON(http.StatusBadRequest, resp)
		}
		dataVersion, err := sqlxposgresdb.GetVersion(OS)
		if err != nil {
			resp := app.ResponseModel{
				Msg:  fmt.Sprintf("%v", err),
				Data: nil,
			}
			return e.JSON(http.StatusBadRequest, resp)
		}

		if dataVersion.Version > Version {
			resp := app.ResponseModel{
				Msg:  "Please Update Your Apps",
				Data: dataVersion.Version,
			}
			return e.JSON(http.StatusHTTPVersionNotSupported, resp)
		}
		if dataVersion.Version < Version {
			resp := app.ResponseModel{
				Msg:  "Version Not Support",
				Data: dataVersion.Version,
			}
			return e.JSON(http.StatusHTTPVersionNotSupported, resp)
		}

		//end

		return next(e)
	}
}
