package interfacesx

import (
	"fmt"
	"io"
)

type Temporary interface {
	Temporary() bool
}

func WriteGreeting(w io.Writer, name string) error {
	_, err := fmt.Fprintf(w, "hello, %s", name)
	return err
}

func AsTemporary(err error) (Temporary, bool) {
	temp, ok := err.(Temporary)
	return temp, ok
}
