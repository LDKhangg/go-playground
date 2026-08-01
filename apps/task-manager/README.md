# Task Manager Capstone

This application turns the Go foundations and exercises into one evolving
project. Each milestone extends the same task-management domain instead of
starting an unrelated example.

## Current Milestone: Context-Aware CRUD HTTP API

The implemented API uses the Go standard library, an application service
layer, and an in-memory, mutex-protected task store. It supports collection
reads/writes, item reads/updates/deletes, strict JSON validation, request
context propagation, and graceful shutdown on `SIGINT`/`SIGTERM`.

Run it from the repository root:

```bash
make run-api
```

The server listens on `http://localhost:8080` by default. Set `PORT` to change
the port:

```bash
PORT=9090 make run-api
```

| Method | Route | Purpose |
| --- | --- | --- |
| `GET` | `/health` | Returns the service health status. |
| `GET` | `/tasks` | Lists tasks and optionally simulates slower work with `?delay=...`. |
| `POST` | `/tasks` | Creates a task from a JSON title. |
| `GET` | `/tasks/{id}` | Returns one task or `404` when missing. |
| `PATCH` | `/tasks/{id}` | Updates `title`, `done`, or both. |
| `DELETE` | `/tasks/{id}` | Removes a task and returns `204`. |

```bash
curl http://localhost:8080/health
curl http://localhost:8080/tasks
curl http://localhost:8080/tasks?delay=250ms
curl -X POST http://localhost:8080/tasks \
  -H 'Content-Type: application/json' \
  -d '{"title":"learn Go"}'
curl http://localhost:8080/tasks/1
curl -X PATCH http://localhost:8080/tasks/1 \
  -H 'Content-Type: application/json' \
  -d '{"title":"ship docs","done":true}'
curl -X DELETE http://localhost:8080/tasks/1
```

Validation notes:

- `POST` and `PATCH` reject malformed JSON or multiple JSON values.
- Titles are trimmed and must not be empty.
- `PATCH` must change at least one field.
- Missing tasks return `404` with a JSON error payload.
- `GET /tasks?delay=<duration>` uses the request context and returns `408`
  when the client deadline or cancellation fires before the simulated work
  finishes.

## Why The Delay Query Exists

The `delay` query is not production-oriented business logic. It is a teaching
tool that makes `context.Context` visible inside the existing API shape.

- It lets learners issue `GET /tasks?delay=2s` and compare it with a client
  timeout or cancellation.
- It shows why handlers should pass `r.Context()` into application work instead
  of ignoring request lifetime.
- It keeps the example inside the current CRUD API instead of introducing a
  separate demo endpoint.

## Next Milestones

These tracks intentionally share the same task-management behavior but are not
implemented yet:

- `task`: a command-line application covering commands, flags, input
  validation, and meaningful exit codes.
- `task-worker`: a long-running worker covering bounded background work,
  `context` cancellation, channels, scheduling, and graceful shutdown.
- `api/proto/`: protobuf source contracts using `proto3` and `go_package`,
  followed later by generated Go messages and gRPC transport.

## Verify

```bash
make test-api
go test -race ./apps/task-manager/...
```
