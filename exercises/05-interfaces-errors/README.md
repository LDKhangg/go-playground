# 05 - Interfaces and Errors

## Goal

Depend on a small behavior and preserve error identity while adding context.

## Concepts

Interfaces, dependency injection, sentinel errors, wrapping with `%w`, and `errors.Is`.

## Exercise

Implement `ValidateTitle` by calling the supplied validator. Return nil on success; otherwise wrap the validator error with useful context.

## Acceptance Criteria

- The validator receives the original title.
- A successful validator produces no error.
- A failed validator produces contextual text while still matching `ErrEmptyTitle` through `errors.Is`.

## Hints

Use `fmt.Errorf` and the `%w` verb. Define interfaces near the code that consumes them.

## Commands

`go test -tags exercise ./exercises/05-interfaces-errors/...`

## Reflection Prompts

Why does wrapping preserve more information than formatting with `%v`? Why is this interface only one method?
