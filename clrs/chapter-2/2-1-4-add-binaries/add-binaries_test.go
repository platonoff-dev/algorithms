package addbinaries

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"
)

func numToBinary(a []int) []bool {
	result := make([]bool, len(a))
	for i := range a {
		if a[i] == 1 {
			result[i] = true
		}
	}

	return result
}

func binaryToNum(a []bool) []int {
	result := make([]int, len(a))
	for i := range a {
		if a[i] {
			result[i] = 1
		}
	}

	return result
}

func TestAddBinaries(t *testing.T) {
	cases := []struct {
		name     string
		input_a  []int
		input_b  []int
		expected []int
	}{
		{
			name:     "simple_0_0",
			input_a:  []int{0},
			input_b:  []int{0},
			expected: []int{0, 0},
		},
		{
			name:     "simple_1_0",
			input_a:  []int{1},
			input_b:  []int{0},
			expected: []int{0, 1},
		},
		{
			name:     "simple_0_1",
			input_b:  []int{0},
			input_a:  []int{1},
			expected: []int{0, 1},
		},
		{
			name:     "simple_1_1",
			input_a:  []int{1},
			input_b:  []int{1},
			expected: []int{1, 0},
		},
		{
			name:     "simple_11_11",
			input_a:  []int{1, 1},
			input_b:  []int{1, 1},
			expected: []int{1, 1, 0},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := AddBinaries(numToBinary(tc.input_a), numToBinary(tc.input_b))
			if slices.Compare(binaryToNum(result), tc.expected) != 0 {
				t.Errorf("Failed:\ngot:\t%v\nwant:\t%v", binaryToNum(result), tc.expected)
			}
		})
	}
}

func generateBits(length int) []bool {
	arr := make([]bool, length)
	for i := range arr {
		arr[i] = (rand.Int() % 2) == 1
	}
	return arr
}

func generateBitsInt(length int) []int {
	arr := make([]int, length)
	for i := range arr {
		arr[i] = rand.Int() % 2
	}
	return arr
}

func BenchmarkAddBinaries(b *testing.B) {
	sizes := []int{10_000, 100_000, 1_000_000, 10_000_000}
	for _, size := range sizes {
		name := fmt.Sprintf("%s_%d", b.Name(), size)
		b.Run(name, func(b *testing.B) {
			ina := generateBits(size)
			inb := generateBits(size)
			b.ResetTimer()
			AddBinaries(ina, inb)
		})
	}
}

func BenchmarkAddBinariesAsInts(b *testing.B) {
	sizes := []int{10_000, 100_000, 1_000_000, 10_000_000}
	for _, size := range sizes {
		name := fmt.Sprintf("%s_%d", b.Name(), size)
		b.Run(name, func(b *testing.B) {
			ina := generateBitsInt(size)
			inb := generateBitsInt(size)
			b.ResetTimer()
			AddBinariesAsInts(ina, inb)
		})
	}
}
