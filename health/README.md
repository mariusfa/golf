# Health

This package provides a standard health check endpoint for web applications.

## Features

- Simple HTTP GET endpoint at `/health`
- Returns HTTP 200 status with "OK" response
- Used by load balancers and monitoring systems to verify application health

## Usage

```go
import (
	"net/http"
	"github.com/mariusfa/golf/health"
)

func main() {
	router := http.NewServeMux()
	health.RegisterRoute(router)
	http.ListenAndServe(":8080", router)
}
```

## Response

```
GET /health
HTTP/1.1 200 OK
Content-Type: text/plain

OK
```