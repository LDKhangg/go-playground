# Foundational Go Curriculum Design

## Purpose

Transform the existing seven exercise chapters into a self-guided, English-language curriculum for learning foundational Go syntax and idioms. The repository remains solution-free: learners write the code, use the tests to validate it, and record their progress in Git.

This work improves learning material only. It does not add backend, database, deployment, authentication, or framework lessons.

## Goals

- Explain each topic before the learner is asked to implement it.
- Make every exercise README sufficient to understand the relevant syntax without leaving the repository.
- Teach the meaning of common compiler and `gopls` diagnostics in context.
- Keep the existing starter code and acceptance tests as the practice mechanism.
- Keep challenge solutions out of the repository.
- Remove editor diagnostics that are style noise or an artifact of opt-in exercise tests.
- Publish one focused commit and push for every changed file so Git history is a granular learning timeline.

## Non-Goals

- Providing worked solutions, hidden solution branches, or answer files.
- Adding more chapters beyond the current seven.
- Extending the task API or beginning the backend roadmap.
- Suppressing Go compiler or type-checking diagnostics.
- Enforcing a project-wide lint standard while the repository contains deliberately incomplete starters.

## Curriculum Structure

Each chapter README will follow this consistent order:

1. **Goal**: the concrete outcome for the chapter.
2. **Concepts**: a short list of terms introduced in the chapter.
3. **Syntax Primer**: the smallest useful syntax reference, with compilable examples.
4. **Mental Model**: an explanation of what the Go runtime or compiler treats as state, value, reference, control flow, or behavior for that topic.
5. **Annotated Examples**: small examples with commentary for the lines that matter.
6. **Common Diagnostics**: realistic compiler or `gopls` messages, what they mean, and the next debugging action.
7. **Exercise**: the existing implementation task.
8. **Acceptance Criteria**: observable behavior already asserted by the tagged challenge test.
9. **Hints**: focused direction without implementation answers.
10. **Verify**: exact commands and expected result after completing the chapter.
11. **Reflection Prompts**: questions that require the learner to explain the concept.

All prose, headings, diagnostic explanations, and examples are English. Go keywords, commands, paths, and identifier names retain their standard English spelling.

## Chapter Coverage

| Chapter | Required teaching coverage |
| --- | --- |
| 01 Basics | package declarations, `main`, comments, variables, constants, basic types, zero values, conversions, comparisons, `if`, `switch`, and `for` |
| 02 Functions | declarations, parameters, return values, multiple returns, scope, errors, pointers, `&`, and `*` |
| 03 Collections | arrays, slices, maps, `range`, `append`, `len`, map membership, copy behavior, and nil versus empty values |
| 04 Structs and Methods | struct literals, fields, constructors, method declarations, value and pointer receivers, mutation, composition, and encapsulation |
| 05 Interfaces and Errors | implicit interface satisfaction, narrow interfaces, `error`, sentinel errors, `fmt.Errorf`, `%w`, and `errors.Is` |
| 06 Testing | package `testing`, test naming, tables, subtests, assertions with standard library tools, coverage, and benchmarks |
| 07 Concurrency | goroutines, channels, channel closure, `select`, `sync.Mutex`, critical sections, `context.Context`, cancellation, and the race detector |

Examples must directly support the chapter exercise and must not introduce a later chapter's core concept without explaining it first.

## Diagnostic Experience

The current yellow diagnostics in starter files are mostly static analysis rather than syntax failures. For example, Chapter 04 declares fields before its unfinished constructors and methods use them; `staticcheck` reports those fields as unused even though the learner is expected to use them.

The Neovim Go configuration will therefore:

- Retain `gopls` compiler and type diagnostics.
- Disable `gopls` `staticcheck` diagnostics.
- Stop running `golangci-lint` automatically on Go buffer entry, writes, and insert-mode exit.
- Retain the existing explicit Go lint command (`<Space>tl`) for use after an exercise is complete.
- Apply `-tags=exercise` to `gopls` only when its workspace root is this repository, so tagged `exercise_test.go` files are analyzed as package members instead of reported as excluded files.

The chapter diagnostics section will distinguish errors that block compilation from style and static-analysis warnings. It will explicitly explain the `//go:build exercise` workflow and the command needed to run a chapter test.

## Exercise Workflow

Challenge tests continue to use `//go:build exercise` until the learner completes their chapter. This keeps `go test ./...` green while allowing an active chapter to fail productively with:

```bash
go test -tags exercise ./exercises/<chapter>/...
```

After completion, the learner removes the build constraint from that chapter's test, runs `make check`, updates the root progress table and learning log, then commits the learning milestone. The existing root README and exercise guide will be updated only where needed to stay consistent with the expanded lesson format.

## Git Workflow

Every created, modified, or deleted file is committed separately and pushed immediately to `origin/main`. Each commit message describes that one file's learning purpose, for example:

```text
docs: expand structs and methods lesson
```

Commits must contain only their one intended path. This deliberately favors a detailed public learning history over a compact implementation history.

## Verification

- Markdown links and command paths in every changed lesson are checked manually.
- `go test ./...` remains green.
- Each chapter's active test command is run with `-tags exercise`; intentionally incomplete chapters may fail only at their documented unmet acceptance criteria.
- The changed Neovim configuration is loaded headlessly and checked for Lua errors.
- A tagged exercise test opened through `gopls` in this repository no longer receives the excluded-file diagnostic.
- `git diff --check` passes before every per-file commit.

## Completion Criteria

- All seven chapter READMEs use the documented lesson structure and cover their required syntax.
- The repository contains no exercise solutions.
- The yellow Chapter 04 starter warnings caused by automatic static analysis no longer appear, while real `gopls` compiler and type errors remain visible.
- Opening an `exercise_test.go` file does not produce the build-tag exclusion message.
- Every changed file has its own pushed commit on `main`.
- The baseline Go test suite passes.
