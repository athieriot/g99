package exercise

func Last(a []int) int {
	switch {
	case a == nil: return 0
	case len(a) == 0: return 0
	default: return a[len(a) - 1]
	}
}
