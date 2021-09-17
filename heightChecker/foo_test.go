package foo

import "testing"

type TestCases struct {
	i1       []int
	expected int
}

func TestLeetcode(t *testing.T) {
	var testCases []TestCases
	testCases = append(testCases, TestCases{[]int{1, 1, 4, 2, 1, 3}, 3})
	testCases = append(testCases, TestCases{[]int{5, 1, 2, 3, 4}, 5})
	testCases = append(testCases, TestCases{[]int{1, 2, 3, 4, 5}, 0})
	testCases = append(testCases, TestCases{[]int{1}, 0})
	testCases = append(testCases, TestCases{[]int{}, 0})

	for _, tc := range testCases {
		actual := heightChecker(tc.i1)
		if actual != tc.expected {
			t.Errorf("heightChecker()) == %v, expected %v", actual, tc.expected)
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
