package mail

var (
	QQMailSmtpAddress = "smtp.qq.com"
	QQMailSmtpPort = "587"
)

type SmtpType int

var (
	QQSmtp SmtpType = 1
)

type MailType string

var (
	TextMailType MailType = "Content-Type: text/plain; charset=UTF-8"
	HtmlMailType MailType = "Content-Type: text/html; charset=UTF-8"
)


