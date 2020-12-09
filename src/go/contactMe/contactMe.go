package contactMe

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

var runningClient *http.Client
var runningGMailService *gmail.Service

func init() {
    fmt.Println("INIT RUNNING!")
    runningClient, runningGMailService = getMyStuff()
    //setupCredentials()
}

func exists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getMyStuff() (*http.Client, *gmail.Service) {
	// Reads in our credentials
	secret, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Printf("Error Reading Credentials: %v", err)
	}

	// Creates a oauth2.Config using the secret
	// The second parameter is the scope, in this case we only want to send email
	conf, err := google.ConfigFromJSON(secret, gmail.GmailSendScope)
	if err != nil {
		log.Printf("Error Creating CONF: %v", err)
	}

	if exists("token.json") {
	    fmt.Println("Token does exist!")
		tok, err := tokenFromFile("token.json")
		if err != nil {
			log.Printf("Error Reading Token: %v", err)
		}

		// Create the *http.Client using the access token
		runningClient = conf.Client(oauth2.NoContext, tok)

		// Create a new gmail service using the client
		gmailService, err := gmail.New(runningClient)
		if err != nil {
			log.Printf("Error Creating GMail Service: %v", err)

		}
		return runningClient, gmailService
	} else {
	    fmt.Println("Token doesn't exist! ")
        tokFile := "token.json"
        tok := getTokenFromWeb(conf)

        saveToken(tokFile, tok)

        // Create the *http.Client using the access token
        runningClient = conf.Client(oauth2.NoContext, tok)

        // Create a new gmail service using the client
        gmailService, err := gmail.New(runningClient)
        if err != nil {
            log.Printf("Error Creating GMail Service: %v", err)

        }
        return runningClient, gmailService
	}
}

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

func setupCredentials() {
        b, err := ioutil.ReadFile("credentials.json")
        if err != nil {
                log.Fatalf("Unable to read client secret file: %v", err)

        }

        // If modifying these scopes, delete your previously saved token.json.
        config, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
        if err != nil {
                log.Fatalf("Unable to parse client secret file to config: %v", err)

        }
        client := getClient(config)

        srv, err := gmail.New(client)
        if err != nil {
                log.Fatalf("Unable to retrieve Gmail client: %v", err)

        }

        user := "me"
        r, err := srv.Users.Labels.List(user).Do()
        if err != nil {
                log.Fatalf("Unable to retrieve labels: %v", err)

        }
        if len(r.Labels) == 0 {
                fmt.Println("No labels found.")
        }
        fmt.Println("Labels:")
        for _, l := range r.Labels {
                fmt.Printf("- %s\n", l.Name)
        }
}

func SendEmail(name, email, phone, message string) error {
 	// New message for our gmail service to send
 	var fullMessage gmail.Message

 	// Compose the message
 	messageStr := []byte(
 		"From: maxintosh.mailer@gmail.com\r\n" +
        "To: "+ email + "\r\n" +
        "Subject: Maxintosh Contact Request\r\n\r\n" +
		"Name: " + name + "\n" +
		"Email: " + email + "\n" +
		"Phone: " + phone + "\n" +
		"Message: " + message)

 	// Place messageStr into message.Raw in base64 encoded format
	fullMessage.Raw = base64.URLEncoding.EncodeToString(messageStr)

 	// Send the message
 	_, err := runningGMailService.Users.Messages.Send("me", &fullMessage).Do()
 	if err != nil {
 		log.Printf("Error: %v", err)
 		return err
 	} else {
 		fmt.Println("Message sent!")
 		return nil
 	}
}
