/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Time complexity: O(n)
// Space complexity: O(h), h = treedepth

func maxPathSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    ans := math.MinInt32
    maxSum(root, &ans)
    return ans
}

func maxSum(root *TreeNode, ans* int) int {
    if root == nil {
        return 0
    }
    leftSum := max(0, maxSum(root.Left, ans))
    rightSum := max(0, maxSum(root.Right, ans))
    thisSum := leftSum + rightSum + root.Val
    
    *ans = max(thisSum, *ans)
    return max(leftSum,rightSum) + root.Val
}

func max(x int, y int) int {
    if x>y {
        return x
    } else {
        return y
    }
}