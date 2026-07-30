# Task Manager Capstone

This application turns the Go foundations and exercises into one evolving
project. Each milestone extends the same task-management domain instead of
starting an unrelated example.

## Current Milestone: HTTP Endpoints

The implemented API uses the Go standard library and an in-memory,
mutex-protected task store.

Run it from the repository root:

```bash
make run-api
```

The server listens on `http://localhost:8080`.

| Method | Route | Purpose |
| --- | --- | --- |
| `GET` | `/health` | Returns the service health status. |
| `GET` | `/tasks` | Lists tasks. |
| `POST` | `/tasks` | Creates a task from a JSON title. |

```bash
curl http://localhost:8080/health
curl http://localhost:8080/tasks
curl -X POST http://localhost:8080/tasks \
  -H 'Content-Type: application/json' \
  -d '{"title":"learn Go"}'
```

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
go test ./apps/task-manager/...
go test -race ./apps/task-manager/...
```
