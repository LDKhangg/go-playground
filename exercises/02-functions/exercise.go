package functions

import "errors"

var ErrDivideByZero = errors.New("cannot divide by zero")

func Divide(dividend, divisor int) (quotient, remainder int, err error) {
	return 0, 0, nil
}
