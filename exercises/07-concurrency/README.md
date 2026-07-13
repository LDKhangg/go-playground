# 07 - Concurrency

## Goal

Make shared mutable state safe when many goroutines use it.

## Concepts

Goroutines, wait groups, mutexes, critical sections, and the race detector.

## Exercise

Make `Counter` safe and correct when 100 goroutines increment it concurrently. Both `Increment` and `Value` must synchronize access to `value`.

## Acceptance Criteria

- The final value is exactly 100.
- `go test -race` reports no data race.
- The counter uses `sync.Mutex` or `sync.RWMutex`, not sleeps or timing assumptions.

## Hints

Put the lock inside `Counter` beside the state it protects. Keep each critical section small and use `defer` when it makes unlocking easier to verify.

## Commands

`go test -race -tags exercise ./exercises/07-concurrency/...`

## Reflection Prompts

Why does the wait group not protect the counter itself? What behavior does the race detector find that a normal assertion may miss?
