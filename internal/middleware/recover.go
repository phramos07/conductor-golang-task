package middleware

import (
	"errors"
	"log"
	"net/http"
)

const (
	unknownErrorStr = "Unknown error."
)

// Internal method that deals with error messages
func recoverInternal(w http.ResponseWriter) {
	var err error
	r := recover()
	if r != nil {
		switch t := r.(type) {
		case string:
			err = errors.New(t)
		case error:
			err = t
		default:
			err = errors.New(unknownErrorStr)
		}
		log.Printf("Panic: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// RecoverMiddleware ...
func RecoverMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer recoverInternal(w)
		h.ServeHTTP(w, r)
	})
}
