# App log package

This is a package for general logging.

## Usage

Here is an example of how to use this package.
Init logger in main to be used everywhere in the app.

```go
// In main function
appName := "todo"
applog.AppLog = applog.NewAppLogger(appName)

// Used elsewhere in the app
applog.AppLog.Info("This is an info message")
```