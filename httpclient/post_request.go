package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type PostRequest struct {
	RequestId string
	Headers   map[string]string
	Url       string
	Body      any
}

func NewPostRequest(requestId string, headers map[string]string, url string, body any) *PostRequest {
	return &PostRequest{RequestId: requestId, Headers: headers, Url: url, Body: body}
}

func (c *HttpClient) PostJson(request *PostRequest, responseDto any) error {
	httpAction := func() error {
		return c.postJsonPlain(request, responseDto)
	}
	bulkheadAction := func() error {
		return c.bulkhead.Execute(httpAction)
	}
	circuitbreakerAction := func() error {
		return c.circuitbreaker.Execute(bulkheadAction)
	}
	if err := circuitbreakerAction(); err != nil {
		return fmt.Errorf("httpclient failed to post: %w", err)
	}
	return nil
}

func (c *HttpClient) postJsonPlain(request *PostRequest, responseDto any) error {
	body, err := json.Marshal(request.Body)
	if err != nil {
		return fmt.Errorf("failed to marshal body: %w", err)
	}
	req, err := http.NewRequest("POST", request.Url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}
	c.logger.RequestInfo(request.RequestId, "POST", request.Url, string(body))
	start := time.Now()
	resp, err := c.client.Do(req)
	duratonMs := time.Since(start).Milliseconds()
	if err != nil {
		c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), 0, "")
		return fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, "")
		return fmt.Errorf("%s from %s", resp.Status, request.Url)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, "")
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if len(bodyBytes) == 0 {
		c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, "")
		return nil
	}

	if err := json.Unmarshal(bodyBytes, responseDto); err != nil {
		c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, string(bodyBytes))
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, string(bodyBytes))
	return nil
}
