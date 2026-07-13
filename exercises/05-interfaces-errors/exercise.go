package interfaceserrors

import "errors"

var ErrEmptyTitle = errors.New("title must not be empty")

type TitleValidator interface {
	Validate(title string) error
}

func ValidateTitle(validator TitleValidator, title string) error {
	return nil
}
