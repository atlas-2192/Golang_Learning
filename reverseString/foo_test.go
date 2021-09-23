package foo

import "testing"

type TestCases struct {
	i1       []byte
	expected int
}

func TestLeetcode(t *testing.T) {
	var testCases []TestCases
	testCases = append(testCases, TestCases{[]byte{'h', 'e', 'l', 'l', 'o'}, 0})
	testCases = append(testCases, TestCases{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, 0})

	for _, tc := range testCases {
		reverseString(tc.i1)
		// if !Equal(actual, tc.expected) {
		// 	t.Errorf("moveZeroes()) == %v, expected %v", tc.i1, tc.expected)
		// }
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
