package routergroup

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChainMiddleware_chainMiddleware(t *testing.T) {
	mw1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "mw1-before-")
			next.ServeHTTP(w, r)
			fmt.Fprint(w, "-mw1-after")
		})
	}
	mw2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "mw2-before-")
			next.ServeHTTP(w, r)
			fmt.Fprint(w, "-mw2-after")
		})
	}
	mw3 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "mw3-before-")
			next.ServeHTTP(w, r)
			fmt.Fprint(w, "-mw3-after")
		})
	}
	final := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "final")
	})

	wrapped := chainMiddleware(final, mw1, mw2, mw3)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	wrapped.ServeHTTP(w, r)
	expected := "mw1-before-mw2-before-mw3-before-final-mw3-after-mw2-after-mw1-after"

	if w.Body.String() != expected {
		t.Errorf("Expected %q but got %q", expected, w.Body.String())
	}
}

func TestRouterGroup_handleFunc(t *testing.T) {
	mux := http.NewServeMux()
	rg := NewRouterGroup("/api", mux)

	middleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "middleware-")
			next.ServeHTTP(w, r)
		})
	}

	rg.Use(middleware)

	rg.HandleFunc("GET", "/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello")
	})

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/hello", nil)
	mux.ServeHTTP(w, r)

	expected := "middleware-hello"
	if w.Body.String() != expected {
		t.Errorf("Expected %q but got %q", expected, w.Body.String())
	}
}

func TestRouterGroup_Use(t *testing.T) {
	mux := http.NewServeMux()
	rg := NewRouterGroup("/api", mux)

	if len(rg.Middlewares) != 0 {
		t.Errorf("Expected %d, but got %d", 0, len(rg.Middlewares))
	}

	middleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			fmt.Print("middleware")
			next.ServeHTTP(rw, r)
		})
	}

	rg.Use(middleware)

	if len(rg.Middlewares) != 1 {
		t.Errorf("Expected %d, but got %d", 1, len(rg.Middlewares))
	}
}
