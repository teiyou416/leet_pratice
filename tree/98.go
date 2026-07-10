/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	var check func(node *TreeNode, min, max *int) bool
	check = func(node *TreeNode, min, max *int) bool {
		if node == nil {
			return true
		}

		if min != nil && node.Val <= *min {
			return false
		}
		if max != nil && node.Val >= *max {
			return false
		}
		return check(node.Left, min, &node.Val) && check(node.Right, &node.Val, max)
	}
	return check(root, nil, nil)
}
