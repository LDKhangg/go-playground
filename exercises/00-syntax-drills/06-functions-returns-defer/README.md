# 06 - Functions, Returns, And Defer

## Goal

Return useful values from a function and guarantee cleanup on every code path.

## Concepts

- Parameters and multiple return values
- The `(value, error)` convention
- `defer` for guaranteed cleanup
- Resource lifetimes (readers, files, response bodies)

## Syntax Primer

A function declares output types after the parameter list; callers receive all results at once:

```go
func ReadAndClose(rc io.ReadCloser) ([]byte, error)
```

`defer` schedules a call to run when the surrounding function returns — including early returns and error paths:

```go
defer rc.Close() // runs no matter how ReadAndClose exits
```

## Mental Model

`defer` is a promise: "when this function is done, do this." Cleanup written with `defer` cannot be skipped by a later-added branch, which keeps files, HTTP bodies, and sockets from leaking in long-running servers.

## Annotated Examples

```go
func readConfig(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close() // guaranteed cleanup

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}
```

## Common Diagnostics

- Resource never closed: the `Close` call sits only in the success branch, so error paths leak it. Move it to `defer` immediately after the resource is acquired.
- `io.ReadAll` errors ignored: reading can fail; return `nil, err` instead of a partial result.
- `defer` invoked before the resource exists: the `defer` must come after the call that creates the resource.

## Exercise

Implement `ReadAndClose` so it reads all bytes from the reader and always closes it.

## Acceptance Criteria

- All bytes are read and returned.
- The reader is closed after a successful read.
- The reader is closed when reading fails, and the read error is returned.

## Hints

- `defer rc.Close()` immediately after the function starts.
- Use `io.ReadAll` for the read.
- On read error, return `nil, err`.

## Verify

```bash
gofmt -w exercises/00-syntax-drills/06-functions-returns-defer
go test -tags exercise ./exercises/06-functions-returns-defer/...
```

## Reflection Prompts

Why is `defer` safer than manually closing in several branches? In what order do multiple `defer` calls run?