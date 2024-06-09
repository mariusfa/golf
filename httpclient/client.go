package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mariusfa/golf/httpclient/bulkhead"
	circuitbreaker "github.com/mariusfa/golf/httpclient/circuit-breaker"
)

type logger interface {
	RequestInfo(requestId string, requestPath string, requestMethod string)
	ResponseInfo(requestId string, durationMs string, status int, responseBody string)
}

type HttpClient struct {
	headers        map[string]string
	client         *http.Client
	logger         logger
	bulkhead       *bulkhead.Bulkhead
	circuitbreaker *circuitbreaker.CircuitBreaker
}

func NewHttpClient(logger logger) *HttpClient {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	timeout := 5 * time.Second
	client := &http.Client{Timeout: timeout}
	bulkhead := bulkhead.NewBulkhead(bulkhead.Options{})
	circuitbreaker := circuitbreaker.NewCircuitBreaker(circuitbreaker.Options{})
	return &HttpClient{headers: headers, client: client, logger: logger, bulkhead: bulkhead, circuitbreaker: circuitbreaker}
}

func (c *HttpClient) GetJsonObject(url string, dto any) error {
	httpAction := func() error {
		return c.getJsonObjectPlain(url, dto)
	}
	bulkheadAction := func() error {
		return c.bulkhead.Execute(httpAction)
	}
	circuitbreakerAction := func() error {
		return c.circuitbreaker.Execute(bulkheadAction)
	}
	return circuitbreakerAction()
}

func (c *HttpClient) getJsonObjectPlain(url string, dto any) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}
	requestId := uuid.New().String()
	c.logger.RequestInfo(requestId, url, "GET")
	start := time.Now()
	resp, err := c.client.Do(req)
	duratonMs := time.Since(start).Milliseconds()
	if err != nil {
		c.logger.ResponseInfo(requestId, fmt.Sprintf("%d", duratonMs), 0, "")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.logger.ResponseInfo(requestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, "")
		return errors.New(fmt.Sprintf("%s from %s", resp.Status, url))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.ResponseInfo(requestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, "")
		return err
	}

	bodyString := string(bodyBytes)

	c.logger.ResponseInfo(requestId, fmt.Sprintf("%d", duratonMs), resp.StatusCode, bodyString)

	bodyReader := bytes.NewReader(bodyBytes)

	return json.NewDecoder(bodyReader).Decode(dto)
}
