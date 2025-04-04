package httpclient

import (
	"net/http"
	"time"

	"github.com/mariusfa/golf/httpclient/bulkhead"
	"github.com/mariusfa/golf/httpclient/circuit-breaker"
)

type HttpClient struct {
	client         *http.Client
	bulkhead       *bulkhead.Bulkhead
	circuitbreaker *circuitbreaker.CircuitBreaker
}

func NewHttpClient() *HttpClient {
	timeout := 15 * time.Second
	client := &http.Client{Timeout: timeout}
	bh := bulkhead.NewBulkhead(bulkhead.Options{})
	cb := circuitbreaker.NewCircuitBreaker(circuitbreaker.Options{})
	return &HttpClient{
		client:         client,
		bulkhead:       bh,
		circuitbreaker: cb,
	}
}
