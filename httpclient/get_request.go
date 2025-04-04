package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/mariusfa/golf/logging/transactionlog"
	"io"
	"net/http"
	"time"
)

type GetRequest struct {
	UserId  string
	Headers map[string]string
	Url     string
}

func NewGetRequest(headers map[string]string, url string) *GetRequest {
	return &GetRequest{Headers: headers, Url: url}
}

func (c *HttpClient) GetJson(ctx context.Context, request *GetRequest, dto any) error {
	httpAction := func() error {
		return c.getJsonPlain(ctx, request, dto)
	}
	bulkheadAction := func() error {
		return c.bulkhead.Execute(httpAction)
	}
	circuitbreakerAction := func() error {
		return c.circuitbreaker.Execute(bulkheadAction)
	}

	if err := circuitbreakerAction(); err != nil {
		return fmt.Errorf("httpclient failed to get: %w", err)
	}

	return nil
}

func (c *HttpClient) getJsonPlain(ctx context.Context, request *GetRequest, dto any) error {
	req, err := http.NewRequest("GET", request.Url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}

	transactionlog.RequestInfo(ctx, "GET", request.Url, "")

	start := time.Now()
	resp, err := c.client.Do(req)
	duratonMs := time.Since(start).Milliseconds()

	if err != nil {
		transactionlog.ResponseInfo(ctx, int(duratonMs), 0, "")
		return fmt.Errorf("failed to do request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		transactionlog.ResponseInfo(ctx, int(duratonMs), resp.StatusCode, "")
		return fmt.Errorf("%s from %s", resp.Status, request.Url)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		transactionlog.ResponseInfo(ctx, int(duratonMs), resp.StatusCode, "")
		return fmt.Errorf("failed to read response body: %w", err)
	}

	bodyString := string(bodyBytes)

	transactionlog.ResponseInfo(ctx, int(duratonMs), resp.StatusCode, bodyString)

	bodyReader := bytes.NewReader(bodyBytes)

	if err := json.NewDecoder(bodyReader).Decode(dto); err != nil {
		return fmt.Errorf("failed to decode response body: %w", err)
	}

	return nil
}
