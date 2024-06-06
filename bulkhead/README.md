# Bulkhead

This is a package for bulkhead

## Usage
Here is an example of how to use this package.

```go
action := func() error {
       // function code to be executed
}

b := NewBulkhead(Options{}) // with default options
err := b.Execute(action)

if err != nil {
	t.Error("error should be nil")
}
```

These are the options you can set and their default values:
```go
type Options struct {
	// Maximum number of concurrent requests allowed
	MaxConcurrent int64
}

const (
	DefaultMaxConcurrent = 10
)
```
