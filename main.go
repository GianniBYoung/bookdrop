package main

import (
	"context"
	"os"

	"bookdrop/configa"

	"github.com/charmbracelet/log"
	"github.com/resend/resend-go/v2"
)

func main() {
	configa.Configure()
	log.Error(configa.Config)
	ctx := context.TODO()
	args := os.Args
	apiKey := os.Getenv("RESEND_API_KEY")

	if apiKey == "" {
		log.Fatal("Api Key is missing")
	}

	// pwd, _ := os.Getwd()
	f, err := os.ReadFile(args[1])
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
