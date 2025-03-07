package smtp

import (
	"book-storage/internal/models"
	"book-storage/pkg/email"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/go-gomail/gomail"
	"github.com/google/uuid"
)

type SMTPSender struct {
	from string
	pass string
	host string
	port int
}

func NewSMTPSender(from, pass, host string, port int) (*SMTPSender, error) {
	if !email.IsValid(from) {
		return nil, fmt.Errorf("failed to create smtp sender: %w", models.ErrEmailFormat)
	}

	return &SMTPSender{from: from, pass: pass, host: host, port: port}, nil
}

func (s *SMTPSender) Send(input SendEmailInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", s.from)
	msg.SetHeader("To", input.To)
	msg.SetHeader("Subject", input.Subject)
	msg.SetHeader("Date", time.Now().Format(time.RFC1123Z))
	msg.SetHeader("MIME-Version", "1.0")
	msg.SetHeader("Message-ID", fmt.Sprintf("<%s@%s>", uuid.NewString(), "yandex.ru"))
	msg.SetHeader("Content-Type", "text/html; charset=UTF-8")
	msg.SetBody("text/html", input.Body)

	dialer := gomail.NewDialer(s.host, s.port, s.from, s.pass)
	dialer.TLSConfig = &tls.Config{ServerName: s.host}
	if err := dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("failed to sent email: %w", err)
	}

	return nil
}
