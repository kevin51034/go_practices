/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
    current := head
    for current != nil{
        for current.Next != nil && current.Val == current.Next.Val {
            current.Next = current.Next.Next 
        }
        current = current.Next
    }

    return head
}

// little optimization

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current!=nil && current.Next!=nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}