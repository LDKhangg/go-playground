package testingpractice

import "testing"

func TestClassify(t *testing.T) {
	cases := []struct {
		name  string
		input int
		want  string
	}{
		{
			name:  "negative",
			input: -10,
			want:  "negative",
		},
		{
			name:  "zero",
			input: 0,
			want:  "zero",
		},
		{
			name:  "positive",
			input: 10,
			want:  "positive",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Classify(tc.input)

			if got != tc.want {
				t.Fatalf(
					"Classify(%d) = %q, want %q",
					tc.input,
					got,
					tc.want,
				)
			}
		})
	}
}

func BenchmarkClassify(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Classify(42)
	}
}
