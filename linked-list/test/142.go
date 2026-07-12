package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast

		}
	}
	return nil
}

func main() {

}
