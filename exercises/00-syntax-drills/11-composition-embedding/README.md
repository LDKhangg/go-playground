# 11 - Composition And Embedding

## Goal

Reuse behavior by combining types instead of inheriting from a base class.

## Concepts

- Struct composition
- Embedded fields (anonymous fields)
- Promoted fields and methods
- Go's preference for composition over inheritance

## Syntax Primer

An embedded field is a struct field declared by type name only. Its fields and methods are promoted to the outer struct:

```go
type Task struct {
	Title string
}

type TimedTask struct {
	Task          // embedded: Task's fields and methods are promoted
	DueInHours int
}

task := TimedTask{Task: Task{Title: "ship API"}, DueInHours: 4}
task.Title     // promoted field, == "ship API"
task.Summary() // promoted method from Task
```

## Mental Model

Embedding is not inheritance: there is no class hierarchy, just a value inside a value. The inner type keeps its identity (`task.Task`), while the outer type gains convenient access to its members. Composition keeps the relationship explicit and visible in the type.

## Annotated Examples

```go
type Address struct {
	City string
}

type User struct {
	Name    string
	Address // promoted: user.City works
}

u := User{Name: "Mina", Address: Address{City: "Hanoi"}}
_ = u.City // "Hanoi"
```

## Common Diagnostics

- `unknown field ... in struct literal`: the nested value is filled in by its own type name, e.g. `TimedTask{Task: ..., DueInHours: ...}`.
- `ambiguous selector`: two embedded types promote the same field name; qualify it explicitly.
- Expecting "override" behavior: a method on the outer type shadows a promoted one instead of overriding through a base-class mechanism.

## Exercise

Implement `TimedTask.Label` so it describes the timed task.

## Acceptance Criteria

- The label contains the embedded task's title.
- The label contains the due-hour count.

## Hints

- Read the title through the promoted field (`t.Title`).
- Include the number with `strconv.Itoa(t.DueInHours)` or `fmt.Sprintf`.
- Combine both into one string, e.g. `"ship API due in 4h"`.

## Verify

```bash
gofmt -w exercises/00-syntax-drills/11-composition-embedding
go test -tags exercise ./exercises/00-syntax-drills/11-composition-embedding/...
```

## Reflection Prompts

When would explicit composition be clearer than embedding? How is embedding different from inheritance in a language like Java?