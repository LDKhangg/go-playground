# 02 - Functions

## Goal

Use parameters, multiple return values, early error returns, and a basic pointer.

## Concepts

Parameters, named returns, integer division, remainder, scope, sentinel errors, addresses, and pointer dereferencing.

## Exercise

Implement `Divide` so it returns an integer quotient and remainder. A zero divisor must return `ErrDivideByZero`. Implement `Double` so it updates the integer at the supplied pointer.

## Acceptance Criteria

- `Divide(17, 5)` returns quotient `3`, remainder `2`, and no error.
- `Divide(10, 0)` returns an error matching `ErrDivideByZero`.
- Passing the address of `6` to `Double` changes that value to `12`.

## Hints

Reject zero before using `/` or `%`. A clear explicit return is preferable to relying on bare named returns here. Dereference a pointer with `*` when reading or assigning the value it points to.

## Commands

`go test -tags exercise ./exercises/02-functions/...`

## Reflection Prompts

Why can one Go function return both a result and an error? How does passing `&number` let `Double` change the caller's value?
