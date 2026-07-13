# Go Learning Playground Implementation Plan

> **Amendment:** This historical plan predates the final Chapter 06 workflow decision. Chapter 06 is the sole exception to universal prewritten challenge tests: the learner authors the test suite, so there is no prewritten failing test or meta-test, and its acceptance criteria are checked manually. The current authoritative workflow is in the [root README](../../../README.md), [exercise guide](../../../exercises/README.md), and [Chapter 06 README](../../../exercises/06-testing/README.md). Historical task steps below are retained unchanged.

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build and publish a polished beginner Go curriculum with test-driven exercises and an evolving task API in one public GitHub repository.

**Architecture:** Keep `main.go` as the composition root, `internal/tasks` as the domain/store package, and `internal/httpapi` as the transport package. Add isolated exercise packages whose challenge tests are opt-in through the `exercise` build tag until completed, plus repository-level documentation and CI that keep `main` green.

**Tech Stack:** Go 1.26 standard library, `net/http`, Go `testing`/`httptest`, Make, GitHub Actions, GitHub CLI.

## Global Constraints

- All learning material is written in English.
- The target learner is new to Go.
- The module path is `github.com/LDKhangg/go-playground`.
- Exercise starter code and executable challenge tests are included without implementation answers.
- Challenge tests use the `exercise` build tag until the learner completes a chapter.
- The initial setup adds no database, web framework, deployment platform, or unnecessary third-party dependency.
- The public GitHub repository is `github.com/LDKhangg/go-playground` on branch `main` with an MIT license.
- The root README is polished and readable without excessive badges or decoration.

---

### Task 1: Protect and Normalize the Existing Task API

**Files:**
- Modify: `go.mod:1`
- Modify: `main.go:7-8`
- Modify: `internal/httpapi/handlers.go:8`
- Modify: `internal/tasks/store.go:22-24`
- Create: `internal/httpapi/handlers_test.go`

**Interfaces:**
- Consumes: `tasks.NewStore() *tasks.Store`, `HealthHandler(http.ResponseWriter, *http.Request)`, and `TasksHandler(*tasks.Store) http.HandlerFunc`.
- Produces: the canonical module path and baseline HTTP behavior where an empty task list serializes as `[]`, not `null`.

- [ ] **Step 1: Update the module path and imports**

Change `go.mod` to:

```go
module github.com/LDKhangg/go-playground

go 1.26
```

Replace every `example.com/go-playground/...` import in `main.go` and `internal/httpapi/handlers.go` with `github.com/LDKhangg/go-playground/...`.

- [ ] **Step 2: Write HTTP characterization tests**

Create `internal/httpapi/handlers_test.go`:

```go
package httpapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/LDKhangg/go-playground/internal/tasks"
)

func TestHealthHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/health", nil)

	HealthHandler(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}
	if got := recorder.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("expected JSON content type, got %q", got)
	}
	if got := recorder.Body.String(); got != "{\"status\":\"ok\"}\n" {
		t.Fatalf("expected health JSON, got %q", got)
	}
}

func TestTasksHandlerListsEmptyStoreAsArray(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/tasks", nil)

	TasksHandler(tasks.NewStore())(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}
	if got := recorder.Body.String(); got != "[]\n" {
		t.Fatalf("expected empty JSON array, got %q", got)
	}
}

func TestTasksHandlerCreatesTask(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"title":"  learn handlers  "}`))

	TasksHandler(tasks.NewStore())(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, recorder.Code)
	}
	if got := recorder.Body.String(); got != "{\"id\":1,\"title\":\"learn handlers\"}\n" {
		t.Fatalf("expected created task JSON, got %q", got)
	}
}

func TestTasksHandlerRejectsInvalidRequests(t *testing.T) {
	tests := []struct {
		name string
		body string
		want string
	}{
		{name: "invalid JSON", body: `{`, want: "{\"error\":\"invalid json\"}\n"},
		{name: "empty title", body: `{"title":"   "}`, want: "{\"error\":\"title must not be empty\"}\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(tt.body))

			TasksHandler(tasks.NewStore())(recorder, request)

			if recorder.Code != http.StatusBadRequest {
				t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
			}
			if got := recorder.Body.String(); got != tt.want {
				t.Fatalf("expected body %q, got %q", tt.want, got)
			}
		})
	}
}

