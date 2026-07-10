/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	maxLen := -1
	var count func(node *TreeNode) int

	count = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftLen := count(node.Left)
		rightLen := count(node.Right)

		currLen := leftLen + rightLen
		if currLen > maxLen {
			maxLen = currLen
		}
		return max(leftLen, rightLen) + 1
	}
	count(root)
	return maxLen
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
