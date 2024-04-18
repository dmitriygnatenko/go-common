package smtp

import (
	"fmt"
	"net/smtp"
	"strings"
)

type auth struct {
	username string
	password string
}

func (a auth) Start(_ *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", nil, nil
}

func (a auth) Next(req []byte, more bool) ([]byte, error) {
	command := strings.ToLower(strings.TrimSuffix(strings.TrimSpace(string(req)), ":"))

	if more {
		if command == "username" {
			return []byte(fmt.Sprintf("%s", a.username)), nil
		}
		if command == "password" {
			return []byte(fmt.Sprintf("%s", a.password)), nil
		}

		return nil, fmt.Errorf("unexpected server challenge: %s", command)
	}

	return nil, nil
}

func getAuth(username, password string) smtp.Auth {
	return &auth{username, password}
}
