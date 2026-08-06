# Practice Exercises With Examples

Hands-on companion to the interview-prep modules. Each exercise has a
**Problem**, a **Hint**, a complete **Example Solution** you can run, and a
short **Why It Works**. Try the problem yourself first, then compare with the
example.

Prerequisites: Go 1.22+ (the repo targets Go 1.26). Save any example to a file
like `main.go` and run it with `go run main.go`; tests with `go test`.

## Contents

1. [Core Fundamentals](#1-core-fundamentals)
2. [Concurrency And Synchronization](#2-concurrency-and-synchronization)
3. [Backend And HTTP](#3-backend-and-http)
4. [Architecture And Testing](#4-architecture-and-testing)
5. [Traps](#5-traps)

---

## 1. Core Fundamentals

### Exercise 1.1 — Zero values and pointer mutation

**Problem:** Write a `DefaultProfile` function that returns a `Profile` struct
(Name, Active, Attempts) holding only zero values, and an `Increment` function
that adds 1 through an `*int` pointer. Then print both results.

**Hint:** A struct literal with no fields (`Profile{}`) is already all zeros.
Check `nil` before dereferencing.

**Example solution** (`main.go`):

```go
package main

import "fmt"

type Profile struct {
	Name     string
	Active   bool
	Attempts int
}

func DefaultProfile() Profile { return Profile{} }

func Increment(v *int) error {
	if v == nil {
		return fmt.Errorf("nil pointer")
	}
	*v++
	return nil
}

func main() {
	p := DefaultProfile()
	fmt.Printf("%q %v %d\n", p.Name, p.Active, p.Attempts) // "" false 0

	n := 4
	_ = Increment(&n)
	fmt.Println(n) // 5
}
```

**Why it works:** Every declared variable is a valid box with a zero value, so
`Profile{}` needs no assignments. `&n` hands `Increment` the address of `n`;
`*v++` writes through it, reaching the caller's variable.

**Try it:** `go run main.go`. Then make `Increment(nil)` return the error
instead of crashing.

---

### Exercise 1.2 — Slice aliasing and `append`

**Problem:** Predict the output of this program, then run it:

```go
package main

import "fmt"

func main() {
	a := make([]int, 3, 6)
	a[0], a[1], a[2] = 1, 2, 3

	b := append(a, 4)
	b[0] = 99

	fmt.Println(a, b)

	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
	}
}
```

**Hint:** `append` writes into the backing array when capacity allows — that is
why `b[0] = 99` also changed `a`.

**Example output:**

```text
[99 2 3] [99 2 3 4]
len=1 cap=1
len=2 cap=2
len=3 cap=4
len=4 cap=4
len=5 cap=8
...
```

**Why it works:** `b := append(a, 4)` fits in `a`'s spare capacity, so `b`
shares the same backing array. When capacity runs out, `append` allocates a new
array (capacity doubling) and the original slice is untouched from then on —
which is why you must always use the result: `s = append(s, v)`.

**Follow-up exercise:** Implement `InsertAt(values []int, index, value int)
([]int, error)` so `InsertAt([]int{1, 3}, 1, 2)` returns `[1 2 3]`, and reject
out-of-range indices with an error. Example solution:

```go
func InsertAt(values []int, index, value int) ([]int, error) {
	if index < 0 || index > len(values) {
		return nil, fmt.Errorf("index %d out of range", index)
	}
	out := make([]int, 0, len(values)+1)
	out = append(out, values[:index]...)
	out = append(out, value)
	out = append(out, values[index:]...)
	return out, nil
}
```

---

### Exercise 1.3 — Maps, word count, and comma-ok

**Problem:** Write `CountWords(text string) map[string]int` and
`LookupScore(scores map[string]int, name string) (int, bool)`, then print the
results.

**Hint:** `strings.Fields` splits on whitespace; `counts[word]++` on a map
created with `make` never panics.

**Example solution:**

```go
package main

import (
	"fmt"
	"strings"
)

func CountWords(text string) map[string]int {
	counts := make(map[string]int)
	for _, word := range strings.Fields(text) {
		counts[word]++
	}
	return counts
}

func LookupScore(scores map[string]int, name string) (int, bool) {
	score, ok := scores[name]
	return score, ok
}

func main() {
	fmt.Println(CountWords("go go tests")) // map[go:2 tests:1]

	scores := map[string]int{"mina": 10}
	fmt.Println(LookupScore(scores, "mina")) // 10 true
	fmt.Println(LookupScore(scores, "kai"))  // 0 false
}
```

**Why it works:** a missing key reads as the zero value (`0`), so `ok` is the
only reliable presence signal. Writing to a nil map panics — `make` avoids
that.

**Try it:** change `CountWords` to write into `var counts map[string]int`
(nil) and observe the panic.

---

## 2. Concurrency And Synchronization

### Exercise 2.1 — Race-free counter

**Problem:** Ten goroutines increment a shared counter 1000 times each. Make
the final value exactly 10000 with no data race.

**Hint:** guard the counter with `sync.Mutex`; lock reads too.

**Example solution:**

```go
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var c Counter
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				c.Inc()
			}
		}()
	}
	wg.Wait()
	fmt.Println(c.Value()) // 10000
}
```

**Why it works:** `Add` happens before the goroutines start; `defer
c.mu.Unlock()` guarantees the lock is released even if `Inc` grows a `return`.
`Value` locks too, because a read racing a write is still a race.

**Try it:** `go run -race main.go` — no race output. Then remove the locks,
run with `-race`, and read the report.

---

### Exercise 2.2 — Stop a goroutine with context

**Problem:** Start a worker that reads jobs from a channel and stops when a
context is canceled. Demonstrate the timeout path with a 100ms deadline.

**Hint:** `select` on `ctx.Done()` and on the channel receive.

**Example solution:**

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, jobs <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopping:", ctx.Err())
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println("job", job)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	jobs := make(chan int) // nobody sends: cancellation wins
	worker(ctx, jobs)
}
```

**Example output:**

```text
stopping: context deadline exceeded
```

**Why it works:** the worker is not killed — it cooperates by watching
`ctx.Done()`. Cancellation is a separate stop signal from channel close; `ok`
handles the closed-channel case, `ctx.Done()` the cancellation case.

**Follow-up exercise:** add a `producer` goroutine that sends 5 jobs and
closes the channel, then explain which path ends the worker first.

---

### Exercise 2.3 — WaitGroup + channel fan-in

**Problem:** Square each number in `[1,2,3,4,5]` in a goroutine and print the
sum (55) without a mutex.

**Hint:** one goroutine per number, results into a buffered channel, sum after
`wg.Wait()` and `close(out)`.

**Example solution:**

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	out := make(chan int, len(nums))
	var wg sync.WaitGroup

	for _, n := range nums {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			out <- v * v
		}(n)
	}

	wg.Wait()
	close(out)

	total := 0
	for v := range out {
		total += v
	}
	fmt.Println(total) // 55
}
```

**Why it works:** each goroutine owns one result, so the channel is the
handoff and no shared mutable state exists — no mutex needed. The buffer size
equals the worker count, so sends never block. `close(out)` lets `range`
terminate after all workers finished.

---

## 3. Backend And HTTP

### Exercise 3.1 — Minimal task API

**Problem:** Build a standard-library server with `GET /health` and
`POST /tasks` that validates JSON strictly (unknown fields rejected, blank
titles rejected) and returns the created task.

**Hint:** `http.ServeMux` with method patterns (`"POST /tasks"`),
`DisallowUnknownFields`, and `json.NewEncoder(w).Encode(task)`.

**Example solution:**

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var (
	tasks  = make(map[string]Task)
	nextID = 1
)

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}

func handleCreateTask(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
	}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&input); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	title := strings.TrimSpace(input.Title)
	if title == "" {
		http.Error(w, "title must not be empty", http.StatusBadRequest)
		return
	}

	id := fmt.Sprintf("task-%d", nextID)
	nextID++
	task := Task{ID: id, Title: title}
	tasks[id] = task

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(task)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handleHealth)
	mux.HandleFunc("POST /tasks", handleCreateTask)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Println("listening on :8080")
	log.Fatal(server.ListenAndServe())
}
```

**Try it:**

```bash
go run main.go
curl http://localhost:8080/health
curl -X POST http://localhost:8080/tasks \
  -H 'Content-Type: application/json' \
  -d '{"title":"learn Go"}'
