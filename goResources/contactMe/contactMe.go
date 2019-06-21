package contactMe

import (
	"gopkg.in/mail.v2"
)


func SendEmail(name, email, phone, message string) error {
	m := mail.NewMessage()
	m.SetHeader("From", "maxintosh.mailer@gmail.com")
	m.SetHeader("To", "maxwellrmorin@gmail.com")
	m.SetAddressHeader("Cc", email, "Mailer")
	m.SetHeader("Subject", "Maxintosh Contact Request")
	m.SetBody("text/html", message)

	d := mail.NewDialer("smtp.gmail.com", 587, "maxintosh.mailer", "17DesertHawk!")
	d.StartTLSPolicy = mail.MandatoryStartTLS

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	} else {
		return nil
	}
}
