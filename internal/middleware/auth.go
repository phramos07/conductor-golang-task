package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	authHeader = "Authorization"
	authEnvVar = "AUTHORIZATION"
)

func auth(token string) bool {
	authToken := os.Getenv(authEnvVar)
	fmt.Println("authToken: ", authToken, "token: ", token)

	return authToken == token
}

// AuthorizationMiddleware ...
func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(authHeader)

		if auth(token) {
			log.Printf("Authorized request")
			next.ServeHTTP(w, r)
		} else {
			log.Printf("Unauthorized request.")
			http.Error(w, "Forbidden.", http.StatusForbidden)
		}
	})
}
