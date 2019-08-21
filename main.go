/**
 * https://sysadmins.co.za/golang-building-a-basic-web-server-in-go/
 * https://itnext.io/building-a-web-http-server-with-go-6554029b4079
 */

package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler) // @see handler.go in order to discover handler function

	http.ListenAndServe(":8080", mux)
}
