# 01 - Go Core Fundamentals

The questions in this module come up in almost every Go interview. Practice the
exercises in `exercises/00-syntax-drills/01` through `15` and
`exercises/01-basics` for hands-on versions of these ideas.

## 1.1 Types And Zero Values

**Q: What is a zero value, and why does Go have them?**

Every variable declared without an explicit initializer gets a type-appropriate
default: `0` for numeric types, `""` for strings, `false` for bools, `nil` for
pointers, slices, maps, channels, and functions, and a zero-valued struct for
struct types. Go never leaves memory uninitialized, which removes a whole class
of "uninitialized memory" bugs.

```go
var count int     // 0
var name string   // ""
var err error     // nil
```

Practice: `exercises/00-syntax-drills/01-variables-zero-values`.

## 1.2 Pointers And Values

**Q: When does a pointer matter in Go?**

Go is call-by-value: parameters are copies. A pointer (`*T`) copies the address
instead, so writes through it reach the caller's variable. Use pointers when a
function must mutate caller state, or when a large struct would be expensive to
copy. Modern Go style returns new values rather than mutating through pointers
when that reads better.

```go
func Increment(v *int) { *v++ }

n := 4
Increment(&n) // n == 5
```

Practice: `exercises/00-syntax-drills/07-pointers`,
`exercises/02-functions`.

**Q: Can you take the address of a map element?**

No. `&m["key"]` does not compile, because map elements move when the map grows,
so the address would not stay valid. Range over a copy if you need addresses.

## 1.3 Slices

**Q: What is a slice, really?**

A slice is a three-word descriptor: a pointer to a backing array, a length, and
a capacity. Slicing (`s[1:3]`) creates a new descriptor that may share the same
backing array, so two slices can observe each other's writes.

```go
a := []int{1, 2, 3, 4}
b := a[1:3]      // len 2, cap 3, same backing array
b[0] = 99        // a[1] is now 99 too
```

**Q: How does `append` work?**

`append` writes into the backing array when capacity allows; otherwise it
allocates a new, larger array (typically doubling), copies the elements, and
returns a slice pointing at the new array. That is why you must always use the
result: `s = append(s, v)`.

```go
s := make([]int, 0, 2)
s = append(s, 1, 2) // fits in cap 2, no allocation
s = append(s, 3)    // allocates a new array (cap ~4)
```

**Q: `nil` slice vs empty slice?**

Both have `len == 0`. A `nil` slice has a nil pointer; an empty slice points at
a real (often shared) zero-length array. `json.Marshal` turns a nil slice into
`null` and an empty slice into `[]` — the one behavioral difference that
matters in practice.

Practice: `exercises/00-syntax-drills/08-arrays-slices`,
`exercises/03-collections`.

## 1.4 Maps

**Q: How are maps implemented?**

A Go map is a hash table with buckets. Keys hash into buckets; collisions are
handled by chaining (overflow buckets). Reads and writes are amortized O(1).
When a map grows it allocates new buckets and rehashes — which is why you
cannot take addresses of map values and why concurrent writes can corrupt the
structure.

**Q: Is a map safe for concurrent use?**

No. Concurrent reads and writes can panic or corrupt data. Guard with a mutex
or use `sync.Map` for narrow read-heavy cases (and even then, a mutex is
usually clearer).

```go
m := make(map[string]int)
m["go"] = 1
v, ok := m["go"]   // ok == false for a missing key
delete(m, "go")
```

**Q: Why is iteration order random?**

Go deliberately randomizes map iteration so programs do not depend on an
order the language does not guarantee. Never rely on map order; sort keys when
order matters.

**Q: What happens when you write to a nil map?**

`m["x"] = 1` on a nil map panics (`assignment to entry in nil map`). Reading
from a nil map is safe and returns the zero value. Always `make` the map
before writing.

Practice: `exercises/00-syntax-drills/09-maps`.

## 1.5 Strings

**Q: What is a string in Go?**

An immutable sequence of bytes. UTF-8 text is the common case, but Go does not
enforce it. `len(s)` counts bytes, not characters. Ranging over a string
decodes runes safely; indexing returns a byte.

