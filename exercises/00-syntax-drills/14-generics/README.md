# 14 - Generics

## Goal

Reuse one algorithm across several comparable types.

## Syntax

Type parameters and `comparable` constraints.

## What It Does

Checks whether a slice contains a target value.

## Why It Matters

Generics remove duplication without hiding behavior behind `interface{}`.

## Mental Model

The function is written once, then specialized by the compiler for each type use.

## Annotated Example

```go
func Contains[T comparable](items []T, target T) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}
```

## Common Mistakes

- Using generics when one concrete type would be simpler.
- Choosing a constraint that is too broad or too narrow.

## Exercise

Implement `Contains`.

## Acceptance Criteria

- Works for `[]int`.
- Works for `[]string`.
- Returns false when missing.

## Verify

```bash
go test -tags exercise ./exercises/00-syntax-drills/14-generics/...
```

## Reflection

When is a generic helper better than two copied functions?
