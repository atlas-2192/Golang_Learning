package foo

import (
	"math"
	"testing"
)

type TestCases struct {
	i1       float64
	i2       int
	expected float64
}

func TestLeetcode(t *testing.T) {
	var testCases []TestCases
	// testCases = append(testCases, TestCases{2.00000, 10, 1024.00000})
	// testCases = append(testCases, TestCases{2.10000, 3, 9.26100})
	// testCases = append(testCases, TestCases{2.00000, -2, 0.25000})
	testCases = append(testCases, TestCases{0.00001, 2147483647, 0.25000})

	for _, tc := range testCases {
		actual := myPow(tc.i1, tc.i2)
		if !almostEqual(actual, tc.expected) {
			t.Errorf("checkIfExist()) == %v, expected %v", actual, tc.expected)
		}
	}
}

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
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
