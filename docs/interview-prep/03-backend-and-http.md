# 03 - Backend And HTTP

Junior backend interviews in Go assume you can build a small HTTP service with
the standard library. Study `apps/task-manager` in this repository as the
working example: it is a standard-library HTTP API with health checks,
collection/item endpoints, strict JSON validation, context-aware work, and
graceful shutdown.

## 3.1 net/http Basics

**Q: How do you serve HTTP with the standard library?**

`http.Handler` is the core interface — anything with `ServeHTTP(w, r)`.
`http.HandlerFunc` adapts a plain function. `http.ServeMux` routes by pattern
(Go 1.22+ supports methods and wildcards):

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /health", handleHealth)     // method + path pattern
mux.HandleFunc("POST /tasks", handleCreateTask) // 1.22+ syntax

server := &http.Server{
	Addr:    ":8080",
	Handler: mux,
}
err := server.ListenAndServe() // blocks; returns ErrServerClosed on Shutdown
```

**Q: Handler vs HandlerFunc — why the distinction?**

`Handler` is an interface (behavior); `HandlerFunc` is a function type that
implements it. Middleware and muxes operate on `Handler`, and you can adapt any
function with `http.HandlerFunc(fn)`. This one little indirection is what makes
the whole middleware ecosystem work.

**Q: How do you read and write a request?**

```go
body, err := io.ReadAll(r.Body)   // always read and close bodies
defer r.Body.Close()
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(task)
```

Always close request bodies. Respect `http.MaxBytesReader` for body size
limits.

## 3.2 Middleware

**Q: What is middleware?**

A function that wraps a `Handler`, runs logic before/after the inner handler,
then delegates. Middleware composes: logging, auth, recovery, CORS.

```go
func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s took %s", r.Method, r.URL.Path, time.Since(start))
	})
}

server.Handler = withLogging(mux)
```

**Q: How do you carry per-request values (user ID, request ID) through middleware?**

Put them in the request context and read them in handlers:

```go
ctx := context.WithValue(r.Context(), keyUserID, id)
next.ServeHTTP(w, r.WithContext(ctx))

// in the handler:
userID, ok := r.Context().Value(keyUserID).(string)
```

Never store request-scoped state in package-level variables — concurrent
requests will race.

## 3.3 JSON And Validation

**Q: Marshal vs Encode; Unmarshal vs Decode?**

`json.Marshal`/`json.Unmarshal` work on bytes; `json.NewEncoder`/`NewDecoder`
stream from/to an `io.Reader`/`io.Writer`. For handlers, the streaming pair is
idiomatic.

**Q: How do you validate JSON strictly?**

Decode once into a struct, use tags, then validate the content — decoding
success is not validation:

```go
dec := json.NewDecoder(r.Body)
dec.DisallowUnknownFields() // reject fields the struct does not know

var input struct {
	Title string `json:"title"`
}
if err := dec.Decode(&input); err != nil {
	writeError(w, http.StatusBadRequest, "invalid JSON")
	return
}
if strings.TrimSpace(input.Title) == "" {
	writeError(w, http.StatusBadRequest, "title must not be empty")
	return
}
```

**Q: What is the `json:",omitempty"` / `json:"-"` tags?**

`omitempty` skips zero-valued fields on output; `"-"` excludes the field from
JSON entirely.

Practice: `exercises/00-syntax-drills/15-json-basics`.

## 3.4 HTTP Error Handling

**Q: How do you map domain errors to HTTP status codes?**

Keep the handler thin: service returns an error; the handler maps it to a
status. Use sentinel errors or a typed error so mapping stays stable.

```go
if errors.Is(err, tasks.ErrNotFound) {
	writeError(w, http.StatusNotFound, "task not found")
	return
}
if err != nil {
	writeError(w, http.StatusInternalServerError, "internal error")
	return
}
```

**Q: Why return generic messages for 500s?**

Internal details (stack traces, SQL strings) leak information. Log the detail
server-side; return a safe message to the client.

## 3.5 Graceful Shutdown

**Q: What is graceful shutdown, and how do you implement it?**

On a shutdown signal, stop accepting new work, finish in-flight requests, then
exit. Otherwise connections drop mid-request.

```go
server := &http.Server{Addr: ":8080", Handler: mux}

go func() {
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}()

stop := make(chan os.Signal, 1)
signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
<-stop

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
if err := server.Shutdown(ctx); err != nil {
	log.Printf("shutdown: %v", err)
}
```

`ListenAndServe` returns `http.ErrServerClosed` after `Shutdown` — treat that
as normal, not as a failure.

## 3.6 Handler Safety And Timeouts

**Q: What can go wrong when handlers share state?**

Anything mutable at package scope. Two requests can mutate the same map or
slice concurrently → data race or panic. Keep state behind the service layer
with synchronization, or design handlers as stateless.

**Q: Which timeouts should an HTTP server set?**

`ReadTimeout` (reading the request), `WriteTimeout` (writing the response),
`ReadHeaderTimeout` (always — slow-loris protection), and `IdleTimeout` for
keep-alive connections. A handler doing long work should also watch
`r.Context()` and stop when the client disconnects.

**Q: How do you limit request body size?**

`r.Body = http.MaxBytesReader(w, r.Body, 1<<20)` — a 1 MiB cap — before
decoding.

## 3.7 Quick Answers For Common Questions

- **Why standard library and not a framework?** For junior-level services
  `net/http` is sufficient, well-tested, and dependency-free; the stdlib is the
  safe default.
- **GET vs POST?** GET reads, is idempotent and cacheable; POST creates and is
  not idempotent. PUT/PATCH update (PUT replaces, PATCH partially modifies);
  DELETE removes.
- **What is an idempotent request?** Repeating it has the same effect as doing
  it once — retries become safe.
- **How do you rate-limit?** Simplest junior answer: per-IP token bucket or a
  middleware counting requests in a mutex-protected map with timestamps.
- **What is CORS?** A browser-enforced policy deciding which origins may call
  your API; the server answers with `Access-Control-Allow-*` headers.

## Practice Check

1. Write a `GET /health` handler and a `POST /tasks` handler using only `net/http`.
2. Explain why middleware receives and returns `http.Handler`, not a function.
3. How would you reject JSON with unknown fields, and why does it matter?
4. What does `Shutdown` do that simply exiting does not?
5. Which timeouts would you set on a production server, and why?
