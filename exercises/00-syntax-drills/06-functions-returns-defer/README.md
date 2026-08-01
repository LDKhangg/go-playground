# 06 - Functions, Returns, And Defer

## Goal

Return useful values and always clean up a resource.

## Syntax

Parameters, multiple returns, and `defer`.

## What It Does

Reads a resource and closes it exactly once.

## Why It Matters

Files, HTTP bodies, and sockets must be closed even on error paths.

## Mental Model

`defer` schedules cleanup now so you do not forget it later.

## Annotated Example

```go
file, err := os.Open(name)
if err != nil {
	return nil, err
}
defer file.Close()
```

## Common Mistakes

- Closing the resource only in the success branch.
- Forgetting to return the read error.

## Exercise

Implement `ReadAndClose`.

## Acceptance Criteria

- Reads all bytes.
- Closes the reader.
- Returns the read error if reading fails.

## Verify

```bash
go test -tags exercise ./exercises/00-syntax-drills/06-functions-returns-defer/...
```

## Reflection

Why is `defer` safer than manually closing in several branches?
