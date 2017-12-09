package exercise

func Length(a []int) int {
	if a == nil {
		return 0
	}

	var size int
	for range a {
		size++
	}

	return size
}
