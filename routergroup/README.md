# Router Groups

A `RouterGroup` can be used to group a collection of endpoints under the same prefix. 
It can also be used to define a set of `Middleware` that should apply to all endpoints
registered on this group.

## Usage

Example with public and private route groups:
```go
router := http.NewServeMux()

// Public routes (no authentication needed)
publicGroup := routergroup.NewRouterGroup("/api/public", router)
publicGroup.Use(middleware.RequestIdMiddleware)
publicGroup.Use(middleware.AccessLogMiddleware)

// Private routes (authentication required)
authParams := middleware.NewAuthParams(jwtSecret, userRepo)
privateGroup := routergroup.NewRouterGroup("/api/private", router)
privateGroup.Use(middleware.RequestIdMiddleware)
privateGroup.Use(middleware.AccessLogMiddleware)
privateGroup.Use(func(next http.Handler) http.Handler {
    return middleware.Auth(next, authParams)
})

// Register routes
publicGroup.HandleFunc("GET", "/health", healthHandler)
privateGroup.HandleFunc("GET", "/users", getUsersHandler)
```

Custom middleware:
```go
logTime := func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("Request took: %v", time.Since(start))
    })
}

apiGroup := routergroup.NewRouterGroup("/api", router)
apiGroup.Use(logTime)
apiGroup.HandleFunc("GET", "/items", getItemsHandler)
```
