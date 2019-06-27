package contactMe

import (
	"gopkg.in/mail.v2"
	"strings"
)


func SendEmail(name, email, phone, message string) error {
	messageBody := []string{}
	messageBody = append(messageBody, "Name: "+name)
	messageBody = append(messageBody, "Email: "+email)
	messageBody = append(messageBody, "Phone: "+phone)
	messageBody = append(messageBody, "Message: "+message)
	messageBodyStr := strings.Join(messageBody, "\n")

	m := mail.NewMessage()
	m.SetHeader("From", "maxintosh.mailer@gmail.com")
	m.SetHeader("To", "maxwellrmorin@gmail.com")
	m.SetAddressHeader("Cc", email, "Mailer")
	m.SetHeader("Subject", "Maxintosh Contact Request")
	m.SetBody("text/plain", messageBodyStr)

	d := mail.NewDialer("smtp.gmail.com", 587, "maxintosh.mailer", "17DesertHawk!")
	d.StartTLSPolicy = mail.MandatoryStartTLS

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	} else {
		return nil
	}
}
