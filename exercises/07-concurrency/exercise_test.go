//go:build exercise

package concurrency

import (
	"context"
	"errors"
	"sync"
	"testing"
)

func TestCounterConcurrentIncrement(t *testing.T) {
	var counter Counter
	start := make(chan struct{})
	var workers sync.WaitGroup

	workers.Add(2)
	go func() {
		defer workers.Done()
		<-start
		for range 1000 {
			counter.Increment()
		}
	}()
	go func() {
		defer workers.Done()
		<-start
		for range 1000 {
			_ = counter.Value()
		}
	}()

	close(start)
	workers.Wait()
	if got := counter.Value(); got != 1000 {
		t.Fatalf("Counter.Value() = %d, want 1000", got)
	}
}

func TestSumReadsValuesFromChannel(t *testing.T) {
	values := make(chan int, 3)
	for _, value := range []int{2, 4, 6} {
		values <- value
	}
	close(values)

	got, err := Sum(context.Background(), values)
	if err != nil {
		t.Fatalf("Sum returned error: %v", err)
	}
	if got != 12 {
		t.Fatalf("Sum = %d, want 12", got)
	}
}

func TestSumStopsWhenContextIsCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	values := make(chan int)
	got, err := Sum(ctx, values)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Sum error = %v, want %v", err, context.Canceled)
	}
	if got != 0 {
		t.Fatalf("Sum after cancellation = %d, want 0", got)
	}
}
