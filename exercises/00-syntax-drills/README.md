# 00 - Syntax Drills

## Goal

Practice one Go syntax idea at a time before combining those ideas in the larger
chapter exercises.

## How This Track Works

Each folder in `00-syntax-drills` focuses on a single skill:

- zero values and declarations
- type inference and conversion
- strings and runes
- branching and loops
- functions, pointers, and `defer`
- slices, maps, methods, interfaces, errors, generics, and JSON

Every drill follows the same lesson format used elsewhere in the repository:

1. Goal
2. Syntax
3. What it does
4. Why it matters
5. Mental model
6. Annotated example
7. Common mistakes
8. Exercise
9. Acceptance criteria
10. Verify
11. Reflection

## Verify A Single Drill

```bash
go test -tags exercise ./exercises/00-syntax-drills/<drill>/...
```

## Verify The Whole Track

```bash
make drills
```

## Finish A Drill

1. Remove the `//go:build exercise` line from that drill's `exercise_test.go`.
2. Run `make check`.
3. Record what you learned in `docs/learning-log.md`.
4. Commit one coherent learning milestone.

## Drill Order

1. `01-variables-zero-values`
2. `02-type-inference-conversion`
3. `03-strings-runes`
4. `04-if-switch`
5. `05-for-range`
6. `06-functions-returns-defer`
7. `07-pointers`
8. `08-arrays-slices`
9. `09-maps`
10. `10-structs-methods`
11. `11-composition-embedding`
12. `12-interfaces-assertions`
13. `13-errors`
14. `14-generics`
15. `15-json-basics`
