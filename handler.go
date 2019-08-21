package main

import (
	"log"
	"net/http"
)

const message = "SUCCESS: e-mail has been sent."

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	email, err := initialize(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ERROR: " + err.Error()))

		return
	}

	err = email.send()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ERROR - SMTP Server Error"))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))

	log.Println("root function handler was executed")
}
