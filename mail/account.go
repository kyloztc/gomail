package mail

import (
	"fmt"
	"net/smtp"
	"strings"
)

type EmailAccount struct {
	MailAddress string
	Passwd string
	SendUserName string
}

func (a *EmailAccount) SendMail(
	smtpHost string, smtpPort string, toMailAddress []string, subject string, content EmailContent, mailType MailType,
) error {
	auth := smtp.PlainAuth("", a.MailAddress, a.Passwd, smtpHost)
	toMail := strings.Join(toMailAddress, ";")
	msg := fmt.Sprintf("To: %s\r\nFrom: %s<%s>\r\nSubject: %s\r\n%s\r\n\r\n%s",
		toMail, a.SendUserName, a.MailAddress, subject, mailType, content.GetContent())
	smtpAddress := smtpHost + ":" + smtpPort
	err := smtp.SendMail(smtpAddress, auth, a.MailAddress, toMailAddress, []byte(msg))
	return err
}