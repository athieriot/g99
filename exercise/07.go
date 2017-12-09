package exercise

func Flatten(a [][]int) []int {
	result := make([]int, 0)

	for _, l := range a {
		for _, v := range l {
			result = append(result, v)
		}
	}

	return result
}