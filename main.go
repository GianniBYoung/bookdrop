package main

import (
	"context"
	"path/filepath"
	"strings"

	"bookdrop/configa"
	"os"

	"github.com/charmbracelet/log"
	"github.com/resend/resend-go/v2"
)

func main() {
	log.SetReportTimestamp(false)
	log.SetReportCaller(false)
	if os.Getenv("BOOKDROP_DEBUG") != "" {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug Enabled!")
	}

	if len(os.Args) < 2 {
		log.Fatal("No Books Attached! Please Provide a Path!")
	}

	configa.Configure()

	ctx := context.TODO()
	args := os.Args

	log.Debugf("Unmarshalled yaml config:\n%+v\n", configa.Config)

	log.Debug(configa.Config.ApiKey)
	if configa.Config.ApiKey == "" {
		log.Fatal("Api Key is missing")
	}

	// pwd, _ := os.Getwd()
	f, err := os.ReadFile(args[1])
	if err != nil {
		log.Fatal(err)
	}

	fileName := filepath.Base(args[1])
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))

	client := resend.NewClient(configa.Config.ApiKey)

	// Create attachments objects
	BookAttachment := &resend.Attachment{
		Content:  f,
		Filename: fileName,
	}

	params := &resend.SendEmailRequest{
		To:          []string{configa.Config.DefaultReciever},
		From:        configa.Config.DefaultSender,
		Text:        "Book Drop!",
		Html:        "<strong>Books attached to the email Boss !!</strong>",
		Subject:     "Book Drop Incoming!",
		Attachments: []*resend.Attachment{BookAttachment},
	}

	sent, err := client.Emails.SendWithContext(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(sent.Id)
}
