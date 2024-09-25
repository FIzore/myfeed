package api

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func main() {
	targetMail := "2671636676@qq.com"
	smtpServer := "smtp.gmail.com"
	emailAddr := "mapmyfeed@gmail.com"
	emailKey := "kulrxhariiczwlvi"

	em := email.NewEmail()
	em.From = fmt.Sprintf("Gmail<%s>", emailAddr)
	em.To = []string{targetMail}
	em.Subject = "Test Email"
	em.Text = []byte("Hello, this is a test email from Go.")
	em.Send(smtpServer+":587", smtp.PlainAuth("", emailAddr, emailKey, smtpServer))
}
