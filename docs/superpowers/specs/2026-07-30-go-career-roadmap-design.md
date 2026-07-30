# Go Career Roadmap Design

## Goal

Turn this repository into one progressive, career-oriented Go learning path:

```text
syntax -> exercises -> core engineering -> endpoints -> Go applications -> deployment
```

The repository will continue to teach syntax through the existing exercises,
then apply that knowledge to one evolving Task Manager application. It will
not introduce a separate unrelated backend sample.

## Current State

The repository already has seven syntax exercise chapters and a small HTTP
Task API at the repository root. Keeping that API at root makes it appear to
be required before a learner has reached the relevant material. The API is
valuable, so this work moves it into a later Task Manager milestone rather
than deleting it.

## Learning Path

### Phase 1: Go Foundations

Keep the existing chapters as the entry point, ordered as follows:

1. Basics: declarations, types, control flow, loops, and conversions.
2. Functions: parameters, returns, errors, and pointers.
3. Collections: arrays, slices, maps, and iteration.
4. Structs and methods: data modeling, receivers, and composition instead of
   inheritance.
5. Interfaces and errors: behavior-based design, error values, and wrapping.
6. Testing: table-driven tests, subtests, and failure-driven development.
7. Concurrency: goroutines, channels, `context.Context`, synchronization, and
   the race detector.

The root dashboard and roadmap must make Chapter 07 visible as a required
concurrency phase, not an optional side topic.

### Phase 2: Core Go Engineering

Add roadmap milestones before the application introduces framework-specific
concepts:

1. Project layout, packages, modules, and dependency boundaries.
2. Error handling, contextual error wrapping, and error classification.
3. Structured logging, configuration through environment variables, and
   graceful shutdown.
4. Testing workflow: unit tests, HTTP handler tests, race detection, and
   coverage as a signal rather than a target.

These milestones will explain why each practice exists and link back to the
syntax chapters that introduced its language constructs.

### Phase 3: Task Manager Capstone

Move the existing Task API out of repository root to `apps/task-manager/` and
grow it through explicit milestones. A learner builds one recognizable
application while each stage adds only one new concern.

1. HTTP: expose the existing health and task endpoints using `net/http`.
2. Concurrency: add bounded background work, cancellation, and graceful
   shutdown; verify with `go test -race`.
3. Operations: add structured logs and environment-based configuration.
4. Persistence: replace the in-memory store through a repository interface
   while preserving handler behavior.
5. Protobuf and gRPC: define versioned messages and service methods in a
   `.proto` contract, generate Go code, and expose a gRPC transport alongside
   the HTTP API.
6. Delivery: containerize the app, run quality checks in CI, and document a
   local deployment path.

The first implementation plan covers the roadmap and directory reorganization
only. It must not prematurely implement all of these application milestones.

### Phase 4: Go Applications Beyond Endpoints

After learners can build and test HTTP endpoints, the roadmap must explicitly
show two non-API application paths where Go is a strong fit.

1. CLI and terminal application: build a `task` command that creates, lists,
   completes, and filters tasks. It starts with standard input/output and
   command-line flags, then may add a terminal UI after the CLI workflow is
   stable. This teaches executable packaging, input validation, exit codes,
   user-facing errors, and reuse of application logic without HTTP.
2. Long-running worker application: build a `task-worker` process that claims
   bounded background work, observes `context.Context` cancellation, reports
   structured progress, and exits gracefully. This teaches goroutines,
   channels, worker-pool limits, scheduling, and operational behavior without
   making the learner write another CRUD API.

Both programs share Task Manager domain and application behavior rather than
duplicating task rules or storage access. The CLI may call the application
layer directly at first; a later optional integration milestone may make it an
HTTP or gRPC client.

Desktop and mobile GUI development are valid Go-adjacent paths but are not
core curriculum milestones: they require a separate UI toolkit and would
dilute the backend, tooling, and systems strengths this repository teaches.

## Repository Layout

The target top-level layout is:

```text
.
├── exercises/                 # Chapters 01-07 and syntax-focused exercises
├── apps/
│   └── task-manager/          # Evolving capstone application
│       ├── cmd/               # api, task CLI, and worker entry points as introduced
│       ├── internal/          # Application-only packages
│       ├── api/               # HTTP and later protobuf/gRPC contracts
│       └── README.md          # Milestone guide and prerequisites
├── README.md                  # Entry dashboard and learning order
├── ROADMAP.md                 # Detailed phases and completion criteria
├── Makefile                   # Repository-wide quality commands
└── .github/workflows/ci.yml   # Formatting, vet, test, and race checks
```

The exact subdirectories under `apps/task-manager/` are created only when a
milestone needs them. No empty framework, database, or generated-code
directories are added in advance.

## Documentation Contract

`README.md` becomes the first page for learners:

- state the recommended learning order;
- distinguish exercises from the capstone application;
- identify what can be skipped until a later phase;
- link to each chapter and capstone milestone.

`ROADMAP.md` becomes the detailed checklist:

- list all phases in prerequisite order;
- explain the capability gained at each milestone;
- include verification commands and a definition of done;
- identify Chapters 01-07, including concurrency, as prerequisites for the
  matching capstone work.

`apps/task-manager/README.md` becomes the application guide:

- state the current milestone and commands that work today;
- list completed and upcoming milestones without implying unfinished features
  are implemented;
- document the API, CLI, and worker programs separately as each is introduced;
- document each transport separately when HTTP and gRPC coexist;
- provide tool prerequisites only at the protobuf milestone.

Exercise lessons remain English, contain explanations and hints, and do not
add completed solutions for unfinished starter APIs.

## Protobuf and gRPC Contract

Protobuf is deliberately introduced after the learner understands the HTTP
version and package boundaries.

The future protobuf milestone must:

- place source contracts under `apps/task-manager/api/proto/`;
- declare `syntax = "proto3"` and an explicit `option go_package` in every
  contract;
- generate Go messages with `protoc --go_out` using `protoc-gen-go`;
- generate service stubs separately with the gRPC Go plugin when a service is
  added;
- commit source `.proto` files and generated Go code together;
- retain HTTP behavior while the gRPC transport is introduced, so transports
  share application behavior rather than duplicate it.

No protobuf compiler, plugin, or generated code is added by the first
reorganization implementation.

## Quality and Verification

Every completed milestone must state its verification command. Repository-wide
verification remains:

```bash
make fmt-check
make vet
make test
make race
```

The implementation plan must preserve these commands after moving the current
Task API. It must add focused commands for the Task Manager and preserve the
existing exercise workflow:

```bash
go test ./...
go test -tags=exercise ./exercises/<chapter>/...
```

The tagged exercise tests that correspond to intentionally unfinished starter
code remain expected to fail until a learner implements the exercise. Such a
failure must be explained as unmet acceptance behavior, not a Go syntax error.

## Scope Boundaries

This design does not:

- add a web frontend or mobile client;
- replace the standard library HTTP implementation with a framework now;
- add a database, Docker image, gRPC dependency, protobuf generated code, or
  deployment configuration now;
- add desktop or mobile GUI dependencies;
- provide exercise answer files;
- change the public behavior of the current Task HTTP API during its move.

Those items are sequenced as later capstone milestones after their prerequisite
material is documented.
