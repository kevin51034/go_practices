/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Time complexity: O(n)
// Space complexity: O(n)

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    return max(maxDepth(root.Left), maxDepth(root.Right)) +1
}

func max(x int, y int) int {
    if x>y {
        return x
    } else{
        return y
    }
}