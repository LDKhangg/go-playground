# 01 - Variables And Zero Values

## Goal

Know exactly what a declared variable holds before you assign anything to it.

## Concepts

- `var` declarations with explicit types
- Short declarations with `:=`
- Zero values: `0` for numbers, `""` for strings, `false` for bools
- Struct types with named fields
- Composite literals

## Syntax Primer

A `var` declaration states the type and reserves a name. A short declaration uses `:=` and lets Go infer the type from the right-hand side.

```go
var count int          // zero value: 0
var name string        // zero value: ""
var active bool        // zero value: false
total := 100           // inferred int
```

A struct is a named collection of fields. A struct literal like `Profile{}` builds a value whose fields all hold their zero values.

```go
type Profile struct {
	Name     string
	Active   bool
	Attempts int
}

profile := Profile{} // Name: "", Active: false, Attempts: 0
```

## Mental Model

A declared variable is always a valid box: it has a type and a value from the moment it exists. Go never leaves memory uninitialized, so you can reason about defaults instead of wondering what an unset variable contains.

## Annotated Examples

```go
var retries int       // retries == 0
var label string      // label == ""
enabled := true       // short declaration, type inferred as bool
_ = enabled           // `_` discards a value the compiler would call unused
```

Reading a variable before assignment is safe because the zero value is already there. That is why `DefaultProfile` can simply return `Profile{}` and every field already holds its documented default.

## Common Diagnostics

- `declared and not used`: every local variable must be read somewhere. Use it, remove it, or discard it with `_ =`.
- `mismatched types ...`: assigning a value of one type to a variable of another type is an error; declare the right type or convert explicitly.
- `undefined: Profile`: the type must be declared in the same package before use.

## Exercise

Implement `DefaultProfile` so it returns a profile with zero-valued fields.

## Acceptance Criteria

- `Name` is `""`.
- `Active` is `false`.
- `Attempts` is `0`.

## Hints

- Return a struct literal with no fields filled in.
- Do not hand-assign defaults; the zero value is the point of this exercise.

## Verify

```bash
gofmt -w exercises/00-syntax-drills/01-variables-zero-values
go test -tags exercise ./exercises/00-syntax-drills/01-variables-zero-values/...
```

## Reflection Prompts

Why are zero values useful in API and configuration code? What does the `Profile{}` literal imply about the fields you did not write?
