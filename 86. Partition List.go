/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// create a new list to store number greater than x
// and concatenate two list afterward

func partition(head *ListNode, x int) *ListNode {
    smallDummy := &ListNode{Val:0, Next:head}
    largeDummy := &ListNode{Val:0}
    head = smallDummy
    tail := largeDummy
    
    for head.Next != nil {
        if head.Next.Val < x {
            head = head.Next
        } else {
            tail.Next = head.Next
            tail = tail.Next
            head.Next = head.Next.Next
        }
    }
    tail.Next = nil
    head.Next = largeDummy.Next
    return smallDummy.Next
    
}