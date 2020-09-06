/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func insertionSortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    dummy := &ListNode{}
    cur := head
    prev := dummy
    for cur != nil {
        next := cur.Next
        for prev.Next!=nil && prev.Next.Val < cur.Val {
            prev = prev.Next
        }
        cur.Next = prev.Next
        prev.Next = cur
        cur = next
        prev = dummy
    }
    return dummy.Next
}