package httpclient

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestPutJsonObject(t *testing.T) {
	// Test without response body
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("PUT", "http://localhost:8080", httpmock.NewStringResponder(204, ""))
	fakeLogger := &fakeLogger{}

	client := NewHttpClient(fakeLogger)
	requestDto := MyDto{Name: "Crazy Test"}

	requestId := "test"
	url := "http://localhost:8080"
	headers := map[string]string{"Content-Type": "application/json"}

	putRequest := NewPutRequest(requestId, headers, url, requestDto)
	err := client.PutJson(putRequest, nil)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if fakeLogger.RequestId == "" {
		t.Errorf("fakeLogger.RequestId is empty")
	}
	if fakeLogger.RequestPath != "http://localhost:8080" {
		t.Errorf("Expected http://localhost:8080, got %v", fakeLogger.RequestPath)
	}
	if fakeLogger.RequestMethod != "PUT" {
		t.Errorf("Expected PUT, got %v", fakeLogger.RequestMethod)
	}
	if fakeLogger.DurationMs == "" {
		t.Errorf("fakeLogger.DurationMs is empty")
	}
	if fakeLogger.Status != 204 {
		t.Errorf("Expected 204, got %v", fakeLogger.Status)
	}
	if fakeLogger.RequestBody == "" {
		t.Errorf("fakeLogger.RequestBody is empty")
	}
}

func TestPutJsonObjectWithResponse(t *testing.T) {
	// Test with response body
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("PUT", "http://localhost:8080", httpmock.NewJsonResponderOrPanic(201, MyDto{Name: "Crazy Test"}))
	fakeLogger := &fakeLogger{}

	client := NewHttpClient(fakeLogger)
	requestDto := MyDto{Name: "Crazy Test"}

	requestId := "test"
	url := "http://localhost:8080"
	headers := map[string]string{"Content-Type": "application/json"}

	var responseDto MyDto
	putRequest := NewPutRequest(requestId, headers, url, requestDto)
	err := client.PutJson(putRequest, &responseDto)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if fakeLogger.RequestId == "" {
		t.Errorf("fakeLogger.RequestId is empty")
	}
	if fakeLogger.RequestPath != "http://localhost:8080" {
		t.Errorf("Expected http://localhost:8080, got %v", fakeLogger.RequestPath)
	}
	if fakeLogger.RequestMethod != "PUT" {
		t.Errorf("Expected PUT, got %v", fakeLogger.RequestMethod)
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
