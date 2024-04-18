package logger

import (
	"bytes"
	"encoding/json"
	"errors"

	"git.dmitriygnatenko.ru/dima/go-common/smtp"
)

type EmailWriter struct {
	recipient string
	subject   string
	smtp      *smtp.SMTP
}

func NewEmailWriter(smtp *smtp.SMTP, recipient string, subject string) (*EmailWriter, error) {
	if len(recipient) == 0 {
		return nil, errors.New("empty recipient")
	}

	if smtp == nil {
		return nil, errors.New("empty smtp client")
	}

	return &EmailWriter{
		recipient: recipient,
		subject:   subject,
		smtp:      smtp,
	}, nil
}

func (w EmailWriter) Write(p []byte) (int, error) {
	var out bytes.Buffer

	if err := json.Indent(&out, p, "", "    "); err != nil {
		return 0, err
	}

	err := w.smtp.Send(w.recipient, w.subject, out.String(), false)
	return 0, err
}
