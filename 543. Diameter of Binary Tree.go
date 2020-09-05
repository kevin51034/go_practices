/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    result := 0
    dfs(root, &result)
    return result
}

func dfs(root *TreeNode, result *int) int {
    if root == nil {
        return 0
    }
    leftLength := dfs(root.Left, result)
    rightLength := dfs(root.Right, result)
    *result = max(leftLength + rightLength, *result)
    return max(leftLength, rightLength) + 1
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}