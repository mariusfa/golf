# Logging Utils

This package provides internal utilities for the logging packages. It is not intended for direct use in applications.

## Components

### Context Extraction
Extracts user and request data from HTTP context:
```go
func ExtractFromContext(ctx context.Context) (username string, requestId string)
```

### Structured Logger Configuration
Provides a pre-configured slog logger with custom field mappings:
- `time` → `timestamp`
- `level` → `log_level` 
- `msg` → `payload`

```go
func NewSlogger() *slog.Logger
func ReplaceDefaultKeys(groups []string, attr slog.Attr) slog.Attr
```

## Usage

This package is used internally by:
- `accesslog` - HTTP access logging
- `applog` - Application logging
- `tracelog` - Request tracing
- `transactionlog` - Transaction logging

For application logging, use the specific logging packages instead of these utilities directly.