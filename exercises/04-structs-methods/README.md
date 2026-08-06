# 04 - Structs and Methods

## Goal

Model state with structs, compose one type from another, and change state through methods.

## Concepts

- A struct groups named fields into one value.
- Constructors are ordinary functions that return an initialized value.
- Methods attach behavior to a named type through a receiver.
- Pointer receivers can mutate the original value.
- Composition lets one struct hold another and delegate behavior to it.

## Syntax Primer

```go
type Note struct {
	text string
	done bool
}

func (n *Note) MarkDone() {
	n.done = true
}

func (n Note) IsDone() bool {
	return n.done
}
```

Fields beginning with a lowercase letter are unexported: code outside this package cannot access them directly. `*Note` means a pointer to a `Note`; use it when a method must update the caller's value.

## Mental Model

A struct is a record of related state. A value receiver gets a copy of that record, so changing its fields changes only the copy. A pointer receiver receives the address of the original record, so its changes persist. A `Project` holding a `Task` is composition: the project owns a task rather than inheriting from it.

## Annotated Examples

```go
type Timer struct {
	seconds int
}

func NewTimer(seconds int) Timer {
	return Timer{seconds: seconds} // Named-field literal is explicit.
}

func (t *Timer) Tick() {
	if t.seconds > 0 {
		t.seconds-- // Updates the original Timer.
	}
}
```

```go
type Session struct {
	timer Timer
}

func (s *Session) TickTimer() {
	s.timer.Tick() // Delegate rather than duplicate Timer logic.
}
```

## Common Diagnostics

- State does not change after calling a method: check whether the method uses a value receiver when it needs `*Type`.
- `cannot refer to unexported field`: a lowercase field belongs to another package and must be accessed through exported behavior.
- `cannot use X as Y`: confirm that a constructor returns the exact struct type expected by the caller.
- Yellow unused-field warnings in an unfinished starter are static analysis, not Go syntax errors; this repository config disables that noise while you build the exercise.

## Exercise

Make `NewTask` preserve the title, make `Complete` mutate the task, and make `IsComplete` report current state. Then make `Project` compose a current `Task`: initialize it in `NewProject` and delegate the project's completion methods to that task.

## Acceptance Criteria

- A new task preserves its title and starts incomplete.
- Calling `Complete` changes that same task to complete.
- A new project contains a current task with the supplied title.
- Completing a project's current task is visible through `IsCurrentComplete`.

## Hints

- Build a struct value with a literal such as `Task{title: title}`.
- A mutating method needs a pointer receiver; a read-only method can use a value receiver here.
- Call the `Task` behavior from `Project` instead of reimplementing the state transition.

## Verify

Run:

```bash
gofmt -w exercises/04-structs-methods
go test -tags exercise ./exercises/04-structs-methods/...
```

The starter intentionally fails this command until you implement all methods. A failing assertion describes unmet exercise behavior; it is not a syntax or package-loading error.

## Reflection Prompts

- Why would a value receiver fail to persist a completion change?
- How does composition let `Project` reuse `Task` behavior?
- What behavior should remain private to a type rather than exposed as a field?
