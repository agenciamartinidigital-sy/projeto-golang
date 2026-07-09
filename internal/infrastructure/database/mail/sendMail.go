package mail

import (
	"fmt"
	"os"
	"projeto-golang/internal/domain/campaign"

	"gopkg.in/gomail.v2"
)

func SendEmail(campaign *campaign.Campaign) error {
	fmt.Println("Sending email...")

	// Propriedades do servidor
	d := gomail.NewDialer(os.Getenv("EMAIL_SMTP"), 587, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))

	var emails []string
	for _, contact := range campaign.Contacts {
		emails = append(emails, contact.Email)
	}

	// Mensagem
	msg := gomail.NewMessage()
	msg.SetHeader("From", os.Getenv("EMAIL_USER"))
	msg.SetHeader("To", emails...)
	msg.SetHeader("Subject", campaign.Name)
	msg.SetBody("text/html", campaign.Content)

	return d.DialAndSend(msg)
}
