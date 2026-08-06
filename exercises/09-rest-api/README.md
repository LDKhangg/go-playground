# 09 - REST API

## Goal

Build HTTP handlers for a small REST API with the standard library: decode
and validate JSON, talk to a store, and map outcomes to status codes.

## Concepts

- REST verbs and resources (`POST` create, `GET` list)
- `http.HandlerFunc` and method patterns on `http.ServeMux`
- Strict JSON decoding with `DisallowUnknownFields`
- Validating content after decoding
- Mapping errors to status codes
- Testing handlers with `httptest`

## Syntax Primer

A handler takes a request, works through a dependency (here a `Store`), writes
a status code, and encodes JSON:

```go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(task)
```

Decode with the streaming decoder so you can reject unknown fields:

```go
dec := json.NewDecoder(r.Body)
dec.DisallowUnknownFields()
if err := dec.Decode(&input); err != nil {
	http.Error(w, "invalid JSON", http.StatusBadRequest)
	return
}
```

## Mental Model

A handler is a thin translator: request in, domain call, HTTP response out.
Decoding success is not validation — always validate the decoded content
afterward. Status codes are the API's contract: `201` created, `400` client
error, `404` missing, `500` unexpected.

## Annotated Examples

```go
func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s took %s", r.Method, r.URL.Path, time.Since(start))
	})
}
```

Middleware wraps a `Handler` and delegates with `next.ServeHTTP`. The handler
does not need to know it exists.

## Common Diagnostics

- `invalid character ... looking for beginning of value`: the body is not valid
  JSON; your decode returned the error, which the handler must turn into `400`.
- `json: unknown field "extra"`: `DisallowUnknownFields` rejected the payload —
  the desired strict behavior.
- Handler writes the status after encoding, so headers change too late: set
  `Content-Type` and `WriteHeader` before `Encode`.
- `nil pointer dereference`: a nil store was passed; the handler uses the
  injected store, never a global.

## Exercise

Implement `CreateTask` and `ListTasks`.

- `CreateTask`: trim and decode the title, reject blank or unknown or
  malformed input with `400`, ask the store to `Create`, then respond `201`
  with the created task as JSON.
- `ListTasks`: list through the store and return `200` with the collection.

## Acceptance Criteria

- A valid payload returns `201` with an assigned `id` and JSON content type.
- Blank titles, unknown fields, and malformed JSON return `400`.
- No handler writes a response for a missing store method other than the ones above.
- `ListTasks` returns `200` with an array containing every stored task.

## Hints

- Trim the title with `strings.TrimSpace` and compare to `""` after decoding.
- Call `store.Create(Task{Title: title})` and encode the returned task.
- The store assigns the `id`; the handler must not invent JSON keys.

## Verify

```bash
gofmt -w exercises/09-rest-api
go test -tags exercise ./exercises/09-rest-api/...
```

The starter intentionally fails until the handlers are implemented. A failing
assertion is unmet exercise behavior, not a syntax error.

## Reflection Prompts

Why is `DisallowUnknownFields` a security feature? Where should validation
live: in the handler or the store — and why? How would you add a
`GET /tasks/{id}` handler and a `404` mapping?