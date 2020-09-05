/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func longestUnivaluePath(root *TreeNode) int {
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
    l, r := 0, 0
    if root.Left != nil && root.Val == root.Left.Val {
        l = leftLength+1
    }
    if root.Right != nil && root.Val == root.Right.Val {
        r = rightLength+1
    }
    *result = max(*result, l+r)
    return max(l,r)
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}