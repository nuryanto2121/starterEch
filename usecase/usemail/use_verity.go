package usemail

import (
	templateemail "property/framework/pkg/app/mail"
	util "property/framework/pkg/utils"
	"strings"
)

type Verify struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	ButtonLink string `json:"button_link"`
}

func (f *Verify) SendVerify() error {
	subjectEmail := "Aktivasi Account"

	err := util.SendEmail(f.Email, subjectEmail, getVerifybody(f))
	if err != nil {
		return err
	}
	return nil

}

func getVerifybody(f *Verify) string {
	forgotMail := templateemail.VerifyEmail

	forgotMail = strings.ReplaceAll(forgotMail, `{Name}`, f.Name)
	forgotMail = strings.ReplaceAll(forgotMail, `{ButtonLink}`, f.ButtonLink)
	return forgotMail
}
