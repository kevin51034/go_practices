/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Time complexity: O(n1 + n2)
// Space complexity: O(n1 + n2)

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    leaf1 := make([]int, 0)
    leaf2 := make([]int, 0)
    dfs(root1, &leaf1)
    dfs(root2, &leaf2)
    if len(leaf1) != len(leaf2) {
        return false
    }
    for i:=0; i<len(leaf1); i++ {
        if leaf1[i] != leaf2[i] {
            return false
        }
    }
    return true
}

func dfs(root *TreeNode, leaf *[]int) {
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil {
        *leaf = append(*leaf, root.Val)
    }
    dfs(root.Left, leaf)
    dfs(root.Right, leaf)
}
