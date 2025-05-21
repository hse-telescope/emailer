package email

import (
	"context"
	"fmt"
	"net/smtp"
)

const mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

// Provider ...
type Provider struct {
	addr      string
	auth      smtp.Auth
	fromEmail string
}

// NewEmailProvider ...
func NewEmailProvider(
	addr string,
	auth smtp.Auth,
	fromEmail string,
) Provider {
	return Provider{
		addr:      addr,
		auth:      auth,
		fromEmail: fromEmail,
	}
}

func setupMessage(headers map[string]string, mime string, body string) string {
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\n", k, v)
	}
	message += mime + body

	return message
}

// SendEmail ...
func (p Provider) SendEmail(ctx context.Context, emailTo string, message Message) error {
	return smtp.SendMail(
		p.addr,
		p.auth,
		p.fromEmail,
		[]string{emailTo},
		[]byte(setupMessage(
			map[string]string{
				"Subject": message.Subject,
			},
			mime,
			message.Body,
		)),
	)
}
