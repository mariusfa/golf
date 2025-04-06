package middleware

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mariusfa/golf/auth"
	"github.com/mariusfa/golf/request"
)

func helloAuthHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
}

type fakeUserRepository struct {
	UserList map[string]auth.AuthUser
}

func (fur *fakeUserRepository) FindAuthUserById(id string) (auth.AuthUser, error) {
	user, exists := fur.UserList[id]
	if !exists {
		return auth.AuthUser{}, errors.New("User not found")
	}
	return user, nil
}

type fakeLogger struct {
	InfoMessage  string
	ErrorMessage string
}

func (fl *fakeLogger) Info(message string, requestId string)  { fl.InfoMessage = message }
func (fl *fakeLogger) Error(message string, requestId string) { fl.ErrorMessage = message }

func setDummyAuthContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestIdCtx := request.NewRequestIdCtx("")
		ctx = context.WithValue(ctx, request.RequestIdCtxKey, requestIdCtx)

		sessionCtx := &request.SessionCtx{}
		ctx = context.WithValue(ctx, request.SessionCtxKey, sessionCtx)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func TestFindUser(t *testing.T) {
	fakeLogger := &fakeLogger{}
	fakeUserRepo := &fakeUserRepository{
		UserList: map[string]auth.AuthUser{
			"123": {Id: "123", Name: "John"},
		},
	}
	token, _ := auth.CreateToken("123", "secret", nil)
	headerValue := "Bearer " + token

	authParams := NewAuthParams("secret", fakeUserRepo, fakeLogger)
	handler := helloAuthHandler()
	handlerWithMiddleware := Auth(handler, authParams)
	handlerWithMiddleware = setDummyAuthContext(handlerWithMiddleware)

	router := http.NewServeMux()
	router.Handle("/hello", handlerWithMiddleware)

	req := httptest.NewRequest("GET", "/hello", nil)
	req.Header.Set("Authorization", headerValue)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}

	if fakeLogger.ErrorMessage != "" {
		t.Errorf("authParams.Logger.ErrorMessage is not empty")
	}

	if fakeLogger.InfoMessage != "User authenticated" {
		t.Errorf("authParams.Logger.InfoMessage is not 'User authenticated'")
	}
}

func TestMissingHeader(t *testing.T) {
	fakeLogger := &fakeLogger{}
	authParams := NewAuthParams("secret", &fakeUserRepository{}, fakeLogger)
	handler := helloAuthHandler()
	handlerWithMiddleware := Auth(handler, authParams)

	router := http.NewServeMux()
	router.Handle("/hello", handlerWithMiddleware)

	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}

	if fakeLogger.ErrorMessage != "Missing Authorization header" {
		t.Errorf("authParams.Logger.ErrorMessage is not 'Missing Authorization header'")
	}

	if fakeLogger.InfoMessage != "" {
		t.Errorf("authParams.Logger.InfoMessage is not empty")
	}
}

func TestMissingBearerString(t *testing.T) {
	fakeLogger := &fakeLogger{}
	authParams := NewAuthParams("secret", &fakeUserRepository{}, fakeLogger)
	handler := helloAuthHandler()
	handlerWithMiddleware := Auth(handler, authParams)

	router := http.NewServeMux()
	router.Handle("/hello", handlerWithMiddleware)

	req := httptest.NewRequest("GET", "/hello", nil)
	req.Header.Set("Authorization", "Bearer ")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}

	if fakeLogger.ErrorMessage != "Missing token" {
		t.Errorf("authParams.Logger.ErrorMessage is not 'Missing token'")
	}

	if fakeLogger.InfoMessage != "" {
		t.Errorf("authParams.Logger.InfoMessage is not empty")
	}
}

func TestMalformedHeader(t *testing.T) {
	fakeLogger := &fakeLogger{}
	authParams := NewAuthParams("secret", &fakeUserRepository{}, fakeLogger)
	handler := helloAuthHandler()
	handlerWithMiddleware := Auth(handler, authParams)

	router := http.NewServeMux()
	router.Handle("/hello", handlerWithMiddleware)

	req := httptest.NewRequest("GET", "/hello", nil)
	req.Header.Set("Authorization", "bad")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}

	if fakeLogger.ErrorMessage != "Invalid Authorization header" {
		t.Errorf("authParams.Logger.ErrorMessage is not 'Invalid Authorization header'")
	}

	if fakeLogger.InfoMessage != "" {
		t.Errorf("authParams.Logger.InfoMessage is not empty")
	}
}

func TestInvalidToken(t *testing.T) {
	fakeLogger := &fakeLogger{}
	fakeUserRepo := &fakeUserRepository{
		UserList: map[string]auth.AuthUser{
			"123": {Id: "123", Name: "John"},
		},
	}
	token := "123"
	headerValue := "Bearer " + token

	authParams := NewAuthParams("secret", fakeUserRepo, fakeLogger)
	handler := helloAuthHandler()
	handlerWithMiddleware := Auth(handler, authParams)

	router := http.NewServeMux()
	router.Handle("/hello", handlerWithMiddleware)

	req := httptest.NewRequest("GET", "/hello", nil)
	req.Header.Set("Authorization", headerValue)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusUnauthorized)
	}

	if fakeLogger.ErrorMessage != "Error parsing token" {
		t.Errorf("authParams.Logger.ErrorMessage is not 'Error parsing token'")
	}

	if fakeLogger.InfoMessage != "" {
		t.Errorf("authParams.Logger.InfoMessage is not empty")
	}
}

func TestMissingUser(t *testing.T) {
	fakeLogger := &fakeLogger{}
	fakeUserRepo := &fakeUserRepository{
		UserList: map[string]auth.AuthUser{
			"123": {Id: "123", Name: "John"},
		},
	}
	token, _ := auth.CreateToken("1234", "secret", nil)
	headerValue := "Bearer " + token

	authParams := NewAuthParams("secret", fakeUserRepo, fakeLogger)
	handler := helloAuthHandler()
	handlerWithMiddleware := Auth(handler, authParams)

	router := http.NewServeMux()
	router.Handle("/hello", handlerWithMiddleware)

	req := httptest.NewRequest("GET", "/hello", nil)
	req.Header.Set("Authorization", headerValue)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusUnauthorized)
	}

	if fakeLogger.ErrorMessage != "Error finding user" {
		t.Errorf("authParams.Logger.ErrorMessage is not 'Error finding user'")
	}

	if fakeLogger.InfoMessage != "" {
		t.Errorf("authParams.Logger.InfoMessage is not empty")
	}
}
