/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    var pre *ListNode
    
    for head != nil {
        post := head.Next
        head.Next = pre
        pre = head
        head = post
    }
    // just return pre because it's head after reversing
    return pre
}