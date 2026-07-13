# Go Learning Playground Design

## Purpose

Turn the existing task API into a public, beginner-friendly Go learning repository. The repository should combine guided exercises with one application that grows as the learner gains new skills. It should also make every completed lesson easy to preserve in Git and publish to GitHub.

The target learner is new to Go. All learning material will be written in English, while code and technical terminology will follow standard Go conventions.

## Goals

- Give a beginner an obvious starting point and an ordered learning path.
- Teach each core concept through an exercise with executable tests.
- Apply learned concepts to the existing task API over a series of milestones.
- Keep the learner responsible for writing solutions rather than shipping completed answers.
- Record progress through a learning log, small commits, and a public GitHub history.
- Keep the baseline project formatted, tested, race-checked, and continuously integrated.
- Present the repository with a polished, readable GitHub README.

## Non-Goals

- Providing solutions for the exercises.
- Building the final production task service during initial setup.
- Adding a database, web framework, deployment platform, or unnecessary third-party dependencies during initial setup.
- Covering advanced Go internals before the beginner curriculum is complete.

## Repository Structure

```text
go-playground/
├── .github/workflows/ci.yml
├── .gitignore
├── CONTRIBUTING.md
├── LICENSE
├── Makefile
├── README.md
├── ROADMAP.md
├── docs/
│   ├── learning-log.md
│   └── superpowers/specs/
├── exercises/
│   ├── README.md
│   ├── 01-basics/
│   ├── 02-functions/
│   ├── 03-collections/
│   ├── 04-structs-methods/
│   ├── 05-interfaces-errors/
│   ├── 06-testing/
│   └── 07-concurrency/
├── internal/
│   ├── httpapi/
│   └── tasks/
├── go.mod
└── main.go
```

The existing `main.go` and `internal` packages remain the central project. This keeps the module simple and lets `go test ./...` cover both the application and completed exercises.

## README Experience

The root README acts as the repository's landing page and study dashboard. It will include:

- A concise project identity and explanation of the learning model.
- A small, purposeful badge row for Go version and CI status.
- Prerequisites and a quick-start path that works after cloning.
- A progress table linking each chapter to its exercise and project milestone.
- Common commands and API examples.
- A concise architecture overview and repository map.
- Instructions for recording progress and making learning commits.
- Links to the roadmap, exercise guide, learning log, and contribution guide.

The visual style should be clean and distinctive without excessive badges, generated artwork, or decorative noise.

## Curriculum

The initial curriculum contains seven ordered chapters:

1. Basics: variables, constants, primitive types, control flow, and loops.
2. Functions: parameters, return values, multiple returns, scope, and basic pointers.
3. Collections: arrays, slices, maps, range, and safe copying.
4. Structs and methods: custom types, constructors, receiver methods, and composition.
5. Interfaces and errors: small interfaces, sentinel errors, wrapping, and error inspection.
6. Testing: focused tests, table-driven tests, subtests, coverage, and introductory benchmarks.
7. Concurrency: goroutines, channels, mutexes, context cancellation, and the race detector.

Each chapter contains an English README with these sections:

- Goal
- Concepts
- Exercise
- Acceptance Criteria
- Hints
- Commands
- Reflection prompts

Each exercise includes starter code and executable tests but no implementation answer. Tests communicate observable requirements and produce useful expected-versus-actual failures. Exercise package names and paths remain independent so learners can run one chapter at a time.

## Exercise and CI Workflow

Future exercises must not make the default branch permanently fail before the learner begins them. Each chapter therefore starts with a buildable starter implementation that returns zero values or a clearly documented placeholder result. Its challenge tests use an explicit opt-in build tag named `exercise` until the chapter is started.

The workflows are:

- Baseline check: `go test ./...` tests the application and completed exercise code.
- Active exercise: `go test -tags exercise ./exercises/<chapter>/...` runs that chapter's challenge tests.
- Completion: the learner removes that chapter's `exercise` build constraint so its tests join the baseline suite, then updates the progress table and learning log in the same milestone commit.

This keeps `main` green while preserving test-driven, solution-free exercises. The exercise guide will explain the build-tag transition explicitly.

## Task API Milestones

The existing in-memory API is the application thread connecting the chapters. The roadmap will introduce these milestones in order:

1. Understand and test the existing health, list, and create behavior.
2. Add complete task CRUD operations and domain validation.
3. Add handler tests with `httptest` and clearer HTTP error responses.
4. Add configuration, request context, and graceful shutdown.
5. Introduce persistence behind a small store interface.
6. Implement database storage and migrations.
7. Add middleware, structured logging, and observability.
8. Package with Docker and deploy after local behavior is well tested.

Only the baseline protection needed for the current API is part of initial setup. Later milestones remain learning work for the repository owner.

## Application Boundaries

- `main.go` owns composition and process startup.
- `internal/tasks` owns task data, validation, and storage behavior.
- `internal/httpapi` owns HTTP decoding, status mapping, headers, and JSON encoding.
- Domain code returns explicit sentinel or wrapped errors for expected failures.
- HTTP handlers translate known domain errors into stable status codes and JSON errors.
- Mutable in-memory state remains synchronized, and callers receive copies rather than internal slices.

No framework abstraction is introduced while the standard library remains sufficient.

## Developer Commands

The Makefile provides stable entry points:

- `make run`: run the task API.
- `make fmt`: format Go source.
- `make test`: run baseline tests.
- `make race`: run tests with the race detector.
- `make vet`: run static analysis with `go vet`.
- `make check`: run formatting verification, vet, tests, and race tests.

The README also shows the underlying Go commands so the learner understands what Make invokes.

## Testing and Quality

Initial setup preserves current store tests and adds HTTP handler coverage for:

- Health success response.
- Empty task list response.
- Successful task creation.
- Invalid JSON.
- Empty title validation.
- Unsupported HTTP methods and the `Allow` header.

GitHub Actions runs formatting verification, `go vet ./...`, `go test ./...`, and `go test -race ./...` for pushes and pull requests. The workflow uses the Go version declared by `go.mod` where supported and caches Go build data through the official setup action.

## Git and GitHub

- The module path becomes `github.com/LDKhangg/go-playground`.
- Git uses `main` as the default branch.
- The GitHub repository is public at `github.com/LDKhangg/go-playground`.
- The repository uses the MIT license.
- Initial work is split into focused commits for the approved design, baseline project, learning scaffold, exercises/documentation, and CI as practical.
- Later learning commits should represent one completed exercise or one coherent project milestone.
- The contribution guide provides a simple commit format such as `learn: complete slices exercise` or `feat: add task deletion` without enforcing additional tooling.

## Learning Log

`docs/learning-log.md` contains a repeatable entry format:

- Date and chapter or milestone.
- Concepts learned.
- What was implemented.
- A problem encountered and how it was solved.
- Related commit hash after the commit exists.
- One follow-up question or topic to revisit.

The log is intentionally manual. Writing the summary is part of the learning process.

## Initial Completion Criteria

- A new learner can clone the repository, understand the README, run the API, run baseline checks, and identify the first exercise.
- The README, roadmap, exercise guide, learning log, and contribution guide agree on one workflow.
- Seven chapter directories contain English instructions, starter code, and opt-in challenge tests without solutions.
- The current application has baseline store and HTTP handler tests.
- Formatting, vet, normal tests, and race tests pass locally and in GitHub Actions.
- The module path matches the GitHub repository.
- The public GitHub repository exists and `main` is pushed successfully.
