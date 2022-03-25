package mail_manager

import (
	"fmt"
	mail2 "gomail/mail"
	"testing"
)

func TestMailManager_SendEmail(t *testing.T) {
	mail := ""
	pwd := ""
	account := &mail2.EmailAccount{
		MailAddress:  mail,
		Passwd:       pwd,
		SendUserName: mail,
	}
	mailManager := NewMailManager()
	err := mailManager.AddSmtpServer(mail2.QQSmtp, account)
	if err != nil {
		fmt.Printf("add smtp server error|%v\n", err)
		return
	}
	htmlContent := mail2.NewHtmlContent("<h1>这是一封测试邮件520</h1>")
	err = mailManager.SendEmail([]string{
		"",
	}, "你好，我是gomail", htmlContent, htmlContent.GetContentType())
	if err != nil {
		fmt.Printf("send email error|%v\n", err)
	}
}
