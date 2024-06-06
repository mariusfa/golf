package metrics

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMetricsRoute(t *testing.T) {
	router := http.NewServeMux()
	RegisterRoute(router)

	req, err := http.NewRequest("GET", "/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
