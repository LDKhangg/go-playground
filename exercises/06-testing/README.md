# 06 - Testing

## Goal

Write a table-driven unit test, subtests, coverage, and a basic benchmark.

## Concepts

The `testing` package, test tables, `t.Run`, failure messages, coverage, and benchmarks.

## Exercise

In `exercise_test.go`, write `TestClassify` as a table with `negative`, `zero`, and `positive` named cases. Add `BenchmarkClassify` that repeatedly calls `Classify(42)` using `b.N`.

This chapter is the intentional workflow exception: you are writing the tests, so no prewritten challenge test fails before you begin. The initial tagged commands succeed without running a test or benchmark. Assess your work against every acceptance criterion below.

## Acceptance Criteria

- The test has three named subtests and all pass.
- Failures print the input, actual value, and expected value.
- Coverage includes every branch in `Classify`.
- `go test -bench .` discovers and runs `BenchmarkClassify`.

## Hints

Keep test cases in a slice of anonymous structs and run each case with `t.Run`. Call `b.ResetTimer()` immediately before the benchmark loop.

## Commands

```bash
go test -tags exercise -cover ./exercises/06-testing/...
go test -tags exercise -bench . ./exercises/06-testing/...
```

## Reflection Prompts

What duplication does a test table remove? Why should a benchmark avoid setup work inside its measured loop?
