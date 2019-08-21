package main

import (
	"errors"
	"os"
)

type Credentials struct {
	hostname string
	username string
	password string
	port     string
}

func readCredentials() (*Credentials, error) {
	hostname := os.Getenv("EMAIL_HOSTNAME")
	if len(hostname) < 1 {
		return nil, errors.New("Missed EMAIL_HOSTNAME on .env")
	}

	username := os.Getenv("EMAIL_USERNAME")
	if len(username) < 1 {
		return nil, errors.New("Missed EMAIL_USERNAME on .env")
	}

	password := os.Getenv("EMAIL_PASSWORD")
	if len(password) < 1 {
		return nil, errors.New("Missed EMAIL_PASSWORD on .env")
	}

	port := os.Getenv("EMAIL_PORT")
	if len(port) < 1 {
		return nil, errors.New("Missed EMAIL_PORT on .env")
	}

	credentials := Credentials{hostname: hostname, username: username, password: password, port: port}

	return &credentials, nil
}
