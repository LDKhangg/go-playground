package loops

func SumUntil(values []int, stopAt int) int {
	total := 0

	for _, value := range values {
		if value != stopAt {
			total += value
		} else {
			return total
		}
	}
	return total
}
