# Circuit breaker

This is a package for circuit breaking.

## Usage
Here is an example of how to use this package.
```go
action := func() error {
       // function code to be executed
}

cb := NewCircuitBreaker(Options{}) // with default options
err := cb.Execute(action)

if err != nil {
	t.Error("error should be nil")
}
```

These are the options you can set and their default values:
```go
type Options struct {
	// Number of failed requests required to set the state to open
	FailureThreshold int
	// Number of successful requests required to set the state to closed
	RecoveryThreshold int
	// Number of concurrent test requests allowed in half open state
	TestRequestsAllowed int
	// Timeout for open state after which the state is set to half open
	ResetTimeout time.Duration
	// Interval after which the failure count is reset
	FailureResetInterval time.Duration
}

const (
	DefaultFailureThreshold     = 5
	DefaultRecoveryThreshold    = 3
	DefaultTestRequestsAllowed  = 3
	DefaultResetTimeout         = 5 * time.Second
	DefaultFailureResetInterval = 60 * time.Second
)
```