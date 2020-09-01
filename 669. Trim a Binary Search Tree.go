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

func trimBST(root *TreeNode, L int, R int) *TreeNode {
    if root == nil {
        return root
    }
    if root.Val > R {
        return trimBST(root.Left, L, R)
    } else if root.Val < L {
        return trimBST(root.Right, L, R)
    }
    root.Left = trimBST(root.Left, L, R)
    root.Right = trimBST(root.Right, L, R)
    return root
}