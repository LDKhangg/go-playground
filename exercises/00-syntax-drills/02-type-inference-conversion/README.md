# 02 - Type Inference And Conversion

## Goal

Convert external string input into a typed Go value.

## Syntax

`:=`, explicit conversion, and `strconv.Atoi`.

## What It Does

Turns a string like `"8080"` into an integer port.

## Why It Matters

HTTP servers, config files, and CLI flags arrive as strings.

## Mental Model

Go infers a type from the right-hand side, but conversions stay explicit.

## Annotated Example

```go
raw := "8080"
port, err := strconv.Atoi(raw)
_ = port
_ = err
```

## Common Mistakes

- Expecting Go to convert strings to ints automatically.
- Ignoring the conversion error.

## Exercise

Implement `ParsePort`.

## Acceptance Criteria

- `"8080"` becomes `8080`.
- Non-numeric input returns an error.
- Negative values return an error.

## Verify

```bash
go test -tags exercise ./exercises/00-syntax-drills/02-type-inference-conversion/...
```

## Reflection

Why is explicit conversion safer than automatic conversion?
