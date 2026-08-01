# 08 - Object-Oriented Design

## Goal

Model a small task service using Go's object-oriented building blocks: structs,
methods, interfaces, composition, and dependency injection.

## Concepts

Structs, methods, pointer receivers, interfaces, composition, constructors,
dependency injection, and service boundaries.

## Syntax Primer

Go does not use classes or inheritance as its main design tools. A `struct`
stores data, a method adds behavior, and an interface describes a capability.
Composition means a larger type uses smaller types instead of inheriting from
them.

```go
type Clock interface {
    Now() time.Time
}

type Task struct {
    Title string
    Done  bool
}

func (t *Task) MarkDone(now time.Time) {
    t.Done = true
}
```

Use pointer receivers when a method must mutate a value or when copying the
struct would be wasteful. Inject dependencies through constructors so the
service can be tested with fakes instead of real infrastructure.

```go
type Service struct {
    repo  Repository
    clock Clock
}

func NewService(repo Repository, clock Clock) *Service {
    return &Service{repo: repo, clock: clock}
}
```

## Mental Model

Go object-oriented design is "behavior around data" rather than "class
hierarchies." A `Task` owns its own state transitions. A `TaskService`
coordinates collaborators. Interfaces stay small and describe only what a
consumer needs. This keeps code explicit, testable, and easy to replace with
fakes.

## Annotated Examples

This is an example of composition, not the exercise solution:

```go
type Notifier interface {
    Send(message string) error
}

type ReminderService struct {
    notifier Notifier
}

func (s ReminderService) Remind(task Task) error {
    return s.notifier.Send("remember: " + task.Title)
}
```

`ReminderService` does not care whether the notifier sends email, logs to the
terminal, or writes to a queue. It only cares about the behavior promised by
the interface.

## Common Diagnostics

- `cannot use X as Y value in argument`: the concrete type does not satisfy the interface you declared.
- `method has pointer receiver`: a value receiver and pointer receiver are not interchangeable in every call site.
- `nil pointer dereference`: a dependency was never initialized before use.
- If a service feels hard to test, the dependency boundary is probably too wide.

## Exercise

Implement `Task.MarkDone`, `TaskService.Create`, and `TaskService.Complete`.
Use the injected repository, clock, and ID generator instead of hardcoding
behavior.

## Acceptance Criteria

- `Create` trims the title and rejects empty input with `ErrEmptyTitle`.
- `Create` uses the injected `IDGenerator` and `Clock` to build the new task.
- `Create` saves the new task through the injected repository.
- `MarkDone` sets `Done` to `true` and records the completion timestamp.
- `Complete` loads the task, marks it done with the injected clock, persists the update, and returns the updated task.
- `Complete` returns `ErrTaskNotFound` when the repository cannot find the task.

## Hints

Keep interfaces small. Let `Task` own its state transition and let the service
coordinate storage and dependencies. Trim strings before validating them.

## Verify

```bash
gofmt -w exercises/08-object-oriented-design
go test -tags exercise ./exercises/08-object-oriented-design/...
```

## Reflection Prompts

Why is `Task.MarkDone` better as a method than as a free function? Which
dependencies belong in the service constructor, and which should stay outside
the type entirely?
