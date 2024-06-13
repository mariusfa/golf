package httpclient

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestDelete(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", "http://localhost:8080/1", httpmock.NewStringResponder(204, ""))

	fakeLogger := &fakeLogger{}

	client := NewHttpClient(fakeLogger)

	requestId := "test"
	url := "http://localhost:8080/1"
	headers := map[string]string{}
	getRequest := NewDeleteRequest(requestId, headers, url)
	err := client.Delete(getRequest)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if fakeLogger.RequestId == "" {
		t.Errorf("fakeLogger.RequestId is empty")
	}
	if fakeLogger.RequestPath != "http://localhost:8080/1" {
		t.Errorf("Expected http://localhost:8080/1, got %v", fakeLogger.RequestPath)
	}
	if fakeLogger.RequestMethod != "DELETE" {
		t.Errorf("Expected DELETE, got %v", fakeLogger.RequestMethod)
	}

	if fakeLogger.DurationMs == "" {
		t.Errorf("fakeLogger.DurationMs is empty")
	}
	if fakeLogger.Status != http.StatusNoContent {
		t.Errorf("Expected %v, got %v", http.StatusOK, fakeLogger.Status)
	}
}
