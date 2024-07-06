package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakeAccessLogger struct {
	DurationMs    int
	Status        int
	RequestPath   string
	RequestMethod string
	RequestId     string
}

func (fal *fakeAccessLogger) Info(durationMs int, status int, requestPath string, requestMethod string, requestId string) {
	fal.DurationMs = durationMs
	fal.Status = status
	fal.RequestPath = requestPath
	fal.RequestMethod = requestMethod
	fal.RequestId = requestId
}

func helloAccessHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
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
