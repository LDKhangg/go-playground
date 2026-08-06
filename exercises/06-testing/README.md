# 06 - Testing

## Goal

Write a table-driven unit test, subtests, coverage check, and a basic benchmark.

## Concepts

- Tests live in files ending with `_test.go`.
- A table collects multiple inputs and expected outputs in one structure.
- `t.Run` gives each table entry its own named subtest.
- `go test -cover` reports which statements executed.
- Benchmarks use `b.N` to let Go choose an appropriate iteration count.

## Syntax Primer

```go
func TestExample(t *testing.T) {
	cases := []struct {
		name string
		input int
		want string
	}{
		{name: "zero", input: 0, want: "zero"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Classify(tc.input)
			if got != tc.want {
				t.Fatalf("Classify(%d) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
```

Tests receive `*testing.T`; benchmarks receive `*testing.B`. `%d` prints an integer and `%q` prints a quoted string, making failures easier to read.

## Mental Model

A table is a compact specification: each row says what the program should do for one input. Subtests make a failed row obvious in command output. Coverage tells you which paths ran, not whether your assertions were good. A benchmark measures repeated work, so setup should be outside the measured loop.

## Annotated Examples

```go
func BenchmarkExample(b *testing.B) {
	b.ResetTimer() // Exclude setup above this line from measurement.
	for i := 0; i < b.N; i++ {
		_ = Classify(42)
	}
}
```

```go
for _, tc := range cases {
	tc := tc // Keep this iteration's value if the subtest later uses t.Parallel.
	t.Run(tc.name, func(t *testing.T) {
		// Assert this case.
	})
}
```

## Common Diagnostics

- `undefined: testing`: add `import "testing"` in the test file.
- `TestXxx has wrong signature`: a test must be `func TestXxx(t *testing.T)`.
- A command reports `[no test files]`: the file name must end in `_test.go`.
- Coverage misses a branch: add a table row whose input reaches that branch.

## Exercise

In `exercise_test.go`, write `TestClassify` as a table with `negative`, `zero`, and `positive` named cases. Add `BenchmarkClassify` that repeatedly calls `Classify(42)` using `b.N`.

## Acceptance Criteria

- The test has three named subtests and all pass.
- Failures print the input, actual value, and expected value.
- Coverage includes every branch in `Classify`.
- `go test -bench .` discovers and runs `BenchmarkClassify`.

## Hints

- Use a slice of anonymous structs for the cases.
- Run each row with `t.Run(tc.name, ...)`.
- Call `b.ResetTimer()` immediately before the benchmark loop.

## Verify

Run:

```bash
gofmt -w exercises/06-testing
go test -tags exercise -cover ./exercises/06-testing/...
go test -tags exercise -bench . ./exercises/06-testing/...
```

This chapter intentionally starts with no prewritten challenge test. The commands can pass before you write the required test and benchmark, so inspect every acceptance criterion rather than treating a green exit code as completion.

## Reflection Prompts

- What duplication does a test table remove?
- Why should a benchmark avoid setup work inside its measured loop?
- Why can full coverage still miss a bug?
