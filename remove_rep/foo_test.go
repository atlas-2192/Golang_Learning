package removerep

import "testing"

type TestCases struct {
	input    []int
	expected int
}

func TestMaxCount(t *testing.T) {
	var testCases []TestCases
	testCases = append(testCases, TestCases{[]int{1, 2}, 2})
	testCases = append(testCases, TestCases{[]int{1, 1, 2}, 2})
	testCases = append(testCases, TestCases{[]int{0, 1, 2, 3, 4}, 5})
	testCases = append(testCases, TestCases{[]int{}, 0})
	testCases = append(testCases, TestCases{nil, 0})
	for _, tc := range testCases {
		actual := removeDuplicates(tc.input)
		if actual != tc.expected {
			t.Errorf("removeDuplicates(%v) == %v, expected %v", tc.input, actual, tc.expected)
		}
	}
}

//refer https://github.com/halfrost/LeetCode-Go/blob/d5599936b7d4b7370e29a03f0218bcdcbf8d143b/leetcode/0598.Range-Addition-II/README.md
