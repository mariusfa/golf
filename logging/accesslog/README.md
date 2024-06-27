# Access log package

This is a package for access logging in a web app.

## Usage

Here is an example of how to use this package.
Init logger in main to be used everywhere in the app.

```go
// In main function
appName := "todo"
accesslog.SetAppName(appName)

// Used elsewhere in the app
duration := 100
status := 200
accesslog.Info(duration, status, "/path", "GET")
```
