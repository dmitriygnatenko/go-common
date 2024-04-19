package logger

import (
	"bytes"
	"encoding/json"
	"errors"
)

type EmailWriter struct {
	recipient  string
	subject    string
	smtpClient SMTPClient
}

func NewEmailWriter(smtpClient SMTPClient, recipient string, subject string) (*EmailWriter, error) {
	if len(recipient) == 0 {
		return nil, errors.New("empty recipient")
	}

	if smtpClient == nil {
		return nil, errors.New("empty smtp client")
	}

	return &EmailWriter{
		recipient:  recipient,
		subject:    subject,
		smtpClient: smtpClient,
	}, nil
}

func (w EmailWriter) Write(p []byte) (int, error) {
	var out bytes.Buffer

	if err := json.Indent(&out, p, "", "    "); err != nil {
		return 0, err
	}

	err := w.smtpClient.Send(w.recipient, w.subject, out.String(), false)
	return 0, err
}
