package mail

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	SMTPAUTHADDRESS   = "smtp.gmail.com"
	SMTPSERVERADDRESS = "smtp.gmail.com:587"
)

type SenderGmail struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

type SenderEmail interface {
	SendEmail(
		subject string,
		body string,
		to []string,
		cc []string,
		bcc []string,
		attachedFiles []string,
	) error
}

func NewSenderGmail(name string, fromEmailAddress string, fromEmailPassword string) SenderEmail {

	return &SenderGmail{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *SenderGmail) SendEmail(subject string, body string, to []string, cc []string, bcc []string, attachedFiles []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(body) //content
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	for _, file := range attachedFiles {
		_, err := e.AttachFile(file)
		if err != nil {
			return fmt.Errorf("failed to attach file %s: %w", file, err)
		}

	}

	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, SMTPAUTHADDRESS)

	err := e.Send(SMTPSERVERADDRESS, smtpAuth)

	return err
}
