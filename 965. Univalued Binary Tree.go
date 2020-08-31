/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isUnivalTree(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return dfs(root, root.Val)
}

func dfs(root *TreeNode, univalue int) bool {
    if root == nil {
        return true
    }
    if root.Val != univalue {
        return false
    }
    return dfs(root.Left, univalue) && dfs(root.Right, univalue)
}