curl -X POST http://localhost:8080/tasks \
  -H 'Content-Type: application/json' \
  -d '{"title":"  "}'        # -> 400
curl -X POST http://localhost:8080/tasks \
  -H 'Content-Type: application/json' \
  -d '{"title":"x","extra":1}' # -> 400, unknown field
```

**Why it works:** the mux routes by method and path; the decoder rejects
malformed and unknown fields; validation happens after decoding because a
well-formed payload can still hold a blank title. Note: `tasks` is a package
map shared by handlers — a real service would guard it with a mutex (see
Exercise 2.1) or keep it behind a service type.

**Follow-up exercise:** wrap the mux with a logging middleware (print method,
path, duration) and add graceful shutdown with `signal.Notify` +
`server.Shutdown(ctx)`.

---

## 4. Architecture And Testing

### Exercise 4.1 — Table-driven test

**Problem:** Write `Classify(n int) string` (negative/zero/positive) and a
table-driven test with three named subtests.

**Hint:** slice of anonymous structs `{name, input, want}` + `t.Run`.

**Example solution** (`main.go` and `main_test.go`):

```go
package main

func Classify(n int) string {
	switch {
	case n < 0:
		return "negative"
	case n > 0:
		return "positive"
	default:
		return "zero"
	}
}
```

```go
package main

import "testing"

