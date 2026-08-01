//go:build exercise

package functions

import (
	"errors"
	"io"
	"strings"
	"testing"
)

type trackingReadCloser struct {
	reader io.Reader
	closed bool
}

func (t *trackingReadCloser) Read(p []byte) (int, error) {
	return t.reader.Read(p)
}

func (t *trackingReadCloser) Close() error {
	t.closed = true
	return nil
}

type failingReader struct{}

func (f failingReader) Read(_ []byte) (int, error) { return 0, errors.New("boom") }

func TestReadAndClose(t *testing.T) {
	rc := &trackingReadCloser{reader: strings.NewReader("go")}
	data, err := ReadAndClose(rc)
	if err != nil {
		t.Fatalf("ReadAndClose returned error: %v", err)
	}
	if string(data) != "go" {
		t.Fatalf("data = %q, want %q", string(data), "go")
	}
	if !rc.closed {
		t.Fatal("expected reader to be closed")
	}
}

func TestReadAndCloseReturnsReadError(t *testing.T) {
	rc := &trackingReadCloser{reader: failingReader{}}
	if _, err := ReadAndClose(rc); err == nil {
		t.Fatal("expected read error")
	}
	if !rc.closed {
		t.Fatal("expected reader to be closed on error")
	}
}
