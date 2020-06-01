package contfileupload

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	isafileupload "property/framework/interface/sa/sa_file_upload"
	"property/framework/pkg/setting"
	util "property/framework/pkg/utils"

	"github.com/fsetiawan29/healthy_food/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ContFileUpload :
type ContFileUpload struct {
	useSaFileUpload isafileupload.UseCase
}

func NewContFileUpload(e *echo.Echo, useSaFileUpload isafileupload.UseCase) {
	cont := &ContFileUpload{
		useSaFileUpload: useSaFileUpload,
	}

	e.Static("/upload", "wwwroot")
	r := e.Group("/fileupload")
	// Configure middleware with custom claims
	var screet = setting.FileConfigSetting.App.JwtSecret
	config := middleware.JWTConfig{
		Claims:     &util.Claims{},
		SigningKey: []byte(screet),
	}

	r.Use(middleware.JWTWithConfig(config))
	r.POST("", cont.CreateImage)

}
func (i *ContFileUpload) CreateImage(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	images := form.File["images"]

	var result []string

	for _, image := range images {
		// Source
		src, err := image.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		if _, err := os.Stat("./upload"); os.IsNotExist(err) {
			err = os.Mkdir("./upload", os.ModePerm)
		}

		fileNameAndUnix := fmt.Sprintf("%d_%s", util.GetTimeNow().Unix(), image.Filename)

		// Destination
		dst, err := os.Create(fmt.Sprintf("./upload/%s", fileNameAndUnix))
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		r := c.Request()
		fileName := fmt.Sprintf("%s://%s/upload/%s", c.Scheme(), r.Host, fileNameAndUnix)
		// err = i.imageUsecase.CreateImage(ctx, &models.Image{
		// 	Name: fileName,
		// })
		// if err != nil {
		// 	return err
		// }
		result = append(result, fileName)

	}

	return c.JSON(http.StatusCreated, models.ResponseImage(http.StatusCreated, result))
}