```go
s := "café"         // 5 bytes
len(s)              // 5, not 4
for _, r := range s { // r is a rune (code point)
	_ = r
}
```

For building strings in a loop, use `strings.Builder` — repeated `+=` creates
new strings and allocates each time.

Practice: `exercises/00-syntax-drills/03-strings-runes`.

## 1.6 Interfaces

**Q: What is an interface value under the hood?**

An interface value is a two-word struct: a pointer to the concrete type and a
pointer to the concrete value (a "fat pointer"). That is why an interface
holding a nil pointer is NOT a nil interface:

```go
var p *int      // nil pointer
var i any = p   // i != nil, because i holds a typed value
```

**Q: How is an interface satisfied?**

Implicitly: a type satisfies an interface when it implements all its methods.
There is no `implements` keyword. This keeps interfaces small and lets
consumers define the behavior they need (see `exercises/05-interfaces-errors`).

**Q: Type assertion vs type switch?**

```go
v, ok := i.(string)   // assertion with comma-ok; no panic
switch x := i.(type) { // type switch
case string:
	_ = x
}
```

Assert without comma-ok and a mismatch panics. Always use the comma-ok form
unless you are sure.

Practice: `exercises/00-syntax-drills/12-interfaces-assertions`.

## 1.7 Structs, Tags, And Comparison

**Q: When can you compare structs with `==`?**

Structs whose fields are all comparable (numbers, strings, bools, pointers,
other comparable structs). Structs containing slices or maps cannot be
compared; use `reflect.DeepEqual` or `slices.Equal` (Go 1.21+) for slices.

**Q: What are struct tags for?**

Tags are strings attached to fields, read via reflection — most famously by
`encoding/json`:

```go
type Task struct {
	Title string `json:"title"`
}
```

Practice: `exercises/00-syntax-drills/10-structs-methods`,
`exercises/00-syntax-drills/15-json-basics`.

## 1.8 Errors

**Q: What makes Go error handling different?**

Errors are ordinary values returned explicitly, not thrown. The convention is
`(result, error)` with `nil` meaning success. Sentinel errors provide stable
identity:

```go
var ErrEmptyTitle = errors.New("empty title")
```

Wrap to add context while preserving identity with `%w`, and detect with
`errors.Is` (wrappers) or `errors.As` (types):

```go
if err != nil {
	return fmt.Errorf("create task %q: %w", title, err)
}
if errors.Is(err, ErrEmptyTitle) { ... }
```

Never compare error strings; messages are for humans, identity is for code.

Practice: `exercises/00-syntax-drills/13-errors`,
`exercises/05-interfaces-errors`.

## 1.9 Memory And GC Basics

**Q: Stack or heap?**

The compiler decides with escape analysis: values that do not outlive the
function stay on the stack; values the compiler cannot prove local (returned,
stored in a heap-allocated container, captured by a closure/goroutine) are
allocated on the heap. Go has a concurrent, non-compacting garbage collector;
latency matters more than throughput in its design.

**Q: How can you reduce allocations?**

Reuse buffers (`bytes.Buffer`), pre-size slices and maps with capacity hints
(`make([]T, 0, n)`), avoid string concatenation in loops, and avoid boxing
(e.g. `interface{}` wrapping).

## 1.10 Defer

**Q: How does `defer` behave?**

`defer` schedules a call for when the surrounding function returns. Arguments
are evaluated when the `defer` statement runs, not when the call executes.
Defers run LIFO (last scheduled, first executed) — so pair acquire/release
operations and stack them in reverse order.

```go
func ReadAndClose(rc io.ReadCloser) ([]byte, error) {
	defer rc.Close() // runs even on the error path
	return io.ReadAll(rc)
}
```

Practice: `exercises/00-syntax-drills/06-functions-returns-defer`.

## Practice Check

1. Explain `append` in one sentence, including why `s = append(...)` is required.
2. Why can't you take the address of a map element?
3. What is the difference between a nil interface and an interface holding a nil pointer?
4. When does `json.Marshal` produce `null` instead of `[]`?
5. Write a one-line explanation of why map iteration order is randomized.
