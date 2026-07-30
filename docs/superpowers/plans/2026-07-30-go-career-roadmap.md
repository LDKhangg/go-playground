# Go Career Roadmap Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Reorganize the existing Task API into a later Task Manager capstone and make the repository's published learning order include Go CLI and worker applications after endpoint development.

**Architecture:** Keep one root Go module (`github.com/LDKhangg/go-playground`) so `go test ./...` continues to cover both exercises and the capstone. Move the current API to `apps/task-manager/`, with its executable at `cmd/api` and private packages below the app's `internal/` directory. Document future CLI and worker entry points without creating their code before their prerequisites are taught.

**Tech Stack:** Go 1.26 standard library, `net/http`, `testing`, `httptest`, GNU Make, GitHub Actions.

## Global Constraints

- Preserve the current HTTP contract: `GET /health`, `GET /tasks`, and `POST /tasks` retain their responses and status codes.
- Keep a single root `go.mod`; do not create a nested module under `apps/task-manager/`.
- Keep `go test ./...`, `make fmt-check`, `make vet`, `make test`, and `make race` functional from repository root.
- Keep all syntax exercises under `exercises/`, English, solution-free, and unchanged by this reorganization.
- Do not add a framework, database, Docker, protobuf/gRPC dependencies, generated code, desktop/mobile GUI dependencies, or deployment configuration.
- Future protobuf contracts belong at `apps/task-manager/api/proto/`, use `syntax = "proto3"`, and set `option go_package`; this plan only documents that contract.
- Commit and push each documentation file independently. The cross-package source relocation is one atomic commit because its executable, packages, imports, and tests must compile together.
- Never stage unrelated untracked workflow documents.

## File Structure

- Create: `apps/task-manager/README.md` — current HTTP milestone, future CLI/worker tracks, and verification commands.
- Move: `main.go` -> `apps/task-manager/cmd/api/main.go` — API executable.
- Move: `internal/tasks/store.go` -> `apps/task-manager/internal/tasks/store.go` — Task domain and concurrency-safe in-memory store.
- Move: `internal/tasks/store_test.go` -> `apps/task-manager/internal/tasks/store_test.go` — store characterization tests.
- Move: `internal/httpapi/handlers.go` -> `apps/task-manager/internal/httpapi/handlers.go` — HTTP transport.
- Move: `internal/httpapi/handlers_test.go` -> `apps/task-manager/internal/httpapi/handlers_test.go` — HTTP contract tests.
- Modify: `Makefile` — point `run` at the API executable and add an explicit `run-api` alias.
- Modify: `README.md` — dashboard that orders exercises, endpoint work, CLI, worker, then delivery.
- Modify: `ROADMAP.md` — milestone checklist and verification commands for all phases.

---

### Task 1: Establish the Relocation Baseline

**Files:**
- Test: `internal/httpapi/handlers_test.go`
- Test: `internal/tasks/store_test.go`

**Interfaces:**
- Consumes: current `tasks.NewStore() *tasks.Store`, `(*tasks.Store).Add(string) (tasks.Task, error)`, `(*tasks.Store).List() []tasks.Task`, `httpapi.HealthHandler`, and `httpapi.TasksHandler(*tasks.Store) http.HandlerFunc`.
- Produces: a recorded baseline for the exact behavior that the relocated packages must retain.

- [ ] **Step 1: Run the existing HTTP contract tests before moving files**

Run:

```bash
go test ./internal/httpapi ./internal/tasks
```

Expected: both packages pass; `GET /health` returns `{"status":"ok"}`, an empty `GET /tasks` returns `[]`, and `POST /tasks` trims a valid title while invalid JSON and an empty title return `400`.

- [ ] **Step 2: Run the full baseline quality suite**

Run:

```bash
make check
go test -tags=exercise ./exercises/01-basics/...
go test -tags=exercise ./exercises/02-functions/...
go test -tags=exercise ./exercises/03-collections/...
go test -tags=exercise ./exercises/06-testing/...
```

