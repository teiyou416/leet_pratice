/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	var buildTree func(left, right int) *TreeNode

	buildTree = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := left + (right-left)/2
		root := &TreeNode{Val: nums[mid]}

		root.Left = buildTree(left, mid-1)
		root.Right = buildTree(mid+1, right)

		return root
	}
	return buildTree(0, len(nums)-1)
}
