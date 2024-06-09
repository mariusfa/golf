package circuitbreaker

import (
	"errors"
	"sync"
	"testing"
	"time"
)

// Happy path test
func TestPerformAction(t *testing.T) {
	actionPerformed := false
	action := func() error {
		actionPerformed = true
		return nil
	}

	cb := NewCircuitBreaker(Options{})
	err := cb.Execute(action)

	if err != nil {
		t.Error("error should be nil")
	}

	if !actionPerformed {
		t.Error("action should have been performed")
	}

	if cb.state != Closed {
		t.Error("state should be closed")
	}

	if cb.failureCount != 0 {
		t.Error("failure count should be 0")
	}
}

// Test that the circuit breaker opens after the failure threshold is reached
func TestOpenCircuit(t *testing.T) {
	action := func() error {
		return errors.New("error")
	}

	cb := NewCircuitBreaker(Options{
		FailureThreshold: 2,
	})

	cb.Execute(action)
	cb.Execute(action)
	cb.Execute(action)

	if cb.state != Open {
		t.Error("state should be open")
	}
}

// Test that failure count is reset after the failure reset interval
func TestFailureCountReset(t *testing.T) {
	action := func() error {
		return errors.New("error")
	}

	cb := NewCircuitBreaker(Options{
		FailureThreshold:     3,
		FailureResetInterval: 1 * time.Millisecond,
	})

	cb.Execute(action)
	cb.Execute(action)
	time.Sleep(2 * time.Millisecond) // waits for reset of failure count
	cb.Execute(action)

	if cb.state != Closed {
		t.Error("state should be closed")
	}

	if cb.failureCount != 1 {
		t.Error("failure count should be 1")
	}
}

// Test going from open to half open state
func TestOpenStateToHalfOpen(t *testing.T) {
	action := func() error {
		return errors.New("error")
	}

	cb := NewCircuitBreaker(Options{
		FailureThreshold: 1,
		ResetTimeout:     1 * time.Millisecond,
	})

	cb.Execute(action)
	cb.Execute(action)
	time.Sleep(10 * time.Millisecond) // waits for half open state

	if cb.state != HalfOpen {
		t.Error("state should be half open, state is: ", cb.state)
	}
}

// Test going from half open to open state
func TestHalfOpenToOpen(t *testing.T) {
	action := func() error {
		return errors.New("error")
	}

	cb := NewCircuitBreaker(Options{
		FailureThreshold: 1,
		ResetTimeout:     1 * time.Millisecond,
	})

	cb.Execute(action)
	cb.Execute(action)
	time.Sleep(10 * time.Millisecond) // waits for half open state
	cb.Execute(action)

	if cb.state != Open {
		t.Error("state should be open, state is: ", cb.state)
	}
}

// Test going from half open to closed state
func TestHalfOpenToClosed(t *testing.T) {
	errorAction := func() error {
		return errors.New("error")
	}

	successAction := func() error {
		return nil
	}

	cb := NewCircuitBreaker(Options{
		FailureThreshold:  1,
		RecoveryThreshold: 2,
		ResetTimeout:      1 * time.Millisecond,
	})

	cb.Execute(errorAction)
	cb.Execute(errorAction)
	time.Sleep(10 * time.Millisecond) // waits for half open state
	err := cb.Execute(successAction)  // half open state
	if err != nil {
		t.Error("error should be nil, error is: ", err)
	}

	if cb.state != HalfOpen {
		t.Error("state should be half open, state is: ", cb.state)
	}

	cb.Execute(successAction) // reach recovery threshold

	if cb.state != Closed {
		t.Error("state should be closed, state is: ", cb.state)
	}
}

// Test concurrent execution in half open state
func TestConcurrentExecutionHalfOpen(t *testing.T) {
	errorAction := func() error {
		return errors.New("error")
	}

	longAction := func() error {
		time.Sleep(10 * time.Millisecond)
		return nil
	}

	cb := NewCircuitBreaker(Options{
		FailureThreshold:    1,
		TestRequestsAllowed: 2,
		ResetTimeout:        1 * time.Millisecond,
	})

	cb.Execute(errorAction)
	cb.Execute(errorAction)
	time.Sleep(10 * time.Millisecond) // waits for half open state

	// execute 3 long actions concurrently. 2 should succeed and 1 should fail
	var wg sync.WaitGroup
	wg.Add(3)

	listOfErrors := make([]error, 3)

	go func() {
		defer wg.Done()
		err := cb.Execute(longAction)
		listOfErrors[0] = err
	}()

	go func() {
		defer wg.Done()
		err := cb.Execute(longAction)
		listOfErrors[1] = err
	}()

	go func() {
		defer wg.Done()
		err := cb.Execute(longAction)
		listOfErrors[2] = err
	}()
	wg.Wait()

	errorInHalfOpenAllowed := false
	successCount := 0
	for _, err := range listOfErrors {
		if err == ErrorCircuitHalfOpenAllowedRequests {
			errorInHalfOpenAllowed = true
		} else if err == nil {
			successCount++
		}
	}

	if !errorInHalfOpenAllowed {
		t.Error("error should be ErrorCircuitHalfOpenAllowedRequests")
	}

	if successCount != 2 {
		t.Error("success count should be 2")
	}
}
