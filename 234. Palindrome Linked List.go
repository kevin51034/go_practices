/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func isPalindrome(head *ListNode) bool {
    if head == nil {
        return true
    }
    
    middlePoint := findMiddlePoint(head)
    tail := reverseList(middlePoint.Next)
    middlePoint.Next = nil
    
    
    for head!=nil && tail!=nil {
        if head.Val != tail.Val {
            return false
        }
        head = head.Next
        tail = tail.Next
    }
    return true
}

func reverseList(head *ListNode) *ListNode {
    var pre *ListNode
    for head!=nil {
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
    for fast!=nil && fast.Next!=nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}