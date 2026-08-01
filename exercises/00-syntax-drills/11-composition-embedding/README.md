# 11 - Composition And Embedding

## Goal

Reuse behavior by combining types instead of inheriting from a base class.

## Syntax

Nested structs and embedded fields.

## What It Does

Builds a timed task from a plain task.

## Why It Matters

Go favors composition over inheritance.

## Mental Model

Embedding promotes fields and methods; composition keeps relationships explicit.

## Annotated Example

```go
type TimedTask struct {
	Task
	DueInHours int
}
```

## Common Mistakes

- Treating embedding as a class hierarchy.
- Hiding where data really lives.

## Exercise

Implement `Label`.

## Acceptance Criteria

- Reuses the embedded task title.
- Includes the due-hour information.

## Verify

```bash
go test -tags exercise ./exercises/00-syntax-drills/11-composition-embedding/...
```

## Reflection

When would explicit composition be clearer than embedding?
