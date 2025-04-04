package httpclient

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetJsonObject(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "http://localhost:8080",
		func(request *http.Request) (*http.Response, error) {
			if request.Header.Get("Accept") != "application/json" {
				t.Errorf("Expected application/json, got %v", request.Header.Get("Accept"))
			}
			return httpmock.NewJsonResponse(http.StatusOK, httpmock.File("./mocks/name.json"))
		})
	var dto MyDto
	headers := map[string]string{"Accept": "application/json"}
	getRequest := NewGetRequest(headers, "http://localhost:8080")
	err := NewHttpClient().GetJson(context.Background(), getRequest, &dto)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if dto.Name != "Crazy Test" {
		t.Errorf("Expected Crazy Test, got %v", dto.Name)
	}
}

func TestGetJsonList(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"GET",
		"http://localhost:8080",
		httpmock.NewJsonResponderOrPanic(http.StatusOK, httpmock.File("./mocks/names.json")))

	var dto []MyDto

	headers := map[string]string{"Accept": "application/json"}
	getRequest := NewGetRequest(headers, "http://localhost:8080")
	err := NewHttpClient().GetJson(context.Background(), getRequest, &dto)

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
}
