package exercise

func Last(a []int) int {
	switch {
	case a == nil: return -1
	case len(a) == 0: return -1
	default: return a[len(a) - 1]
	}
}
