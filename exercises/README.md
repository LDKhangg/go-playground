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
