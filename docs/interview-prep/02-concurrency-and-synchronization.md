# 02 - Concurrency And Synchronization

Concurrency is the single most-tested Go topic in interviews. Practice with
`exercises/07-concurrency` and run everything with the race detector.

## 2.1 Goroutines And The Scheduler

**Q: What is a goroutine?**

A lightweight thread managed by the Go runtime: it starts with a small (a few
KB) growable stack and is multiplexed onto OS threads. Thousands of goroutines
are normal; thousands of OS threads are not.

**Q: How does Go schedule goroutines?**

The GMP model: **G**oroutines are scheduled onto **P** (processors, runqueues)
which run on **M**achines (OS threads). The runtime uses M:N scheduling with
work stealing — idle Ps steal work from busier ones. `GOMAXPROCS` caps the
number of Ps (defaults to the number of logical CPUs) and therefore the number
of goroutines that run in parallel.

**Q: Why does a goroutine not block the world when it blocks?**

Blocking operations (channel sends/receives, `time.Sleep`, I/O) park the
goroutine and let the runtime run another goroutine on the same thread — that
is why concurrent servers with blocking I/O still scale.

## 2.2 Channels

**Q: Buffered vs unbuffered channels?**

An unbuffered channel synchronizes: the sender blocks until a receiver is
ready, so the send and receive happen "at the same time." A buffered channel
holds up to `cap` values: sends block only when the buffer is full, receives
block only when it is empty.

```go
ch := make(chan int)     // unbuffered — synchronized handoff
buf := make(chan int, 4) // buffered — up to 4 queued values
```

**Q: What happens when a channel is closed?**

- Receives keep working until the buffer drains, then return the zero value with `ok == false`.
- `range` over a channel loops until it is closed.
- Sending to a closed channel panics.
- Closing a closed channel panics.
- The sender closes; the receiver never should.

```go
for v := range ch { // stops when ch is closed and drained
	_ = v
}
v, ok := <-ch // ok == false means closed and empty
```

**Q: How do you stop a goroutine cleanly?**

Pass a `context.Context` or use a dedicated stop channel and `select` on it.
Never kill goroutines; coordinate with them.

```go
for {
	select {
	case <-ctx.Done():
		return
	case v, ok := <-jobs:
		if !ok {
			return
		}
		process(v)
	}
}
```

Practice: `exercises/07-concurrency` (`Sum` receives from a channel until it
closes and selects on `ctx.Done()`).

## 2.3 select

**Q: What does `select` do?**

Waits until one of its cases can proceed. If several can, one is chosen
randomly — the only randomness in Go's channel model, and a common source of
"works sometimes" bugs. A `default` case makes the select non-blocking.

```go
select {
case <-ctx.Done():
	return 0, ctx.Err()
case value, ok := <-values:
	if !ok {
		return total, nil
	}
	total += value
case <-time.After(5 * time.Second):
	return 0, errors.New("timeout")
}
```

## 2.4 sync

**Q: Mutex vs channel — which do you use?**

"Share memory by communicating" is the slogan, but the practical rule: use
channels to pass ownership of data between goroutines; use a mutex to protect
shared state that many goroutines touch (like a counter). Both are idiomatic;
picking the clearer one for the situation is the skill being tested.

**Q: How do you use a mutex correctly?**

Lock, then `defer Unlock` immediately so early returns cannot skip the unlock.
Protect reads too — a read racing with a write is still a data race. Use
`sync.RWMutex` when reads dominate.

```go
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
```

Practice: `exercises/07-concurrency` (`Counter` must be safe for concurrent
increment and read).

**Q: Common `sync` mistakes?**

- Copying a mutex after use (see it in `go vet` warnings) — mutexes carry state.
- `WaitGroup.Add` called inside the goroutine instead of before it starts — the
  wait may start before `Add`.
- Calling `Done` more times than `Add` — panics with a negative counter.

```go
var wg sync.WaitGroup
wg.Add(2)               // Add before starting work
go func() { defer wg.Done(); work() }()
go func() { defer wg.Done(); work() }()
wg.Wait()               // blocks until both Done calls
```

**Q: What about `sync.Once`, `sync.Map`, `sync.Pool`?**

- `sync.Once` runs a function exactly once (lazy singleton init).
- `sync.Map` is a specialized concurrent map; a mutex-guarded map is usually clearer.
- `sync.Pool` caches reusable objects to reduce allocations; objects may be discarded anytime — never rely on it for correctness.

## 2.5 context

**Q: What is `context.Context` for?**

Carrying cancellation, deadlines, and request-scoped values through an
operation tree. Rules:

- `context.Background()` is the root; `context.TODO()` says "should be wired later."
- Pass context as the first parameter; never store it in a struct.
- When a parent is canceled, all derived contexts cancel too (`WithCancel`, `WithTimeout`, `WithDeadline`, `WithValue`).
- Check `ctx.Err()` after calling a cancelable function.

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel() // release resources even on the happy path

result, err := doWork(ctx)
if err != nil {
	return err
}
```

**Q: How do you make work cancelable?**

Choose the operation that can observe cancellation — usually a `select` on
`ctx.Done()` around channel receives, or passing `ctx` to `net/http` and
database calls. `context.Canceled` is returned as the error when cancellation
wins.

Practice: `exercises/07-concurrency` (`Sum` returns `context.Canceled` for an
already-canceled context).

## 2.6 The Race Detector

**Q: What is a data race, and how do you find one?**

A data race is two goroutines accessing the same memory without
synchronization, at least one write. Symptoms are often intermittent.
`go test -race` instruments the binary and reports races — make it part of your
normal workflow:

```bash
go test -race ./...
```

The race detector only reports races it observes; it is evidence of absence of
races in the tested runs, not a proof of safety.

**Q: Common causes of races?**

- Two goroutines mutating the same map.
- Reading a counter while another goroutine increments it.
- Closing a channel while another goroutine sends.
- Slicing shared state in one goroutine and writing it in another.

## 2.7 Deadlock, Livelock, Starvation — Quick Definitions

- **Deadlock**: everyone is waiting for something that never arrives — classic
  symptom is `fatal error: all goroutines are asleep - deadlock!`.
- **Livelock**: goroutines keep acting but make no progress (e.g., two
  goroutines forever passing a resource back and forth).
- **Starvation**: a goroutine never gets scheduled or never gets the lock it
  needs; the program may look alive but one participant is stuck.

## Practice Check

1. Explain the difference between a buffered and an unbuffered channel with a one-line example each.
2. Who should close a channel, and what happens if you send to a closed one?
3. Why must `Value()` lock even though it only reads?
4. How do you stop a goroutine without killing it?
5. What does `go test -race` report, and what does it not prove?
