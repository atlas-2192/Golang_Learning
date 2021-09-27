package foo

import "testing"

type TestCases struct {
	i1       int
	expected int
}

func TestLeetcode(t *testing.T) {
	var testCases []TestCases
	testCases = append(testCases, TestCases{1, 1})
	testCases = append(testCases, TestCases{0, 1})
	testCases = append(testCases, TestCases{2, 2})
	testCases = append(testCases, TestCases{3, 3})
	testCases = append(testCases, TestCases{3, 3})

	for _, tc := range testCases {
		actual := climbStairs(tc.i1)
		if actual != tc.expected {
			t.Errorf("reverseList()) == %v, expected %v", tc.i1, tc.expected)
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
