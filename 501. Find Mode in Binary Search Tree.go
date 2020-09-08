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
func findMode(root *TreeNode) []int {
    result := make([]int, 0)
    count := make(map[int]int)
    dfs(root, &count)
    max := 0
    for key,val := range count {
        if val > max {
            result = []int{key}
            max = val
        } else if val == max {
            result = append(result, key)
        }
    }
    return result
}

func dfs(root *TreeNode, count *map[int]int) {
    if root == nil {
        return
    }
    (*count)[root.Val]++
    dfs(root.Left, count)
    dfs(root.Right, count)
}

// optimize
// Time complexity: O(n)
// Space complexity: O(1) // O(n) for stack
func findMode(root *TreeNode) []int {
    result := []int{}
    if root == nil {
        return result
    }
    count := 0
    maxCount := 0
    prevVal := math.MinInt64
    inorder(root, &prevVal, &count, &maxCount, &result)
    return result
}

func inorder(root *TreeNode, prevVal, count, maxCount *int, result *[]int) {
    if root == nil {
        return
    }
    inorder(root.Left, prevVal, count, maxCount, result)
    visit(root.Val, prevVal, count, maxCount, result)
    *prevVal = root.Val
    inorder(root.Right, prevVal, count, maxCount, result)
}

func visit(val int, prevVal, count, maxCount *int, result *[]int) {
    if val == *prevVal {
        *count++
    } else {
        *count = 1
    }
    
    if *count > *maxCount {
        *maxCount = *count
        *result = []int{val}
    } else if *count == *maxCount {
        *result = append(*result, val)
    }
}