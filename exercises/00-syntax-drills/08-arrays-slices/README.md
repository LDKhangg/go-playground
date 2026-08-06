# 08 - Arrays And Slices

## Goal

Insert and remove values in a slice while keeping the API safe for invalid indices.

## Concepts

- Arrays as fixed-length values
- Slices as windows over an underlying array
- `append` and the `...` spread operator
- Index validation and error returns
- Avoiding unintended mutation of the caller's slice

## Syntax Primer

A slice is a descriptor pointing into an array. Slicing an existing slice creates a new window that shares backing storage:

```go
values := []int{1, 2, 3}
left := values[:1]   // [1]
right := values[1:]  // [2, 3]
```

`append` grows a slice, allocating new backing storage when needed:

```go
out := append(left, right...) // spreads right's elements
```

`append` can also write into shared storage, so building a fresh result slice is the way to leave the caller's data untouched.

## Mental Model

A slice is a small descriptor (pointer, length, capacity), not the data itself. Appending may reuse existing capacity or allocate a new array — which is exactly why helpers must decide whether they mutate the caller's slice or return independent storage.

## Annotated Examples

```go
func merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	result = append(result, a...)
	result = append(result, b...)
	return result
}
```

## Common Diagnostics

- `index out of range` panic: a slice expression or `values[index]` used an index beyond `len(values)`. Validate before indexing.
- `cannot use values (variable of type []int) as [4]int`: slices and arrays are distinct types; the signature tells you which one is expected.
- Caller's slice changed after calling a helper: `append` reused the shared backing array. Build a new result slice instead.

## Exercise

Implement `InsertAt` and `RemoveAt` so they return new slices and reject invalid indices with an error.

## Acceptance Criteria

- `InsertAt([]int{1, 3}, 1, 2)` returns `[]int{1, 2, 3}`.
- `RemoveAt([]int{1, 2, 3}, 1)` returns `[]int{1, 3}`.
- An out-of-range index (inserting at `3` into a length-1 slice, or removing at `-1`) returns an error.

## Hints

- For `InsertAt`, valid indices run `0` through `len(values)` inclusive; for `RemoveAt`, `0` through `len(values)-1`.
- Build the result with slice expressions around the index, then `append` the pieces into fresh storage.
- Return `nil, err` on invalid input before doing any slicing.

## Verify

```bash
gofmt -w exercises/08-arrays-slices
go test -tags exercise ./exercises/08-arrays-slices/...
```

## Reflection Prompts

Why can slice helpers accidentally share backing arrays? When does `append` reuse existing capacity, and why does that matter for callers?