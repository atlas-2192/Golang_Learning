package foo

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// var depth int
// var MDepth int

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Left)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

//https://leetcode.com/explore/featured/card/recursion-i/250/principle-of-recursion/1681/
