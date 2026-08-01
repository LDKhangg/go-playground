# 01 - Variables And Zero Values

## Goal

Understand what Go gives you before you assign any value.

## Syntax

`var`, short declaration `:=`, named struct fields, and zero values.

## What It Does

Go initializes every variable to a predictable zero value.

## Why It Matters

You can reason about defaults without uninitialized memory bugs.

## Mental Model

A declared variable is always a valid box with a type and a default value.

## Annotated Example

```go
var count int        // 0
var name string      // ""
active := false      // inferred bool
_ = active
```

## Common Mistakes

- Assuming an `int` starts as `nil`.
- Forgetting that a struct field also gets a zero value.

## Exercise

Implement `DefaultProfile` so it returns a zero-valued profile.

## Acceptance Criteria

- `Name` is `""`.
- `Active` is `false`.
- `Attempts` is `0`.

## Verify

```bash
go test -tags exercise ./exercises/00-syntax-drills/01-variables-zero-values/...
```

## Reflection

Why are zero values useful in API and config code?
