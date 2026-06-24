/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{Val:0,Next:head}    
    prev := dummy
    for prev.Next != nil && prev.Next.Next != nil{
	    node1 ,node2 := prev.Next,prev.Next.Next

	    prev.Next = node2
	    node1.Next = node2.Next 
	    node2.Next = node1


	    prev = node1

    }
    return dummy.Next
}
