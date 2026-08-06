# 05 - For And Range

## Goal

Use Go's single loop construct to process a collection and stop at a sentinel.

## Concepts

- `for range` over a slice
- The `_` index discard
- Early `return` as a stop signal
- Accumulator variables
- Looping over an empty slice

## Syntax Primer

Go has one loop keyword: `for`. The `range` form visits every element of a slice, exposing the index and the value:

```go
total := 0
for _, value := range values {
	total += value
}
```

A function can leave the loop early with `return`, which stops processing and hands control back with a result:

```go
for _, value := range values {
	if value == stop {
		return total
	}
}
```

## Mental Model

`range` hands you each value in sequence; the loop body decides what to do with it. `return` is the loop's stop signal. Ranging over a `nil` slice is legal and simply visits zero elements, which makes loops naturally safe for empty input.

## Annotated Examples

```go
func firstPositive(values []int) int {
	for _, value := range values {
		if value > 0 {
			return value
		}
	}
	return -1 // no positive value found
}
```

## Common Diagnostics

- `declared and not used: index`: use `_` for the index when you only need the value.
- Wrong totals: resetting or misplacing the accumulator inside the loop so it is overwritten every iteration.
- Sentinel not stopping the loop: forgetting to check for the stop condition inside the body.

## Exercise

Implement `SumUntil` so it adds values from the slice until it meets `stopAt`.

## Acceptance Criteria

- Values before the sentinel are summed.
- The loop stops when `stopAt` appears, and `stopAt` itself is not added.
- An empty slice returns `0`.

## Hints

- Accumulate in a `total := 0` variable outside the loop.
- Compare each value to `stopAt`; on a match return the current total immediately.
- If the loop finishes without a match, return the total.

## Verify

```bash
gofmt -w exercises/00-syntax-drills/05-for-range
go test -tags exercise ./exercises/00-syntax-drills/05-for-range/...
```

## Reflection Prompts

Why is a sentinel-based loop common in parser and stream code? What should `SumUntil` return when `stopAt` is never found?