# 04 - If And Switch

## Goal

Choose clear branches for score classification.

## Syntax

`if`, chained conditions, and expression `switch`.

## What It Does

Turns a numeric score into a label.

## Why It Matters

Validation and state transitions often depend on ordered branching.

## Mental Model

Go checks conditions top to bottom and stops at the first matching path.

## Annotated Example

```go
switch {
case score < 0:
	return "invalid"
case score < 50:
	return "fail"
}
```

## Common Mistakes

- Checking broad conditions before narrow boundary cases.
- Using `switch` without thinking about order.

## Exercise

Implement `ClassifyScore`.

## Acceptance Criteria

- Below `0` or above `100` is `"invalid"`.
- `0-49` is `"fail"`.
- `50-84` is `"pass"`.
- `85-100` is `"excellent"`.

## Verify

```bash
go test -tags exercise ./exercises/00-syntax-drills/04-if-switch/...
```

## Reflection

What made `switch` clearer than several unrelated `if` blocks?
