# 04 - If And Switch

## Goal

Choose the clearest form of branching for ordered category checks.

## Concepts

- `if` and `else if` chains
- Expression `switch` with no operand
- Ordered conditions and short-circuiting
- Boundary values in range checks

## Syntax Primer

An expression `switch` with no operand tests each `case` condition top to bottom and runs the first one that is true:

```go
switch {
case score < 0:
	return "invalid"
case score < 50:
	return "fail"
}
```

The order matters: once a case matches, later cases are never evaluated, so `score < 50` already excludes negative scores when it comes after `score < 0`.

## Mental Model

Go checks conditions top to bottom and stops at the first match. Writing boundaries in ascending order turns a chain of comparisons into a readable table: each case handles exactly one contiguous range, and `default` catches everything else.

## Annotated Examples

```go
func speedLabel(kmh int) string {
	switch {
	case kmh <= 40:
		return "city"
	case kmh <= 90:
		return "road"
	default:
		return "highway"
	}
}
```

Note how `default` catches everything the earlier ranges did not cover — the same role it plays for out-of-range input.

## Common Diagnostics

- Wrong results at exact boundaries (49 vs 50): off-by-one in a `<` or `<=`; recheck which side of the boundary each test value falls on.
- `missing return`: every code path must return; a `default` case guarantees one.
- `syntax error: unexpected case`: the expression `switch` must have no operand, and `case` labels appear directly inside the `switch` body.

## Exercise

Implement `ClassifyScore` so a numeric score maps to a label.

## Acceptance Criteria

- Below `0` or above `100` is `"invalid"`.
- `0` through `49` is `"fail"`.
- `50` through `84` is `"pass"`.
- `85` through `100` is `"excellent"`.

## Hints

- Order the cases from the narrowest range upward.
- Let `default` return `"invalid"` for anything above 100 (and anything the earlier cases missed).
- Verify boundary values 0, 49, 50, 84, 85, and 100 mentally against the tests.

## Verify

```bash
gofmt -w exercises/00-syntax-drills/04-if-switch
go test -tags exercise ./exercises/00-syntax-drills/04-if-switch/...
```

## Reflection Prompts

What made `switch` clearer than several unrelated `if` blocks? Where does the boundary between `pass` and `excellent` actually live in the tests?
