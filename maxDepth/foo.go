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

var depth int
var MDepth int

func maxDepth(root *TreeNode) int {
	if depth > MDepth {
		MDepth = depth
	}

	if root.Left != nil {
		depth++
		maxDepth(root.Left)
	}

	if root.Right != nil {
		depth++
		maxDepth(root.Right)
	}

	depth--

	if root == nil {
		return 0
	}
	return MDepth + 1
}

//https://leetcode.com/explore/featured/card/recursion-i/250/principle-of-recursion/1681/
