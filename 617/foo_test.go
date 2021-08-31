package main

import "testing"

type TestCases struct {
	m, n     int
	input    [][]int
	expected int
}

func TestMaxCount(t *testing.T) {
	var testCases []TestCases
	testCases = append(testCases, TestCases{3, 3, [][]int{{2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}}, 4})
	testCases = append(testCases, TestCases{3, 3, [][]int{}, 9})
	testCases = append(testCases, TestCases{3, 3, [][]int{{2, 2}, {3, 3}}, 4})

	for _, tc := range testCases {
		actual := maxCount(tc.m, tc.n, tc.input)
		if actual != tc.expected {
			t.Errorf("maxCount(%v) == %v, expected %v", tc.input, actual, tc.expected)
		}
	}
}

//refer https://github.com/halfrost/LeetCode-Go/blob/d5599936b7d4b7370e29a03f0218bcdcbf8d143b/leetcode/0598.Range-Addition-II/README.md
