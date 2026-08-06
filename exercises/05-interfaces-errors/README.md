# 05 - Interfaces and Errors

## Goal

Depend on a small behavior and preserve error identity while adding context.

## Concepts

- An interface describes behavior, not data.
- A type satisfies an interface implicitly by implementing its methods.
- Dependency injection passes a collaborator into the code that needs it.
- Sentinel errors provide a stable identity for a known condition.
- `%w` wraps an error so callers can still use `errors.Is`.

## Syntax Primer

```go
type TitleValidator interface {
	Validate(title string) error
}

var ErrEmptyTitle = errors.New("title must not be empty")

if err := validator.Validate(title); err != nil {
	return fmt.Errorf("validate title %q: %w", title, err)
}
```

An `error` is an interface value. `if err := ...; err != nil` creates `err` only for the `if` statement and its branches. Use `%w`, not `%v`, when callers must be able to identify the original error later.

## Mental Model

An interface is a small contract: "give me any value that can validate a title." The consuming function should know only that contract, which makes it easier to test and replace collaborators. Wrapping an error adds a breadcrumb about where it failed while retaining the original error inside it.

## Annotated Examples

```go
type PrefixChecker struct{}

func (PrefixChecker) Validate(name string) error {
	if name == "" {
		return ErrEmptyTitle
	}
	return nil
}
```

```go
func report(err error) string {
	if errors.Is(err, ErrEmptyTitle) {
		return "Please enter a title."
	}
	return "Unexpected validation failure."
}
```

`PrefixChecker` needs no declaration that it implements `TitleValidator`; its method signature is enough.

## Common Diagnostics

- `does not implement TitleValidator`: the supplied type is missing `Validate(string) error` or its signature differs.
- `undefined: fmt`: import `fmt` before calling `fmt.Errorf`.
- `errors.Is` returns false after adding context: make sure the formatting verb is `%w`, not `%v`.
- A tagged test fails with a missing error context: this is an unmet exercise behavior, not a syntax error.

## Exercise

Implement `ValidateTitle` by calling the supplied validator. Return `nil` on success; otherwise wrap the validator error with useful context while preserving its identity.

## Acceptance Criteria

- The validator receives the original title.
- A successful validator produces no error.
- A failed validator produces contextual text while still matching `ErrEmptyTitle` through `errors.Is`.

## Hints

- Define the interface at the point where it is consumed.
- Store the validator result in an `if` initializer.
- Include the title or validation operation in the message, then use `%w` for the original error.

## Verify

Run:

```bash
gofmt -w exercises/05-interfaces-errors
go test -tags exercise ./exercises/05-interfaces-errors/...
```

The starter intentionally fails until `ValidateTitle` calls and wraps its validator. That assertion failure is separate from a compiler or gopls syntax diagnostic.

## Reflection Prompts

- Why does wrapping preserve more information than formatting with `%v`?
- Why is this interface only one method?
- What changes if the validation function directly constructs a concrete validator instead?
