package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type GetRequest struct {
	RequestId string
	UserId    string
	Headers   map[string]string
	Url       string
}

func NewGetRequest(requestId string, headers map[string]string, url string) *GetRequest {
	return &GetRequest{RequestId: requestId, Headers: headers, Url: url}
}

func (c *HttpClient) GetJson(request *GetRequest, dto any) error {
	httpAction := func() error {
		return c.getJsonPlain(request, dto)
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

func (c *HttpClient) getJsonPlain(request *GetRequest, dto any) error {
	req, err := http.NewRequest("GET", request.Url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}
	c.logger.RequestInfo(request.RequestId, "GET", request.Url, "", request.UserId)
	start := time.Now()
	resp, err := c.client.Do(req)
	duratonMs := time.Since(start).Milliseconds()
	if err != nil {
		c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), 0, "", request.UserId)
		return fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, "", request.UserId)
		return fmt.Errorf("%s from %s", resp.Status, request.Url)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, "", request.UserId)
		return fmt.Errorf("failed to read response body: %w", err)
	}

	bodyString := string(bodyBytes)

	c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, bodyString, request.UserId)

	bodyReader := bytes.NewReader(bodyBytes)

	if err := json.NewDecoder(bodyReader).Decode(dto); err != nil {
		return fmt.Errorf("failed to decode response body: %w", err)
	}

	return nil
}
