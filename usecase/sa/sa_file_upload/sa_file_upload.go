package usesafileupload

import (
	"context"
	"log"
	"math"
	"os"
	isafileupload "property/framework/interface/sa/sa_file_upload"
	"property/framework/models"
	sa_models "property/framework/models/sa"
	"property/framework/pkg/file"
	util "property/framework/pkg/utils"
	"reflect"
	"time"

	uuid "github.com/satori/go.uuid"
)

type useSaFileUpload struct {
	repoSaFileUpload isafileupload.Repository
	contexTimeOut    time.Duration
}

// NewUseSaFileUpload :
func NewUseSaFileUpload(a isafileupload.Repository, timeout time.Duration) isafileupload.UseCase {
	return &useSaFileUpload{
		repoSaFileUpload: a,
		contexTimeOut:    timeout,
	}
}

func (u *useSaFileUpload) GetBySaFileUpload(ctx context.Context, fileuploadID uuid.UUID) (sa_models.SaFileUpload, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()

	var (
		result = sa_models.SaFileUpload{}
		err    error
	)

	result, err = u.repoSaFileUpload.GetBySaFileUpload(ctx, fileuploadID)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (u *useSaFileUpload) GetList(ctx context.Context, queryparam models.ParamList) (models.ResponseModelList, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()
	var (
		result      = models.ResponseModelList{}
		tfileupload = sa_models.SaFileUpload{}
		err         error
	)

	/*membuat Where like dari struct*/
	if queryparam.Search != "" {
		value := reflect.ValueOf(tfileupload)
		types := reflect.TypeOf(&tfileupload)
		queryparam.Search = util.GetWhereLikeStruct(value, types, queryparam.Search, "")
	}
	result.Data, err = u.repoSaFileUpload.GetList(ctx, queryparam)
	if err != nil {
		return result, err
	}

	result.Total, err = u.repoSaFileUpload.CountFileUploadList(ctx, queryparam)
	if err != nil {
		return result, err
	}

	// d := float64(result.Total) / float64(queryparam.PerPage)
	result.LastPage = int(math.Ceil(float64(result.Total) / float64(queryparam.PerPage)))
	result.Page = queryparam.Page

	return result, nil
}

func (u *useSaFileUpload) CreateSaFileUpload(ctx context.Context, fileuploadData *sa_models.SaFileUpload) error {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()
	var (
		err error
	)

	fileuploadData.UpdatedBy = fileuploadData.CreatedBy
	fileuploadData.CreatedAt = util.GetTimeNow()
	fileuploadData.UpdatedAt = util.GetTimeNow()
	err = u.repoSaFileUpload.CreateSaFileUpload(ctx, fileuploadData)
	if err != nil {
		return err
	}
	return nil
}

func (u *useSaFileUpload) UpdateSaFileUpload(ctx context.Context, fileuploadData *sa_models.SaFileUpload) error {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()
	var (
		err error
	)
	fileuploadData.UpdatedAt = util.GetTimeNow()
	err = u.repoSaFileUpload.UpdateSaFileUpload(ctx, fileuploadData)
	if err != nil {
		return err
	}
	return nil

}

func (u *useSaFileUpload) DeleteSaFileUpload(ctx context.Context, fileuploadID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, u.contexTimeOut)
	defer cancel()
	//Delete File in Folder
	dataFile, _ := u.repoSaFileUpload.GetBySaFileUpload(ctx, fileuploadID)
	//directory api
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	dir_file := dir + dataFile.FilePath
	log.Printf(dir_file)
	//check file klo ada delete
	if notExist := file.CheckNotExist(dir_file); notExist != true {
		err = os.Remove(dir_file)
		if err != nil {
			return err
		}
	}

	// delete data db
	err = u.repoSaFileUpload.DeleteSaFileUpload(ctx, fileuploadID)
	if err != nil {
		return err
	}
	return nil
}
