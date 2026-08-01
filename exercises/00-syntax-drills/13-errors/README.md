# 13 - Errors

## Goal

Return domain-specific errors instead of vague strings.

## Syntax

Sentinel errors, `fmt.Errorf`, and `errors.Is`.

## What It Does

Validates a task title.

## Why It Matters

Callers need stable ways to detect invalid input.

## Mental Model

An error value is part of the function contract, not an afterthought.

## Annotated Example

```go
if strings.TrimSpace(title) == "" {
	return ErrEmptyTitle
}
```

## Common Mistakes

- Comparing only error strings.
- Returning one giant generic error for all invalid states.

## Exercise

Implement `ValidateTitle`.

## Acceptance Criteria

- Blank titles fail.
- Overly long titles fail.
- Valid titles succeed.

## Verify

```bash
go test -tags exercise ./exercises/00-syntax-drills/13-errors/...
```

## Reflection

Why is `errors.Is` more stable than string matching?
