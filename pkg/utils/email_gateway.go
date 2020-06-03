package util

import (
	"net/mail"
	"property/framework/pkg/setting"

	"github.com/go-gomail/gomail"
)

//SendEail :
func SendEail(to string, subject string, htmlBody string, txtBody string) error {

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

	m.SetBody("text/plain", txtBody)
	m.AddAlternative("text/html", htmlBody)

	d := gomail.NewDialer(smtp.Server, smtp.Port, smtp.User, smtp.Passwd)

	return d.DialAndSend(m)

}
