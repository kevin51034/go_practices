/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// use Floyd's cycle detection algorithm
// see https://cs.stackexchange.com/questions/10360/floyds-cycle-detection-algorithm-determining-the-starting-point-of-cycle
// above link explain why this method works

func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    fast := head.Next
    slow := head
    
    for fast!=nil && fast.Next!= nil {
        if fast == slow {
            slow = head
            fast = fast.Next
            for slow != fast {
                fast = fast.Next
                slow = slow.Next
            }
            return fast
		}
		fast = fast.Next.Next
        slow = slow.Next
    }
    return nil
}