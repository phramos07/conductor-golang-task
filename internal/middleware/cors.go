package middleware

import (
	"net/http"
)

const (
	accessControlOriginHeader      = "Access-Control-Allow-Origin"
	accessControlOriginHeaderValue = "*"
	accessControlHeader            = "Access-Control-Allow-Headers"
	accessControlHeaderValue       = "Content-Type, api_key, Authorization, Origin, X-Requested-With, Accept"
)

// CorsMiddleware ...
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(accessControlOriginHeader, accessControlOriginHeaderValue)
		w.Header().Set(accessControlHeader, accessControlHeaderValue)

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}
