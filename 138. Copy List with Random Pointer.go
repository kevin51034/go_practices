/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

// use clone solution
// Time complexity: O(n)
// Space complexity: O(1)


func copyRandomList(head *Node) *Node {
    if head == nil {
        return head
    }
    cur := head
    for cur != nil {
        tmp := cur.Next
        clone := &Node{Val:cur.Val, Next:cur.Next}
        cur.Next = clone
        cur = tmp
    }
    
    cur = head
    
    for cur!=nil {
        if cur.Random != nil {
			// need to point to cur.Random.Next clone
			// or it will point to original list
			// cur.Random is original list cur.Random.Next is clone
            cur.Next.Random = cur.Random.Next
        }
        cur = cur.Next.Next
    }
    
    cur = head
    cloneHead := head.Next
    for cur != nil && cur.Next != nil {
        tmp := cur.Next
        cur.Next = cur.Next.Next
        cur = tmp
    }
    
    return cloneHead

}