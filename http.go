package main

import (
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!\n"))
	log.Println("root function handler was executed")
}

func serve() {
	http.HandleFunc("/", root)
	http.ListenAndServe(":8080", nil)
}
