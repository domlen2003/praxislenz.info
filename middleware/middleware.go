package middleware

import (
	"log"
	"net/http"
	"time"
)

//RecoverHandler handles recovering of errors and gives error-code when unable to recover
func RecoverHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := recover(); err != nil {
			log.Printf("panic: %+v", err)
			http.Error(w, http.StatusText(500), 500)
		}
		next.ServeHTTP(w, r)
	})
}

//LoggerHandler logs every incoming request and its according serving time
func LoggerHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		log.Printf("<< %s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}
