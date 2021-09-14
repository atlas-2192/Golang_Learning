package foo

import "testing"

type TestCases struct {
	input    []int
	expected int
}

func TestMaxCount(t *testing.T) {
	var testCases []TestCases
	testCases = append(testCases, TestCases{[]int{555, 901, 482, 1771}, 1})
	testCases = append(testCases, TestCases{[]int{12, 345, 2, 6, 7896}, 2})
	testCases = append(testCases, TestCases{[]int{00, 123, 12}, 1})
	testCases = append(testCases, TestCases{[]int{0}, 0})
	testCases = append(testCases, TestCases{[]int{1}, 0})
	testCases = append(testCases, TestCases{[]int{}, 0})

	for _, tc := range testCases {
		actual := findNumbers(tc.input)
		if actual != tc.expected {
			t.Errorf("findMaxConsecutiveOnes(%v) == %v, expected %v", tc.input, actual, tc.expected)
		}
	}
}

//refer https://github.com/halfrost/LeetCode-Go/blob/d5599936b7d4b7370e29a03f0218bcdcbf8d143b/leetcode/0598.Range-Addition-II/README.md
