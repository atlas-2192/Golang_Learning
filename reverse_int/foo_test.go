package reverseint

import "testing"

type TestCases struct {
	input    int
	expected int
}

func TestMaxCount(t *testing.T) {
	var testCases []TestCases
	testCases = append(testCases, TestCases{123, 321})
	testCases = append(testCases, TestCases{120, 21})
	testCases = append(testCases, TestCases{-321, -123})

	for _, tc := range testCases {
		actual := reverse(tc.input)
		if actual != tc.expected {
			t.Errorf("reverse(%v) == %v, expected %v", tc.input, actual, tc.expected)
		}
	}
}

//refer https://github.com/halfrost/LeetCode-Go/blob/d5599936b7d4b7370e29a03f0218bcdcbf8d143b/leetcode/0598.Range-Addition-II/README.md
