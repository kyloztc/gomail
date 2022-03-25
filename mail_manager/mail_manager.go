package mail_manager

import (
	"errors"
	"gomail/mail"
)

type MailManager struct {
	SmtpServerList []mail.SmtpServer
}

func NewMailManager() *MailManager {
	return &MailManager{
		SmtpServerList: make([]mail.SmtpServer, 0),
	}
}

func (m *MailManager) AddSmtpServer(smtpType mail.SmtpType, accountList ...*mail.EmailAccount) error {
	factory, ok := mail.SmtpFactories[smtpType]
	if !ok {
		return errors.New("smtp type not support")
	}
	smtpServer, err := factory(accountList...)
	if err != nil {
		return err
	}
	m.SmtpServerList = append(m.SmtpServerList, smtpServer)
	return nil
}

func (m *MailManager) SendEmail(
	toMailAddress []string, subject string, content mail.EmailContent, mailType mail.MailType,
) error {
	errMsg := "send mail failed: \n"
	for _, smtpServer := range m.SmtpServerList {
		err := smtpServer.SendMail(toMailAddress, subject, content, mailType)
		if err != nil {
			errMsg += "use smtp: " + smtpServer.GetName() + "\n"
			errMsg += err.Error()
			continue
		}
		return nil
	}
	return errors.New(errMsg)
}

