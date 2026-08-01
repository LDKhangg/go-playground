# 05 - For And Range

## Goal

Use Go's one loop construct to process a collection.

## Syntax

`for`, `range`, `break`, and accumulation.

## What It Does

Adds numbers until a sentinel value appears.

## Why It Matters

Many programs scan values until a stop condition or invalid state appears.

## Mental Model

`range` hands you each value in sequence; `break` stops the loop early.

## Annotated Example

```go
total := 0
for _, value := range values {
	total += value
}
```

## Common Mistakes

- Forgetting to stop once the sentinel appears.
- Reusing the wrong accumulator variable.

## Exercise

Implement `SumUntil`.

## Acceptance Criteria

- Adds values before the sentinel.
- Stops once `stopAt` appears.
- Returns `0` for an empty slice.

## Verify

```bash
go test -tags exercise ./exercises/00-syntax-drills/05-for-range/...
```

## Reflection

Why is a sentinel-based loop common in parser and stream code?
