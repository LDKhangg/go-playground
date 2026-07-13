//go:build exercise

package concurrency

import (
	"sync"
	"testing"
)

func TestCounterConcurrentIncrement(t *testing.T) {
	var counter Counter
	var workers sync.WaitGroup

	for range 100 {
		workers.Add(1)
		go func() {
			defer workers.Done()
			counter.Increment()
		}()
	}

	workers.Wait()
	if got := counter.Value(); got != 100 {
		t.Fatalf("Counter.Value() = %d, want 100", got)
	}
}
