/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    evenhead := head.Next
    cur := head
    toggle := true
    lastNode := head
    for cur != nil && cur.Next != nil {
        tmp := cur.Next
        cur.Next = cur.Next.Next
        if toggle {
            if cur.Next != nil {
                lastNode = cur.Next
            }
        }
        toggle = !toggle
        cur = tmp
    }
    lastNode.Next = evenhead
    return head
}

// solution
// using two list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func oddEvenList(head *ListNode) *ListNode {
    oddhead := &ListNode{Val:0, Next:head}
    odd := oddhead
    evenhead := &ListNode{Val:0, Next:head}
    even := evenhead
    toggle := true
    for head != nil {
        if toggle {
            odd.Next = head
            odd = odd.Next
        } else {
            even.Next = head
            even = even.Next
        }
        toggle = !toggle
        head = head.Next
    }
    even.Next = nil
    odd.Next = evenhead.Next
    return oddhead.Next
}