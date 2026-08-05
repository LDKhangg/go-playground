package pointers

import "errors"

func Increment(value *int) error {
	if value == nil {
		return errors.New("value must not be null")
	}
	*value += 1
	return nil
}
