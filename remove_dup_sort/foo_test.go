package foo

import "testing"

type TestCases struct {
	i1       []int
	expected int
}

func TestLeetcode(t *testing.T) {
	var testCases []TestCases
	testCases = append(testCases, TestCases{[]int{1, 1, 2}, 2})
	testCases = append(testCases, TestCases{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5})
	testCases = append(testCases, TestCases{[]int{}, 0})
	testCases = append(testCases, TestCases{[]int{1, 2, 3}, 3})

	for _, tc := range testCases {
		actual := removeDuplicates(tc.i1)
		if actual != tc.expected {
			t.Errorf("removeDuplicates()) == %v, expected %v", tc.i1, tc.expected)
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
