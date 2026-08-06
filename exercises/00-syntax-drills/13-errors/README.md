# 13 - Errors

## Goal

Return domain-specific errors that callers can detect reliably.

## Concepts

- Errors as values
- Sentinel errors with `errors.New`
- `errors.Is` for comparison
- Validation ordering

## Syntax Primer

A sentinel error is a package-level error value that names a known failure:

```go
var ErrEmptyTitle = errors.New("empty title")
```

Functions return it directly, and callers compare with `errors.Is` rather than by string equality:

```go
if err := ValidateTitle(title); errors.Is(err, ErrEmptyTitle) {
	// handle the known condition
}
```

## Mental Model

An error value is part of the function's contract: the same condition always produces the same sentinel, so callers can branch on identity instead of fragile message text. Validation functions convert bad input into exactly one stable error.

## Annotated Examples

```go
var ErrNegativeBalance = errors.New("balance cannot be negative")

func Withdraw(balance, amount int) error {
	if balance-amount < 0 {
		return ErrNegativeBalance
	}
	return nil
}
```

## Common Diagnostics

- `errors.Is` never matches: returning `fmt.Errorf("...")` without `%w` creates a new error instead of wrapping the sentinel.
- Comparing `err == ErrEmptyTitle` directly: works for bare sentinels but breaks once an error is wrapped; `errors.Is` handles both.
- `undefined: ErrEmptyTitle`: the sentinel must be declared at package level before the function uses it.

## Exercise

Implement `ValidateTitle` so it distinguishes blank titles from overly long ones.

## Acceptance Criteria

- A blank (whitespace-only) title returns `ErrEmptyTitle`.
- A title longer than 80 characters returns `ErrTitleTooLong`.
- A valid title like `"ship API"` returns `nil`.

## Hints

- Trim whitespace and check `strings.TrimSpace(title) == ""`.
- Compare the length with `len(title) > 80`.
- Return the sentinels directly so `errors.Is` matches.

## Verify

```bash
gofmt -w exercises/00-syntax-drills/13-errors
go test -tags exercise ./exercises/00-syntax-drills/13-errors/...
```

## Reflection Prompts

Why is `errors.Is` more stable than string matching? What order should the blank and length checks run in, and why?