package exercise

func Reverse(a []int) []int {
	if a == nil {
		return nil
	}

	reversed := make([]int, 0)
	for i := len(a) -1; i >= 0; i-- {
		reversed = append(reversed, a[i])
	}

	return reversed
}
