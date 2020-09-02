/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Time complexity: O(n)
// Space complexity: O(n) / output can be O(n^2)
func binaryTreePaths(root *TreeNode) []string {
    result := make([]string, 0)
    if root == nil {
        return result
    }
    dfs(root, &result, "")
    return result
}

func dfs(root *TreeNode, result *[]string, curPath string) {
    if root == nil {
        return
    }
    if len(curPath) == 0 {
        curPath += strconv.Itoa(root.Val)
    } else {
        curPath += "->" + strconv.Itoa(root.Val)
    }
    if root.Left == nil && root.Right == nil {
		*result = append(*result, curPath)
		return
    }
    dfs(root.Left, result, curPath)
    dfs(root.Right, result, curPath)
}