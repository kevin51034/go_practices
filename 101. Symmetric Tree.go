/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// solution 1
// recursive

func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return check(root.Left, root.Right)
}

func check(p, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    } else if p != nil && q != nil {
        if p.Val != q.Val {
            return false 
        } else {
            return check(p.Left, q.Right) && check(p.Right, q.Left)
        }
    } else {
        return false
    }
}

// solution2
// iterative

func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root.Left)
    queue = append(queue, root.Right)
    for len(queue)>0 {
        node1 := queue[0]
        node2 := queue[1]
        queue = queue[2:]
        if node1 == nil && node2 == nil {
            continue
        }
        if node1 == nil || node2 == nil {
            return false
        }
        if node1.Val != node2.Val {
            return false
        }
        queue = append(queue, node1.Left)
        queue = append(queue, node2.Right)
        queue = append(queue, node1.Right)
        queue = append(queue, node2.Left)
    }
    return true
}