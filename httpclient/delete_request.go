package httpclient

import (
	"fmt"
	"net/http"
	"time"
)

type DeleteRequest struct {
	RequestId string
	Headers   map[string]string
	Url       string
}

func NewDeleteRequest(requestId string, headers map[string]string, url string) *DeleteRequest {
	return &DeleteRequest{RequestId: requestId, Headers: headers, Url: url}
}

func (c *HttpClient) Delete(request *DeleteRequest) error {
	httpAction := func() error {
		return c.deletePlain(request)
	}
	bulkheadAction := func() error {
		return c.bulkhead.Execute(httpAction)
	}
	circuitbreakerAction := func() error {
		return c.circuitbreaker.Execute(bulkheadAction)
	}
	return circuitbreakerAction()
}

func (c *HttpClient) deletePlain(request *DeleteRequest) error {
	req, err := http.NewRequest("DELETE", request.Url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}
	c.logger.RequestInfo(request.RequestId, "DELETE", request.Url, "")
	start := time.Now()
	resp, err := c.client.Do(req)
	duratonMs := time.Since(start).Milliseconds()
	if err != nil {
		c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), 0, "")
		return fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, "")
		return fmt.Errorf("%s from %s", resp.Status, request.Url)
	}

	c.logger.ResponseInfo(request.RequestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, "")

	return nil
}
