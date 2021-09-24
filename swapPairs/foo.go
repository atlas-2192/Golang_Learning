package foo

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	root := &ListNode{Next: head}
	for prev, node := root, root.Next; node != nil; prev, node = node, node.Next {
		if node.Next == nil {
			break
		}
		swap(prev, node, node.Next)
	}
	return root.Next
}

func swap(pre, node, next *ListNode) {
	node.Next = next.Next
	next.Next = node
	pre.Next = next
}

//https://leetcode.com/explore/featured/card/recursion-i/250/principle-of-recursion/1681/
