# 03 - Collections

## Goal

Use arrays, slices, maps, and `range` to process a sequence while preserving the order in which values first appeared.

## Concepts

- Arrays have a fixed length as part of their type.
- Slices describe a variable-length window over an underlying array.
- Maps associate a key with a value and are useful as sets.
- `range` visits elements in a collection.
- `append` grows a slice and may allocate a new backing array.

## Syntax Primer

```go
numbers := [4]int{2, 4, 6, 8}
words := []string{"go", "maps", "go"}
seen := make(map[string]bool)

for index, word := range words {
	_ = index
	seen[word] = true
}
```

`[4]int` and `[]int` are different types: the first always has four elements; the second can grow or shrink. Looking up a missing `bool` key in a map returns `false`, which makes `map[string]bool` convenient for recording whether a word was seen.

## Mental Model

An array is a box with a fixed number of slots. A slice is a small descriptor pointing at a portion of an array. A map is a lookup table, not an ordered list. Use the input slice to decide order, and use the map only to remember membership.

## Annotated Examples

```go
func countNonEmpty(words []string) int {
	count := 0
	for _, word := range words { // `_` discards the index.
		if word != "" {
			count++
		}
	}
	return count
}
```

```go
func copyNames(names []string) []string {
	result := make([]string, 0, len(names)) // New backing storage.
	for _, name := range names {
		result = append(result, name)
	}
	return result
}
```

## Common Diagnostics

- `cannot use []int as [4]int`: a slice and a fixed-size array are distinct types.
- `assignment to entry in nil map`: initialize a map with `make` or a map literal before writing to it.
- Unexpected ordering after ranging over a map: Go intentionally does not guarantee map iteration order.
- A changed input slice after returning a result: the result reused the input backing array instead of allocating its own storage.

## Exercise

Implement `SumArray` by visiting every value in its `[4]int` argument. Implement `UniqueWords` so it returns each word's first occurrence, keeps input order, and returns independent storage rather than filtering in place.

## Acceptance Criteria

- `SumArray([4]int{2, 4, 6, 8})` returns `20`.
- `UniqueWords([]string{"go", "test", "go", "maps", "test"})` returns `[]string{"go", "test", "maps"}`.
- Mutating the returned slice does not mutate the input slice.

## Hints

- Start a total at its zero value, then add each array value.
- Use `seen[word]` to decide whether to append a word.
- Allocate a separate result slice. Do not overwrite `words` while iterating.

## Verify

Run:

```bash
go test -tags exercise ./exercises/03-collections/...
```

## Reflection Prompts

- Why does an array's length belong to its type while a slice length does not?
- Why does a map help detect duplicates but not preserve their original order?
- When might sharing a backing array be useful, and when is it risky?
