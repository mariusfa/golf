package middleware

type contextKey string

const (
	UserKey = contextKey("user")
)
