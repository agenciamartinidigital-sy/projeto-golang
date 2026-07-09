package mail

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendEmail() error {
	fmt.Println("Sending email...")

	// Propriedades do servidor
	d := gomail.NewDialer(os.Getenv("EMAIL_SMTP"), 587, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))

	// Mensagem
	msg := gomail.NewMessage()
	msg.SetHeader("From", os.Getenv("EMAIL_USER"))
	msg.SetHeader("To", "agenciamartinidigital@gmail.com")
	msg.SetHeader("Subject", "Hello")
	msg.SetBody("text/html", "Hello <b>Luís Martini</b>")

	return d.DialAndSend(msg)
}
