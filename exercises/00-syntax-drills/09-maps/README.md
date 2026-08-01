# 09 - Maps

## Goal

Count values and use the comma-ok form safely.

## Syntax

`map[K]V`, indexing, assignment, and `value, ok := m[key]`.

## What It Does

Builds a word frequency table and looks up named scores.

## Why It Matters

Maps are the standard Go tool for key-based lookup.

## Mental Model

Reading a missing key returns the zero value; comma-ok tells you whether it was present.

## Annotated Example

```go
count := counts[word]
counts[word] = count + 1
score, ok := scores[name]
```

## Common Mistakes

- Assuming a missing key is an error automatically.
- Forgetting to initialize a nil map before writing to it.

## Exercise

Implement `CountWords` and `LookupScore`.

## Acceptance Criteria

- Duplicate words increase the count.
- Missing words have no entry.
- `LookupScore` returns comma-ok results.

## Verify

```bash
go test -tags exercise ./exercises/00-syntax-drills/09-maps/...
```

## Reflection

Why is the comma-ok form clearer than using a magic fallback number?
