package middleware

import (
	"net/http"

	"github.com/google/uuid"
)

func RequestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := r.Header.Get("X-Request-Id")
		if requestId == "" {
			requestId = uuid.New().String()
		}
		r.Header.Set("X-Request-Id", requestId)
		next.ServeHTTP(w, r)
	})
}
