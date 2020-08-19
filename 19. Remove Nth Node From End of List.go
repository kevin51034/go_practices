/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {

    fast := head
    slow := head
    for i:=0; i<n; i++ {
        fast = fast.Next
	}
	// need to consider remove head node
    if fast == nil {
        head = head.Next
        return head
    }
    
    for fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return head
}