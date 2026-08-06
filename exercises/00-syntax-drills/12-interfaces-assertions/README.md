# 12 - Interfaces And Assertions

## Goal

Accept behavior through a small interface and safely inspect optional capabilities with type assertions.

## Concepts

- Interface types as behavior contracts
- `io.Writer` as a dependency
- Implicit satisfaction
- Type assertions with the comma-ok form
- Errors as values that can carry optional capabilities

## Syntax Primer

An interface lists methods; any type with those methods satisfies it implicitly. Writing through `io.Writer` works for byte buffers, files, and network connections alike:

```go
func WriteGreeting(w io.Writer, name string) error {
	_, err := fmt.Fprintf(w, "hello, %s", name)
	return err
}
```

A type assertion asks whether a value also implements a richer interface:

```go
temp, ok := err.(Temporary) // ok is true when err implements Temporary
```

## Mental Model

An interface stores a concrete value plus the methods it exposes. When you accept an `io.Writer`, you promise to use only the `Write` method — which makes any writer a drop-in. Assertions let you probe the concrete type underneath without losing type safety, thanks to the `ok` result.

## Annotated Examples

```go
var buf bytes.Buffer
fmt.Fprintf(&buf, "hello, %s", "Go") // &buf is an io.Writer (its address matters)
fmt.Println(buf.String())            // "hello, Go"

var err error = temporaryError{}
if temp, found := err.(Temporary); found {
	_ = temp.Temporary() // safe: the assertion succeeded
}
```

## Common Diagnostics

- `cannot use buf (variable of type *bytes.Buffer) as io.Writer (missing Write method)`: the value must have a `Write([]byte) (int, error)` method — usually by passing a pointer to a buffer.
- `invalid type assertion: ... does not implement ...`: the target interface is incompatible with the assertion.
- `panic: interface conversion` (without comma-ok): asserting without checking `ok` panics on a mismatch. Always use the comma-ok form.
- Write errors ignored: `Write` can fail; return the `Fprintf` error.

## Exercise

Implement `WriteGreeting` so it writes a greeting to the supplied writer, and `AsTemporary` so it reports whether an error is temporary.

## Acceptance Criteria

- `WriteGreeting(buf, "Mina")` writes exactly `"hello, Mina"`.
- A failing writer produces the writer's error.
- `AsTemporary` returns a matching temporary error for a temporary error and `(nil, false)` for a plain error.

## Hints

- Use `fmt.Fprintf` and return its error.
- In `AsTemporary`, assert with `err.(Temporary)` and return both results.

## Verify

```bash
gofmt -w exercises/00-syntax-drills/12-interfaces-assertions
go test -tags exercise ./exercises/00-syntax-drills/12-interfaces-assertions/...
```

## Reflection Prompts

Why is a small input interface easier to fake in tests? What could go wrong with a type assertion that ignores `ok`?