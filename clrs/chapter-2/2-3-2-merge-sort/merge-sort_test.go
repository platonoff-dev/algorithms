package mergesort

import (
	"math/rand"
	"slices"
	"strconv"
	"testing"
)

func TestMergeSort(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Simple",
			input:    []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "With repeats",
			input:    []int{3, 3, 2, 2, 1, 1},
			expected: []int{1, 1, 2, 2, 3, 3},
		},
		{
			name:     "More complex",
			input:    []int{1, 3, 2, 999, 5},
			expected: []int{1, 2, 3, 5, 999},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			MergeSort(tc.input)
			if !slices.Equal(tc.expected, tc.input) {
				t.Errorf("got:\t%v\nwant:\t%v", tc.input, tc.expected)
			}
		})
	}
}

func generateRandomSlice(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(100000)
	}
	return arr
}

func BenchmarkSelectionSort(b *testing.B) {
	sizes := []int{10, 100, 1000, 10_000, 100_000, 1_000_000, 10_000_000}
	for _, s := range sizes {
		b.Run(strconv.Itoa(s), func(b *testing.B) {
			arr := generateRandomSlice(s)
			b.ResetTimer()
			MergeSort(arr)
		})
	}
}
