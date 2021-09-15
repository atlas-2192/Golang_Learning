package foo

import "testing"

type TestCases struct {
	i1       []int
	val      int
	expected []int
	ext_val  int
}

func TestLeetcode(t *testing.T) {
	var testCases []TestCases
	testCases = append(testCases, TestCases{[]int{3, 2, 2, 3}, 3, []int{2, 2, 0, 0}, 2})
	testCases = append(testCases, TestCases{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{2, 2, 0, 0}, 5})

	for _, tc := range testCases {
		actual := removeElement(tc.i1, tc.val)
		if actual != tc.ext_val {
			t.Errorf("removeElement(%v, %v)) == %v, expected %v", tc.i1, tc.val, actual, tc.ext_val)
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
