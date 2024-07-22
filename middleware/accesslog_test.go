package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mariusfa/golf/auth"
)

type fakeAccessLogger struct {
	DurationMs    int
	Status        int
	RequestPath   string
	RequestMethod string
	RequestId     string
	UserId        string
}

func (fal *fakeAccessLogger) Info(durationMs int, status int, requestPath string, requestMethod string, requestId string, userId string) {
	fal.DurationMs = durationMs
	fal.Status = status
	fal.RequestPath = requestPath
	fal.RequestMethod = requestMethod
	fal.RequestId = requestId
	fal.UserId = userId
}

func helloAccessHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
}

func setDummyUserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctxValue := ctx.Value(UserKey)
		userContext := ctxValue.(*auth.UserContext)
		userContext.Id = "dummy"
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func TestAccessMiddleware(t *testing.T) {
	loggerFake := &fakeAccessLogger{}

	// set loggerfake a negative value to check if it is updated
	loggerFake.DurationMs = -1

	handler := helloAccessHandler()
	handlerWithMiddleware := AccessLogMiddleware(handler, loggerFake)

	router := http.NewServeMux()
	router.Handle("/hello", handlerWithMiddleware)

	req, _ := http.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if loggerFake.Status == 0 {
		t.Errorf("loggerFake.Status is 0")
	}
	if loggerFake.RequestPath == "" {
		t.Errorf("loggerFake.RequestPath is empty")
	}
	if loggerFake.RequestMethod == "" {
		t.Errorf("loggerFake.RequestMethod is empty")
	}
	if loggerFake.DurationMs < 0 {
		t.Errorf("loggerFake.DurationMs is less than 0")
	}
}

func TestAccessMiddlwareWithDummyUser(t *testing.T) {
	loggerFake := &fakeAccessLogger{}

	handler := helloAccessHandler()
	handler = setDummyUserContext(handler)
	handlerWithMiddleware := AccessLogMiddleware(handler, loggerFake)

	router := http.NewServeMux()
	router.Handle("/hello", handlerWithMiddleware)

	req, _ := http.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if loggerFake.UserId != "dummy" {
		t.Errorf("loggerFake.UserId is not dummy")
	}
}
