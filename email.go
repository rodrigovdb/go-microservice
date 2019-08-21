package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/smtp"
)

type Email struct {
	to      string
	from    string
	subject string
	body    string
}

func (e Email) to_s() string {
	return fmt.Sprintf("{ to: \"%s\", from: \"%s\", subject: \"%s\", body: \"%s\" }", e.to, e.from, e.subject, e.body)
}

func (e Email) msg() string {
	return "From: " + e.from + "\n" +
		"To:" + e.to + "\n" +
		"Subject" + e.subject + "\n" +
		e.body

}

func (e Email) send() error {
	credentials, err := readCredentials()

	if err != nil {
		return err
	}

	err = smtp.SendMail(credentials.hostname+":"+credentials.port,
		smtp.PlainAuth("", credentials.username, credentials.password, credentials.hostname),
		e.from, []string{e.to}, []byte(e.body))

	return err
}

func initialize(r *http.Request) (*Email, error) {
	email := Email{to: "", from: "", subject: "", body: ""}

	to := r.URL.Query()["to"]
	if len(to) < 1 {
		return nil, errors.New("URL param 'to' is missing")
	} else {
		email.to = to[0]
	}

	from := r.URL.Query()["from"]
	if len(from) < 1 {
		email.from = "Default From <default@from.com>"
	} else {
		email.from = from[0]
	}

	subject := r.URL.Query()["subject"]
	if len(subject) < 1 {
		return nil, errors.New("URL param 'subject' is missing")
	} else {
		email.subject = subject[0]
	}

	body := r.URL.Query()["body"]
	if len(body) < 1 {
		return nil, errors.New("URL param 'body' is missing")
	} else {
		email.body = body[0]
	}

	return &email, nil
}
