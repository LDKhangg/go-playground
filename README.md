# Go Playground

> Build Go fluency from syntax to production-minded applications, one
> task-management capstone at a time.

[![Go](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![CI](https://github.com/LDKhangg/go-playground/actions/workflows/ci.yml/badge.svg)](https://github.com/LDKhangg/go-playground/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-2f855a.svg)](LICENSE)

This is a solution-free learning repository. Start with focused syntax
exercises, then apply those ideas to one Task Manager capstone that grows into
an HTTP service, CLI application, background worker, and later protobuf/gRPC
transport.

## Start Here

1. Install Go 1.26 or newer.
2. Clone the repository and enter it.
3. Run `make check` to verify the baseline.
4. Read [`exercises/README.md`](exercises/README.md).
5. Begin with [`exercises/00-syntax-drills`](exercises/00-syntax-drills).
6. Record completed work in [`docs/learning-log.md`](docs/learning-log.md) and commit it.

## Career Learning Path

1. **Foundations**: learn variables, control flow, functions, collections, and
   tests through short syntax exercises.
2. **Exercises**: use structs, methods, interfaces, errors, and concurrency in
   increasingly realistic starter code.
3. **Core Go engineering**: practice package boundaries, table-driven tests,
   formatting, `go vet`, the race detector, and `context`.
4. **HTTP endpoints**: run the Task Manager's health, list, and create routes.
5. **Go applications**: build a `task` CLI and a `task-worker` long-running
   process for automation and background work.
6. **Transport and persistence**: add storage, then define protobuf contracts
   and generate Go/gRPC code.
7. **Delivery**: add configuration, structured logging, observability, Docker,
   CI, and deployment.

The complete sequence is in [`ROADMAP.md`](ROADMAP.md).

## Foundations And Exercises

| Chapter | Topic | Project connection |
| --- | --- | --- |
| [00](exercises/00-syntax-drills) | Syntax drills | Practice one Go idea at a time before larger exercises. |
| [01](exercises/01-basics) | Basics | Read values returned by the Task Manager. |
| [02](exercises/02-functions) | Functions | Trace handler and store function calls. |
| [03](exercises/03-collections) | Collections | Inspect the in-memory task list. |
| [04](exercises/04-structs-methods) | Structs and methods | Model task and project domain behavior. |
| [05](exercises/05-interfaces-errors) | Interfaces and errors | Introduce a persistence boundary. |
| [06](exercises/06-testing) | Testing | Expand handler and worker coverage. |
| [07](exercises/07-concurrency) | Concurrency | Build worker pools, cancellation, graceful shutdown, and race-safe code. |

Each chapter README contains **Goal**, **Concepts**, **Syntax Primer**, **Mental
Model**, **Annotated Examples**, **Common Diagnostics**, **Exercise**,
**Acceptance Criteria**, **Hints**, **Verify**, and **Reflection Prompts**. The
examples explain syntax without including completed exercise solutions.

`go test ./...` is the baseline check: it deliberately excludes unfinished
challenge tests. While working through a chapter, run:

```bash
go test -tags exercise ./exercises/<chapter>/...
```

Read the [exercise guide](exercises/README.md) for why this build tag exists
and how to interpret its diagnostics.

## Task Manager Capstone

The current Task Manager milestone is a standard-library HTTP API backed by an
in-memory, concurrency-safe store. Its source and guide are in
[`apps/task-manager/`](apps/task-manager/).

```text
apps/task-manager/cmd/api -> internal/httpapi -> internal/tasks
```

Run the current API:

```bash
make run-api
```

```bash
curl http://localhost:8080/health
curl http://localhost:8080/tasks
curl -X POST http://localhost:8080/tasks \
  -H 'Content-Type: application/json' \
  -d '{"title":"learn Go"}'
```

Go is not only for HTTP services. The later capstone milestones use the same
domain to teach CLI programs, automation, and long-running workers before
moving to protobuf/gRPC and delivery work.

## Commands

| Command | Purpose |
| --- | --- |
| `make drills` | Run the opt-in syntax drill challenge tests. |
| `make run-api` | Start the Task Manager API on port 8080. |
| `make run` | Alias for `make run-api`. |
| `make fmt` | Format all Go code. |
| `make test` | Run baseline tests. |
| `make race` | Run tests with the race detector. |
| `make vet` | Run Go static analysis. |
| `make check` | Run every quality check. |

## Repository Map

```text
exercises/                 Guided, solution-free drills and chapters
apps/task-manager/         Evolving capstone application
apps/task-manager/cmd/api/ HTTP executable
apps/task-manager/internal/ HTTP and task domain packages
docs/                      Learning log, design, and plans
.github/workflows/         Continuous integration
```

## Learning Rules

- Read the goal before reading the tests.
- For Chapters 01-05 and 07, run the challenge test before changing starter code.
- For Chapter 06, author tests in the tagged starter and check every documented acceptance criterion.
- Prefer the smallest code that makes the test pass.
- Explain failures in your own words in the learning log.
- Make one coherent learning commit at a time.
- Do not copy a solution you cannot explain.

## Documentation

- [`ROADMAP.md`](ROADMAP.md): ordered curriculum and capstone milestones
- [`apps/task-manager/README.md`](apps/task-manager/README.md): current application guide
- [`exercises/README.md`](exercises/README.md): exercise workflow
- [`CONTRIBUTING.md`](CONTRIBUTING.md): how to add learning work
- [`docs/learning-log.md`](docs/learning-log.md): progress journal

## License

Released under the [MIT License](LICENSE).
