package mailer

import (
	"errors"
	"fmt"
	"loan-service/pkg/helper"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

type IMailerService interface {
	Send(req *Request) error
}

type Mailer struct {
	email       string
	appPassword string
	host        string
	port        string
}

func New() IMailerService {
	return &Mailer{
		email:       os.Getenv("MAILER_EMAIL"),
		appPassword: os.Getenv("MAILER_APP_PASSWORD"),
		host:        os.Getenv("MAILER_SMTP_HOST"),
		port:        os.Getenv("MAILER_SMTP_PORT"),
	}
}

func (m *Mailer) Send(req *Request) (err error) {
	if len(req.To) == 0 {
		return errors.New("destination must not empty")
	}

	if helper.IsBlank(req.Subject) || helper.IsBlank(req.Text) {
		return errors.New("subject or text must not blank")
	}

	e := email.NewEmail()
	e.From = m.email
	e.To = req.To
	e.Subject = req.Subject
	e.Text = []byte(req.Text)

	if req.Attachment != nil {
		e.Attach(req.Attachment.Content, req.Attachment.Name, req.Attachment.MimeType)
	}

	err = e.Send(fmt.Sprintf("%s:%s", m.host, m.port), smtp.PlainAuth("", m.email, m.appPassword, m.host))

	return
}
