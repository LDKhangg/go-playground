# 10 - Structs And Methods

## Goal

Attach behavior to domain data.

## Syntax

Struct declarations, methods, and pointer receivers.

## What It Does

Renames a task and summarizes it for display.

## Why It Matters

Most Go business logic lives in small concrete types with methods.

## Mental Model

A method is a function with a named receiver value.

## Annotated Example

```go
type Task struct { Title string }
func (t *Task) Rename(title string) { t.Title = title }
```

## Common Mistakes

- Using a value receiver when you need mutation.
- Forgetting to validate input before updating state.

## Exercise

Implement `Rename` and `Summary`.

## Acceptance Criteria

- Blank titles return an error.
- Valid titles are trimmed.
- `Summary` includes ID and title.

## Verify

```bash
go test -tags exercise ./exercises/00-syntax-drills/10-structs-methods/...
```

## Reflection

Why should `Rename` use a pointer receiver while `Summary` can use a value receiver?