Expected: all commands exit successfully. Do not run or change the intentionally unfinished tagged challenge tests in Chapters 04, 05, and 07 as part of this relocation.

### Task 2: Move the Task Manager HTTP Milestone

**Files:**
- Create: `apps/task-manager/cmd/api/main.go`
- Create: `apps/task-manager/internal/tasks/store.go`
- Create: `apps/task-manager/internal/tasks/store_test.go`
- Create: `apps/task-manager/internal/httpapi/handlers.go`
- Create: `apps/task-manager/internal/httpapi/handlers_test.go`
- Delete: `main.go`
- Delete: `internal/tasks/store.go`
- Delete: `internal/tasks/store_test.go`
- Delete: `internal/httpapi/handlers.go`
- Delete: `internal/httpapi/handlers_test.go`

**Interfaces:**
- Consumes: the characterization behavior verified in Task 1.
- Produces: `go run ./apps/task-manager/cmd/api`, `github.com/LDKhangg/go-playground/apps/task-manager/internal/tasks`, and `github.com/LDKhangg/go-playground/apps/task-manager/internal/httpapi`.

- [ ] **Step 1: Move each existing source and test file without changing its package API**

Run:

```bash
mkdir -p apps/task-manager/cmd/api apps/task-manager/internal/httpapi apps/task-manager/internal/tasks
git mv main.go apps/task-manager/cmd/api/main.go
git mv internal/tasks/store.go internal/tasks/store_test.go apps/task-manager/internal/tasks/
git mv internal/httpapi/handlers.go internal/httpapi/handlers_test.go apps/task-manager/internal/httpapi/
```

Keep package names `tasks` and `httpapi`. Keep `Task`, `Store`, `NewStore`, `Add`, `List`, `ErrEmptyTitle`, `HealthHandler`, and `TasksHandler` unchanged.

- [ ] **Step 2: Update imports to the app-local internal packages**

In `apps/task-manager/cmd/api/main.go`, replace the two root imports with:

```go
"github.com/LDKhangg/go-playground/apps/task-manager/internal/httpapi"
"github.com/LDKhangg/go-playground/apps/task-manager/internal/tasks"
```

In `apps/task-manager/internal/httpapi/handlers.go` and its test, use:

```go
"github.com/LDKhangg/go-playground/apps/task-manager/internal/tasks"
```

Do not otherwise change handler or store logic.

- [ ] **Step 3: Verify the moved packages and executable build**

Run:

```bash
go test ./apps/task-manager/...
go build -o /tmp/task-manager-api ./apps/task-manager/cmd/api
go test ./...
```

Expected: all commands exit successfully. The app test output covers `internal/httpapi` and `internal/tasks`; no root `main` package remains.

- [ ] **Step 4: Commit the atomic relocation and push it**

Run:

```bash
git add apps/task-manager/cmd/api/main.go apps/task-manager/internal/httpapi apps/task-manager/internal/tasks main.go internal/httpapi internal/tasks
git commit -m "refactor: move task API into capstone app"
git push origin main
```

Expected: the commit contains only relocated application source and tests. Do not stage untracked workflow documents.

### Task 3: Make the Makefile Run the Relocated API

**Files:**
- Modify: `Makefile:3-6`

**Interfaces:**
- Consumes: the Task 2 executable path `./apps/task-manager/cmd/api`.
- Produces: `make run` and `make run-api` both start the HTTP API.

- [ ] **Step 1: Update the phony target list and commands**

Replace the initial command section with:

```make
.PHONY: run run-api fmt fmt-check test race vet check

run: run-api

run-api:

	go run ./apps/task-manager/cmd/api
```

Leave `fmt`, `fmt-check`, `test`, `race`, `vet`, and `check` unchanged.

- [ ] **Step 2: Verify the target resolves and repository quality commands still work**

Run:

```bash
make -n run
make fmt-check
make vet
make test
make race
```

