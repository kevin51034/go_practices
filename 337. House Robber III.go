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

func rob(root *TreeNode) int {
    return max(dfs(root))
}

func dfs(root *TreeNode) (a, b int) {
    if root == nil {
        return 0, 0
    }
    left0, left1 := dfs(root.Left)
	right0, right1 := dfs(root.Right)
	// doesn't choose current root
	tmp0 := max(left0, left1) + max(right0, right1)
	// choose current root
    tmp1 := root.Val + left0 + right0
    return tmp0, tmp1
}

func max(x, y int) int {
    if x>y {
        return x
    }
    return y
}