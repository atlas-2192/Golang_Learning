package foo

import "testing"

type TestCases struct {
	i1       []int
	expected bool
}

func TestLeetcode(t *testing.T) {
	var testCases []TestCases
	testCases = append(testCases, TestCases{[]int{10, 2, 5, 3}, true})
	testCases = append(testCases, TestCases{[]int{7, 1, 14, 11}, true})
	testCases = append(testCases, TestCases{[]int{}, false})
	testCases = append(testCases, TestCases{[]int{0, 0}, true})
	testCases = append(testCases, TestCases{[]int{3, 1, 7, 11}, false})

	for _, tc := range testCases {
		actual := checkIfExist(tc.i1)
		if actual != tc.expected {
			t.Errorf("checkIfExist()) == %v, expected %v", tc.i1, tc.expected)
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
