# 14 - Generics

## Goal

Write one algorithm that works across several types with type parameters.

## Concepts

- Type parameters `[T any]`
- The `comparable` constraint
- Type inference at call sites
- Avoiding duplication without losing type safety

## Syntax Primer

A generic function declares type parameters in brackets. The `comparable` constraint allows `==` and `!=` inside the body:

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

Callers never write the type argument explicitly: `Contains([]int{1, 2}, 2)` infers `T = int`, and `Contains([]string{"go"}, "go")` infers `T = string`.

## Mental Model

The function is written once and specialized by the compiler for each type it is used with. `comparable` is a promise from the type to the function: "you may compare me with `==`." The result is one body, fully type-checked, with no `interface{}` casting at call sites.

## Annotated Examples

```go
func IndexOf[T comparable](items []T, target T) int {
	for i, item := range items {
		if item == target {
			return i
		}
	}
	return -1
}
```

## Common Diagnostics

- `T does not implement comparable`: the constraint is missing or too weak; use `[T comparable]` when the body compares values.
- `cannot compare T (missing comparable)`: comparing with `==` requires the `comparable` constraint.
- `syntax error: unexpected [`: type parameters come immediately after the function name, before the parameter list.

## Exercise

Implement `Contains` so it reports whether a slice holds a target value.

## Acceptance Criteria

- Works for `[]int` (contains `2` in `[1 2 3]`).
- Works for `[]string` (contains `"go"`).
- Returns `false` when the target is missing.

## Hints

- Keep the `[T comparable]` signature from the starter.
- Loop with `for _, item := range items` and compare `item == target`.
- Return `true` on the first match and `false` after the loop.

## Verify

```bash
gofmt -w exercises/00-syntax-drills/14-generics
go test -tags exercise ./exercises/00-syntax-drills/14-generics/...
```

## Reflection Prompts

When is a generic helper better than two copied functions? Why is `comparable` needed instead of `any`?