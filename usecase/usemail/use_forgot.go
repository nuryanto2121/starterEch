package useforgot

import (
	templateemail "property/framework/pkg/app/mail"
	util "property/framework/pkg/utils"
	"strings"
)

type Forgot struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	ButtonLink string `json:"button_link"`
}

// // Store :
// func (f *Forgot) Store() error {
// 	return redisdb.StoreForgot(f)
// }

func (f *Forgot) Send() error {
	subjectEmail := "Permintaan Pergantian Password"

	err := util.SendEmail(f.Email, subjectEmail, getForgotBody(f))
	if err != nil {
		return err
	}
	return nil

}

func getForgotBody(f *Forgot) string {
	forgotMail := templateemail.ForgotEmail

	forgotMail = strings.ReplaceAll(forgotMail, `{Name}`, f.Name)
	forgotMail = strings.ReplaceAll(forgotMail, `{ButtonLink}`, f.ButtonLink)
	return forgotMail
}
