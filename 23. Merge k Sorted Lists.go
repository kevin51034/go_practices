/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


// solution 1 : Merge with Divide And Conquer
// Time complexity: O(Nlogk) where N is the total number of the nodes and k is the number of linked lists.
// Space complexity: O(1),  O(logk) for call stack

// iteration
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    for len(lists) > 1 {
        mergeList := mergeTwoList(lists[0], lists[1])
        lists = lists[2:]
        lists = append(lists, mergeList)
    }
    return lists[0]
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    cur := dummy
    for l1!=nil && l2!=nil {
        if l1.Val <= l2.Val {
            cur.Next = l1
            l1 = l1.Next
        } else {
            cur.Next = l2
            l2 = l2.Next
        }
        cur = cur.Next
    }
    if l1!=nil {
        cur.Next = l1
    }
    if l2!=nil {
        cur.Next = l2
    }
    return dummy.Next
}


// recursion
func mergeKLists(lists []*ListNode) *ListNode {
    return merge(0, len(lists)-1, lists)
}

func merge(l, r int, lists []*ListNode) *ListNode {
    if l>r {
        return nil
    }
    if l==r {
        return lists[l]
    }
    mid := l + (r-l)/2
    left := merge(l, mid, lists)
    right := merge(mid+1, r, lists)
    return mergeTwoList(left, right)
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    cur := dummy
    for l1!=nil && l2!=nil {
        if l1.Val <= l2.Val {
            cur.Next = l1
            l1 = l1.Next
        } else {
            cur.Next = l2
            l2 = l2.Next
        }
        cur = cur.Next
    }
    if l1!=nil {
        cur.Next = l1
    }
    if l2!=nil {
        cur.Next = l2
    }
    return dummy.Next
}
