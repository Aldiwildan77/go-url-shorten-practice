package middleware

import (
	"log"
	"net/http"
)

// LoggerMiddleware logging and goto next middlewares
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
