package basics

func TicketPrice(age int) int {
	if age < 13 {
		return 5
	} else if age < 65 {
		return 12
	}
	return 7
}

func TotalTicketPrice(ages []int) int {
	total := 0
	for _, age := range ages {
		total += TicketPrice(age)
	}
	return total
}
