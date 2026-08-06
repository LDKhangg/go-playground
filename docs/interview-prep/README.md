# Go Interview Prep — Fresher / Junior

A focused, repo-linked study guide for Go interviews at fresher and junior
levels. Every topic here is backed by an exercise or example elsewhere in this
repository, so you can practice what you read.

## How To Use This Guide

- Read each module top to bottom; every module ends with a **Practice Check**.
- Treat the Q&A sections as oral drills: cover the answer, say it out loud, then check.
- Re-run the linked exercises from `exercises/` before an interview day.
- Re-read `05-top-interview-questions-and-traps.md` the evening before an interview.

## Suggested Roadmap (2 weeks)

| Week | Day | Module |
| --- | --- | --- |
| 1 | 1-2 | `01-go-core-fundamentals.md` |
| 1 | 3-4 | `02-concurrency-and-synchronization.md` |
| 1 | 5-6 | `03-backend-and-http.md` |
| 1 | 7 | Review + code exercises (`exercises/01-basics` ... `exercises/05-interfaces-errors`) |
| 2 | 1-2 | `04-architecture-and-testing.md` |
| 2 | 3-4 | `05-top-interview-questions-and-traps.md` |
| 2 | 5-6 | Build one small API from memory (see `apps/task-manager`) |
| 2 | 7 | Mock interview: answer questions out loud, timed |

## Modules

1. [Go Core Fundamentals](01-go-core-fundamentals.md) — types, slices, maps, strings, interfaces, errors, memory basics.
2. [Concurrency And Synchronization](02-concurrency-and-synchronization.md) — goroutines, channels, `select`, `sync`, `context`, race detector.
3. [Backend And HTTP](03-backend-and-http.md) — `net/http`, middleware, JSON, error handling, graceful shutdown.
4. [Architecture And Testing](04-architecture-and-testing.md) — package layout, dependency injection, repository pattern, table-driven tests, benchmarks.
5. [Top Interview Questions And Traps](05-top-interview-questions-and-traps.md) — the questions interviewers actually ask, and the traps that fail candidates.

## Practice

[Practice Exercises With Examples](practice-exercises.md) — runnable problem +
example-solution pairs for every module (slices, maps, mutex counters, context
cancellation, a small HTTP API, table-driven tests, service dependency
injection, and the classic traps). Try each problem first, then compare with
the example.

## What Interviewers Look For At Junior Level

- **Correctness first**: you can write working Go with the standard library.
- **Mental models**: you can explain how slices, maps, goroutines, and `defer` work, not just use them.
- **Concurrency hygiene**: you know about data races and reach for `go test -race`.
- **Error handling**: you return and propagate errors instead of swallowing them.
- **Testing mindset**: you can write a table-driven test and explain what it verifies.

## The Five-Minute Pitch

Be ready to answer, in under five minutes:

1. What is this repository?
2. Which part are you most proud of and why?
3. What was the hardest bug you fixed and how did you debug it?
4. What do you want to learn next, and what are you doing about it?
