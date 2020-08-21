package middleware

import (
	"github.com/gorilla/handlers"
)

// CorsMiddleware for http handlers
func CorsMiddleware() []handlers.CORSOption {
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	var cors []handlers.CORSOption
	cors = append(cors, allowedHeaders)
	cors = append(cors, allowedOrigins)
	cors = append(cors, allowedMethods)

	return cors
}
