package unique_chars

import "testing"

var cases = []struct {
	input    string
	expected bool
}{
	{"", true},
	{"ab", true},
	{"aa", false},
	{"\x00\x00", false},
	{" !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~", true},
}

func TestUniqueCharBinary(t *testing.T) {
	for _, tc := range cases {
		actual := isUniqueBinary(tc.input)
		if actual != tc.expected {
			t.Fatalf("IsUniqueBinary(%s) expected %v, got %v", tc.input, tc.expected, actual)
		}
	}
}

func BenchmarkUniqueCharBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isUniqueBinary(cases[4].input)
	}
}

func TestUniqueCharArray(t *testing.T) {
	for _, tc := range cases {
		actual := isUniqueArray(tc.input)
		if actual != tc.expected {
			t.Fatalf("IsUniqueArray(%s) expected %v, got %v", tc.input, tc.expected, actual)
		}
	}
}

func BenchmarkUniqueCharArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isUniqueArray(cases[4].input)
	}
}

func TestUniqueCharMap(t *testing.T) {
	for _, tc := range cases {
		actual := isUniqueMap(tc.input)
		if actual != tc.expected {
			t.Fatalf("IsUniqueArray(%s) expected %v, got %v", tc.input, tc.expected, actual)
		}
	}
}

func BenchmarkUniqueCharMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isUniqueMap(cases[4].input)
	}
}
