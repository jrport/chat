package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"log/slog"
	"net/smtp"

	"github.com/goloop/env"
)

const templ string = `
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    {{.Content}}
</body>
</html>
`

type EmailData struct {
	Title   string
	Content template.HTML
}

func NewEmail(title string, content string) *EmailData {
	return &EmailData{
		Title:   title,
		Content: template.HTML(content),
	}
}

type MailCredentials struct {
	Email    string
	AppToken string
	Host     string
	Port     string
}

func GetMailerCredentials() (*MailCredentials, error) {
	if err := env.Load(".env"); err != nil {
		return nil, err
	}

	return &MailCredentials{
		Email:    env.Get("EMAIL"),
		AppToken: env.Get("PASSWORD"),
		Host:     env.Get("HOST"),
		Port:     env.Get("PORT"),
	}, nil
}

func (mc *MailCredentials) Auth() smtp.Auth {
	return smtp.PlainAuth("", mc.Email, mc.AppToken, mc.Host)
}

func WriteEmail(content *EmailData) *[]byte {
	t, err := template.New("messsage").Parse(templ)
	var htmlBuffer bytes.Buffer
	if err != nil {
		slog.Error(fmt.Sprintf("Error on email of content: %v", content.Content))
		return nil
	}

	err = t.Execute(&htmlBuffer, content)
	if err != nil {
		slog.Error(fmt.Sprintf("Error on email of content: %v", content.Content))
		return nil
	}

	res := []byte(
		"Subject: " + content.Title + "\n" +
			"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\"\n\n" +
			htmlBuffer.String(),
	)
	return &res
}
