# 01 - Basics

## Goal

Implement age-based ticket pricing and use a loop to total several tickets.

## Concepts

Variables, constants, integers, comparisons, `if`, `switch`, and `for` loops.

## Syntax Primer

Go gives every variable a type. Use `var` when you want to state that type, and `:=` when Go can infer it from the right-hand side.

```go
var age int = 18
name := "Mina"
const weekdayFee = 9
```

An unassigned `int` has the zero value `0`; an unassigned `string` has `""`. Integer division stays integer division, so convert deliberately when a calculation needs another numeric type: `float64(age)`.

Comparisons produce booleans: `age < 13`, `age >= 65`, and `name == "Mina"`. Use `if` when one branch depends on a condition. Use `switch` when several cases describe the same value or expression. Use `for range` to visit each element in a slice.

```go
for _, age := range ages {
    if age < 18 {
        // Handle one category.
    }
}
```

## Mental Model

A variable is a named box holding a value of one fixed type. A constant is a named value that cannot change. Conditions choose one path through a program; a loop repeats one path for each input item. The order of conditions matters because Go stops at the first matching branch.

## Annotated Examples

This is a small ticket-reporting loop, not the exercise implementation:

```go
ages := []int{11, 25, 70} // A slice holds a variable number of ints.
adults := 0

for _, age := range ages { // `_` ignores the index; `age` receives each value.
    if age >= 18 {
        adults++ // `++` increments an integer variable.
    }
}

switch adults {
case 0:
    fmt.Println("no adult tickets")
default:
    fmt.Println("adult tickets:", adults)
}
```

## Common Diagnostics

- `declared and not used`: every local variable must be read. Remove it, use it, or replace an intentionally ignored value with `_`.
- `mismatched types int and string`: Go does not silently convert between types. Convert explicitly only when the conversion makes sense.
- `syntax error: unexpected ...`: braces and parentheses must balance. Let `gofmt` format the file after each small change.
- A yellow style warning is not the same as a compiler error. Focus first on diagnostics that prevent `go test` from building.

## Exercise

Implement `TicketPrice`. Ages below 13 cost 5, ages from 13 through 64 cost 12, and ages 65 or above cost 7. Then implement `TotalTicketPrice` by looping over every supplied age and adding its ticket price.

## Acceptance Criteria

- Ages 12 and below cost `5`.
- Ages 13 through 64 cost `12`.
- Ages 65 and above cost `7`.
- `TotalTicketPrice([]int{8, 30, 70})` returns `24`.

## Hints

Name the three prices as constants. Check boundaries in an order that makes every age belong to one group. A `for range` loop can visit every age in the slice.

## Verify

```bash
gofmt -w exercises/01-basics
go test -tags exercise ./exercises/01-basics/...
```

## Reflection Prompts

Why does the order of age checks matter? What value should the total loop produce for an empty slice?
