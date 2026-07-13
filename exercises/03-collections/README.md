# 03 - Collections

## Goal

Use a map and slice together to remove duplicates without losing order.

## Concepts

Slices, maps, `range`, `append`, zero values, and independent backing storage.

## Exercise

Implement `UniqueWords`. Keep only each word's first occurrence and return a newly allocated result slice.

## Acceptance Criteria

- `go test go maps test` becomes `go test maps` in that order.
- Mutating the returned slice does not change the input slice.

## Hints

Use a map as a set and append unseen words to a result slice. Do not filter by overwriting the input.

## Commands

`go test -tags exercise ./exercises/03-collections/...`

## Reflection Prompts

Why does iterating over the map lose input order? What is the relationship between a slice and its backing array?
