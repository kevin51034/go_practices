/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// solution 1
// dfs

func hasPathSum(root *TreeNode, sum int) bool {
    resultSum := make([]int, 0)
    dfsSum(root, 0, &resultSum)
    for _,v := range resultSum {
        if v == sum {
            return true
        }
    }
    return false
}

func dfsSum(root *TreeNode, prevSum int, resultSum *[]int) {
    if root == nil {
        return
    }
    curSum := prevSum + root.Val
    if root.Left == nil && root.Right == nil {
        *resultSum = append(*resultSum, curSum)
        return
    }
    dfsSum(root.Left, curSum, resultSum)
    dfsSum(root.Right, curSum, resultSum)
}

// solution 2
// recursive

func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    if root.Left == nil && root.Right == nil {
        return root.Val == sum
    }
    return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}