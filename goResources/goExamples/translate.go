package goExamples

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func TranslateString (stringToTranslate, lang string) (string, error) {

	fmt.Println("String After: ", stringToTranslate)
	fmt.Println("Lang After: ", lang)

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
