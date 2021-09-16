package foo

import "testing"

type TestCases struct {
	i1       []int
	expected bool
}

func TestLeetcode(t *testing.T) {
	var testCases []TestCases
	testCases = append(testCases, TestCases{[]int{2, 1}, false})
	testCases = append(testCases, TestCases{[]int{3, 5, 5}, false})
	testCases = append(testCases, TestCases{[]int{}, false})
	testCases = append(testCases, TestCases{[]int{0, 3, 2, 1}, true})
	testCases = append(testCases, TestCases{[]int{0, 2, 3, 4, 5, 2, 1, 0}, true})

	for _, tc := range testCases {
		actual := validMountainArray(tc.i1)
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
