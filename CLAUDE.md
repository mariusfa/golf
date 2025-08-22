# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is GOLF (Go Library and Framework) - a collection of packages for building web applications in Go. It provides reusable components for common web application concerns like authentication, configuration, database management, HTTP clients, logging, middleware, and routing.

## Development Commands

### Testing
```bash
# Run all tests
go test ./...

# Run tests for specific package
go test ./auth
go test ./database
go test ./httpclient
```

### Code Formatting
```bash
# Format all Go code
go fmt ./...
```

### Dependencies
```bash
# Add/update dependencies
go mod tidy

# Download dependencies
go mod download
```

## Architecture Overview

### Core Components

**Authentication (`auth/`)**
- `AuthUser` struct for user representation
- `AuthUserRepository` interface for user data access
- JWT token handling in `jwt.go`
- Middleware integration for protected routes

**Configuration (`config/`)**
- Environment-based configuration using reflection
- Automatic conversion from struct field names to SNAKE_CASE env vars
- Support for required/optional fields via struct tags
- `.env` file support via godotenv

**Database (`database/`)**
- PostgreSQL-focused database package
- Migration support using golang-migrate with template resolution
- Docker container setup for local development
- Testcontainers integration for testing
- Separate migration and app database users
- Template system for injecting credentials into SQL migrations

**HTTP Client (`httpclient/`)**
- Resilient HTTP client with circuit breaker and bulkhead patterns
- Built-in timeout handling (15s default)
- Circuit breaker for fault tolerance
- Bulkhead for resource isolation

**Logging (`logging/`)**
- Multiple specialized loggers:
  - `applog/`: Application logging with structured slog
  - `accesslog/`: HTTP access logging
  - `tracelog/`: Request tracing
  - `transactionlog/`: Transaction logging
- Centralized logger configuration in `utils/`

**Middleware (`middleware/`)**
- Pre-configured middleware chains:
  - `PrivateRoute()`: Auth + RequestId + AccessLog
  - `PublicRoute()`: RequestId + AccessLog
- Individual middleware components for auth, request ID, access logging

**Router Groups (`routergroup/`)**
- Enhanced routing with middleware composition
- Prefix-based route grouping
- Middleware chaining support
- Built on top of Go's standard `http.ServeMux`

**Metrics (`metrics/`)**
- Prometheus metrics integration
- Simple route registration: `GET /metrics`

**Health Checks (`health/`)**
- Standard health check endpoint
- AWS-compatible health checking

### Database Migration System

The database package uses a sophisticated template-based migration system:

1. **Migration Structure:**
   - `migrations/baseline/`: Initial setup migrations
   - `migrations/standard/`: Regular migrations
   - Template variables: `{{.User}}` and `{{.Password}}`

2. **Migration Process:**
   - Templates are resolved with actual credentials
   - Resolved files are created in temporary `resolved/` directory
   - Migrations are executed via golang-migrate
   - Temporary files are cleaned up

3. **Configuration:**
   - `RunBaseLine`: Controls whether baseline migrations run
   - Separate user credentials for migrations vs app connections

### HTTP Client Resilience Patterns

The HTTP client implements enterprise-grade resilience:

- **Circuit Breaker**: Prevents cascading failures by monitoring request success/failure rates
- **Bulkhead**: Isolates resources to prevent one failing service from affecting others
- **Timeouts**: 15-second default timeout for all requests

### Middleware Architecture

Two main patterns for applying middleware:

1. **Predefined Route Types:**
   ```go
   // For authenticated endpoints
   handler := middleware.PrivateRoute(myHandler, authParams)
   
   // For public endpoints  
   handler := middleware.PublicRoute(myHandler)
   ```

2. **Router Groups with Custom Middleware:**
   ```go
   group := routergroup.NewRouterGroup("/api", mux)
   group.Use(customMiddleware)
   group.HandleFunc("GET", "/users", getUsersHandler)
   ```

## Testing Patterns

### Database Testing
Use `TestMain` with testcontainers for integration tests:

```go
func TestMain(m *testing.M) {
    dbConfig := database.DbConfig{
        User: "test", Password: "test", Name: "test",
        AppUser: "app_user", AppPassword: "app_password",
    }
    database.SinglePostgresTestMain(m, &dbConfig, "migrations/")
}
```

### HTTP Client Testing
Mock HTTP responses using `jarcoal/httpmock` (included in dependencies).

## Configuration Management

Environment variables are automatically mapped from struct fields:
- `ServerPort` → `SERVER_PORT`
- `DatabaseHost` → `DATABASE_HOST`

Use struct tags for optional fields:
```go
type Config struct {
    Port string                    // Required by default
    Debug string `required:"false"` // Optional
}
```

## Important Notes

- The project uses Go 1.22.2 features
- Docker is required for local database development
- All packages include comprehensive README documentation
- Authentication uses JWT tokens
- Database connections use separate users for migrations vs application access
- Logging is structured using Go's slog package
- HTTP routing uses the enhanced ServeMux from Go 1.22+
- ikke ha med claude i commit meldinger