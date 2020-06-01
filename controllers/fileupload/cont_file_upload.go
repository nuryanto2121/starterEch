package contfileupload

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	isafileupload "property/framework/interface/sa/sa_file_upload"
	"property/framework/pkg/app"
	"property/framework/pkg/file"
	"property/framework/pkg/logging"
	"property/framework/pkg/setting"
	util "property/framework/pkg/utils"
	"strings"

	sa_models "property/framework/models/sa"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ContFileUpload :
type ContFileUpload struct {
	useSaFileUpload isafileupload.UseCase
}

// NewContFileUpload :
func NewContFileUpload(e *echo.Echo, useSaFileUpload isafileupload.UseCase) {
	cont := &ContFileUpload{
		useSaFileUpload: useSaFileUpload,
	}

	e.Static("/wwwroot", "wwwroot")
	r := e.Group("/api/fileupload")
	// Configure middleware with custom claims
	var screet = setting.FileConfigSetting.App.JwtSecret
	config := middleware.JWTConfig{
		Claims:     &util.Claims{},
		SigningKey: []byte(screet),
	}

	r.Use(middleware.JWTWithConfig(config))
	r.POST("", cont.CreateImage)

}

// CreateImage :
// @Summary File Upload
// @Security ApiKeyAuth
// @Description Upload file
// @Tags FileUpload
// @Accept  multipart/form-data
// @Produce json
// @Param upload_file formData file true "account image"
// @Param path formData string true "path images"
// @Success 200 {object} app.ResponseModel
// @Router /api/fileupload [post]
func (u *ContFileUpload) CreateImage(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		appE          = app.Res{R: c}
		imageFormList []sa_models.SaFileUpload
		logger        = logging.Logger{}
	)

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	images := form.File["upload_file"]

	pt := form.Value["path"]

	logger.Info(pt)
	//directory api
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	var dir_file = dir + "/wwwroot/uploads"
	var path_file = "/wwwroot/uploads"
	err = file.IsNotExistMkDir(dir_file)
	if err != nil {
		return err
	}

	for i, image := range images {
		// Source
		src, err := image.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// create folder directory if not exist from param
		if pt[i] != "" {

			dirx := dir_file
			for _, val := range strings.Split(pt[i], "/") {
				dirx = dirx + "/" + val
				fmt.Printf(dirx)
				err = file.IsNotExistMkDir(dirx)
				if err != nil {
					return err
				}
			}
			dir_file = fmt.Sprintf("%s/%s", dir_file, pt[i])

			path_file = fmt.Sprintf("%s/%s", path_file, pt[i])
		}

		fileNameAndUnix := fmt.Sprintf("%d_%s", util.GetTimeNow().Unix(), image.Filename)

		// Destination
		dest := fmt.Sprintf("%s/%s", dir_file, fileNameAndUnix)
		dst, err := os.Create(dest)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		// r := c.Request()
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*util.Claims)

		var imageForm sa_models.SaFileUpload
		// fileName := fmt.Sprintf("%s://%s/upload/%s", c.Scheme(), r.Host, fileNameAndUnix)
		imageForm.FileName = fileNameAndUnix
		imageForm.FilePath = fmt.Sprintf("%s/%s", path_file, fileNameAndUnix)
		imageForm.FileType = filepath.Ext(fileNameAndUnix)
		imageForm.CreatedBy = claims.UserName
		imageForm.UpdatedBy = claims.UserName
		err = u.useSaFileUpload.CreateSaFileUpload(ctx, &imageForm)
		// err = i.imageUsecase.CreateImage(ctx, &models.Image{
		// 	Name: fileName,
		// })
		if err != nil {
			return err
		}
		imageFormList = append(imageFormList, imageForm)

	}
	return appE.Response(http.StatusOK, "Ok", imageFormList)

	// return c.JSON(http.StatusCreated, models.ResponseImage(http.StatusCreated, imageFormList))
}
