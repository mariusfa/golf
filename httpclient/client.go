package httpclient

import (
	"net/http"
	"time"

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
