# 05 - Top Interview Questions And Traps

The traps below fail candidates at every seniority level. Go through these the
night before an interview. Each entry is Q + short answer + usually code.

## Fundamentals

1. **Q: What is the zero value?** Every type has a default: `0`, `""`, `false`, `nil`, zero struct. `var x int` is always valid to use.
2. **Q: `:=` vs `var`?** `:=` infers type and requires a new variable; `var` states the type explicitly. `:=` is local scope only.
3. **Q: What is `rune`?** An alias for `int32`, a Unicode code point. `for range` over a string yields runes; indexing yields bytes.
4. **Q: What happens with `&` on a range variable?** In a `for i := range s`, giving `&` of the loop variable (or value copy) points at a fresh copy per iteration since Go 1.22; before 1.22 all iterations shared one variable — the famous loop-variable capture trap.
5. **Q: String concatenation `+=` in a loop?** O(n²) allocations. Use `strings.Builder`.
6. **Q: Difference between array and slice?** Array length is part of its type (`[4]int` vs `[]int`); a slice is a descriptor (ptr, len, cap) over an array.
7. **Q: What happens with `append` when capacity is exceeded?** New larger array is allocated (commonly doubling cap), old elements copied; the original slice still points at the old array.
8. **Q: Why can't you compare `[]int` with `==`?** Slices are not comparable; compare element-by-element or use `slices.Equal` (Go 1.21+).
9. **Q: Map key types?** Any comparable type — strings, ints, structs of comparables, pointers. Not slices/maps/functions.
10. **Q: `%s` vs `%q` vs `%v` vs `%w`?** `%s` raw string, `%q` quoted Go string, `%v` default, `%w` error-wrapping format (only valid for errors, wraps to preserve identity). `%w` is not `%v`.
11. **Q: Errors.Is vs Errors.As?** `Is` target sentinel value (identity); `As` matched a type. Both walk the `Unwrap` chain.
12. **Q: Why should you not compare two `error` values with `==`?** Wrapped errors lose identity to `==`; `errors.Is` handles the chain.
13. **Q: `defer` argument vs call timing?** Arguments are evaluated when defer is scheduled; the function body runs at return. LIFO order.
14. **Q: Named return values + `defer`?** A deferred func can modify named returns — used for log timing or converting panics to errors (careful).
15. **Q: `panic` vs `recover`?** Panic aborts; `recover` in a deferred call can catch it. Not for normal error flow; the Go idiom is error returns.

## Structs & Interfaces

16. **Q: Pointer receiver vs value receiver?** Pointer: mutate; value: copy-safe read. Mixing receivers on the same type is legal but a footgun; the type must be consistent at call sites.
17. **Q: Interface with a value vs pointer?** A type with value receivers satisfies an interface both as value and pointer; pointer receivers only satisfy it as a pointer. If a type's methods are pointer-receiver, `var i SomeIf; i = &v` is needed (cannot use value `v` as once).
18. **Q: Interface zero value?** A nil interface has nil type and nil data. An interface wrapping a nil pointer is NOT nil — it has a type + nil data. "Is `err` nil?" checks can say no when it actually holds a nil `*T` (rare), which is a classic trap.
19. **Q: Empty struct `struct{}`?** Zero bytes; used for signaling (channels `<-chan struct{}`), sets. `map[string]struct{}` is a memory-minimal set.
20. **Q: `any` vs `interface{}`?** `any` is a type alias for `interface{}` — identical. An empty interface holds any value; use it as the smallest possible contract only when the type is genuinely unknown (e.g. `[]any` parameters), not as a design habit.

## Concurrency

21. **Q: Do goroutines run in parallel?** Scheduling of goroutines (M:N) — as many as possible; parallel execution limited by GOMAXPROCS/ P benches. "goroutines are not necessarily parallel" is the precise answer.
22. **Q: Buffered vs unbuffered channel** — sender blocks until receiver ready (synchronized) vs buffered up to cap.
23. **Q: Are channel operations safe?** Channel ops are synchronized (concurrent-safe for send/receive); NOT the same for the data carried. Race on the value is separate.
24. **Q: How to avoid a race in a map?** mutex-protect; never concurrent-write; `-race` catches it.
25. **Q: What happens if you read a closed channel repeatedly?** Returns zero value + ok=false forever; a range just stops. No panic on read; panic only on send to closed/double-close.
26. **Q: What is a "goroutine leak"?** A goroutine waiting on something never completed, e.g., sends on an unbuffered channel nobody will receive. Guard with context / timeouts.
27. **Q: WaitGroup Add inside goroutine bug?** Add must be called before goroutine starts, else Wait may race.
28. **Q: Copying the mutex?** See `go vet` (copylock). Mutexes are not value-safe. Pass pointers; use `sync.Mutex` by reference.
29. **Q: `select` with all channels create?** Random pick among ready cases — not deterministic.
30. **Q: `time.After` in every loop iteration?** It allocates a timer per call; for long loops schedule one ticker and `select` on it instead of a fresh `time.After` each iteration.

## HTTP / Backend

31. **Q: Why read and close the body?** Not reading/closing a request body can stall keep-alive connection reuse and leak resources. Read it (or drain it) and `defer r.Body.Close()`.
32. **Q: Why an explicit `http.Server` instead of the defaults?** `http.ListenAndServe` uses the global `http.DefaultServeMux` with no timeouts. An explicit `http.Server` lets you set timeouts and a scoped mux — safer and testable.
33. **Q: Does `json.NewEncoder(w).Encode(v)` set `Content-Type`?** No — set `w.Header().Set("Content-Type", "application/json")` before writing the status, or clients may guess the type.
34. **Q: Content-Type for JSON on request?** Check `r.Header.Get("Content-Type")` if you care; decode doesn't verify by default.
35. **Q: What is `http.MaxBytesReader`?** Limit body size; protects from huge uploads.
36. **Q: Where does the handler error go?** Return it with status code mapping; log internal error server-side; keep client responses generic for 5xx.
37. **Q: How to serve static files?** `http.FileServer(http.Dir("public"))` behind a mux pattern. Trap: `/` matches everything — mount carefully.
38. **Q: Context of a handler?** `r.Context()` — the request context; it's canceled when client disconnects. Use for DB/call timeouts.
39. **Q: Middleware for CORS?** Set the `Access-Control-Allow-*` headers before calling `next`, and answer `OPTIONS` preflight requests directly. Headers written after `next.ServeHTTP` may be too late.

## Testing

40. **Q: What is the difference between `t.Errorf` and `t.Fatalf`?** Error continues the subtest vs Fatal stops it immediately. For sequential checks Fatal stops flow; Error may let a broken test run on.
41. **Q: Why do a table-driven with `t.Run`?** Named subtests, independent fail/pass, `-run` filtering.
42. **Q: Race detector** — `-race` tool; finds data races; catches — see the test run.
43. **Q: Test a panic?** Recover in the test; Fatal on recovering no panic.
44. **Q: `t.Parallel` caution?** Runs in parallel — shared state races in tests themselves; the loop-capture note applies pre-1.22.
45. **Q: Coverage says 100% but still a bug?** Yes: coverage doesn't assert; untested branches of code logic with bug.

## Final Oral Drill

Answer these out loud in order:

1. What is a slice internally? (header + backing array)
2. When does `append` allocate? (cap exhausted)
3. Why does `defer` LIFO matter? (close order)
4. How do you stop a goroutine? (context)
5. What does the race detector see? (unsynchronized access)
6. How do you structure a service? (constructor dependency injection)
7. Test a handler without a server? (`httptest`)