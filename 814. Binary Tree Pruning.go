/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Time complexity: O(n)
// Space complexity: O(h)

func pruneTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    root.Left = pruneTree(root.Left)
    root.Right = pruneTree(root.Right)
    if root.Val == 1 || root.Left != nil || root.Right != nil {
        return root
    }
    return nil
}