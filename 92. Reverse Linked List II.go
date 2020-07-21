/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseBetween(head *ListNode, m int, n int) *ListNode {
    dummy := &ListNode{Val:0, Next:head}
    head = dummy
    for count:=0; head.Next!=nil && count<m-1; count++ {
        head = head.Next
    }
    
    current := head.Next
    for i:=0; i<n-m;i++ {
        tmp := head.Next
        head.Next = current.Next
        current.Next = current.Next.Next
        head.Next.Next = tmp
    }
    return dummy.Next
}