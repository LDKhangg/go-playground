# Go Career Roadmap

Work from top to bottom. Exercises isolate one idea; the Task Manager capstone
then applies that idea in a program that becomes progressively more realistic.

## Phase 1: Foundations And Exercises

- [ ] 01 Basics: variables, constants, types, conditions, and loops.
- [ ] 02 Functions: parameters, multiple returns, scope, and basic pointers.
- [ ] 03 Collections: arrays, slices, maps, `range`, and safe copies.
- [ ] 04 Structs and methods: custom types, constructors, receivers, and composition.
- [ ] 05 Interfaces and errors: small interfaces, sentinel errors, wrapping, and inspection.
- [ ] 06 Testing: table-driven tests, subtests, coverage, and benchmarks.
- [ ] 07 Concurrency: goroutines, wait groups, mutexes, contexts, and race detection.
- [ ] 08 Object-oriented design: structs plus behavior, composition, dependency injection, and service boundaries.

Run a chapter's opt-in challenge with:

```bash
go test -tags exercise ./exercises/<chapter>/...
```

Starter failures in Chapters 04, 05, 07, and 08 are unmet acceptance criteria, not
syntax errors. `go test ./...` deliberately excludes these unfinished tagged
tests from the baseline.

## Phase 2: Core Go Engineering

- [ ] Organize packages around a clear domain boundary.
- [ ] Use `gofmt`, `go vet`, table-driven tests, and the race detector daily.
- [ ] Propagate `context.Context` through work that can be cancelled.
- [ ] Model validation and errors before adding a transport layer.

Verify with:

```bash
make check
```

## Phase 3: Task Manager HTTP Endpoints

The capstone source is in [`apps/task-manager/`](apps/task-manager/).

- [x] Health endpoint.
- [x] List tasks endpoint.
- [x] Create task endpoint with JSON validation.
- [x] Add complete CRUD operations and domain validation.
- [x] Expand handler coverage with `httptest`.
- [x] Add `PORT` configuration and graceful HTTP shutdown.
- [ ] Propagate request contexts through longer-running application work.

Run and verify the current milestone:

```bash
make run-api
curl http://localhost:8080/health
curl http://localhost:8080/tasks
curl -X PATCH http://localhost:8080/tasks/1 -H 'Content-Type: application/json' -d '{"done":true}'
```

## Phase 4: Go Applications Beyond Endpoints

- [ ] Build `task`, a CLI application for the same Task Manager behavior.
- [ ] Define commands and flags, validate input, write useful stderr messages,
  and return meaningful exit codes.
- [ ] Build `task-worker`, a long-running worker application.
- [ ] Bound worker concurrency, cancel work with `context`, coordinate through
  channels, schedule work, and exit gracefully on shutdown.
- [ ] Test worker cancellation and run the race detector before moving on.

Desktop or mobile UIs are optional paths, not core Go milestones.

## Phase 5: Persistence And Transport Evolution

- [ ] Place persistence behind a small store interface.
- [ ] Add database storage and migrations.
- [ ] Define protobuf contracts in `apps/task-manager/api/proto/`.
- [ ] Use `syntax = "proto3"` and an explicit `option go_package` in each
  `.proto` file.
- [ ] Generate Go messages with `protoc --go_out` and add gRPC only after the
  shared application behavior is tested.
- [ ] Keep HTTP, CLI, worker, and gRPC transports consistent with the same
  domain behavior.

## Phase 6: Delivery

- [ ] Add environment-based configuration and structured logging.
- [ ] Add observability and operational error handling.
- [ ] Package with Docker.
- [ ] Keep CI running format, vet, tests, and the race detector.
- [ ] Deploy the tested application.

## Definition Of Done

A chapter or project milestone is complete when its targeted tests pass,
`make check` passes, the relevant documentation is current, a learning-log
reflection exists when appropriate, and the commit is pushed to GitHub.
