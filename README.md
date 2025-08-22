# Golf - Go Library and Framework

A collection of packages for building robust web applications in Go with built-in resilience patterns, authentication, and comprehensive logging.

## Quick Start

```go
go get github.com/mariusfa/golf
```

```go
import (
    "github.com/mariusfa/golf/middleware"
    "github.com/mariusfa/golf/health"
)

func main() {
    router := http.NewServeMux()
    
    // Health check
    health.RegisterRoute(router)
    
    // Protected routes
    authParams := middleware.NewAuthParams(jwtSecret, userRepo)
    router.Handle("GET /api/users", middleware.PrivateRoute(getUsersHandler, authParams))
    
    // Public routes  
    router.Handle("GET /api/health", middleware.PublicRoute(healthHandler))
    
    http.ListenAndServe(":8080", router)
}
```

## Core Packages

| Package | Purpose |
|---------|---------|
| **auth** | JWT token management and user authentication |
| **config** | Environment-based configuration with struct mapping |
| **database** | PostgreSQL with migrations, Docker, and testcontainers |
| **health** | Standard health check endpoint |
| **httpclient** | Resilient HTTP client with circuit breaker and bulkhead |
| **logging** | Structured logging (access, app, trace, transaction) |
| **metrics** | Prometheus metrics integration |
| **middleware** | Pre-configured middleware chains for auth and logging |
| **routergroup** | Enhanced routing with middleware composition |

## Key Features

- **Resilience**: Circuit breaker and bulkhead patterns in HTTP client
- **Authentication**: JWT-based auth with middleware integration
- **Database**: Template-based migrations with baseline/standard strategy
- **Logging**: Structured JSON logging with multiple specialized loggers
- **Configuration**: Automatic struct field to environment variable mapping
- **Testing**: Comprehensive testcontainers support

## Examples & Development

- Examples: [golf-examples](https://github.com/mariusfa/golf-examples)
- Run tests: `go test ./...`
- Format code: `go fmt ./...`
- Tasks: See [TODO.md](TODO.md)

## Versioning

Currently using minor versions regardless of breaking changes. Will switch to semantic versioning when stable.

