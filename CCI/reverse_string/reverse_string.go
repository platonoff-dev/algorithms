package reverse_string

// Reverse a null terminated string

func reverse(s string) string {
	bytes := []byte(s)
	left := 0
	right := len(s) - 2
	for left < right {
		tmp := bytes[left]
		bytes[left] = bytes[right]
		bytes[right] = tmp
		left++
		right--
	}

	return string(bytes)
}
