/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{Val:0}
    head := dummy
    carry := 0
    for l1 != nil || l2 != nil {
        val1 := 0
        val2 := 0
        if l1 == nil {
            val1 = 0
        } else {
            val1 = l1.Val
        }
        if l2 == nil {
            val2 = 0
        } else {
            val2 = l2.Val
        }
        val := val1 + val2 + carry
        if val >= 10 {
            carry = 1
            val = val%10
        } else {
            carry = 0
        }
        head.Next = &ListNode{Val:val}
        if l1 != nil {
            l1 = l1.Next
        }
        if l2 != nil {
            l2 = l2.Next
        }
        head = head.Next
    }
    if carry > 0 {
        head.Next = &ListNode{Val:1}
    }
    return dummy.Next
}