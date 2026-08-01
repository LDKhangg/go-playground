# 15 - JSON Basics

## Goal

Decode JSON into a Go struct and reject malformed input.

## Syntax

Struct tags, `encoding/json`, and decoder errors.

## What It Does

Parses one task payload.

## Why It Matters

HTTP handlers spend a lot of time translating JSON into typed data.

## Mental Model

JSON decoding is data validation at the boundary between text and code.

## Annotated Example

```go
type Task struct {
	Title string `json:"title"`
}
```

## Common Mistakes

- Forgetting struct tags.
- Ignoring the decoder error.

## Exercise

Implement `DecodeTask`.

## Acceptance Criteria

- Valid JSON decodes successfully.
- Missing or blank title returns an error.
- Malformed JSON returns an error.

## Verify

```bash
go test -tags exercise ./exercises/00-syntax-drills/15-json-basics/...
```

## Reflection

Why should decoding and validation happen before business logic?
