package branching

func ClassifyScore(score int) string {
	switch {
	case score < 0:
		return "invalid"
	case score < 50:
		return "fail"
	case score < 85:
		return "pass"
	case score < 101:
		return "excellent"
	default:
		return "invalid"
	}
}
