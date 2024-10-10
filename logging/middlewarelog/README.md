# Middleware log package

This is a package for general logging inside middlewares.

## Usage

Here is an example of how to use this package.
Init logger in main to be used everywhere in the app.

```go
// In main function
appName := "todo"
middlwarelog.SetAppName(appName)

// Used elsewhere in the app
applog.Info("This is an info message", "requestId")
applog.Error("This is an error message", "requestId")
```
