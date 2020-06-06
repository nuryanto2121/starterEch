package util

import (
	"net/mail"
	"property/framework/pkg/setting"

	"github.com/go-gomail/gomail"
)

//SendEmail :
func SendEmail(to string, subject string, htmlBody string) error {

	smtp := setting.FileConfigSetting.SMTP

	from := mail.Address{
		Name:    smtp.Identity,
		Address: smtp.Sender,
	}
	m := gomail.NewMessage()
	m.Reset()
	m.SetHeader("From", from.String())
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)

	m.SetBody("text/html", htmlBody)
	// m.AddAlternative("text/html", htmlBody)

	d := gomail.NewDialer(smtp.Server, smtp.Port, smtp.User, smtp.Passwd)

	return d.DialAndSend(m)

}