func TestTasksHandlerRejectsUnsupportedMethod(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "/tasks", nil)

	TasksHandler(tasks.NewStore())(recorder, request)

	if recorder.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, recorder.Code)
	}
	if got := recorder.Header().Get("Allow"); got != "GET, POST" {
		t.Fatalf("expected Allow header %q, got %q", "GET, POST", got)
	}
}
```

- [ ] **Step 3: Run the new tests and observe the empty-list failure**

Run: `go test ./internal/httpapi`

Expected: FAIL in `TestTasksHandlerListsEmptyStoreAsArray` because the current nil slice is encoded as `null`.

- [ ] **Step 4: Initialize the store with an empty non-nil slice**

Change `NewStore` in `internal/tasks/store.go` to:

```go
func NewStore() *Store {
	return &Store{
		nextID: 1,
		tasks:  make([]Task, 0),
	}
}
```

- [ ] **Step 5: Verify the baseline application**

Run: `gofmt -w main.go internal/tasks/*.go internal/httpapi/*.go && go test ./... && go test -race ./...`

Expected: all packages pass; no race is reported.

- [ ] **Step 6: Commit the baseline**

```bash
git add go.mod main.go internal
git commit -m "test: establish task API baseline"
```

---

### Task 2: Add Repository Tooling and Continuous Integration

**Files:**
- Create: `.gitignore`
- Create: `LICENSE`
- Create: `Makefile`
- Create: `.github/workflows/ci.yml`

**Interfaces:**
- Consumes: the Go module and tests from Task 1.
- Produces: `make run`, `make fmt`, `make test`, `make race`, `make vet`, and `make check`; CI runs the same quality gates.

- [ ] **Step 1: Add repository ignores**

Create `.gitignore`:

```gitignore
# Build output
/go-playground
/bin/
*.test
*.out

# Local configuration
.env
.env.*
!.env.example

# Editors and operating systems
.idea/
.vscode/
.DS_Store

# Serena local state
.serena/
```

- [ ] **Step 2: Add the MIT license**

Create `LICENSE`:

```text
MIT License

Copyright (c) 2026 LDKhangg

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

- [ ] **Step 3: Add stable development commands**

Create `Makefile`:

```makefile
GO_FILES := $(shell find . -type f -name '*.go' -not -path './.git/*')

.PHONY: run fmt fmt-check test race vet check

run:
	go run .

fmt:
	gofmt -w $(GO_FILES)

fmt-check:
	@unformatted="$$(gofmt -l $(GO_FILES))"; \
	if [ -n "$$unformatted" ]; then \
		printf '%s\n' "$$unformatted"; \
		exit 1; \
	fi

test:
	go test ./...

race:
	go test -race ./...

vet:
	go vet ./...

check: fmt-check vet test race
```

- [ ] **Step 4: Verify Make targets locally**

Run: `make check`

Expected: formatting check, vet, normal tests, and race tests all pass.

- [ ] **Step 5: Add GitHub Actions CI**

Create `.github/workflows/ci.yml`:

```yaml
name: CI

on:
  push:
    branches: [main]
  pull_request:

permissions:
  contents: read

jobs:
  test:
    name: Go quality checks
    runs-on: ubuntu-latest
    steps:
      - name: Check out repository
        uses: actions/checkout@v6

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version-file: go.mod
          cache-dependency-path: go.mod

      - name: Check formatting
        run: make fmt-check

      - name: Vet
        run: make vet

      - name: Test
        run: make test

      - name: Test for races
        run: make race
```

- [ ] **Step 6: Commit tooling**

```bash
git add .gitignore LICENSE Makefile .github/workflows/ci.yml
git commit -m "chore: add Go quality tooling"
```

---

### Task 3: Build the Learning Documentation and GitHub Landing Page

**Files:**
- Replace: `README.md`
- Create: `ROADMAP.md`
- Create: `CONTRIBUTING.md`
- Create: `docs/learning-log.md`
- Create: `exercises/README.md`

**Interfaces:**
- Consumes: commands from Task 2 and the chapter paths defined in the design.
- Produces: one consistent beginner workflow linking the README, roadmap, exercise guide, contribution guide, and learning log.

- [ ] **Step 1: Replace the root README with a polished study dashboard**

Write `README.md` with these exact sections and facts:

```markdown
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
| 01 | Basics | Not started | Read task fields and HTTP status values |
| 02 | Functions | Not started | Trace constructors and handlers |
| 03 | Collections | Not started | Understand task slices and copies |
| 04 | Structs and methods | Not started | Extend the task domain |
| 05 | Interfaces and errors | Not started | Add validation and storage boundaries |
| 06 | Testing | Not started | Expand API characterization tests |
| 07 | Concurrency | Not started | Understand mutexes and cancellation |

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
```

- [ ] **Step 2: Add the ordered roadmap**

Create `ROADMAP.md`:

```markdown
# Learning Roadmap

Work from top to bottom. Exercises teach one concept in isolation; project milestones apply it to the task API.

## Core Curriculum

- [ ] 01 Basics: variables, constants, types, conditions, and loops
- [ ] 02 Functions: parameters, multiple returns, scope, and basic pointers
- [ ] 03 Collections: arrays, slices, maps, range, and safe copies
- [ ] 04 Structs and methods: custom types, constructors, receivers, and composition
- [ ] 05 Interfaces and errors: small interfaces, sentinel errors, wrapping, and inspection
- [ ] 06 Testing: table-driven tests, subtests, coverage, and benchmarks
- [ ] 07 Concurrency: goroutines, wait groups, mutexes, contexts, and race detection

## Task API Milestones

- [x] Understand the baseline health, list, and create routes
- [ ] Add complete CRUD operations and domain validation
- [ ] Expand handler coverage with `httptest`
- [ ] Add configuration, context propagation, and graceful shutdown
- [ ] Place persistence behind a small store interface
- [ ] Add database storage and migrations
- [ ] Add middleware, structured logging, and observability
- [ ] Package with Docker and deploy the tested service

## Definition of Done

A chapter or project milestone is complete when its tests pass, `make check` passes, the README progress table is current, a learning-log reflection exists, one coherent commit records the work, and the commit is pushed to GitHub.
```

- [ ] **Step 3: Document the exercise lifecycle**

Create `exercises/README.md` explaining this exact sequence:

```markdown
# Exercise Guide

Each chapter contains starter code, a lesson, and challenge tests. It does not contain a solution.

## Start a Chapter

1. Read the chapter README.
2. Run `go test -tags exercise ./exercises/<chapter>/...` and read the failure.
3. Implement one acceptance criterion at a time.
4. Format and rerun the tagged test.
5. Run `make check` to protect previously completed work.

## Finish a Chapter

1. Remove the `//go:build exercise` line and the blank line below it from the chapter test.
2. Run `make check`; the challenge test now belongs to the baseline suite.
3. Mark the chapter complete in the root README.
4. Add a reflection to `docs/learning-log.md`.
5. Commit and push the milestone.

## Chapters

1. [Basics](01-basics)
2. [Functions](02-functions)
3. [Collections](03-collections)
4. [Structs and methods](04-structs-methods)
5. [Interfaces and errors](05-interfaces-errors)
6. [Testing](06-testing)
7. [Concurrency](07-concurrency)
```

- [ ] **Step 4: Add contribution and reflection guides**

Create `CONTRIBUTING.md`:

```markdown
# Contributing Learning Work

## Prerequisites

- Go 1.26 or newer
- Git
- Make

## Workflow

1. Pull the latest `main` branch.
2. Create a focused branch such as `learn/collections` or `feat/task-delete`.
3. Run the relevant test before editing.
4. Make the smallest change that satisfies one requirement.
5. Run `make check`.
6. Update the root progress table and learning log when finishing a milestone.
7. Commit and push the branch.

## Completing a Chapter

Remove the `exercise` build constraint from a solved chapter's test, run `make check`, mark the chapter complete in `README.md`, and add a reflection to `docs/learning-log.md`.

## Commit Examples

```text
learn: complete slices exercise
test: cover invalid task payloads
feat: add task deletion
docs: record concurrency notes
```

Keep each commit limited to one exercise, one behavior change, or one documentation milestone.
```

Create `docs/learning-log.md`:

```markdown
# Learning Log

Use this journal to explain what changed in your own words. Add one entry for every completed chapter or project milestone.

## Entry Template

### YYYY-MM-DD - Chapter or milestone

**Concepts learned:**

**What I built:**

**Problem and solution:**

**Commit:**

**Question to revisit:**

## Entries

### 2026-07-13 - Repository baseline

**Concepts learned:** Go package boundaries, an in-memory store protected by a mutex, HTTP handlers, and baseline tests.

**What I built:** A structured learning repository around the existing task API, with repeatable quality commands and continuous integration.

**Problem and solution:** The original project had working code but no ordered curriculum or Git history. The repository now separates exercises from the evolving application and documents one workflow for both.

**Commit:** Add the relevant commit hash after the repository setup is published.

**Question to revisit:** How should persistent storage replace the in-memory store without coupling it to HTTP handlers?
```

- [ ] **Step 5: Check documentation consistency**

Run: `grep -RE "example\\.com/go-playground|[T]ODO|[T]BD" README.md ROADMAP.md CONTRIBUTING.md exercises/README.md docs/learning-log.md`

Expected: no output.

Run: `make check`

Expected: all checks pass.

- [ ] **Step 6: Commit documentation**

```bash
git add README.md ROADMAP.md CONTRIBUTING.md exercises/README.md docs/learning-log.md
git commit -m "docs: add beginner Go learning path"
```

---

### Task 4: Add Chapters 01-03

**Files:**
- Create: `exercises/01-basics/README.md`
- Create: `exercises/01-basics/exercise.go`
- Create: `exercises/01-basics/exercise_test.go`
- Create: `exercises/02-functions/README.md`
- Create: `exercises/02-functions/exercise.go`
- Create: `exercises/02-functions/exercise_test.go`
- Create: `exercises/03-collections/README.md`
- Create: `exercises/03-collections/exercise.go`
- Create: `exercises/03-collections/exercise_test.go`

**Interfaces:**
- Produces: `TicketPrice(age int) int`, `Divide(dividend, divisor int) (quotient, remainder int, err error)`, `ErrDivideByZero`, and `UniqueWords(words []string) []string`.

- [ ] **Step 1: Create Chapter 01 starter and challenge test**

Create `exercises/01-basics/exercise.go`:

```go
package basics

func TicketPrice(age int) int {
	return 0
}
```

Create `exercises/01-basics/exercise_test.go`:

```go
//go:build exercise

package basics

import "testing"

func TestTicketPrice(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want int
	}{
		{name: "child", age: 8, want: 5},
		{name: "adult", age: 30, want: 12},
		{name: "senior", age: 70, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TicketPrice(tt.age); got != tt.want {
				t.Fatalf("TicketPrice(%d) = %d, want %d", tt.age, got, tt.want)
			}
		})
	}
}
```

Create `exercises/01-basics/README.md`:

```markdown
# 01 - Basics

## Goal

Implement age-based ticket pricing with Go's basic types and control flow.

## Concepts

Variables, constants, integers, comparisons, `if`, and `switch`.

## Exercise

Implement `TicketPrice`. Ages below 13 cost 5, ages from 13 through 64 cost 12, and ages 65 or above cost 7.

## Acceptance Criteria

- `TicketPrice(8)` returns `5`.
- `TicketPrice(30)` returns `12`.
- `TicketPrice(70)` returns `7`.

## Hints

Name the three prices as constants. Check boundaries in an order that makes every age belong to one group.

## Commands

`go test -tags exercise ./exercises/01-basics/...`

## Reflection Prompts

Why does the order of age checks matter? When would a `switch` read more clearly than chained `if` statements?
```

- [ ] **Step 2: Verify Chapter 01 is opt-in and failing**

Run: `go test ./exercises/01-basics/...`

Expected: package builds successfully with no active test files.

Run: `go test -tags exercise ./exercises/01-basics/...`

Expected: FAIL for the child case because the starter returns `0`.

- [ ] **Step 3: Create Chapter 02 starter and challenge test**

Create `exercises/02-functions/exercise.go`:

```go
package functions

import "errors"

var ErrDivideByZero = errors.New("cannot divide by zero")

func Divide(dividend, divisor int) (quotient, remainder int, err error) {
	return 0, 0, nil
}
```

Create `exercises/02-functions/exercise_test.go`:

```go
//go:build exercise

package functions

import (
	"errors"
	"testing"
)

func TestDivide(t *testing.T) {
	quotient, remainder, err := Divide(17, 5)
	if err != nil {
		t.Fatalf("Divide returned error: %v", err)
	}
	if quotient != 3 || remainder != 2 {
		t.Fatalf("Divide(17, 5) = (%d, %d), want (3, 2)", quotient, remainder)
	}
}

func TestDivideRejectsZeroDivisor(t *testing.T) {
	_, _, err := Divide(10, 0)
	if !errors.Is(err, ErrDivideByZero) {
		t.Fatalf("Divide(10, 0) error = %v, want %v", err, ErrDivideByZero)
	}
}
```

Create `exercises/02-functions/README.md`:

```markdown
# 02 - Functions

## Goal

Use parameters, multiple return values, and early error returns.

## Concepts

Parameters, named returns, integer division, remainder, scope, and sentinel errors.

## Exercise

Implement `Divide` so it returns an integer quotient and remainder. A zero divisor must return `ErrDivideByZero`.

## Acceptance Criteria

- `Divide(17, 5)` returns quotient `3`, remainder `2`, and no error.
- `Divide(10, 0)` returns an error matching `ErrDivideByZero`.

## Hints

Reject zero before using `/` or `%`. A clear explicit return is preferable to relying on bare named returns here.

## Commands

`go test -tags exercise ./exercises/02-functions/...`

## Reflection Prompts

Why can one Go function return both a result and an error? What would happen without the early zero check?
```

- [ ] **Step 4: Create Chapter 03 starter and challenge test**

Create `exercises/03-collections/exercise.go`:

```go
package collections

func UniqueWords(words []string) []string {
	return nil
}
```

Create `exercises/03-collections/exercise_test.go`:

```go
//go:build exercise

package collections

import (
	"slices"
	"testing"
)

func TestUniqueWordsPreservesFirstOccurrenceOrder(t *testing.T) {
	input := []string{"go", "test", "go", "maps", "test"}
	want := []string{"go", "test", "maps"}

	got := UniqueWords(input)
	if !slices.Equal(got, want) {
		t.Fatalf("UniqueWords(%v) = %v, want %v", input, got, want)
	}
}

func TestUniqueWordsDoesNotReuseInputStorage(t *testing.T) {
	input := []string{"go", "test"}
	got := UniqueWords(input)
	if len(got) != len(input) {
		t.Fatalf("UniqueWords(%v) length = %d, want %d", input, len(got), len(input))
	}
	got[0] = "changed"

	if input[0] != "go" {
		t.Fatalf("UniqueWords modified or reused input storage: input = %v", input)
	}
}
```

Create `exercises/03-collections/README.md`:

```markdown
# 03 - Collections

## Goal

Use a map and slice together to remove duplicates without losing order.

## Concepts

Slices, maps, `range`, `append`, zero values, and independent backing storage.

## Exercise

Implement `UniqueWords`. Keep only each word's first occurrence and return a newly allocated result slice.

## Acceptance Criteria

- `go test go maps test` becomes `go test maps` in that order.
- Mutating the returned slice does not change the input slice.

## Hints

Use a map as a set and append unseen words to a result slice. Do not filter by overwriting the input.

## Commands

`go test -tags exercise ./exercises/03-collections/...`

## Reflection Prompts

Why does iterating over the map lose input order? What is the relationship between a slice and its backing array?
```

- [ ] **Step 5: Verify all three chapter contracts**

Run: `go test ./...`

Expected: baseline suite passes because challenge tests are opt-in.

Run:

```bash
go test -tags exercise ./exercises/01-basics/...
go test -tags exercise ./exercises/02-functions/...
go test -tags exercise ./exercises/03-collections/...
```

Expected: each command fails on its first unmet acceptance criterion without a compile error or panic.

- [ ] **Step 6: Commit Chapters 01-03**

```bash
git add exercises/01-basics exercises/02-functions exercises/03-collections
git commit -m "learn: add Go fundamentals exercises"
```

---

### Task 5: Add Chapters 04-07

**Files:**
- Create: `exercises/04-structs-methods/{README.md,exercise.go,exercise_test.go}`
- Create: `exercises/05-interfaces-errors/{README.md,exercise.go,exercise_test.go}`
- Create: `exercises/06-testing/{README.md,exercise.go,exercise_test.go}`
- Create: `exercises/07-concurrency/{README.md,exercise.go,exercise_test.go}`

**Interfaces:**
- Produces: `NewTask(title string) Task`, `(*Task).Complete()`, `Task.IsComplete() bool`, `TitleValidator`, `ValidateTitle(TitleValidator, string) error`, `ErrEmptyTitle`, `Classify(int) string`, `Counter.Increment()`, and `Counter.Value() int`.

- [ ] **Step 1: Create Chapter 04 struct and method contract**

Create `exercises/04-structs-methods/exercise.go`:

```go
package structsmethods

type Task struct {
	title     string
	completed bool
}

func NewTask(title string) Task {
	return Task{}
}

func (t *Task) Complete() {}

func (t Task) IsComplete() bool {
	return false
}
```

Create `exercises/04-structs-methods/exercise_test.go`:

```go
//go:build exercise

package structsmethods

import "testing"

func TestTaskLifecycle(t *testing.T) {
	task := NewTask("learn methods")
	if task.title != "learn methods" {
		t.Fatalf("NewTask title = %q, want %q", task.title, "learn methods")
	}
	if task.IsComplete() {
		t.Fatal("new task is complete, want incomplete")
	}

	task.Complete()
	if !task.IsComplete() {
		t.Fatal("task remains incomplete after Complete")
	}
}
```

Create `exercises/04-structs-methods/README.md`:

```markdown
# 04 - Structs and Methods

## Goal

Model state with a struct and change it through methods.

## Concepts

Structs, constructors, unexported fields, value receivers, pointer receivers, and encapsulation.

## Exercise

Make `NewTask` preserve the title, make `Complete` mutate the task, and make `IsComplete` report current state.

## Acceptance Criteria

- A new task preserves its title and starts incomplete.
- Calling `Complete` changes that same task to complete.

## Hints

A method that mutates a value needs a pointer receiver. A read-only method can use a value receiver for this small struct.

## Commands

`go test -tags exercise ./exercises/04-structs-methods/...`

## Reflection Prompts

Why would a value receiver fail to persist the completion change? What does an unexported field protect?
```

- [ ] **Step 2: Create Chapter 05 interface and error contract**

Create `exercises/05-interfaces-errors/exercise.go`:

```go
package interfaceserrors

import "errors"

var ErrEmptyTitle = errors.New("title must not be empty")

type TitleValidator interface {
	Validate(title string) error
}

func ValidateTitle(validator TitleValidator, title string) error {
	return nil
}
```

Create `exercises/05-interfaces-errors/exercise_test.go`:

```go
//go:build exercise

package interfaceserrors

import (
	"errors"
	"testing"
)

type validatorFunc func(string) error

func (f validatorFunc) Validate(title string) error { return f(title) }

func TestValidateTitleUsesValidator(t *testing.T) {
	var gotTitle string
	err := ValidateTitle(validatorFunc(func(title string) error {
		gotTitle = title
		return nil
	}), "learn interfaces")

	if err != nil {
		t.Fatalf("ValidateTitle returned error: %v", err)
	}
	if gotTitle != "learn interfaces" {
		t.Fatalf("validator received %q, want %q", gotTitle, "learn interfaces")
	}
}

func TestValidateTitleWrapsValidationError(t *testing.T) {
	err := ValidateTitle(validatorFunc(func(string) error {
		return ErrEmptyTitle
	}), "")

	if !errors.Is(err, ErrEmptyTitle) {
		t.Fatalf("ValidateTitle error = %v, want wrapped %v", err, ErrEmptyTitle)
	}
	if err.Error() == ErrEmptyTitle.Error() {
		t.Fatalf("ValidateTitle error %q has no context", err)
	}
}
```

Create `exercises/05-interfaces-errors/README.md`:

```markdown
# 05 - Interfaces and Errors

## Goal

Depend on a small behavior and preserve error identity while adding context.

## Concepts

Interfaces, dependency injection, sentinel errors, wrapping with `%w`, and `errors.Is`.

## Exercise

Implement `ValidateTitle` by calling the supplied validator. Return nil on success; otherwise wrap the validator error with useful context.

## Acceptance Criteria

- The validator receives the original title.
- A successful validator produces no error.
- A failed validator produces contextual text while still matching `ErrEmptyTitle` through `errors.Is`.

## Hints

Use `fmt.Errorf` and the `%w` verb. Define interfaces near the code that consumes them.

## Commands

`go test -tags exercise ./exercises/05-interfaces-errors/...`

## Reflection Prompts

Why does wrapping preserve more information than formatting with `%v`? Why is this interface only one method?
```

- [ ] **Step 3: Create Chapter 06 testing contract**

Create `exercises/06-testing/exercise.go`:

```go
package testingpractice

func Classify(number int) string {
	if number < 0 {
		return "negative"
	}
	if number > 0 {
		return "positive"
	}
	return "zero"
}
```

Create `exercises/06-testing/exercise_test.go`:

```go
//go:build exercise

package testingpractice
```

Create `exercises/06-testing/README.md`:

```markdown
# 06 - Testing

## Goal

Write a table-driven unit test, subtests, coverage, and a basic benchmark.

## Concepts

The `testing` package, test tables, `t.Run`, failure messages, coverage, and benchmarks.

## Exercise

In `exercise_test.go`, write `TestClassify` as a table with negative, zero, and positive cases. Add `BenchmarkClassify` that repeatedly calls `Classify(42)` using `b.N`.

## Acceptance Criteria

- The test has three named subtests and all pass.
- Failures print the input, actual value, and expected value.
- Coverage includes every branch in `Classify`.
- `go test -bench .` discovers and runs `BenchmarkClassify`.

## Hints

Keep test cases in a slice of anonymous structs. Call `b.ResetTimer()` immediately before the benchmark loop.

## Commands

```bash
go test -tags exercise -cover ./exercises/06-testing/...
go test -tags exercise -bench . ./exercises/06-testing/...
```

## Reflection Prompts

What duplication does a test table remove? Why should a benchmark avoid setup work inside its measured loop?
```

The starter function is complete because this chapter's learning output is the learner-authored test and benchmark.

- [ ] **Step 4: Create Chapter 07 concurrency contract**

Create `exercises/07-concurrency/exercise.go`:

```go
package concurrency

type Counter struct {
	value int
}

func (c *Counter) Increment() {}

func (c *Counter) Value() int {
	return 0
}
```

Create `exercises/07-concurrency/exercise_test.go`:

```go
//go:build exercise

package concurrency

import (
	"sync"
	"testing"
)

func TestCounterConcurrentIncrement(t *testing.T) {
	var counter Counter
	var workers sync.WaitGroup

	for range 100 {
		workers.Add(1)
		go func() {
			defer workers.Done()
			counter.Increment()
		}()
	}

	workers.Wait()
	if got := counter.Value(); got != 100 {
		t.Fatalf("Counter.Value() = %d, want 100", got)
	}
}
```

Create `exercises/07-concurrency/README.md`:

```markdown
# 07 - Concurrency

## Goal

Make shared mutable state safe when many goroutines use it.

## Concepts

Goroutines, wait groups, mutexes, critical sections, and the race detector.

## Exercise

Make `Counter` safe and correct when 100 goroutines increment it concurrently. Both `Increment` and `Value` must synchronize access to `value`.

## Acceptance Criteria

- The final value is exactly 100.
- `go test -race` reports no data race.
- The counter uses `sync.Mutex` or `sync.RWMutex`, not sleeps or timing assumptions.

## Hints

Put the lock inside `Counter` beside the state it protects. Keep each critical section small and use `defer` when it makes unlocking easier to verify.

## Commands

`go test -race -tags exercise ./exercises/07-concurrency/...`

## Reflection Prompts

Why does the wait group not protect the counter itself? What behavior does the race detector find that a normal assertion may miss?
```

- [ ] **Step 5: Verify Chapters 04-07 safely**

Run: `make fmt && go test ./... && go test -race ./...`

Expected: baseline checks pass.

Run:

```bash
go test -tags exercise ./exercises/04-structs-methods/...
go test -tags exercise ./exercises/05-interfaces-errors/...
go test -tags exercise ./exercises/06-testing/...
go test -race -tags exercise ./exercises/07-concurrency/...
```

Expected: Chapters 04, 05, and 07 each report an unmet requirement without a compile error or panic. Chapter 06 passes with `[no tests to run]` until the learner authors tests.

- [ ] **Step 6: Commit Chapters 04-07**

```bash
git add exercises/04-structs-methods exercises/05-interfaces-errors exercises/06-testing exercises/07-concurrency
git commit -m "learn: add applied Go exercises"
```

---

### Task 6: Final Verification and GitHub Publication

**Files:**
- Verify: all tracked repository files
- Modify only if verification reveals a concrete defect

**Interfaces:**
- Consumes: all outputs from Tasks 1-5.
- Produces: a clean local `main` branch and public `LDKhangg/go-playground` GitHub repository with passing CI.

- [ ] **Step 1: Verify formatting, analysis, tests, and races**

Run: `make check`

Expected: command exits 0; all packages pass and the race detector reports no races.

- [ ] **Step 2: Verify challenge tests are discoverable and intentionally incomplete**

Run:

```bash
go test -tags exercise ./exercises/01-basics/...
go test -tags exercise ./exercises/02-functions/...
go test -tags exercise ./exercises/03-collections/...
go test -tags exercise ./exercises/04-structs-methods/...
go test -tags exercise ./exercises/05-interfaces-errors/...
go test -race -tags exercise ./exercises/07-concurrency/...
go test -tags exercise ./exercises/06-testing/...
```

Expected: Chapters 01-05 and 07 reach their challenge tests and fail because starter code has not been solved; no package fails to compile and no test panics. Chapter 06 compiles and passes with `[no tests to run]` because its learner-authored tests do not exist yet.

- [ ] **Step 3: Inspect repository state and history**

Run:

```bash
git status --short --branch
git log --oneline --decorate -10
git diff --check
```

Expected: `main` has only intended changes, history contains focused milestone commits, and `git diff --check` prints nothing.

- [ ] **Step 4: Commit any remaining plan/documentation changes**

```bash
git add docs/superpowers/plans/2026-07-13-go-learning-playground.md
git commit -m "docs: add Go playground implementation plan"
```

If the plan was committed before execution, skip this step rather than creating an empty commit.

- [ ] **Step 5: Create and push the public repository**

First confirm the target does not already exist:

```bash
gh repo view LDKhangg/go-playground
```

Expected before creation: GitHub reports that the repository could not be found. Then run:

```bash
gh repo create LDKhangg/go-playground --public --source=. --remote=origin --push --description "Learn Go through test-driven exercises and an evolving task API."
```

Expected: GitHub prints the repository URL and pushes `main` to `origin`.

If the repository already exists and belongs to `LDKhangg`, add it without replacing any existing remote:

```bash
git remote add origin git@github.com:LDKhangg/go-playground.git
git push -u origin main
```

- [ ] **Step 6: Verify the published repository and CI**

Run:

```bash
gh repo view LDKhangg/go-playground --web=false
gh run list --repo LDKhangg/go-playground --limit 1
git status --short --branch
```

Expected: repository visibility is public, one CI run is listed, and local `main` tracks `origin/main` with no uncommitted changes.
