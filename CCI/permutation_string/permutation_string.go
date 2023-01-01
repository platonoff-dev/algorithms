package permutation_string

// Check if one string is permutation of another

func isPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	counts := make([]int, 256)
	for _, ch := range []byte(s1) {
		counts[ch] += 1
	}

	for _, ch := range []byte(s2) {
		counts[ch] -= 1
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}
