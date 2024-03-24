package main

import (
	"net/http"
)

func main() {

	mw := Chain()

	server := http.Server{
		Addr:    ":8080",
		Handler: mw,
	}

	server.ListenAndServe()
}
