/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findFrequentTreeSum(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    sumMap := make(map[int]int, 0)
    dfs(root, 0, sumMap)
    max := 0
    result := make([]int, 0)
    for key,value := range sumMap {
        if value > max {
            max = value
            result = []int{key}
        } else if value == max{
            result = append(result, key)
        }
    }
    return result
}

func dfs(root *TreeNode, curSum int, sumMap map[int]int) int {
    if root == nil {
        return 0
    }
    curSum = root.Val + dfs(root.Left, curSum, sumMap) + dfs(root.Right, curSum, sumMap)
    sumMap[curSum]++
    return curSum
}