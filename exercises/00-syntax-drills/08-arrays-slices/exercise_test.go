//go:build exercise

package slices

import (
	"reflect"
	"testing"
)

func TestInsertAt(t *testing.T) {
	got, err := InsertAt([]int{1, 3}, 1, 2)
	if err != nil {
		t.Fatalf("InsertAt returned error: %v", err)
	}
	if want := []int{1, 2, 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("InsertAt = %v, want %v", got, want)
	}
}

func TestRemoveAt(t *testing.T) {
	got, err := RemoveAt([]int{1, 2, 3}, 1)
	if err != nil {
		t.Fatalf("RemoveAt returned error: %v", err)
	}
	if want := []int{1, 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("RemoveAt = %v, want %v", got, want)
	}
}

func TestSliceHelpersRejectInvalidIndex(t *testing.T) {
	if _, err := InsertAt([]int{1}, 3, 2); err == nil {
		t.Fatal("expected invalid index error from InsertAt")
	}
	if _, err := RemoveAt([]int{1}, -1); err == nil {
		t.Fatal("expected invalid index error from RemoveAt")
	}
}
