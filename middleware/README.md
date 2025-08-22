# Middleware

This package provides pre-configured middleware chains for common web application patterns.

## Usage

Public routes (RequestId + AccessLog):
```go
handler := middleware.PublicRoute(myHandler)
router.Handle("GET /api/health", handler)
```

Private routes (Auth + RequestId + AccessLog):
```go
authParams := middleware.NewAuthParams(jwtSecret, userRepo)
handler := middleware.PrivateRoute(myHandler, authParams)
router.Handle("GET /api/users", handler)
```

Complete example:
```go
router := http.NewServeMux()
authParams := middleware.NewAuthParams("jwt-secret", userRepository)

// Public endpoints
router.Handle("GET /health", middleware.PublicRoute(healthHandler))
router.Handle("POST /login", middleware.PublicRoute(loginHandler))

// Private endpoints  
router.Handle("GET /users", middleware.PrivateRoute(getUsersHandler, authParams))
router.Handle("POST /users", middleware.PrivateRoute(createUserHandler, authParams))
```

## Authentication

The Auth middleware requires JWT tokens in the Authorization header:
```
Authorization: Bearer <jwt-token>
```