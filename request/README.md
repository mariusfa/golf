# Request

This package provides internal context management utilities for HTTP requests. It is primarily used by other packages in the framework and not intended for direct use.

## Components

### Context Keys
Defines type-safe context keys:
- `SessionCtxKey` - For user session data
- `RequestIdCtxKey` - For request ID tracking

### RequestIdCtx
Manages request ID in context:
```go
type RequestIdCtx struct {
    RequestId string
}
```

### SessionCtx
Manages user session data in context:
```go
type SessionCtx struct {
    Id       string
    Name     string
    Email    string
    Username string
}
```

## Usage

This package is used internally by:
- `middleware` package for request ID and authentication
- `logging` packages for context extraction
- Other framework components that need request context data

For typical application development, use the middleware package instead of accessing these utilities directly.