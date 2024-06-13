package httpclient

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

type MyDto struct {
	Name string `json:"name"`
}

type fakeLogger struct {
	RequestId     string
	RequestPath   string
	RequestMethod string
	DurationMs    string
	Status        int
	ResponseBody  string
}

func (fl *fakeLogger) RequestInfo(requestId string, requestPath string, requestMethod string) {
	fl.RequestId = requestId
	fl.RequestPath = requestPath
	fl.RequestMethod = requestMethod
}

func (fl *fakeLogger) ResponseInfo(requestId string, durationMs string, status int, responseBody string) {
	fl.RequestId = requestId
	fl.DurationMs = durationMs
	fl.Status = status
	fl.ResponseBody = responseBody
}

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
	err := client.GetJsonObject(getRequest, &dto)

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
