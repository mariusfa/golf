package httpclient

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetJsonObject(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "http://localhost:8080", httpmock.NewJsonResponderOrPanic(http.StatusOK, httpmock.File("./mocks/name.json")))

	fakeLogger := &fakeLogger{}

	client := NewHttpClient(fakeLogger)
	var dto MyDto

	requestId := "test"
	url := "http://localhost:8080"
	headers := map[string]string{"Accept": "application/json"}
	getRequest := NewGetRequest(requestId, headers, url)
	err := client.GetJson(getRequest, &dto)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if dto.Name != "Crazy Test" {
		t.Errorf("Expected Crazy Test, got %v", dto.Name)
	}

	if fakeLogger.RequestId == "" {
		t.Errorf("fakeLogger.RequestId is empty")
	}
	if fakeLogger.RequestPath != "http://localhost:8080" {
		t.Errorf("Expected http://localhost:8080, got %v", fakeLogger.RequestPath)
	}
	if fakeLogger.RequestMethod != "GET" {
		t.Errorf("Expected GET, got %v", fakeLogger.RequestMethod)
	}

	if fakeLogger.DurationMs == "" {
		t.Errorf("fakeLogger.DurationMs is empty")
	}
	if fakeLogger.Status != http.StatusOK {
		t.Errorf("Expected %v, got %v", http.StatusOK, fakeLogger.Status)
	}
	if fakeLogger.ResponseBody == "" {
		t.Errorf("fakeLogger.ResponseBody is empty")
	}
}

func TestGetJsonList(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "http://localhost:8080", httpmock.NewJsonResponderOrPanic(http.StatusOK, httpmock.File("./mocks/names.json")))

	fakeLogger := &fakeLogger{}

	client := NewHttpClient(fakeLogger)
	var dto []MyDto

	requestId := "test"

	url := "http://localhost:8080"
	headers := map[string]string{"Accept": "application/json"}
	getRequest := NewGetRequest(requestId, headers, url)
	err := client.GetJson(getRequest, &dto)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if len(dto) != 2 {
		t.Errorf("Expected 2, got %v", len(dto))
	}
	if dto[0].Name != "Crazy Test" {
		t.Errorf("Expected Crazy Test, got %v", dto[0].Name)
	}
	if dto[1].Name != "Crazy Test 2" {
		t.Errorf("Expected Crazy Test 2, got %v", dto[1].Name)
	}

	if fakeLogger.RequestId == "" {
		t.Errorf("fakeLogger.RequestId is empty")
	}
	if fakeLogger.RequestPath != "http://localhost:8080" {
		t.Errorf("Expected http://localhost:8080, got %v", fakeLogger.RequestPath)
	}
	if fakeLogger.RequestMethod != "GET" {
		t.Errorf("Expected GET, got %v", fakeLogger.RequestMethod)
	}

	if fakeLogger.DurationMs == "" {
		t.Errorf("fakeLogger.DurationMs is empty")
	}
	if fakeLogger.Status != http.StatusOK {
		t.Errorf("Expected %v, got %v", http.StatusOK, fakeLogger.Status)
	}
	if fakeLogger.ResponseBody == "" {
		t.Errorf("fakeLogger.ResponseBody is empty")
	}
}
