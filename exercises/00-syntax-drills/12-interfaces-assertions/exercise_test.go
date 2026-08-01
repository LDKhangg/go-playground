//go:build exercise

package interfacesx

import (
	"bytes"
	"errors"
	"testing"
)

type temporaryError struct{}

func (temporaryError) Error() string   { return "temporary" }
func (temporaryError) Temporary() bool { return true }

type failingWriter struct{}

func (failingWriter) Write(_ []byte) (int, error) { return 0, errors.New("write failed") }

func TestWriteGreeting(t *testing.T) {
	buf := new(bytes.Buffer)
	if err := WriteGreeting(buf, "Mina"); err != nil {
		t.Fatalf("WriteGreeting returned error: %v", err)
	}
	if got := buf.String(); got != "hello, Mina" {
		t.Fatalf("buffer = %q, want %q", got, "hello, Mina")
	}
}

func TestWriteGreetingReturnsWriterError(t *testing.T) {
	if err := WriteGreeting(failingWriter{}, "Mina"); err == nil {
		t.Fatal("expected writer error")
	}
}

func TestAsTemporary(t *testing.T) {
	if temp, ok := AsTemporary(temporaryError{}); !ok || !temp.Temporary() {
		t.Fatal("expected temporary assertion to succeed")
	}
	if _, ok := AsTemporary(errors.New("nope")); ok {
		t.Fatal("expected non-temporary error to fail assertion")
	}
}
