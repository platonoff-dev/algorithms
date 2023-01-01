package permutation_string

import (
	"fmt"
	"testing"
)

var cases = []struct {
	input    []string
	expected bool
}{
	{[]string{"", ""}, true},
	{[]string{"a", "a"}, true},
	{[]string{"a", "b"}, false},
	{[]string{"aaa", "aaa"}, true},
	{[]string{"abc", "bca"}, true},
	{[]string{"abc", "bcc"}, false},
	{[]string{"abc", "bccc"}, false},
}

func TestIsPermutation(t *testing.T) {
	for i, tc := range cases {
		name := fmt.Sprintf("[%d]", i)
		t.Run(name, func(t *testing.T) {
			res := isPermutation(tc.input[0], tc.input[1])
			if res != tc.expected {
				t.Errorf("isPermutation(%s, %s) result %v, expected %v", tc.input[0], tc.input[1], res, tc.expected)
			}
		})
	}
}
