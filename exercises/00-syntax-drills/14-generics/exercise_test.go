//go:build exercise

package generics

import "testing"

func TestContains(t *testing.T) {
	if !Contains([]int{1, 2, 3}, 2) {
		t.Fatal("expected int slice to contain 2")
	}
	if !Contains([]string{"go", "tests"}, "go") {
		t.Fatal("expected string slice to contain go")
	}
	if Contains([]string{"go", "tests"}, "rust") {
		t.Fatal("expected missing string to return false")
	}
}
