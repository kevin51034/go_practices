/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// time complexity: O(n)
// space complexity: O(1) // O(n) for recursion stack
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    tmp := root.Left
    root.Left = invertTree(root.Right)
    root.Right = invertTree(tmp)
    return root
}