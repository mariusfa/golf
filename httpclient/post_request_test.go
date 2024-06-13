package httpclient

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestPostJsonObject(t *testing.T) {
	// Test without response body
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "http://localhost:8080", httpmock.NewStringResponder(200, ""))
	fakeLogger := &fakeLogger{}

	client := NewHttpClient(fakeLogger)
	requestDto := MyDto{Name: "Crazy Test"}

	requestId := "test"
	url := "http://localhost:8080"
	headers := map[string]string{"Content-Type": "application/json"}

	postRequest := NewPostRequest(requestId, headers, url, requestDto)
	err := client.PostJson(postRequest, nil)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if fakeLogger.RequestId == "" {
		t.Errorf("fakeLogger.RequestId is empty")
	}
	if fakeLogger.RequestPath != "http://localhost:8080" {
		t.Errorf("Expected http://localhost:8080, got %v", fakeLogger.RequestPath)
	}
	if fakeLogger.RequestMethod != "POST" {
		t.Errorf("Expected POST, got %v", fakeLogger.RequestMethod)
	}
	if fakeLogger.DurationMs == "" {
		t.Errorf("fakeLogger.DurationMs is empty")
	}
	if fakeLogger.Status != 200 {
		t.Errorf("Expected 200, got %v", fakeLogger.Status)
	}
	if fakeLogger.RequestBody == "" {
		t.Errorf("fakeLogger.RequestBody is empty")
	}
}

func TestPostJsonObjectWithResponse(t *testing.T) {
	// Test with response body
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "http://localhost:8080", httpmock.NewJsonResponderOrPanic(201, MyDto{Name: "Crazy Test"}))
	fakeLogger := &fakeLogger{}

	client := NewHttpClient(fakeLogger)
	requestDto := MyDto{Name: "Crazy Test"}

	requestId := "test"
	url := "http://localhost:8080"
	headers := map[string]string{"Content-Type": "application/json"}

	var responseDto MyDto
	postRequest := NewPostRequest(requestId, headers, url, requestDto)
	err := client.PostJson(postRequest, &responseDto)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if fakeLogger.RequestId == "" {
		t.Errorf("fakeLogger.RequestId is empty")
	}
	if fakeLogger.RequestPath != "http://localhost:8080" {
		t.Errorf("Expected http://localhost:8080, got %v", fakeLogger.RequestPath)
	}
	if fakeLogger.RequestMethod != "POST" {
		t.Errorf("Expected POST, got %v", fakeLogger.RequestMethod)
	}
	if fakeLogger.DurationMs == "" {
		t.Errorf("fakeLogger.DurationMs is empty")
	}
	if fakeLogger.Status != 201 {
		t.Errorf("Expected 201, got %v", fakeLogger.Status)
	}
	if fakeLogger.RequestBody == "" {
		t.Errorf("fakeLogger.RequestBody is empty")
	}
	if fakeLogger.ResponseBody == "" {
		t.Errorf("fakeLogger.ResponseBody is empty")
	}
	if responseDto.Name != "Crazy Test" {
		t.Errorf("Expected Crazy Test, got %v", responseDto.Name)
	}
}
