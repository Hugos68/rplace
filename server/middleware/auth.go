package middleware

import (
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader != "" {
			// TODO: Get user from auth header and set it in the context, see https://youtu.be/H7tbjKFSg58?si=bvAzAk-u-1UhYgem&t=687
		}

		next.ServeHTTP(w, r)
	})
}
