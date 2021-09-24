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

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val > root.Val {
		return searchBST(root.Right, val)
	} else if val < root.Val {
		return searchBST(root.Left, val)
	}
	return root
}

//https://leetcode.com/explore/featured/card/recursion-i/250/principle-of-recursion/1681/
