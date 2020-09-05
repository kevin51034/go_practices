/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Time complexity: O(l1 + l2)
// Space complexity: O(max(l1, l2))
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    s1 := make([]int, 0)
    s2 := make([]int, 0)
    for l1!=nil {
        s1 = append(s1, l1.Val)
        l1 = l1.Next
    }
    for l2!=nil {
        s2 = append(s2, l2.Val)
        l2 = l2.Next
    }
    curHead := &ListNode{}
    sum := 0
    for len(s1)>0 || len(s2)>0 {
        if len(s1)>0 {
            sum += s1[len(s1)-1]
            s1 = s1[:len(s1)-1]
        }
        if len(s2)>0 {
            sum += s2[len(s2)-1]
            s2 = s2[:len(s2)-1]
        }
        curHead.Val = sum%10
        head := &ListNode{Val:sum/10, Next:curHead}
        curHead = head
        sum /= 10
    }
    if curHead.Val == 0 {
        return curHead.Next
    }
    return curHead
}