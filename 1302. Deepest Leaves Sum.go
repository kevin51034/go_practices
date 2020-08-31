/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// solution 1
// level order traversal
// Time complexity: O(n)
// Space complexity: O(n)

func deepestLeavesSum(root *TreeNode) int {
    sum := 0
    if root == nil {
        return sum
    }
    queue := []*TreeNode{root}
    for len(queue)>0 {
        length := len(queue)
        sum = 0
        for i:=0; i<length; i++ {
            node := queue[0]
            queue = queue[1:]
            sum += node.Val
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }
    return sum
}

// solutnion 2
// dfs
// Time complexity: O(n)
// Space complexity: O(n) // n for recursion stack size
func deepestLeavesSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    maxlevel, sum := 0, 0
    dfs(root, 0, &maxlevel, &sum)
    return sum
}

func dfs(root *TreeNode, curlevel int, maxlevel, sum *int) {
    if root == nil {
        return
    }
    if curlevel > *maxlevel {
        *maxlevel = curlevel
        *sum = root.Val
    } else if curlevel == *maxlevel {
        *sum += root.Val
    }
    dfs(root.Left, curlevel+1, maxlevel, sum)
    dfs(root.Right, curlevel+1, maxlevel, sum)
    return
}