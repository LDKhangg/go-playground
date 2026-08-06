# 07 - Concurrency

## Goal

Make shared mutable state safe, receive values through a channel, and stop work through context cancellation.

## Concepts

- A goroutine runs a function concurrently with its caller.
- Channels communicate typed values between goroutines.
- A mutex protects a critical section around shared state.
- A wait group waits for known goroutines to finish.
- `context.Context` carries cancellation through an operation.
- The race detector observes unsynchronized conflicting memory access.

## Syntax Primer

```go
var workers sync.WaitGroup
workers.Add(1)
go func() {
	defer workers.Done()
	// Concurrent work.
}()
workers.Wait()
```

```go
select {
case <-ctx.Done():
	return 0, ctx.Err()
case value, ok := <-values:
	if !ok {
		return total, nil
	}
	total += value
}
```

Receiving from a closed channel returns its element type's zero value, so the second result `ok` is necessary when zero is a valid value. A `select` waits until one communication can proceed.

## Mental Model

Goroutines share the process memory, which is useful but unsafe when two goroutines access the same mutable value without coordination. A mutex is a key for the small section that reads or changes that value. A channel is a handoff lane: senders put values in, receivers take values out. Context cancellation is a separate stop signal, not a channel close you own.

## Annotated Examples

```go
type SafeCount struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCount) AddOne() {
	c.mu.Lock()
	defer c.mu.Unlock() // Always release the lock, including early returns.
	c.value++
}
```

```go
func waitForValue(ctx context.Context, values <-chan int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case _, ok := <-values:
		if !ok {
			return nil
		}
		return nil
	}
}
```

## Common Diagnostics

- `fatal error: all goroutines are asleep - deadlock!`: a goroutine is waiting for a send, receive, or lock that cannot happen.
- `concurrent map writes` or race-detector output: shared state needs synchronization or a single owning goroutine.
- `sync: negative WaitGroup counter`: `Done` ran more times than `Add`.
- `go vet` warns about copying a lock value: avoid copying structs that contain a mutex after first use.
- A starter assertion fails because `Counter` or `Sum` is incomplete: that is exercise behavior, not a syntax error.

## Exercise

Make `Counter` safe while one goroutine increments it and another reads it. Both `Increment` and `Value` must synchronize access to `value`. Implement `Sum` so it receives and totals channel values until the channel closes, or returns the context error when cancellation wins.

## Acceptance Criteria

- The final counter value is exactly 1000.
- `go test -race` reports no data race.
- The counter uses `sync.Mutex` or `sync.RWMutex`, not sleeps or timing assumptions.
- `Sum` returns `12` for a channel containing `2`, `4`, and `6`.
- `Sum` returns `context.Canceled` when called with an already canceled context.

## Hints

- Put the lock inside `Counter` beside the state it protects.
- Lock reads as well as writes; a read can race with an increment.
- In `Sum`, use `select` to wait for either `ctx.Done()` or a channel receive.
- Check the channel's `ok` result before adding its value.

## Verify

Run:

```bash
gofmt -w exercises/07-concurrency
go test -race -tags exercise ./exercises/07-concurrency/...
```

The starter intentionally fails until both synchronization and channel behavior are implemented. Use the race detector after the functional assertions pass; it detects a class of timing-dependent bugs that ordinary tests may miss.

## Reflection Prompts

- Why must `Value` lock even though it only reads?
- How does selecting on `ctx.Done()` keep channel work cancelable?
- When is a channel a better fit than protecting shared state with a mutex?
