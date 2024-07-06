package httpclient

import (
	"net/http"
	"time"

	"github.com/mariusfa/golf/httpclient/bulkhead"
	circuitbreaker "github.com/mariusfa/golf/httpclient/circuit-breaker"
)

type logger interface {
	RequestInfo(requestId string, requestMethod string, requestPath string, requestBody string, userId string)
	ResponseInfo(requestId string, durationMs string, status int, responseBody string, userId string)
}

type HttpClient struct {
	client         *http.Client
	logger         logger
	bulkhead       *bulkhead.Bulkhead
	circuitbreaker *circuitbreaker.CircuitBreaker
}

func NewHttpClient(logger logger) *HttpClient {
	timeout := 15 * time.Second
	client := &http.Client{Timeout: timeout}
	bulkhead := bulkhead.NewBulkhead(bulkhead.Options{})
	circuitbreaker := circuitbreaker.NewCircuitBreaker(circuitbreaker.Options{})
	return &HttpClient{client: client, logger: logger, bulkhead: bulkhead, circuitbreaker: circuitbreaker}
}
