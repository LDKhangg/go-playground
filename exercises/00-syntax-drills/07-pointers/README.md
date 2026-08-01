# 07 - Pointers

## Goal

Mutate shared state deliberately through a pointer.

## Syntax

`*T`, `&value`, and nil checks.

## What It Does

Increments an integer through a pointer.

## Why It Matters

Methods and helpers often need to mutate a caller-owned value.

## Mental Model

A pointer stores an address; dereferencing reaches the original value.

## Annotated Example

```go
count := 1
ptr := &count
*ptr++
```

## Common Mistakes

- Dereferencing a nil pointer.
- Using a pointer when a returned copy would be clearer.

## Exercise

Implement `Increment`.

## Acceptance Criteria

- Adds one to the pointed value.
- Returns an error for nil.

## Verify

```bash
go test -tags exercise ./exercises/00-syntax-drills/07-pointers/...
```

## Reflection

When is a pointer better than returning a new value?
