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
// BFS queue
// Time complexity: O(n)
// Space complexity: O(n)

func connect(root *Node) *Node {
    if root == nil {
        return root
    }
    queue := []*Node{root}
    for len(queue) > 0 {
        l := len(queue)

        for i:=0; i<l; i++ {
            node := queue[0]
            queue = queue[1:]
            if node != nil && i != l-1 {
                node.Next = queue[0]
            }
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }
    return root
}

// solution 2
// BFS
// Time complexity: O(n)
// Space complexity: O(1)

func connect(root *Node) *Node {
    if root == nil {
        return root
    }
    cur := root
    var nextFirst *Node
    var prev *Node
    
    for cur != nil {
        if nextFirst == nil {
            if cur.Left != nil {
                nextFirst = cur.Left
            } else if cur.Right != nil {
                nextFirst = cur.Right
            }
        }
        
        if cur.Left != nil {
            if prev != nil {
                prev.Next = cur.Left
            }
            prev = cur.Left
        }
        if cur.Right != nil {
            if prev != nil {
                prev.Next = cur.Right
            }
            prev = cur.Right
        }
        if cur.Next != nil {
            cur = cur.Next
        } else {
            cur = nextFirst
            nextFirst = nil
            prev = nil
        }
    }
    return root
}

