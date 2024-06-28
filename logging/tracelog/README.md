# Trace log package

This is a package for logging inside of app. Typical used for error logging explicit.

## Usage

Here is an example of how to use this package.
Init logger in main to be used everywhere in the app.

```go
// In main function
appName := "todo"
translog.SetAppName(appName)

// Used elsewhere in the app
translog.Info("This is an info message")
```