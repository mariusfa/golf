# Auth

This package provides JWT token management and user authentication interfaces.

## Usage

### JWT Token Creation
```go
// Default 24-hour expiration
token, err := auth.CreateToken(userId, secret, nil)

// Custom expiration
expires := time.Now().Add(2 * time.Hour)
token, err := auth.CreateToken(userId, secret, &expires)
```

### JWT Token Parsing
```go
userId, err := auth.ParseToken(token, secret)
if err != nil {
    // Handle invalid token
}
```

### User Repository Interface
```go
type AuthUserRepository interface {
    FindAuthUserById(userId string) (AuthUser, error)
}
```

## Integration

This package integrates with the middleware package for route protection. See middleware documentation for usage examples.

## Token Details
- Uses HMAC-SHA256 signing
- Contains userId in custom claims
- Default expiration: 24 hours