# 03 - Strings And Runes

## Goal

Work with user-facing text without accidentally splitting multi-byte UTF-8 characters.

## Concepts

- Strings as sequences of bytes
- Runes as Unicode code points
- `for range` decoding strings one rune at a time
- `strings.Fields` to split on whitespace
- `strings.Builder` for efficient text assembly

## Syntax Primer

A Go string is a sequence of bytes. One visible character can occupy several bytes in UTF-8, so indexing by position (`word[0]`) gives you a byte, not necessarily a character. Ranging over a string decodes it safely, one rune at a time:

```go
for _, char := range "café" {
	_ = char // char is a rune, not a byte
}
```

`strings.Fields` splits text on any whitespace and discards runs of spaces, so `"  Mina   Park  "` becomes `["Mina" "Park"]`. `strings.Builder` collects text without repeated allocations:

```go
var sb strings.Builder
sb.WriteRune('M')
sb.String() // "M"
```

## Mental Model

A string is bytes; `range` is the safe lens that decodes those bytes into runes. Whenever you mean "character," iterate — and whenever you mean "raw byte index," be explicit about it.

## Annotated Examples

```go
firstLetter := func(word string) rune {
	for _, r := range word {
		return r // the first decoded rune
	}
	return 0
}

firstLetter("Mina") // 'M'
```

## Common Diagnostics

- `invalid operation: cannot index ...`: indexing a string yields a `byte`; if you need the character, range instead.
- Misrendered output with accented or non-ASCII names: byte indexing split a multi-byte rune; use `range`.
- `undefined: strings` or `undefined: sb`: add the `strings` import before calling its functions.

## Exercise

Implement `Initials` so it returns the first letter of every word in a name.

## Acceptance Criteria

- `Initials("Mina Park")` returns `"MP"`.
- Leading and repeated spaces are ignored.
- A one-word name returns one initial.

## Hints

- Split the name with `strings.Fields` first.
- For each word, range over it and write only the first rune into a `strings.Builder`.
- Return the builder's `String()`.

## Verify

```bash
gofmt -w exercises/00-syntax-drills/03-strings-runes
go test -tags exercise ./exercises/00-syntax-drills/03-strings-runes/...
```

## Reflection Prompts

When is byte indexing correct, and when is rune iteration safer? Why does `strings.Fields` make the "extra spaces" test pass for free?
