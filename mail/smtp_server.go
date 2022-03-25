package mail

import (
	"errors"
	"strconv"
	"strings"
)

var (
	accountNilError = errors.New("mail account nil")
	mailDomainError = errors.New("mail domain name error")
)

var SmtpFactories map[SmtpType]smtpInitFunc

type smtpInitFunc func(accountList ...*EmailAccount) (SmtpServer, error)

func init() {
	SmtpFactories = map[SmtpType]smtpInitFunc{
		QQSmtp: newQQMailSmtpServer,
	}
}

type SmtpServer interface {
	GetName() string
	GetAddress() string
	AddAccount(account []*EmailAccount) error
	RemoveAccount(mailAddress string)
	SendMail(toMailAddress []string, subject string, content EmailContent, mailType MailType) error
}

type QQMailSmtpServer struct {
	SmtpAddress string
	SmtpPort string
	AccountList []*EmailAccount
	index int
}

func newQQMailSmtpServer(accountList ...*EmailAccount) (SmtpServer, error) {
	qqMailSmtpServer := &QQMailSmtpServer{
		SmtpAddress: QQMailSmtpAddress,
		SmtpPort:    QQMailSmtpPort,
		index:       0,
	}
	err := qqMailSmtpServer.AddAccount(accountList)
	if err != nil {
		return nil, err
	}
	return qqMailSmtpServer, nil
}

func (q *QQMailSmtpServer) GetName() string {
	return "QQ mail smtp server"
}

func (q *QQMailSmtpServer) GetAddress() string {
	return q.SmtpAddress + ":" + q.SmtpPort
}

func (q *QQMailSmtpServer) AddAccount(account []*EmailAccount) error {
	if q.AccountList == nil {
		q.AccountList = make([]*EmailAccount, 0)
	}
	for _, account := range account {
		if !strings.HasSuffix(account.MailAddress, "@qq.com") {
			return mailDomainError
		}
		q.AccountList = append(q.AccountList, account)
	}
	return nil
}

func (q *QQMailSmtpServer) RemoveAccount(mailAddress string) {
	return
}

func (q *QQMailSmtpServer) SendMail(
	toMailAddress []string, subject string, content EmailContent, mailType MailType,
) error {
	if q.AccountList == nil || len(q.AccountList) == 0 {
		return accountNilError
	}
	retryTimes := len(q.AccountList) * 3
	errMsg := ""
	for i := 0; i < retryTimes; i ++ {
		mailAccount := q.AccountList[q.index]
		q.index = (q.index + 1) % len(q.AccountList)
		err := mailAccount.SendMail(q.SmtpAddress, q.SmtpPort, toMailAddress, subject, content, mailType)
		if err != nil {
			errMsg += "retry times " + strconv.Itoa(i) + ":" +  mailAccount.MailAddress + ":" + err.Error() + "\n"
			continue
		}
		return nil
	}
	return errors.New(errMsg)
}


