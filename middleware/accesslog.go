package middleware

import (
	"net/http"
	"time"
)

type customResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *customResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func newCustomedResponseWriter(w http.ResponseWriter) *customResponseWriter {
	// Defaults to 200
	return &customResponseWriter{w, http.StatusOK}
}

func AccessLogMiddleware(next http.Handler, logger AccessLoggerPort) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		crw := newCustomedResponseWriter(w)

		next.ServeHTTP(w, r)

		durationMs := time.Since(start).Milliseconds()
		requestId := r.Header.Get("X-Request-Id")
		logger.Info(int(durationMs), crw.status, r.URL.Path, r.Method, requestId)
	})
}

type AccessLoggerPort interface {
	Info(durationMs int, status int, requestPath string, requestMethod string, requestId string)
}
