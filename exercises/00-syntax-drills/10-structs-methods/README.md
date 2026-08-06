# 10 - Structs And Methods

## Goal

Attach behavior to domain data and choose the right receiver for the job.

## Concepts

- Struct declarations with named fields
- Methods with receiver parameters
- Pointer receivers for mutation
- Value receivers for read-only behavior
- Validating input before updating state

## Syntax Primer

A method declares a receiver between `func` and the method name. A pointer receiver can change the caller's value; a value receiver operates on a copy:

```go
type Task struct {
	ID    int
	Title string
}

func (t *Task) Rename(title string) error {
	// t.Title changes the original Task
	return nil
}

func (t Task) Summary() string {
	return t.Title // reads a copy — fine for reading
}
```

## Mental Model

A method is a function with a named receiver value. The receiver decides who owns the state you touch: pointer receivers reach the original struct, value receivers work on a copy. When a method must change the struct, the pointer receiver is required.

## Annotated Examples

```go
type Note struct {
	text string
}

func (n *Note) Clear() {
	n.text = "" // the caller's Note changes
}

func (n Note) Length() int {
	return len(n.text) // read-only: value receiver is fine
}
```

## Common Diagnostics

- State unchanged after a call: the method has a value receiver where a pointer receiver is required.
- `cannot use task (variable of type Task) as *Task`: calling a pointer-receiver method on a non-addressable value; use a variable and pass its address.
- `missing return`: every code path through the method must produce the declared result.

## Exercise

Implement `Task.Rename` (pointer receiver) to trim and validate a new title, and `Task.Summary` (value receiver) to describe the task.

## Acceptance Criteria

- `Rename("  ship API  ")` sets `Title` to `"ship API"`.
- `Rename("   ")` returns an error; the title stays unchanged.
- `Summary` returns a non-empty string that includes the task ID and title.

## Hints

- Trim the input with `strings.TrimSpace` and reject it when it is empty.
- `Rename` needs the pointer receiver; `Summary` reads only and can use the value receiver.
- Format the summary with `fmt.Sprintf("%d: %s", t.ID, t.Title)` or similar.

## Verify

```bash
gofmt -w exercises/00-syntax-drills/10-structs-methods
go test -tags exercise ./exercises/00-syntax-drills/10-structs-methods/...
```

## Reflection Prompts

Why should `Rename` use a pointer receiver while `Summary` can use a value receiver? What happens if you forget to trim before validating?