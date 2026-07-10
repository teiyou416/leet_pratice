/**
* Definition for singly-linked list.
* type ListNode struct {
	*     Val int
	*     Next *ListNode
	* }
*/
func isPalindrome(head *ListNode) bool {
	stack := []int{}
	for ; head != nil; head = head.Next {
		stack = append(stack, head.Val)
	}
	n := len(stack)

	for i, v := range stack {
		if v != stack[n-i-1] {
			return false
		}
	}
	return true
}
