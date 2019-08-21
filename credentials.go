package main

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Credentials struct {
	hostname string
	username string
	password string
	port     string
}

func readCredentials() (*Credentials, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	hostname := os.Getenv("EMAIL_HOSTNAME")
	log.Println("hostname: " + hostname)
	if len(hostname) < 1 {
		return nil, errors.New("Missed EMAIL_HOSTNAME on .env")
	}

	username := os.Getenv("EMAIL_USERNAME")
	log.Println("username: " + username)
	if len(username) < 1 {
		return nil, errors.New("Missed EMAIL_USERNAME on .env")
	}

	password := os.Getenv("EMAIL_PASSWORD")
	log.Println("password: " + password)
	if len(password) < 1 {
		return nil, errors.New("Missed EMAIL_PASSWORD on .env")
	}

	port := os.Getenv("EMAIL_PORT")
	log.Println("port: " + port)
	if len(port) < 1 {
		return nil, errors.New("Missed EMAIL_PORT on .env")
	}

	credentials := Credentials{hostname: hostname, username: username, password: password, port: port}

	return &credentials, nil
}
