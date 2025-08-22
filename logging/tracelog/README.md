# Trace Log

This package provides request tracing and error logging for applications.

## Usage

Initialize the logger in main and use throughout the app:

```go
// In main function
appName := "todo"
tracelog.SetAppName(appName)

// Used elsewhere in the app for tracing requests
tracelog.Info("Processing user request")
tracelog.Error("Database connection failed")
```

## Purpose

Typically used for:
- Request tracing and debugging
- Error logging with context
- Application flow monitoring