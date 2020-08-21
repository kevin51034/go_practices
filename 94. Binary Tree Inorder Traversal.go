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

func inorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    if root == nil {
        return result
    }
    
    stack := make([]*TreeNode, 0)
    for len(stack)>0 || root!=nil {
        for root!=nil {
            stack = append(stack, root)
            root = root.Left
        }
    
        result = append(result, stack[len(stack)-1].Val)
        root = stack[len(stack)-1].Right
        stack = stack[:len(stack)-1]
    }
    
    return result
}