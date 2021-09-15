package foo

import "testing"

type TestCases struct {
	i1       *[]int
	m        int
	i2       []int
	n        int
	expected []int
}

func TestLeetcode(t *testing.T) {
	var testCases []TestCases
	testCases = append(testCases, TestCases{&[]int{1, 2, 3}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}})
	testCases = append(testCases, TestCases{&[]int{1}, 1, []int{}, 0, []int{1}})

	for _, tc := range testCases {
		merge(*tc.i1, tc.m, tc.i2, tc.n)
		if !Equal(*tc.i1, tc.expected) {
			t.Errorf("merge()) == %v, expected %v", tc.i1, tc.expected)
		}
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
