package goExamples

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

type GoogleInfo struct {
	Type                    string
	ProjectID               string
	PrivateKeyID            string
	ClientEmail             string
	ClientID                string
	AuthURI                 string
	TokenURI                string
	AuthProviderX509CertURL string
	ClientX509CertURL       string
}

func getGoogleInfo() (GoogleInfo, error) {
	file, err := ioutil.ReadFile("googleKey.json")
	if err != nil {
		fmt.Println("Error reading JSON file: ", err)
	}
	data := GoogleInfo{}
	err = json.Unmarshal([]byte(file), &data)
	return data, nil
}

func TranslateString (stringToTranslate, lang string) (string, error) {

	fmt.Println("String After: ", stringToTranslate)
	fmt.Println("Lang After: ", lang)


	// Brody, this contains all of the googleKey.json informaiton. Please use this?
	// Thank you!
	googleInfo, err := getGoogleInfo()
	if err != nil {
		fmt.Println("Error obtaining our Google Info! ", err.Error())
		return "", err
	}

	fmt.Println("Google Info: ", googleInfo)

	ctx := context.Background()

	// Creates a client.
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the text to translate.
	text := stringToTranslate
	// Sets the target language.
	target, err := language.Parse(lang)
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}

	// Translates the text into French.
	translations, err := client.Translate(ctx, []string{text}, target, nil)
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
	}

	fmt.Printf("Text: %v\n", text)
	fmt.Printf("Translation: %v\n", translations[0].Text)

	return translations[0].Text, nil
}
