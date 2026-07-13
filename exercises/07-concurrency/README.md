# 07 - Concurrency

## Goal

Make shared mutable state safe, receive values through a channel, and stop work through context cancellation.

## Concepts

Goroutines, channels, wait groups, mutexes, critical sections, `context.Context`, cancellation, and the race detector.

## Exercise

Make `Counter` safe while one goroutine increments it and another reads it. Both `Increment` and `Value` must synchronize access to `value`. Implement `Sum` so it receives and totals channel values until the channel closes, or returns the context error when cancellation wins.

## Acceptance Criteria

- The final counter value is exactly 1000.
- `go test -race` reports no data race.
- The counter uses `sync.Mutex` or `sync.RWMutex`, not sleeps or timing assumptions.
- `Sum` returns `12` for a channel containing `2`, `4`, and `6`.
- `Sum` returns `context.Canceled` when called with an already canceled context.

## Hints

Put the lock inside `Counter` beside the state it protects, and lock reads as well as writes. In `Sum`, use `select` to wait for either `ctx.Done()` or a value and its open/closed status from the channel.

## Commands

`go test -race -tags exercise ./exercises/07-concurrency/...`

## Reflection Prompts

Why must `Value` lock even though it only reads? How does selecting on `ctx.Done()` keep channel work cancelable?
