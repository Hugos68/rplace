package main

import (
	"net/http"

	"github.com/hugos68/rplace/middleware"
)

func main() {
	router := http.NewServeMux()

	loadRoutes(router)

	mw := middleware.Chain(middleware.Logging)

	server := http.Server{
		Addr:    ":8080",
		Handler: mw(router),
	}

	server.ListenAndServe()
}

func loadRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
}
