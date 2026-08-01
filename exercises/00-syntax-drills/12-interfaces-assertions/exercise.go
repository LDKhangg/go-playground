package interfacesx

import "io"

type Temporary interface {
	Temporary() bool
}

func WriteGreeting(w io.Writer, name string) error {
	panic("TODO")
}

func AsTemporary(err error) (Temporary, bool) {
	panic("TODO")
}
