# 04 - Architecture And Testing

Interviewers ask "how would you structure this?" almost as often as "what does
this do?". Study `exercises/08-object-oriented-design` (starers, dependency
injection) and `exercises/06-testing` (table-driven tests) for the concrete
versions of these ideas.

## 4.1 Package Layout

**Q: How do you lay out a Go service?**

Small, dependency-lean services do fine with `cmd/` for executables and
`internal/` for code nobody outside the module should import:

```text
apps/task-manager/
  cmd/api/            -> main: wires everything, starts the server
  internal/httpapi/   -> handlers, middleware
  internal/tasks/     -> domain types + service (no HTTP imports)
```

`internal/` is enforced by the compiler: nothing outside the module can import
it. Keep `main` thin — it should glue dependencies, not contain logic.

**Q: What belongs where?**

- HTTP concerns (requests, responses, status codes) → handler package.
- Business rules (validation, state transitions, invariants) → domain/service package.
- Persistence → repository package behind an interface.
- `main` → composition and configuration only.

## 4.2 Dependency Injection

**Q: What is dependency injection in Go?**

A service takes its collaborators as constructor arguments instead of
constructing them itself. This makes behavior replaceable (fakes in tests) and
makes dependencies explicit.

```go
type TaskService struct {
	repo  Repository
	clock Clock
	ids   IDGenerator
}

func NewTaskService(repo Repository, clock Clock, ids IDGenerator) *TaskService {
	return &TaskService{repo: repo, clock: clock, ids: ids}
}
```

**Q: Why inject a `Clock` of all things?**

So tests can control time. `stubClock` returns a fixed timestamp instead of
`time.Now()`, and tests assert exact timestamps. If the code called
`time.Now()` directly, that assertion would be impossible.

Practice: `exercises/08-object-oriented-design`.

**Q: Where to define interfaces — provider or consumer?**

At the consumer. Define the smallest interface the caller needs; the concrete
type satisfies it implicitly. This avoids coupling to implementations you do
not use. Example: `exercises/05-interfaces-errors` defines `TitleValidator`
where it is consumed.

## 4.3 Repository Pattern

**Q: What is the repository pattern?**

An interface around persistence so the domain does not know whether data lives
in memory, a database, or a file:

```go
type Repository interface {
	Save(task Task) error
	FindByID(id string) (Task, error)
}
```

The service calls `repo.Save(...)`. In tests you supply a `fakeRepository`,
and later the real implementation can be swapped without touching the service.

## 4.4 Service Boundaries

**Q: How do you keep service boundaries clean?**

Each unit answers three questions: what does it do? how do you use it? what does
it depend on? If a type is hard to test, the boundary is usually too wide.
`Task` owning its own state transitions (`MarkDone`) and the service
coordinating collaborators is the shape to aim for.

## 4.5 Table-Driven Tests

**Q: What is a table-driven test and why prefer it?**

One test function, a slice of `{name, input, want}` cases, each run as its own
subtest. It removes copy-paste, makes failures nameable, and documents the
behavior contract in one place.

```go
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

Practice: `exercises/06-testing`.

**Q: When do you need `tc := tc`?**

Only when the subtest runs in parallel (`t.Parallel()`). Before Go 1.22 the
loop variable was shared, so subtests racing could read the wrong row; since
1.22 each iteration has its own variable, but the `tc := tc` habit is still
harmless and appears in older code.

## 4.6 Testing HTTP Handlers

**Q: How do you test a handler without a running server?**

Use `httptest` — `httptest.NewRecorder` + a real `http.Request`:

```go
req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
rec := httptest.NewRecorder()

myHandler.ServeHTTP(rec, req)

if rec.Code != http.StatusOK {
	t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
}
```

## 4.7 Mocks, Stubs, Fakes

**Q: Do you need a mocking library?**

No. Your own interfaces plus tiny hand-written fakes are clearer, dependency-
free, and test exactly what the caller cares about. A stub is a canned answer;
a fake is a mini-implementation with behavior. Start with fakes and stubs; add a
mock library only when hand-writing becomes tedious.

```go
type validatorFunc func(string) error

func (f validatorFunc) Validate(title string) error { return f(title) }

// use: ValidateTitle(validatorFunc(func(s string) error { ... }), input)
```

Practice: `exercises/05-interfaces-errors` (the test file uses exactly this
`validatorFunc` trick).

## 4.8 Benchmarks

**Q: How does a benchmark work?**

`func BenchmarkXxx(b *testing.B)` loops `b.N` times; Go picks `N` to get a
stable sample. Setup must sit outside the measured loop — call `b.ResetTimer()`
after setup.

```go
func BenchmarkClassify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Classify(42)
	}
}
```

Run with `go test -bench .`. Be careful comparing benchmark runs across
machines; use `-benchmem` to see allocations.

## 4.9 Coverage

`go test -cover` reports which statements executed. It proves execution, not
correctness: 100% coverage can still miss bugs, because coverage does not check
assertions, and branches the input never takes are invisible. Prefer meaningful
assertions over chasing a percentage.

## 4.10 Common Architecture Questions

- **Composition vs inheritance:** Go has no classes or inheritance. Features/behavior is achieved through struct composition and embedding ("has-a"), not class hierarchies.
- **Why small interfaces:** an interface one method (e.g., `io.Writer`) is easy to fake and leaves the concrete type free.
- **Why constructors returning concrete types:** idiomatic Go; the interface is defined at the consumer, so a factory function returning an interface is often an anti-pattern.
- **Why not a framework for everything:** the standard library plus a few small packages is often enough; each dependency is a maintenance and security cost.
- **How do you know a boundary is too wide:** if faking it is painful or a unit does three different jobs, split it.

## Practice Check

1. Why put dependencies in a constructor instead of `time.Now()`/`&MyRepo{}` inside a method?
2. Write the `Repository` interface your task service would need, then a `fakeRepository`.
3. Test a handler with `httptest` without a socket.
4. In your own words: composition vs inheritance in Go.
5. What does coverage prove and not prove?