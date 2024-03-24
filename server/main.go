package main

import (
	"net/http"

	"github.com/hugos68/rplace/middleware"
)

func main() {

	mw := middleware.Chain()

	server := http.Server{
		Addr:    ":8080",
		Handler: mw,
	}

	server.ListenAndServe()
}
