package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mariusfa/golf/request"
)

func helloAccessHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
}

func setDummyContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		sessionCtx := ctx.Value(request.SessionCtxKey).(*request.SessionCtx)
		sessionCtx.Id = "dummy"

		next.ServeHTTP(w, r)
	})
}

func TestAccessMiddleware(t *testing.T) {
	handler := helloAccessHandler()
	handlerWithMiddleware := AccessLogMiddleware(handler)

	router := http.NewServeMux()
	router.Handle("/hello", handlerWithMiddleware)

	req, _ := http.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// TODO: Andreas. Her ble det sjekket mye på loggeren før, må erstattes med annen testing
	/*	if loggerFake.Status == 0 {
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
	*/
}

func TestAccessMiddlwareWithDummyUser(t *testing.T) {
	handler := helloAccessHandler()
	handler = setDummyContext(handler)
	handlerWithMiddleware := AccessLogMiddleware(handler)

	router := http.NewServeMux()
	router.Handle("/hello", handlerWithMiddleware)

	req, _ := http.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// TODO: Andreas. Her ble det sjekket mye på loggeren før, må erstattes med annen testing
	/*	if loggerFake.UserId != "dummy" {
			t.Errorf("loggerFake.UserId is not dummy")
		}
	*/
}
