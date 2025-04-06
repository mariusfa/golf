package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mariusfa/golf/request"
)

func setDummyRequestContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestIdCtx := request.NewRequestIdCtx("")
		ctx = context.WithValue(ctx, request.RequestIdCtxKey, requestIdCtx)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func requestIdHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestIdCtx := r.Context().Value(request.RequestIdCtxKey).(*request.RequestIdCtx)

		w.Write([]byte(requestIdCtx.RequestId))
	})
}

func TestGetRequestIdFromIncomingHeader(t *testing.T) {
	handler := requestIdHandler()
	handlerWithMiddleware := RequestIdMiddleware(handler)
	handlerWithMiddleware = setDummyRequestContext(handlerWithMiddleware)

	router := http.NewServeMux()
	router.Handle("/request", handlerWithMiddleware)

	req := httptest.NewRequest("GET", "/request", nil)
	req.Header.Set("X-Request-Id", "123")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if body := w.Body.String(); body != "123" {
		t.Errorf("handler returned unexpected body: got %v want %v", body, "123")
	}
}

func TestGenerateRequestIdIfNotProvided(t *testing.T) {
	handler := requestIdHandler()
	handlerWithMiddleware := RequestIdMiddleware(handler)
	handlerWithMiddleware = setDummyRequestContext(handlerWithMiddleware)

	router := http.NewServeMux()
	router.Handle("/request", handlerWithMiddleware)

	req := httptest.NewRequest("GET", "/request", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if body := w.Body.String(); body == "" {
		t.Errorf("handler returned empty body")
	}
}
