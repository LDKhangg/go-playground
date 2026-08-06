# 07 - Pointers

## Goal

Mutate a caller-owned value deliberately through a pointer, safely.

## Concepts

- Pointer types `*T`
- Address-of `&` and dereference `*`
- Nil pointers and nil checks
- Why mutation needs a pointer instead of a copy

## Syntax Primer

`&value` produces a pointer to a variable; `*pointer` reads or writes the value it points to:

```go
count := 1
ptr := &count        // ptr is *int
*ptr = *ptr + 1      // same as count = count + 1
```

Because a pointer can also be `nil` (pointing at nothing), dereferencing requires a check first:

```go
if value == nil {
	return errors.New("value must not be null")
}
*value += 1
```

## Mental Model

A pointer stores an address. Passing a pointer hands the function the address of your variable instead of a copy, so writes through the pointer reach the original. Arguments passed by value are copies; changing a copy never changes the caller's variable.

## Annotated Examples

```go
func Reset(number *int) {
	if number == nil {
		return
	}
	*number = 0
}
```

## Common Diagnostics

- `panic: runtime error: invalid memory address or nil pointer dereference`: dereferencing without a nil check. Validate the pointer first.
- `invalid operation: cannot indirect value ...`: `*x` only compiles when `x` has a pointer type.
- Caller's value unchanged: the function received a copy or wrote only to a local variable.

## Exercise

Implement `Increment` so it adds one to the integer at the supplied pointer.

## Acceptance Criteria

- `Increment` adds one to the pointed value (a variable holding `4` becomes `5`).
- Passing `nil` returns an error instead of crashing.

## Hints

- Check `value == nil` first and return an error.
- Dereference with `*value` when assigning: `*value = *value + 1`.

## Verify

```bash
gofmt -w exercises/00-syntax-drills/07-pointers
go test -tags exercise ./exercises/07-pointers/...
```

## Reflection Prompts

When is a pointer better than returning a new value? Why does Go make pointer passing explicit rather than implicit?