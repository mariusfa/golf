package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
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
		t.Errorf("authParams.Logger.ErrorMesasge is not 'Missing Authorization header'")
	}

	if fakeLogger.InfoMessage != "" {
		t.Errorf("authParams.Logger.InfoMessage is not empty")
	}
}

func TestMissingBearerString(t *testing.T) {
	t.Errorf("Test not implemented")
}

func TestInvalidToken(t *testing.T) {
	t.Errorf("Test not implemented")
}

func TestMissingUser(t *testing.T) {
	t.Errorf("Test not implemented")
}
