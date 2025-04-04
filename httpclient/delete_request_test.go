package httpclient

import (
	"context"
	"github.com/mariusfa/golf/request"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestDelete(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", "http://localhost:8080/1", httpmock.NewStringResponder(204, ""))

	client := NewHttpClient()

	url := "http://localhost:8080/1"
	headers := map[string]string{}
	getRequest := NewDeleteRequest(headers, url)
	ctx := request.WithRequestIdCtx(context.Background(), "123")
	err := client.Delete(ctx, getRequest)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
