/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	count := 0
	result := 0
	found := false
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || found {
			return
		}
		dfs(node.Left)

		if !found {
			count++
			if count == k {
				result = node.Val
				found = true
				return
			}
		}
		dfs(node.Right)
	}
	dfs(root)
	return result
}	
