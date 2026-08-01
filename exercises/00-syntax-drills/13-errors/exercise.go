package errvalues

import "errors"

var ErrEmptyTitle = errors.New("empty title")
var ErrTitleTooLong = errors.New("title too long")

func ValidateTitle(title string) error {
	panic("TODO")
}
