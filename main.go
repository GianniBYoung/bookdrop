package main

import (
	"context"
	"os"

	"github.com/GianniBYoung/configa"
	"github.com/charmbracelet/log"
	"github.com/resend/resend-go/v2"
)

type Config struct {
	defaultSender   string `yaml:"defaultSender"`
	defaultReciever string `yaml:"defaultReciever"`
	apikey          string `yaml:"apikey"`
	DebugMode       bool   `yaml:"debugMode"`
}

func main() {
	configa.MainConfiga()
	ctx := context.TODO()
	apiKey := os.Getenv("RESEND_API_KEY")

	if apiKey == "" {
		log.Fatal("Api Key is missing")
	}

	// Read attachment file
	pwd, _ := os.Getwd()
	f, err := os.ReadFile(pwd + "/aaaa.epub")
	if err != nil {
		log.Fatal(err)
	}

	client := resend.NewClient(apiKey)

	// Create attachments objects
	BookAttachment := &resend.Attachment{
		Content:  f,
		Filename: "a.epub",
	}

	params := &resend.SendEmailRequest{
		To:          []string{"younggianniguy@gmail.com", "whitepapergianni@kindle.com"},
		From:        "kindle@mancys-metal.xyz",
		Text:        "This is a book!",
		Html:        "<strong>email with attachments !!</strong>",
		Subject:     "Automate the Boring Stuff with Python!",
		Attachments: []*resend.Attachment{BookAttachment},
	}

	sent, err := client.Emails.SendWithContext(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(sent.Id)
}
