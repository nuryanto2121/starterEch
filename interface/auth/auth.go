package iauth

import (
	"context"
	"property/framework/models"
	sa_models "property/framework/models/sa"
)

type Usecase interface {
	Login(ctx context.Context, dataLogin *models.LoginForm) (output models.Output)
	ForgotPassword(ctx context.Context, dataForgot *models.ForgotForm) (output models.Output)
	Register(ctx context.Context, dataRegister *sa_models.SaClient) (output models.Output) //(err error) //(result map[string]interface{}, err error)
	ResetPassword(ctx context.Context, dataReset *models.ResetPasswd) (output models.Output)
	Verify(ctx context.Context, dataVerify *models.ResetPasswd) (output models.Output)
}
