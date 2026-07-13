package testingpractice

func Classify(number int) string {
	if number < 0 {
		return "negative"
	}
	if number > 0 {
		return "positive"
	}
	return "zero"
}
