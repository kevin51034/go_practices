/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    tmp := head.Next.Next
    head.Next.Next = head
    next := head.Next
    head.Next = swapPairs(tmp)
    return next
}