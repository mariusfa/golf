package httpclient

import (
	"context"
	"fmt"
	"github.com/mariusfa/golf/logging/transactionlog"
	"net/http"
	"time"
)

type DeleteRequest struct {
	Headers map[string]string
	Url     string
}

func NewDeleteRequest(headers map[string]string, url string) *DeleteRequest {
	return &DeleteRequest{Headers: headers, Url: url}
}

func (c *HttpClient) Delete(ctx context.Context, request *DeleteRequest) error {
	httpAction := func() error {
		return c.deletePlain(ctx, request)
	}
	bulkheadAction := func() error {
		return c.bulkhead.Execute(httpAction)
	}
	circuitbreakerAction := func() error {
		return c.circuitbreaker.Execute(bulkheadAction)
	}
	if err := circuitbreakerAction(); err != nil {
		return fmt.Errorf("httpclient failed to delete: %w", err)
	}
	return nil
}

func (c *HttpClient) deletePlain(ctx context.Context, request *DeleteRequest) error {
	req, err := http.NewRequest("DELETE", request.Url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}

	transactionlog.RequestInfo(ctx, "DELETE", request.Url, "")

	start := time.Now()
	resp, err := c.client.Do(req)
	duratonMs := time.Since(start).Milliseconds()

	if err != nil {
		transactionlog.ResponseInfo(ctx, int(duratonMs), 0, "")
		return fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		transactionlog.ResponseInfo(ctx, int(duratonMs), resp.StatusCode, "")
		return fmt.Errorf("%s from %s", resp.Status, request.Url)
	}

	transactionlog.ResponseInfo(ctx, int(duratonMs), resp.StatusCode, "")

	return nil
}
