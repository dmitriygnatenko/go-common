package smtp

import (
	"errors"
	"fmt"
	"net/smtp"
)

type SMTP struct {
	config Config
}

func NewSMTP(c Config) (*SMTP, error) {
	if len(c.username) == 0 {
		return nil, errors.New("empty username")
	}

	if len(c.password) == 0 {
		return nil, errors.New("empty password")
	}

	if len(c.host) == 0 {
		c.host = defaultHost
	}

	if c.port == 0 {
		c.port = defaultPort
	}

	return &SMTP{config: c}, nil
}

func (s SMTP) Send(recipient string, subject string, content string, html bool) error {
	contentType := "text/plain"
	if html {
		contentType = "text/html"
	}

	msg := []byte("To: " + recipient + "\r\n" +
		"From: " + s.config.username + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: " + contentType + "; charset=\"UTF-8\"" + "\n\r\n" +
		content + "\r\n")

	return smtp.SendMail(
		fmt.Sprintf("%s:%d", s.config.host, s.config.port),
		getAuth(s.config.username, s.config.password),
		s.config.username,
		[]string{recipient},
		msg,
	)
}
