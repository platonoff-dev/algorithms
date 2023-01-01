package reverse_string

import (
	"fmt"
	"testing"
)

var tests = []struct {
	input    string
	expected string
}{
	{"\x00", "\x00"},
	{"ab\x00", "ba\x00"},
	{"aa\x00", "aa\x00"},
}

func TestRevers(t *testing.T) {
	for i, test := range tests {
		name := fmt.Sprintf("[%d]", i)
		t.Run(name, func(t *testing.T) {
			res := reverse(test.input)
			if res != test.expected {
				t.Errorf("reverse(%s) result %s, expected %s", test.input, res, test.expected)
			}
		})
	}
}
