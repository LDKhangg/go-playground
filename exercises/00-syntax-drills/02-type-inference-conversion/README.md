# 02 - Type Inference And Conversion

## Goal

Turn external string input into a typed Go value, and be explicit about conversions.

## Concepts

- Type inference with `:=`
- Explicit numeric conversions
- `strconv.Atoi` for string-to-int parsing
- The `(value, error)` return convention

## Syntax Primer

`:=` infers the variable's type from the right-hand side. Go does not convert between types silently, so string-to-number parsing is an explicit operation that can fail.

```go
raw := "8080"               // inferred string
port, err := strconv.Atoi(raw)
if err != nil {
	return 0, err
}
```

When types differ, Go requires an explicit conversion: `int8(age)`, `float64(count)`. A `string` to an `int` is not a conversion but a parse, which is why `strconv.Atoi` returns an `error` as well as the value.

## Mental Model

Go infers types but never guesses conversions. Text arriving from the outside world (config files, CLI flags, HTTP headers) is just bytes until you parse it; parsing produces a value and a possible error, and the error is part of the contract.

## Annotated Examples

```go
width := "1920"                 // string
px, err := strconv.Atoi(width)  // px is an int
if err != nil {
	return err
}
_ = px

size := float64(px) // explicit conversion to a new type
```

## Common Diagnostics

- `cannot use raw (variable of type string) as int value in ...`: there is no implicit conversion; call `strconv.Atoi`.
- `mismatched types int and string`: convert explicitly before combining.
- `strconv.Atoi: parsing "abc": invalid syntax`: the returned error is information, not a crash — check it.
- Ignoring the second return value hides parse failures.

## Exercise

Implement `ParsePort` so a valid numeric string becomes an `int`, and any invalid input produces an error.

## Acceptance Criteria

- `ParsePort("8080")` returns `(8080, nil)`.
- Non-numeric input such as `"abc"` returns an error.
- Negative values such as `"-1"` return an error.

## Hints

- Parse with `strconv.Atoi` and return `(0, err)` on failure.
- Check the parsed value for `val < 0` and return a descriptive error.
- Return the value and `nil` only at the end.

## Verify

```bash
gofmt -w exercises/00-syntax-drills/02-type-inference-conversion
go test -tags exercise ./exercises/00-syntax-drills/02-type-inference-conversion/...
```

## Reflection Prompts

Why is explicit conversion safer than automatic conversion? What would happen if `ParsePort` silently defaulted invalid input to `0`?
