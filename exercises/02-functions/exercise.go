package functions

import "errors"

var ErrDivideByZero = errors.New("cannot divide by zero")

func Divide(dividend, divisor int) (quotient, remainder int, err error) {
	if divisor == 0 {
		return 0, 0, ErrDivideByZero
	}

	quotient = dividend / divisor

	remainder = dividend % divisor

	return quotient, remainder, nil
}

func Double(number *int) { *number = *number * 2 }
