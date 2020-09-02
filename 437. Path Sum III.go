/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// Time complexity: O(n^2)
// Space complexity: O(n)
func pathSum(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }
    result := 0
    result = dfs(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
    return result
}

func dfs(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }
    result := 0
    if root.Val == sum {
        result++
    }
    result += dfs(root.Left, sum-root.Val)
    result += dfs(root.Right, sum-root.Val)
    return result
}

// solution 2
// recursive
// Time complexity: O(n)
// Space complexity: O(h)

// We can calculate any sum of partial paths from A to B by (sum from root to B) - (sum from root to A).
// So just by storing all sums from root to each node, we can calculate any sum of possible path in O(1).

// sumMap might greater than 1 because if there is a node with value 0, 
// there would be multiple ans in this tree(with and without this node, or even there might be negative node)
// sumMap need to subtract 1 when returning or other subtree would affect current subtree

func pathSum(root *TreeNode, sum int) int {
    return dfs(root, sum, 0, map[int]int{0:1})
}

func dfs(root *TreeNode, sum, curSum int, sumMap map[int]int) int {
    if root == nil {
        return 0
    }
    curSum += root.Val
    count := sumMap[curSum - sum]
    sumMap[curSum]++
    count += dfs(root.Left, sum, curSum, sumMap)
    count += dfs(root.Right, sum, curSum, sumMap)
    sumMap[curSum]--
    return count
}

