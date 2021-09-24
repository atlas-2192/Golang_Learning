package foo

import "testing"

type TestCases struct {
	i1       *ListNode
	expected *ListNode
}

func TestLeetcode(t *testing.T) {
	var testCases []TestCases
	testCases = append(testCases, TestCases{
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil}}}},
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  3,
						Next: nil}}}}})

	for _, tc := range testCases {
		actual := swapPairs(tc.i1)
		if !EqualNote(actual, tc.expected) {
			t.Errorf("swapPairs()) == %v, expected %v", tc.i1, tc.expected)
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
