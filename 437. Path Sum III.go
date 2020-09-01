/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// Time complexity: O(n^2)
// Space complexity: O(n)
func pathSum(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }
    result := 0
    result = dfs(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
    return result
}

func dfs(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }
    result := 0
    if root.Val == sum {
        result++
    }
    result += dfs(root.Left, sum-root.Val)
    result += dfs(root.Right, sum-root.Val)
    return result
}