# 03 - Strings And Runes

## Goal

Handle user-facing text without accidentally splitting bytes.

## Syntax

Strings, runes, slicing, and `for range` over text.

## What It Does

Builds initials from a person's name.

## Why It Matters

User names and text can contain multi-byte UTF-8 characters.

## Mental Model

A string is bytes; `range` decodes runes safely.

## Annotated Example

```go
for _, r := range "Go" {
	_ = r
}
```

## Common Mistakes

- Indexing bytes when you mean characters.
- Assuming every visible character is one byte.

## Exercise

Implement `Initials`.

## Acceptance Criteria

- `"Mina Park"` returns `"MP"`.
- Leading and repeated spaces are ignored.
- A one-word name returns one initial.

## Verify

```bash
go test -tags exercise ./exercises/00-syntax-drills/03-strings-runes/...
```

## Reflection

When is byte indexing correct, and when is rune iteration safer?
