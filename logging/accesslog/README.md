# Access Log

This package provides HTTP access logging for web applications.

## Usage

Initialize the logger in main and use in HTTP handlers:

```go
import "github.com/mariusfa/golf/logging/accesslog"

// In main function
appName := "todo"
accesslog.SetAppName(appName)

// Used in HTTP middleware or handlers
duration := 150 // milliseconds
status := 200
path := "/api/users"
method := "GET"

accesslog.Info(duration, status, path, method)
```

## Purpose

Specifically designed for:
- HTTP request/response logging
- Performance monitoring (request duration)
- Access pattern analysis
- API usage tracking

## Integration

This package is automatically used by the middleware package for request logging.
