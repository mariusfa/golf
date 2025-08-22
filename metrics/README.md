# Metrics

This package provides Prometheus metrics integration for web applications.

## Usage

Register the metrics endpoint:
```go
import (
    "net/http"
    "github.com/mariusfa/golf/metrics"
)

func main() {
    router := http.NewServeMux()
    metrics.RegisterRoute(router)
    http.ListenAndServe(":8080", router)
}
```

The metrics endpoint will be available at `GET /metrics` and returns Prometheus-formatted metrics data.

## Integration

This endpoint is commonly used by monitoring systems like Prometheus to scrape application metrics.