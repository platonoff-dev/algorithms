package mergesort

func merge(arr []int, left, middle, right int) {
	l1 := make([]int, middle-left+1)
	l2 := make([]int, right-middle)

	for i := 0; i < len(l1); i++ {
		l1[i] = arr[left+i]
	}

	for i := 0; i < len(l2); i++ {
		l2[i] = arr[middle+i+1]
	}

	i := 0
	j := 0
	for k := left; k <= right; k++ {
		if i == len(l1) {
			arr[k] = l2[j]
			j++
			continue
		}

		if j == len(l2) {
			arr[k] = l1[i]
			i++
			continue
		}

		if l1[i] <= l2[j] {
			arr[k] = l1[i]
			i++
		} else {
			arr[k] = l2[j]
			j++
		}
	}
}

func mergeSort(arr []int, left, right int) {
	if left < right {
		middle := (left + right) / 2
		mergeSort(arr, left, middle)
		mergeSort(arr, middle+1, right)
		merge(arr, left, middle, right)
	}
}

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}
