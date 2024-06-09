package bulkhead

import (
	"sync"
	"testing"
	"time"
)

// happy path test
func TestBulkhead(t *testing.T) {
	action := func() error {
		return nil
	}
	b := NewBulkhead(Options{MaxConcurrent: 1})

	err := b.Execute(action)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestConcurrentActions(t *testing.T) {
	action := func() error {
		time.Sleep(time.Millisecond)
		return nil
	}

	b := NewBulkhead(Options{MaxConcurrent: 2})

	var wg sync.WaitGroup
	wg.Add(2)

	listOfErrors := make([]error, 2)

	go func() {
		defer wg.Done()
		err := b.Execute(action)
		listOfErrors[0] = err
	}()

	go func() {
		defer wg.Done()
		err := b.Execute(action)
		listOfErrors[1] = err
	}()

	wg.Wait()

	// no errors
	for _, err := range listOfErrors {
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	}
}

func TestBulkheadFull(t *testing.T) {
	action := func() error {
		time.Sleep(time.Millisecond)
		return nil
	}

	b := NewBulkhead(Options{MaxConcurrent: 1})

	var wg sync.WaitGroup
	wg.Add(2)

	listOfErrors := make([]error, 2)

	go func() {
		defer wg.Done()
		err := b.Execute(action)
		listOfErrors[0] = err
	}()

	go func() {
		defer wg.Done()
		err := b.Execute(action)
		listOfErrors[1] = err
	}()

	wg.Wait()

	errorCount := 0
	for _, err := range listOfErrors {
		if err != nil && err == ErrBulkheadFull {
			errorCount++
		}
	}

	if errorCount != 1 {
		t.Errorf("Expected 1, got %v", errorCount)
	}
}
