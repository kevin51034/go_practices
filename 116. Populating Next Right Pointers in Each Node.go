/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// solution 1
// BFS + queue
// Time complexity: O(n)
// Space complexity: O(n)

func connect(root *Node) *Node {
    if root == nil {
        return root
    }
    queue := make([]*Node, 0)
    queue = append(queue, root)
    
    for len(queue) > 0 {
        l := len(queue)
        
        for i:=0; i<l; i++ {
            node := queue[0]
            queue = queue[1:]
            if node.Left == nil {
                continue
            }
            node.Left.Next = node.Right
            if node.Next != nil {
                node.Right.Next = node.Next.Left
            }
            queue = append(queue, node.Left)
            queue = append(queue, node.Right)
        }
    }
    return root
}

// solution 2
// recursive
// Time complexity: O(n)
// Space complexity: O(1)

func connect(root *Node) *Node {
    if root == nil || root.Left == nil {
        return root
    }
    root.Left.Next = root.Right
    if root.Next != nil {
        root.Right.Next = root.Next.Left
    }
    connect(root.Left)
    connect(root.Right)
    return root
}