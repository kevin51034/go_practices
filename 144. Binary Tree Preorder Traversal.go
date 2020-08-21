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

func preorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    travel(root, &result)
    return result
}

func travel(root *TreeNode, result * []int) {
    if root == nil {
        return
    }
    *result = append(*result, root.Val)
    travel(root.Left, result)
    travel(root.Right, result)
} 

// solution 2
// iterative

func preorderTraversal(root *TreeNode) []int {
    stack := make([]*TreeNode, 0)
    result := make([]int, 0)
    if root == nil {
        return result
    }
    stack = append(stack, root)
    for len(stack) != 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if node != nil {
            result = append(result, node.Val)
        }
        if node.Right != nil {
            stack = append(stack, node.Right)
        }
        if node.Left != nil {
            stack = append(stack, node.Left)
        }
    }
    return result
}