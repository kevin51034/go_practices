/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return
    }
    middle := findMiddlePoint(head)
    
    tail := middle.Next
    middle.Next = nil
    
    reverseTail := reverseList(tail)
    
    result := mergeTwoList(head, reverseTail)
    head = result
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{Val:0}
    head := dummy
    toggle := true
    
    for l1!=nil && l2!=nil {
        if toggle {
            head.Next = l1
            l1 = l1.Next
        } else {
            head.Next = l2
            l2 = l2.Next
        }
        head = head.Next
        toggle = !toggle
    }
    
    for l1 != nil {
        head.Next = l1
        head = head.Next
        l1 = l1.Next
    }
    for l2 != nil {
        head.Next = l2
        head = head.Next
        l2 = l2.Next
    }
    return dummy.Next
}

func reverseList(head *ListNode) *ListNode {
    var pre *ListNode
    
    for head != nil {
        next := head.Next
        head.Next = pre
        pre = head
        head = next
    }
    return pre
} 

func findMiddlePoint(head *ListNode) *ListNode {
    fast := head.Next
    slow := head
    for fast.Next!=nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}