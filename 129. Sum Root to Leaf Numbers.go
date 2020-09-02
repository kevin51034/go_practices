/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func sumNumbers(root *TreeNode) int {
    if root == nil {
        return 0
    }
    list := make([]int, 0)
    dfs(root, &list, 0)
    result := 0
    for _,v := range list {
        result += v
    }
    return result
}

func dfs(root *TreeNode, list *[]int, sum int) {
    if root == nil {
        return
    }
    curSum := sum*10 + root.Val
    if root.Left == nil && root.Right == nil {
        *list = append(*list, curSum)
    }
    dfs(root.Left, list, curSum)
    dfs(root.Right, list, curSum)
}


// optimize
func sumNumbers(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return dfs(root, 0)
}

func dfs(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }
    sum = sum*10 + root.Val
    if root.Left == nil && root.Right == nil {
        return sum
    }
    return dfs(root.Left, sum) + dfs(root.Right, sum)
}