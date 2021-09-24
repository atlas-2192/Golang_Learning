package foo

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var cur, next, prev *ListNode
	cur = head
	prev = nil
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func reverse(pre, node, next *ListNode) {
	next.Next = node
	node.Next = pre
}

//https://leetcode.com/explore/featured/card/recursion-i/250/principle-of-recursion/1681/
