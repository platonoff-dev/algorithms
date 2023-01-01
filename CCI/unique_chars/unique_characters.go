package unique_chars

//Implement an algorithm to determine if a string has all unique characters. What if you can not use data structures.
func isUniqueBinary(s string) bool {
	// Let's suppose that empty string returns true
	if len(s) == 0 {
		return true
	}

	// Converting to list of runes required only if support of unicode is required
	// chars := []rune(s)
	chars := s
	res := chars[0]
	for i := 1; i < len(chars); i++ {
		res = res ^ chars[i]
	}

	return res != 0
}

// Disadvantage of solution: it doesn't support unicode
func isUniqueArray(s string) bool {
	counts := make([]int, 256)
	for i := 0; i < len(s); i++ {
		counts[s[i]] += 1
	}

	for _, count := range counts {
		if count > 1 {
			return false
		}
	}

	return true
}

// Support unicode but slow
func isUniqueMap(s string) bool {
	counts := map[rune]int{}

	for _, ch := range s {
		counts[ch] += 1
	}

	for _, count := range counts {
		if count > 1 {
			return false
		}
	}

	return true
}
