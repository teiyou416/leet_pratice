/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var check func(u, v *TreeNode) bool
	check = func(u, v *TreeNode) bool {
		if u == nil && v == nil {
			return true
		}
		if u == nil || v == nil || u.Val != v.Val {
			return false
		}
		return check(u.Left, v.Right) && check(u.Right, v.Left)
	}
	return check(root.Left, root.Right)
}

