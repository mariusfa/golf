package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func requestIdHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := r.Header.Get("X-Request-Id")
		w.Write([]byte(requestId))
	})
}

func TestGetRequestIdFromIncomingHeader(t *testing.T) {
	handler := requestIdHandler()
	handlerWithMiddleware := RequestIdMiddleware(handler)

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
