/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 0 {
        return head
    }
    count := 0
    last := head
    cur := head
    for cur != nil {
        count++
        last = cur
        cur = cur.Next
    }
    k = k%count
    if k == 0 {
        return head
    }
    cur = head
    for i:=0; i<count-k-1; i++ {
        cur = cur.Next
    }
    tmp := cur.Next
    cur.Next = nil
    last.Next = head
    
    return tmp
}