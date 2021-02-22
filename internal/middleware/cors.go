package middleware

import (
	"net/http"
)

const (
	accessControlOriginHeader = "Access-Control-Allow-Origin"
)

// CorsMiddleware ...
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(accessControlOriginHeader, "*")
		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}
