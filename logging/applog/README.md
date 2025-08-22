# App Log

This package provides general application logging with structured output.

## Usage

Initialize the logger in main and use throughout the app:

```go
import "github.com/mariusfa/golf/logging/applog"

// In main function
appName := "todo"
applog.SetAppName(appName)

// Used elsewhere in the app
applog.Info("Application started successfully")
applog.Debug("Processing user input", map[string]interface{}{
    "userId": "123",
    "action": "login",
})
applog.Error("Failed to connect to database")
```

## Purpose

General purpose logging for:
- Application lifecycle events
- Debug information during development
- General information and error messages
- Structured logging with JSON output