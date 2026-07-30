# 02 - Functions

## Goal

Use parameters, multiple return values, early error returns, and a basic pointer.

## Concepts

Parameters, named returns, integer division, remainder, scope, sentinel errors, addresses, and pointer dereferencing.

## Syntax Primer

A function declares its input types after parameter names and its output types after the parameter list.

```go
func remainder(left, right int) int {
    return left % right
}

func lookup(id int) (string, error) {
    return "", nil
}
```

Go functions can return several values. `error` is an interface value; the convention is `nil` when no error occurred. Create a stable sentinel error with `errors.New`. Use `&value` to obtain an address and `*pointer` to read or write the value at that address.

## Mental Model

Function parameters are local names. Passing an `int` copies its value, while passing `*int` copies an address that still points at the caller's original integer. Multiple results let a function return its useful result and its failure state together. Return early when an input makes the normal operation invalid.

## Annotated Examples

Call the starter APIs by receiving all their results before using them:

```go
quotient, remainder, err := Divide(17, 5)
if err != nil {
    return err // Stop before using invalid results.
}
fmt.Println(quotient, remainder)

number := 6
Double(&number) // `&number` passes the address, not another copy of 6.
fmt.Println(number)
```

The example demonstrates calling the functions. Your exercise is to make their behavior meet the tests.

## Common Diagnostics

- `assignment mismatch`: the number of variables on the left must match the number of values returned, unless you intentionally discard one with `_`.
- `invalid operation: cannot indirect`: `*value` only works when `value` is a pointer type such as `*int`.
- An error that is ignored can hide a failed operation. Check `err` before using a result that depends on it.
- `integer divide by zero` is a runtime failure. Validate the divisor before applying `/` or `%`.

## Exercise

Implement `Divide` so it returns an integer quotient and remainder. A zero divisor must return `ErrDivideByZero`. Implement `Double` so it updates the integer at the supplied pointer.

## Acceptance Criteria

- `Divide(17, 5)` returns quotient `3`, remainder `2`, and no error.
- `Divide(10, 0)` returns an error matching `ErrDivideByZero`.
- Passing the address of `6` to `Double` changes that value to `12`.

## Hints

Reject zero before using `/` or `%`. A clear explicit return is preferable to relying on bare named returns here. Dereference a pointer with `*` when reading or assigning the value it points to.

## Verify

```bash
gofmt -w exercises/02-functions
go test -tags exercise ./exercises/02-functions/...
```

## Reflection Prompts

Why can one Go function return both a result and an error? How does passing `&number` let `Double` change the caller's value?
