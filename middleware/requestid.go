package middleware

import "net/http"

func RequestIdMiddleware(next http.Handler, logger AccessLoggerPort) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO
		next.ServeHTTP(w, r)
	})
}
