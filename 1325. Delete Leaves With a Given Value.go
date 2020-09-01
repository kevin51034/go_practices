/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//Time complexity: O(n)
//Space complexity: O(n)

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    if root == nil {
        return root
    }
    root.Left = removeLeafNodes(root.Left, target)
    root.Right = removeLeafNodes(root.Right, target)
    if root.Left == nil && root.Right == nil && root.Val == target {
        return nil
    }
    return root
}