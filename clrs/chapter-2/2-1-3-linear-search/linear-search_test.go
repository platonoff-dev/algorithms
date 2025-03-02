package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		x    int
		want int
	}{
		{"Edge left", []int{1, 2, 3}, 1, 0},
		{"Regular", []int{1, 2, 3}, 2, 1},
		{"Edge right", []int{1, 2, 3}, 3, 2},
		{"Not in array", []int{1, 2, 3}, 4, -1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LinearSearch(tc.arr, tc.x)
			if got != tc.want {
				t.Errorf("got: %d, want: %d", got, tc.want)
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

func BenchmarkLinearSearch(b *testing.B) {
	elements := []int{10, 100, 1000, 10000, 100000}
	for _, n := range elements {
		name := fmt.Sprintf("%s_%d", b.Name(), n)
		b.Run(name, func(b *testing.B) {
			arr := generateRandomSlice(n)
			target := arr[len(arr)-1]

			b.ResetTimer() // Reset timer to exclude setup time
			for i := 0; i < b.N; i++ {
				_ = LinearSearch(arr, target)
			}
		})
	}
}
