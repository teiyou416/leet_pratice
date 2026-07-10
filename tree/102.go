/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		var currL []int
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			currL = append(currL, curr.Val)
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		result = append(result, currL)

	}
	return result
}
