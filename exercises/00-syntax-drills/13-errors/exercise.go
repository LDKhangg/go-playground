package errvalues

import (
	"errors"
	"strings"
)

var (
	ErrEmptyTitle   = errors.New("empty title")
	ErrTitleTooLong = errors.New("title too long")
)

func ValidateTitle(title string) error {
	trimmedStr := strings.TrimSpace(title)
	if trimmedStr == "" {
		return ErrEmptyTitle
	}
	if len(trimmedStr) > 80 {
		return ErrTitleTooLong
	}
	return nil
}
