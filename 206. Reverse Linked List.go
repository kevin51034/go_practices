/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    var prev *ListNode
    for head!=nil {
        tmp := head.Next
        head.Next = prev
        prev = head
        head = tmp
    }
    // just return prev because it's head after reversing
    return prev
}

// recursive
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := reverse(head)
    head.Next = nil
    return newHead
}

func reverse(cur *ListNode) *ListNode {
    if cur == nil {
        return cur
    }
    newHead := reverse(cur.Next)
    if newHead == nil {
        return cur
    }
    cur.Next.Next = cur
    return newHead
}