Expected: dry-run prints `go run ./apps/task-manager/cmd/api`; all four quality commands exit successfully.

- [ ] **Step 3: Commit and push the Makefile change**

Run:

```bash
git add Makefile
git commit -m "build: run the task manager API"
git push origin main
```

### Task 4: Add the Task Manager Milestone Guide

**Files:**
- Create: `apps/task-manager/README.md`

**Interfaces:**
- Consumes: Task 2 app path and Task 3 `make run-api` target.
- Produces: the authoritative guide for the capstone's current HTTP milestone and planned CLI/worker/gRPC stages.

- [ ] **Step 1: Write the guide with only implemented commands marked as available**

Include these sections and content:

```markdown
# Task Manager

## Current Milestone

The HTTP API is the current implemented milestone.

## Run the HTTP API

make run-api

## HTTP Routes

GET /health
GET /tasks
POST /tasks

## Next Application Paths

1. `task` CLI: create, list, complete, and filter tasks.
2. `task-worker`: bounded background work, cancellation, and graceful shutdown.

## Later Transport

Protobuf contracts will live in `api/proto/`; gRPC is not implemented yet.

## Verify

go test ./apps/task-manager/...
go test -race ./apps/task-manager/...
```

Add curl examples for the three current HTTP routes. State explicitly that CLI, worker, persistence, protobuf, gRPC, Docker, and deployment are roadmap items, not available commands.

- [ ] **Step 2: Verify documented commands and routes**

Run:

```bash
go test ./apps/task-manager/...
go test -race ./apps/task-manager/...
```

Expected: both commands exit successfully. Start `make run-api` separately and use the documented curls to confirm the current routes if an interactive terminal is available.

- [ ] **Step 3: Commit and push the guide**

Run:

```bash
git add apps/task-manager/README.md
git commit -m "docs: add task manager guide"
git push origin main
```

### Task 5: Rewrite the Root Dashboard Around Learning Order

**Files:**
- Modify: `README.md`

**Interfaces:**
- Consumes: exercise paths, `apps/task-manager/README.md`, `make run-api`, and `ROADMAP.md`.
- Produces: a root README that separates syntax exercises from endpoint, CLI, and worker capstone phases.

- [ ] **Step 1: Replace the API-first wording and learning-path table**

Make the introductory learning sequence exactly:

```text
Foundations -> exercises -> core Go engineering -> HTTP endpoints -> CLI application -> worker application -> protobuf/gRPC -> delivery
```

Keep the Chapter 01-07 table, but make its project connection column point only to the later relevant milestone rather than asking a beginner to inspect the current API. Keep Chapter 07 explicitly connected to worker pools, cancellation, graceful shutdown, and `go test -race`.

- [ ] **Step 2: Replace `## Task API` with `## Task Manager Capstone`**

Link to `apps/task-manager/README.md`. List only the HTTP API as implemented. List CLI (`task`) and worker (`task-worker`) as later application milestones. State that Go is being taught for services, command-line tools, automation, and concurrent workers, not only REST APIs.

- [ ] **Step 3: Update commands and repository map**

Document `make run-api` as the API command and include focused test commands:

```bash
go test ./apps/task-manager/...
go test -race ./apps/task-manager/...
```

Replace root `internal/` entries in the repository map with `apps/task-manager/`. Keep exercise commands and build-tag explanation intact.

- [ ] **Step 4: Verify all local links resolve**

Run:

```bash
test -f apps/task-manager/README.md
grep -q 'apps/task-manager/README.md' README.md
grep -q 'exercises/07-concurrency' README.md
git diff --check -- README.md
```

Expected: all commands exit successfully and the README contains both the capstone link and Chapter 07 link.

- [ ] **Step 5: Commit and push the dashboard**

Run:

```bash
git add README.md
git commit -m "docs: order the Go learning path"
git push origin main
```

### Task 6: Rewrite the Detailed Roadmap

**Files:**
- Modify: `ROADMAP.md`

