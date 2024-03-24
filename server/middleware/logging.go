package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		stop := time.Now()
		log.Printf("%s %s %v", r.Method, r.URL.Path, stop.Sub(start))
	})
}
