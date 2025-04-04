package httpclient

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestPutJsonObject(t *testing.T) {
	// Test without response body
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("PUT", "http://localhost:8080",
		func(h *http.Request) (*http.Response, error) {
			if h.Header.Get("Content-Type") != "application/json" {
				t.Errorf("Expected Content-Type application/json, got %s", h.Header.Get("Content-Type"))
			}

			return httpmock.NewStringResponse(200, ""), nil
		})

	headers := map[string]string{"Content-Type": "application/json"}
	putRequest := NewPutRequest(headers, "http://localhost:8080", MyDto{Name: "Crazy Test"})
	ctx := context.Background()

	err := NewHttpClient().PutJson(ctx, putRequest, nil)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestPutJsonObjectWithResponse(t *testing.T) {
	// Test with response body
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("PUT", "http://localhost:8080",
		func(h *http.Request) (*http.Response, error) {
			if h.Header.Get("Content-Type") != "application/json" {
				t.Errorf("Expected Content-Type application/json, got %s", h.Header.Get("Content-Type"))
			}

			return httpmock.NewJsonResponse(200, MyDto{Name: "Hello!"})
		})

	headers := map[string]string{"Content-Type": "application/json"}
	var responseDto MyDto
	putRequest := NewPutRequest(headers, "http://localhost:8080", MyDto{Name: "Crazy Test"})
	err := NewHttpClient().PutJson(context.Background(), putRequest, &responseDto)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if responseDto.Name != "Hello!" {
		t.Errorf("Expected Crazy Test, got %v", responseDto.Name)
	}
}
