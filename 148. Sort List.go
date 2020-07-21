/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// use metge sort

func sortList(head *ListNode) *ListNode {
    result := mergeSort(head)
    return result
}

func mergeSort(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    middle := findMiddlePoint(head)
    
    tail := middle.Next
    middle.Next = nil
    
    left := mergeSort(head)
    right := mergeSort(tail)
    
    result := mergeTwoList(left, right)
    return result
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{Val:0}
    head := dummy
    
    for l1!=nil && l2!=nil {
        if l1.Val < l2.Val {
            head.Next = l1
            l1 = l1.Next
        } else {
            head.Next = l2
            l2 = l2.Next
        }
        head = head.Next
    }
    
    for l1!=nil {
        head.Next = l1
        l1 = l1.Next
        head = head.Next
    }
    for l2!=nil {
        head.Next = l2
        l2 = l2.Next
        head = head.Next
    }
    
    return dummy.Next
}

func findMiddlePoint(head *ListNode) *ListNode {
    fast := head.Next
    slow := head
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}


