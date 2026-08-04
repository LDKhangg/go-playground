package conversion

import (
	"errors"
	"strconv"
)

func ParsePort(raw string) (int, error) {
	val, err := strconv.Atoi(raw)
	if err != nil {
		return 0, err
	}
	if val < 0 {
		return 0, errors.New("port cannot be negative")
	}
	return val, nil
}
