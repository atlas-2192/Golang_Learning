package foo

import "testing"

type TestCases struct {
	i1       *TreeNode
	expected int
}

func TestLeetcode(t *testing.T) {
	var testCases []TestCases
	testCases = append(testCases, TestCases{
		i1: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   15,
					Left:  nil,
					Right: nil},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil}}},
		expected: 2})
	testCases = append(testCases, TestCases{
		i1: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil}},
		expected: 1})
	for _, tc := range testCases {
		actual := maxDepth(tc.i1)
		if actual != tc.expected {
			t.Errorf("maxDepth()) == %v, expected %v", actual, tc.expected)
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

func EqualNote(a, b *ListNode) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}

	if a.Val == b.Val {
		if a.Next == nil && b.Next == nil {
			return true
		}
		return EqualNote(a.Next, b.Next)
	}

	return false
}
