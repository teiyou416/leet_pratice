
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil && head.Next == nil {
		return
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	head2 := slow.Next
	slow.Next = nil
	head2 = reverse(head2)

	p1 := head
	p2 := head2
	for p2 != nil {
		next1 := p1.Next
		p1.Next = p2
		p1 = next1
		next2 := p2.Next
		p2.Next = p1
		p2 = next2
	}

}
func reverse(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode = nil
	for curr != nil {
		next := curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}
	return prev
}
