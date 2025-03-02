package main

// 0 0 0        -> 0 0
// 0 1 1        -> 1 0
// 0 1 0, 0 0 1 -> 0 1
// 1 1 1        -> 1 1

//    [1, 1, 1]
//    [1, 1, 1]
// [1, 1, 1, 1]

func xor(a, b bool) bool {
	return (a || b) && !(a && b)
}

// 0 c1 -> c1 0
// 1 c1 -> c1 1

func AddBinaries(a []bool, b []bool) []bool {
	c := make([]bool, len(a)+1)

	carry := false
	for i := len(a) - 1; i >= 0; i-- {
		c[i+1] = xor(xor(a[i], b[i]), carry)
		carry = a[i] && b[i] || carry
	}
	c[0] = carry

	return c
}

// Experiment: Operate on numbers instead of booleans to compare performance
func AddBinariesAsInts(a []int, b []int) []int {
	c := make([]int, len(a)+1)

	carry := 0
	for i := len(a) - 1; i >= 0; i-- {
		c[i+1] = (a[i] + b[i] + carry) % 2
		carry = (a[i] + b[i] + carry) / 2
	}
	c[0] = carry

	return c
}
