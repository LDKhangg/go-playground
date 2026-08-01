# 08 - Arrays And Slices

## Goal

Insert and remove values from a slice safely.

## Syntax

Slice expressions, `append`, and index validation.

## What It Does

Builds a new slice around an insertion or removal point.

## Why It Matters

Most Go collection work happens with slices.

## Mental Model

A slice is a window over an underlying array; appending can allocate a new one.

## Annotated Example

```go
left := values[:index]
right := values[index:]
out := append(left, right...)
```

## Common Mistakes

- Accepting an out-of-range index.
- Mutating the caller's slice unexpectedly.

## Exercise

Implement `InsertAt` and `RemoveAt`.

## Acceptance Criteria

- Insert keeps original order around the new value.
- Remove drops one element.
- Invalid indices return an error.

## Verify

```bash
go test -tags exercise ./exercises/00-syntax-drills/08-arrays-slices/...
```

## Reflection

Why can slice helpers accidentally share backing arrays?
