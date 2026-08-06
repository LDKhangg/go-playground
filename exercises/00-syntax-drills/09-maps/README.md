# 09 - Maps

## Goal

Count occurrences with a map and read entries safely with the comma-ok form.

## Concepts

- Map types `map[K]V`
- Assignment, indexing, and reading by key
- Reading a missing key returns the zero value
- The comma-ok form `value, ok := m[key]`
- Nil maps and `make`

## Syntax Primer

A map associates keys with values. Indexing a missing key returns the value type's zero value — which is ambiguous, so the comma-ok form also reports presence:

```go
counts := make(map[string]int)
counts["go"] = counts["go"] + 1

score, ok := scores["mina"] // ok is false when the key is absent
```

Maps must be initialized before writing. A literal works too:

```go
scores := map[string]int{"mina": 10}
```

## Mental Model

A map is a lookup table, not an ordered list. Reading is always safe: a missing key yields the zero value, and comma-ok tells you whether the key was actually there — one idiom that removes a whole class of "key not found" crashes.

## Annotated Examples

```go
func recordVisit(seen map[string]int, name string) {
	seen[name]++ // increments the counter for name
}

if count, ok := seen["admin"]; !ok {
	// handle the absence deliberately
}
```

## Common Diagnostics

- `assignment to entry in nil map`: the map was never initialized; use `make` or a literal first.
- Wrong count for a missing key: reading a missing key returns `0`, not an error — decide whether `0` means "present with zero" or "absent".
- `cannot use map[string]int as map[string]string`: map types are exact; keys and values must match the signature.

## Exercise

Implement `CountWords` to build a word-frequency map, and `LookupScore` to read a score with comma-ok semantics.

## Acceptance Criteria

- `CountWords("go go tests")` counts `"go"` twice and `"tests"` once.
- `LookupScore` returns the score and `true` for a present key.
- `LookupScore` returns `(0, false)` for a missing key.

## Hints

- Split the text with `strings.Fields` and loop over the words.
- Increment with `counts[word] = counts[word] + 1` (or `counts[word]++`).
- In `LookupScore`, read with `score, ok := scores[name]` and return both values.

## Verify

```bash
gofmt -w exercises/00-syntax-drills/09-maps
go test -tags exercise ./exercises/00-syntax-drills/09-maps/...
```

## Reflection Prompts

Why is the comma-ok form clearer than using a magic fallback number? What happens if you write to a map before initializing it?