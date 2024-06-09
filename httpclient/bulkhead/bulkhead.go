package bulkhead

// import semaphore
import (
	"errors"

	"golang.org/x/sync/semaphore"
)

type Options struct {
	// Maximum number of concurrent requests allowed
	MaxConcurrent int64
}

const (
	DefaultMaxConcurrent = 10
)

var (
	ErrBulkheadFull = errors.New("bulkhead is full")
)

type Bulkhead struct {
	sem *semaphore.Weighted
}

func NewBulkhead(opt Options) *Bulkhead {
	if opt.MaxConcurrent == 0 {
		opt.MaxConcurrent = DefaultMaxConcurrent
	}

	return &Bulkhead{
		sem: semaphore.NewWeighted(opt.MaxConcurrent),
	}
}

func (b *Bulkhead) Execute(action func() error) error {
	if !b.sem.TryAcquire(1) {
		return ErrBulkheadFull
	}

	defer b.sem.Release(1)

	return action()
}
