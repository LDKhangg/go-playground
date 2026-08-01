# 12 - Interfaces And Assertions

## Goal

Accept behavior through interfaces and safely inspect optional capabilities.

## Syntax

Interface types, `io.Writer`, and type assertions.

## What It Does

Writes a greeting and checks whether an error is temporary.

## Why It Matters

Small interfaces keep code testable and decoupled.

## Mental Model

An interface stores a concrete value plus the methods it exposes.

## Annotated Example

```go
if temp, ok := err.(Temporary); ok {
	_ = temp
}
```

## Common Mistakes

- Defining interfaces before knowing the caller's real need.
- Using unsafe assertions without checking `ok`.

## Exercise

Implement `WriteGreeting` and `AsTemporary`.

## Acceptance Criteria

- Writes `hello, <name>` to the writer.
- Returns the writer error if writing fails.
- Returns comma-ok style assertion results.

## Verify

```bash
go test -tags exercise ./exercises/00-syntax-drills/12-interfaces-assertions/...
```

## Reflection

Why is a small input interface easier to fake in tests?
