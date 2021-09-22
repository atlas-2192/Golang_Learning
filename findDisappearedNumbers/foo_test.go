package foo

import "testing"

type TestCases struct {
	i1       []int
	expected []int
}

func TestLeetcode(t *testing.T) {
	var testCases []TestCases
	testCases = append(testCases, TestCases{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}})
	testCases = append(testCases, TestCases{[]int{1, 1}, []int{2}})
	testCases = append(testCases, TestCases{[]int{2, 2, 3, 1}, []int{4}})

	for _, tc := range testCases {
		actual := findDisappearedNumbers(tc.i1)
		if !Equal(actual, tc.expected) {
			t.Errorf("findDisappearedNumbers()) == %v, expected %v, input %v", actual, tc.expected, tc.i1)
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
