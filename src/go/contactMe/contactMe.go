package contactMe

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mail.v2"
	"io/ioutil"
	"strings"
)

type GoogleInfo struct {
	Password            string
}

func getGoogleInfo() (GoogleInfo, error) {
	file, err := ioutil.ReadFile("googleEmailPassword.json")
	if err != nil {
		fmt.Println("Error reading JSON file: ", err)
	}
	data := GoogleInfo{}
	err = json.Unmarshal([]byte(file), &data)
	return data, nil
}

func SendEmail(name, email, phone, message string) error {
	googleInfo, err := getGoogleInfo()
	if err != nil {
		fmt.Println("Error obtaining our Google Email Info! ", err.Error())
		return err
	}
	fmt.Println("Password: ", googleInfo.Password)

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

	d := mail.NewDialer("smtp.gmail.com", 587, "maxintosh.mailer", googleInfo.Password)
	d.StartTLSPolicy = mail.MandatoryStartTLS

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
	    fmt.Println("Error: ", err.Error())
		return err
	} else {
		return nil
	}
}
