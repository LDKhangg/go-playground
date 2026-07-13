# 02 - Functions

## Goal

Use parameters, multiple return values, and early error returns.

## Concepts

Parameters, named returns, integer division, remainder, scope, and sentinel errors.

## Exercise

Implement `Divide` so it returns an integer quotient and remainder. A zero divisor must return `ErrDivideByZero`.

## Acceptance Criteria

- `Divide(17, 5)` returns quotient `3`, remainder `2`, and no error.
- `Divide(10, 0)` returns an error matching `ErrDivideByZero`.

## Hints

Reject zero before using `/` or `%`. A clear explicit return is preferable to relying on bare named returns here.

## Commands

`go test -tags exercise ./exercises/02-functions/...`

## Reflection Prompts

Why can one Go function return both a result and an error? What would happen without the early zero check?
