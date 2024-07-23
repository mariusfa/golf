package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mariusfa/golf/auth"
)

func helloAuthHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
}

type fakeUserRepository struct {
	UserList map[string]*User
}

func (fur *fakeUserRepository) FindById(id string) (*User, error) {
	user, exists := fur.UserList[id]
	if !exists {
		return &User{}, errors.New("User not found")
	}
	return user, nil
}

type fakeLogger struct {
	InfoMessage  string
	ErrorMessage string
}

func (fl *fakeLogger) Info(message string, requestId string) { fl.InfoMessage = message }
func (fl *fakeLogger) Error(message string, requestId string) { fl.ErrorMessage = message }

func TestFindUser(t *testing.T) {
	fakeLogger := &fakeLogger{}
	fakeUesrRepo := &fakeUserRepository{
		UserList: map[string]*User{
			"123": {Id: "123", Name: "John"},
		},
	}
	token, _ := auth.CreateToken("123", "secret", nil)
	headerValue := "Bearer " + token

	authParams := NewAuthParams("secret", fakeUesrRepo, fakeLogger)
	handler := helloAuthHandler()
	handlerWithMiddleware := Auth(handler, authParams)

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
	t.Errorf("Test not implemented")
}

func TestMissingUser(t *testing.T) {
	t.Errorf("Test not implemented")
}