func TestClassify(t *testing.T) {
	cases := []struct {
		name  string
		input int
		want  string
	}{
		{name: "negative", input: -1, want: "negative"},
		{name: "zero", input: 0, want: "zero"},
		{name: "positive", input: 1, want: "positive"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Classify(tc.input); got != tc.want {
				t.Fatalf("Classify(%d) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
```

**Why it works:** each row is a mini contract; `t.Run` names the failure so a
red run says exactly which input broke. `t.Fatalf` prints input, actual, and
expected — no guessing.

**Try it:** `go test -v .` and `go test -cover .`. Add a case with input 0 and
watch coverage hit the `default` branch.

---

### Exercise 4.2 — Service with dependency injection and a fake

**Problem:** A `TaskService.Create(title)` must trim the title, reject blanks,
assign an ID from an injected `IDGenerator`, and save through an injected
`Repository`. Write the service and a test using fakes (no mock library).

**Hint:** constructor injection; the test defines tiny fake structs.

**Example solution** (`service.go` + `service_test.go`):

```go
package main

import (
	"errors"
	"strings"
)

var ErrEmptyTitle = errors.New("title must not be empty")

type Task struct {
	ID    string
	Title string
}

type Repository interface {
	Save(task Task) error
}

type IDGenerator interface {
	NextID() string
}

type TaskService struct {
	repo Repository
	ids  IDGenerator
}

func NewTaskService(repo Repository, ids IDGenerator) *TaskService {
	return &TaskService{repo: repo, ids: ids}
}

func (s *TaskService) Create(title string) (Task, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return Task{}, ErrEmptyTitle
	}
	task := Task{ID: s.ids.NextID(), Title: title}
	if err := s.repo.Save(task); err != nil {
		return Task{}, err
	}
	return task, nil
}
```

```go
package main

import (
	"errors"
	"testing"
)

type fakeRepo struct {
	saved []Task
}

func (f *fakeRepo) Save(task Task) error {
	f.saved = append(f.saved, task)
	return nil
}

type fakeIDs struct{ next string }

func (f fakeIDs) NextID() string { return f.next }

func TestCreate(t *testing.T) {
	repo := &fakeRepo{}
	service := NewTaskService(repo, fakeIDs{next: "task-1"})

	task, err := service.Create("  ship API  ")
	if err != nil {
		t.Fatalf("Create returned error: %v", err)
	}
	if task.Title != "ship API" {
		t.Fatalf("title = %q, want %q", task.Title, "ship API")
	}
	if task.ID != "task-1" {
		t.Fatalf("id = %q, want %q", task.ID, "task-1")
	}
	if len(repo.saved) != 1 {
		t.Fatalf("saved = %d, want 1", len(repo.saved))
	}

	if _, err := service.Create("   "); !errors.Is(err, ErrEmptyTitle) {
		t.Fatalf("expected ErrEmptyTitle, got %v", err)
	}
}
```

**Why it works:** the service never calls `time.Now()` or constructs a repo —
it uses whatever it was given. The fakes are hand-written, dependency-free, and
assert exactly what the service is responsible for.

**Try it:** `go test -v .`

---

## 5. Traps

### Exercise 5.1 — Loop variable capture

**Problem:** Predict the output before running. Then run on Go 1.22+.

```go
package main

import "fmt"

func main() {
	var fns []func()
	for i := 0; i < 3; i++ {
		i := i // per-iteration copy
		fns = append(fns, func() { fmt.Println(i) })
	}
	for _, f := range fns {
		f()
	}
}
```

**Example output:** `0`, `1`, `2` (in any order — one per line).

**Why it works:** before Go 1.22, the loop variable was shared and all closures
saw the final value (`3`); the `i := i` idiom created a fresh copy per
iteration. Since Go 1.22, `for` loop variables are per-iteration, so the idiom
is optional but still common in older codebases. Removing `i := i` still prints
`0 1 2` on 1.22+.

---

### Exercise 5.2 — The nil-interface trap

**Problem:** Predict the output, then run:

```go
package main

import "fmt"

func main() {
	var p *int            // nil pointer
	var err error = p     // interface holding a typed nil
	fmt.Println(err == nil) // ?
}
```

**Example output:** `false`.

**Why it works:** an interface value is (type, value). `err` holds the type
`*int` with a nil pointer, so the interface itself is not nil — only its data
is. This is why a function returning `nil` explicitly must return `nil`
(untyped), and why interfaces wrapping typed nils can silently skip `err ==
nil` checks.

---

### Exercise 5.3 — Defer order and argument timing

**Problem:** Predict the print order, then run:

```go
package main

import "fmt"

func main() {
	defer fmt.Println("third")
	defer fmt.Println("second")
	defer fmt.Println("first")
}
```

**Example output:**

```text
first
second
third
```

**Why it works:** defers run LIFO — the last scheduled runs first. That is the
correct order for paired resources: `defer file.Close()` right after
`os.Open`, so a later `defer` (e.g. unlocking) runs before the file closes.
Also note arguments are evaluated when `defer` is scheduled, not when it runs:

```go
x := 1
defer fmt.Println(x) // prints 1
x = 2                // too late to change the deferred argument
```

## How To Keep Going

- Re-solve each exercise a week later without looking at the examples.
- Explain every `Why it works` block out loud in Vietnamese or English — if
  you cannot, re-read the linked module.
- Wire the module maps back to the repo exercises:
  `exercises/00-syntax-drills` → fundamentals, `exercises/07-concurrency` →
  concurrency, `apps/task-manager` → HTTP/architecture.
