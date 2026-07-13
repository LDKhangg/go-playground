# Go Playground

> Learn Go by solving focused exercises, growing a real HTTP API, and recording every milestone in Git.

[![Go](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![CI](https://github.com/LDKhangg/go-playground/actions/workflows/ci.yml/badge.svg)](https://github.com/LDKhangg/go-playground/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-2f855a.svg)](LICENSE)

This repository is a beginner learning workspace, not a collection of finished answers. Each chapter gives you a goal, starter code, hints, and opt-in tests. The task API grows alongside the exercises so each concept is used in a real program.

## Start Here

1. Install Go 1.26 or newer.
2. Clone the repository and enter it.
3. Run `make check` to verify the baseline.
4. Read [`exercises/README.md`](exercises/README.md).
5. Begin with [`exercises/01-basics`](exercises/01-basics).
6. Record completed work in [`docs/learning-log.md`](docs/learning-log.md) and commit it.

## Learning Path

| Chapter | Topic | Status | Project connection |
| --- | --- | --- | --- |
| [01](exercises/01-basics) | Basics | Not started | [Read the current task API](#task-api) |
| [02](exercises/02-functions) | Functions | Not started | [Trace the current task API](#task-api) |
| [03](exercises/03-collections) | Collections | Not started | [Inspect the current in-memory store](#task-api) |
| [04](exercises/04-structs-methods) | Structs and methods | Not started | [Extend CRUD and domain validation](ROADMAP.md#task-api-milestones) |
| [05](exercises/05-interfaces-errors) | Interfaces and errors | Not started | [Introduce a persistence boundary](ROADMAP.md#task-api-milestones) |
| [06](exercises/06-testing) | Testing | Not started | [Expand handler coverage](ROADMAP.md#task-api-milestones) |
| [07](exercises/07-concurrency) | Concurrency | Not started | [Add context and graceful shutdown](ROADMAP.md#task-api-milestones) |

See [`ROADMAP.md`](ROADMAP.md) for later API, persistence, observability, and deployment milestones.

## Daily Workflow

```bash
make check
go test -tags exercise ./exercises/01-basics/...
gofmt -w exercises/01-basics
go test -tags exercise ./exercises/01-basics/...
git add exercises/01-basics README.md docs/learning-log.md
git commit -m "learn: complete basics exercise"
git push
```

When a chapter is complete, remove `//go:build exercise` from its challenge test so `go test ./...` includes it from then on. Update the status table and learning log in the same commit.

## Task API

The project includes a standard-library HTTP API backed by an in-memory, concurrency-safe task store.

```text
main.go -> internal/httpapi -> internal/tasks
```

Run it with `make run`, then try:

```bash
curl http://localhost:8080/health
curl http://localhost:8080/tasks
curl -X POST http://localhost:8080/tasks \
  -H 'content-type: application/json' \
  -d '{"title":"learn Go"}'
```

Current routes:

| Method | Path | Purpose |
| --- | --- | --- |
| `GET` | `/health` | Check service health |
| `GET` | `/tasks` | List tasks |
| `POST` | `/tasks` | Create a task |

## Commands

| Command | Purpose |
| --- | --- |
| `make run` | Start the API on port 8080 |
| `make fmt` | Format all Go code |
| `make test` | Run baseline tests |
| `make race` | Run tests with the race detector |
| `make vet` | Run Go static analysis |
| `make check` | Run every quality check |

## Repository Map

```text
exercises/         Guided, solution-free chapters
internal/tasks/    Task domain and in-memory storage
internal/httpapi/  HTTP handlers and JSON responses
docs/              Learning log, design, and plans
.github/workflows/ Continuous integration
```

## Learning Rules

- Read the goal before reading the tests.
- Run a challenge test before changing starter code.
- Prefer the smallest code that makes the test pass.
- Explain failures in your own words in the learning log.
- Make one coherent learning commit at a time.
- Do not copy a solution you cannot explain.

## Documentation

- [`ROADMAP.md`](ROADMAP.md): ordered curriculum and project milestones
- [`exercises/README.md`](exercises/README.md): exercise workflow
- [`CONTRIBUTING.md`](CONTRIBUTING.md): how to add learning work
- [`docs/learning-log.md`](docs/learning-log.md): progress journal

## License

Released under the [MIT License](LICENSE).
