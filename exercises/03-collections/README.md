# 03 - Collections

## Goal

Work with arrays and `range`, then use a map and slice together to remove duplicates without losing order.

## Concepts

Arrays, slices, maps, `range`, `append`, zero values, and independent backing storage.

## Exercise

Implement `SumArray` by ranging over its fixed-size array. Implement `UniqueWords` by keeping only each word's first occurrence and returning a newly allocated result slice.

## Acceptance Criteria

- `SumArray([4]int{2, 4, 6, 8})` returns `20`.
- `go test go maps test` becomes `go test maps` in that order.
- Mutating the returned slice does not change the input slice.

## Hints

Use `range` to visit array values. For unique words, use a map as a set and append unseen words to a result slice. Do not filter by overwriting the input.

## Commands

`go test -tags exercise ./exercises/03-collections/...`

## Reflection Prompts

How does an array's length differ from a slice's length? Why does iterating over the map lose input order?
