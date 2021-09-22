package foo

import "testing"

type TestCases struct {
	i1       []int
	expected int
}

func TestLeetcode(t *testing.T) {
	var testCases []TestCases
	testCases = append(testCases, TestCases{[]int{3, 2, 1}, 1})
	testCases = append(testCases, TestCases{[]int{1, 2}, 2})
	testCases = append(testCases, TestCases{[]int{2, 2, 3, 1}, 1})
	testCases = append(testCases, TestCases{[]int{1}, 1})
	testCases = append(testCases, TestCases{[]int{}, 0})

	for _, tc := range testCases {
		actual := thirdMax(tc.i1)
		if actual != tc.expected {
			t.Errorf("thirdMax()) == %v, expected %v, input %v", actual, tc.expected, tc.i1)
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
