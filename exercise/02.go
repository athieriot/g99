package exercise

func Penultimate(a []int) int {
	if len(a) > 1 {
		return Last(a[:len(a) - 1])
	}

	return -1
}