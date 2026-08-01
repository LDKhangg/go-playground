//go:build exercise

package mapsx

import "testing"

func TestCountWords(t *testing.T) {
	counts := CountWords("go go tests")
	if counts["go"] != 2 {
		t.Fatalf("counts[go] = %d, want 2", counts["go"])
	}
	if counts["tests"] != 1 {
		t.Fatalf("counts[tests] = %d, want 1", counts["tests"])
	}
}

func TestLookupScore(t *testing.T) {
	scores := map[string]int{"mina": 10}
	if score, ok := LookupScore(scores, "mina"); !ok || score != 10 {
		t.Fatalf("LookupScore(mina) = (%d, %t), want (10, true)", score, ok)
	}
	if _, ok := LookupScore(scores, "kai"); ok {
		t.Fatal("expected missing key to return ok=false")
	}
}