**Interfaces:**
- Consumes: the root dashboard's phase order and the Task Manager paths established by Tasks 2-4.
- Produces: a checklist with prerequisites, capabilities, verification commands, and clear implemented versus future states.

- [ ] **Step 1: Replace the two-section roadmap with ordered phases**

Use these phase headings, in this order:

```markdown
## Phase 1: Foundations and Exercises
## Phase 2: Core Go Engineering
## Phase 3: Task Manager HTTP Endpoints
## Phase 4: Go Applications Beyond Endpoints
## Phase 5: Persistence and Transport Evolution
## Phase 6: Delivery
```

Phase 1 lists Chapters 01-07, explicitly naming concurrency primitives and the race detector. Phase 2 lists packages/modules, errors, logging/configuration, graceful shutdown, and testing. Phase 3 marks health/list/create endpoints as complete and CRUD/validation coverage as future. Phase 4 lists `task` CLI and `task-worker` with worker-pool limits, cancellation, scheduling, and graceful exit. Phase 5 lists persistence, then protobuf/gRPC with `.proto`, `proto3`, `go_package`, code generation, and shared application behavior. Phase 6 lists Docker, CI, and deployment.

- [ ] **Step 2: Add phase-specific verification commands**

Include these exact commands in the relevant phases:

```bash
go test -tags=exercise ./exercises/<chapter>/...
go test ./apps/task-manager/...
go test -race ./apps/task-manager/...
make check
```

State that unfinished tagged exercise failures represent unmet learning acceptance criteria, not syntax failures.

- [ ] **Step 3: Update the definition of done**

Require passing focused tests and `make check`, current README/roadmap status, a learning-log reflection for completed learner milestones, one focused commit, and a push to GitHub. Do not mark future CLI, worker, persistence, protobuf/gRPC, or delivery work complete.

- [ ] **Step 4: Verify the roadmap documents all required tracks**

Run:

```bash
grep -E -q 'CLI|task`' ROADMAP.md
grep -q 'task-worker' ROADMAP.md
grep -q 'protobuf' ROADMAP.md
grep -q 'Concurrency' ROADMAP.md
git diff --check -- ROADMAP.md
```

Expected: all commands exit successfully.

- [ ] **Step 5: Commit and push the roadmap**

Run:

```bash
git add ROADMAP.md
git commit -m "docs: map Go career milestones"
git push origin main
```

### Task 7: Run Final Regression Verification

**Files:**
- Verify only: all modified and moved files.

**Interfaces:**
- Consumes: all changes from Tasks 1-6.
- Produces: evidence that the relocation preserves the HTTP contract and every documented root quality command.

- [ ] **Step 1: Run all root and Task Manager checks**

Run:

```bash
make fmt-check
make vet
make test
make race
go test ./apps/task-manager/...
go test -race ./apps/task-manager/...
go test -tags=exercise ./exercises/01-basics/...
go test -tags=exercise ./exercises/02-functions/...
go test -tags=exercise ./exercises/03-collections/...
go test -tags=exercise ./exercises/06-testing/...
```

Expected: every command exits successfully.

- [ ] **Step 2: Verify API behavior from the relocated executable**

Run in one terminal:

```bash
make run-api
```

Run in another terminal:

```bash
curl http://localhost:8080/health
curl http://localhost:8080/tasks
curl -X POST http://localhost:8080/tasks -H 'content-type: application/json' -d '{"title":"learn Go"}'
```

Expected responses are `{"status":"ok"}`, `[{"id":1,"title":"learn go"}]`, and a created task with ID `2` and title `learn Go`; the executable intentionally seeds `learn go` at startup.

- [ ] **Step 3: Verify repository state and remote synchronization**

Run:

```bash
git diff --check
git status --short
git log --oneline -10
git rev-parse HEAD
git ls-remote origin refs/heads/main
```

Expected: no unintended tracked changes, unrelated untracked workflow documents remain unstaged, and local `HEAD` matches the remote `main` SHA.
