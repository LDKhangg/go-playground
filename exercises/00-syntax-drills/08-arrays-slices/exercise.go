package slices

import "fmt"

func InsertAt(values []int, index int, value int) ([]int, error) {
	if index < 0 || index > len(values) {
		return nil, fmt.Errorf(
			"index %d out of range [0,%d]",
			index,
			len(values),
		)
	}
	newArr := make([]int, 0, len(values)+1)

	newArr = append(newArr, values[:index]...)
	newArr = append(newArr, value)
	newArr = append(newArr, values[index:]...)

	return newArr, nil
}

func RemoveAt(values []int, index int) ([]int, error) {
	if index < 0 || index > len(values) {
		return nil, fmt.Errorf(
			"index %d out of range [0,%d]",
			index,
			len(values),
		)
	}
	newArr := make([]int, 0, len(values)-1)

	newArr = append(newArr, values[:index]...)
	newArr = append(newArr, values[index+1:]...)
	return newArr, nil
}
