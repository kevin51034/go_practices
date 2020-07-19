/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// solution 1
// Time complexity: O(nlogn)
// = O(n) + 2*O(n/2) + 4*O(n/4) + ...
// ** is not O(n), but it can be optimized to O(n) **

// Space complexity: O(n)
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    leftDepth := depth(root.Left)
    rightDepth := depth(root.Right)
    
    return abs(leftDepth, rightDepth) <=1 && isBalanced(root.Left) && isBalanced(root.Right)
}


func depth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    return max(depth(root.Left), depth(root.Right)) +1
}

func max(x int, y int) int {
    if x>y {
        return x
    } else {
        return y
    }
}

func abs(x int, y int) int {
    if x>y {
        return x-y
    } else {
        return y-x
    }
}


// solution2
// Time complexity: O(n)
// Space complexity: O(n)

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    balance := true
    depth(root, &balance)
    
    return balance
}

func depth(root *TreeNode, balance * bool) int{
    if root == nil {
        return 0
    }
    leftDepth := depth(root.Left, balance)
    rightDepth := depth(root.Right, balance)
    if abs(leftDepth, rightDepth)>1 {
        *balance = false
        return -1
    }
    
    return max(leftDepth, rightDepth) + 1
}

func abs(x int, y int) int {
    if x>y {
        return x-y
    } else {
        return y-x
    }
}
func max(x int, y int) int {
    if x>y {
        return x
    } else {
        return y
    }
}