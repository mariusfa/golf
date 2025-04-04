package httpclient

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestPostJsonObject(t *testing.T) {
	// Test without response body
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "http://localhost:8080",
		func(h *http.Request) (*http.Response, error) {
			if h.Header.Get("Content-Type") != "application/json" {
				t.Errorf("Expected Content-Type application/json, got %s", h.Header.Get("Content-Type"))
			}

			return httpmock.NewStringResponse(200, ""), nil
		})

	requestDto := MyDto{Name: "Crazy Test"}
	url := "http://localhost:8080"
	headers := map[string]string{"Content-Type": "application/json"}

	postRequest := NewPostRequest(headers, url, requestDto)
	err := NewHttpClient().PostJson(context.Background(), postRequest, nil)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestPostJsonObjectWithResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "http://localhost:8080",
		func(h *http.Request) (*http.Response, error) {
			if h.Header.Get("Content-Type") != "application/json" {
				t.Errorf("Expected Content-Type application/json, got %s", h.Header.Get("Content-Type"))
			}

			return httpmock.NewJsonResponse(200, MyDto{Name: "Hello!"})
		})

	headers := map[string]string{"Content-Type": "application/json"}
	var responseDto MyDto
	postRequest := NewPostRequest(headers, "http://localhost:8080", MyDto{Name: "Crazy Test"})
	err := NewHttpClient().PostJson(context.Background(), postRequest, &responseDto)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if responseDto.Name != "Hello!" {
		t.Errorf("Expected Crazy Test, got %v", responseDto.Name)
	}
}
