package mail

import (
	"fmt"
	"testing"
)

func TestQQMailSmtpServer_SendMail(t *testing.T) {
	mail := ""
	pwd := ""
	account := &EmailAccount{
		MailAddress:  mail,
		Passwd:       pwd,
		SendUserName: mail,
	}
	smtpServer, err := newQQMailSmtpServer(account)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	textContent := NewTextContent("这是一封测试邮件")
	err = smtpServer.SendMail([]string{
		"",
	}, "你好，我是gomail", textContent, textContent.GetContentType())
	if err != nil {
		fmt.Printf("send maul error|%v\n", err)
	}
}
