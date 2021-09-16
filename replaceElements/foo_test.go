package foo

import "testing"

type TestCases struct {
	i1       []int
	expected []int
}

func TestLeetcode(t *testing.T) {
	var testCases []TestCases
	testCases = append(testCases, TestCases{[]int{17, 18, 5, 4, 6, 1}, []int{18, 6, 6, 6, 1, -1}})
	testCases = append(testCases, TestCases{[]int{400}, []int{-1}})

	for _, tc := range testCases {
		actual := replaceElements(tc.i1)
		if !Equal(actual, tc.expected) {
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
