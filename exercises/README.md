# Exercise Guide

Each chapter contains starter code and a lesson. Chapters 01-05 and 07 include challenge tests without solutions. In Chapter 06, writing the tests is the exercise, so its tagged `exercise_test.go` starts empty.

## Start a Chapter

1. Read the chapter README.
2. Run `go test -tags exercise ./exercises/<chapter>/...`. Chapters 01-05 and 07 fail on an unmet requirement; Chapter 06 intentionally succeeds without tests until you author them.
3. Implement one acceptance criterion at a time.
4. Format and rerun the tagged test.
5. Run `make check` to protect previously completed work.

For Chapter 06, use its README acceptance criteria to assess the table-driven test, subtests, diagnostics, coverage, and benchmark. There is no prewritten meta-test.

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
