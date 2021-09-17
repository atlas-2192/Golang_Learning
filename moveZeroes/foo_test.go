package foo

import "testing"

type TestCases struct {
	i1       []int
	expected []int
}

func TestLeetcode(t *testing.T) {
	var testCases []TestCases
	testCases = append(testCases, TestCases{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}})

	for _, tc := range testCases {
		moveZeroes(tc.i1)
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
