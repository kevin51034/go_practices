/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pathSum(root *TreeNode, sum int) [][]int {
    result := make([][]int, 0)
    if root == nil {
        return result
    }
    stack := []int{}
    dfs(root, sum, &result, stack)
    return result
}

func dfs(root *TreeNode, sum int, result *[][]int, stack []int) {
    if root == nil {
        return 
    }
    stack = append(stack, root.Val)
    if root.Left == nil && root.Right == nil && root.Val == sum {
		// **stack need to use copy to result or the value would be change**
        tmp := make([]int, len(stack))
        copy(tmp, stack)
        *result = append(*result, tmp)
        stack = stack[:len(stack)-1]
    }
    dfs(root.Left, sum-root.Val, result, stack)
    dfs(root.Right, sum-root.Val, result, stack)
    return
}