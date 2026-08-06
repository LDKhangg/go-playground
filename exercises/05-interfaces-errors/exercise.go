package interfaceserrors

import (
	"errors"
	"fmt"
)

var ErrEmptyTitle = errors.New("title must not be empty")

type TitleValidator interface {
	Validate(title string) error
}

func ValidateTitle(validator TitleValidator, title string) error {
	if err := validator.Validate(title); err != nil {
		return fmt.Errorf("Validate title %q failed: %w", title, err)
	}
	return nil
}
