/**
* Definition for singly-linked list.
* type ListNode struct {
	*     Val int
	*     Next *ListNode
	* }
*/
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

	}

	head2 := slow.Next
	slow.Next = nil

	p1 := head
	head2 = reverse(head2)
	p2 := head2
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return true

}
func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
