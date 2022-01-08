package hamsg

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type Mailer struct {
	Account  string
	Psw      string
	SmtpHost string
	SmtpPort int
}

func NewMailer(account string, psw string, smtpHost string, smtpPort int) *Mailer {
	return &Mailer{
		Account:  account,
		Psw:      psw,
		SmtpHost: smtpHost,
		SmtpPort: smtpPort,
	}
}

func (mailer *Mailer) Send(to string, title string, content string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", mailer.Account)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", title)
	msg.SetBody("text/plain", content)
	d := gomail.NewDialer(mailer.SmtpHost, mailer.SmtpPort, mailer.Account, mailer.Psw)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(msg)
	return err
}
