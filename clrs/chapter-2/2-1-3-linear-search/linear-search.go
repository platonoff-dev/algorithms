package linearsearch

func LinearSearch(arr []int, x int) int {
	for i, item := range arr {
		if item == x {
			return i
		}
	}

	return -1
}
