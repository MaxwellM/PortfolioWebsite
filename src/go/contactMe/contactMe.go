package contactMe

import (
	"fmt"
	"gopkg.in/mail.v2"
	"strings"

	"PortfolioWebsite/src/go/common"

    "encoding/json"
    "log"
    "net/http"
    "os"
	"io/ioutil"

    "golang.org/x/net/context"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "google.golang.org/api/gmail/v1"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
        // The file token.json stores the user's access and refresh tokens, and is
        // created automatically when the authorization flow completes for the first
        // time.
        tokFile := "token.json"
        tok, err := tokenFromFile(tokFile)
        if err != nil {
                tok = getTokenFromWeb(config)
                saveToken(tokFile, tok)
        }
        return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
        authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
        fmt.Printf("Go to the following link in your browser then type the "+
                "authorization code: \n%v\n", authURL)

        var authCode string
        if _, err := fmt.Scan(&authCode); err != nil {
                log.Fatalf("Unable to read authorization code: %v", err)
        }

        tok, err := config.Exchange(context.TODO(), authCode)
        if err != nil {
                log.Fatalf("Unable to retrieve token from web: %v", err)
        }
        return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
        f, err := os.Open(file)
        if err != nil {
                return nil, err
        }
        defer f.Close()
        tok := &oauth2.Token{}
        err = json.NewDecoder(f).Decode(tok)
        return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
        fmt.Printf("Saving credential file to: %s\n", path)
        f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
        if err != nil {
                log.Fatalf("Unable to cache oauth token: %v", err)
        }
        defer f.Close()
        json.NewEncoder(f).Encode(token)
}

func SendEmail(name, email, phone, message string) error {
        b, err := ioutil.ReadFile("credentials.json")
        if err != nil {
                log.Fatalf("Unable to read client secret file: %v", err)
                return err
        }

        // If modifying these scopes, delete your previously saved token.json.
        config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
        if err != nil {
                log.Fatalf("Unable to parse client secret file to config: %v", err)
                return err
        }
        client := getClient(config)

        srv, err := gmail.New(client)
        if err != nil {
                log.Fatalf("Unable to retrieve Gmail client: %v", err)
                return err
        }

        user := "me"
        r, err := srv.Users.Labels.List(user).Do()
        if err != nil {
                log.Fatalf("Unable to retrieve labels: %v", err)
                return err
        }
        if len(r.Labels) == 0 {
                fmt.Println("No labels found.")
        }
        fmt.Println("Labels:")
        for _, l := range r.Labels {
                fmt.Printf("- %s\n", l.Name)
        }
        return nil
}

func SendEmail2(name, email, phone, message string) error {
    googleEmailInfo, err := common.ReadJsonFile("googleEmailPassword.json")
    if err != nil {
        fmt.Println("Error reading JSON file!")
        return err
    }
    googleEmailPassword := googleEmailInfo["Password"].(string)

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

	d := mail.NewDialer("smtp.gmail.com", 587, "maxintosh.mailer", googleEmailPassword)
	d.StartTLSPolicy = mail.MandatoryStartTLS

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
	    fmt.Println("Error: ", err.Error())
		return err
	} else {
		return nil
	}
}
