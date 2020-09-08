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
func kthSmallest(root *TreeNode, k int) int {
    sortList := make([]int, 0)
    inorder(root, &sortList)
    return sortList[k-1]
}

func inorder(root *TreeNode, sortList *[]int) {
    if root == nil {
        return
    }
    inorder(root.Left, sortList)
    *sortList = append(*sortList, root.Val)
    inorder(root.Right, sortList)
    return
}

// optimize
// Time complexity: O(n)
// Space complexity: O(1) // O(n) for stack
// count need to use pointer or answer would be wrong
func kthSmallest(root *TreeNode, k int) int {
    if root == nil {
        return 0
    }
    result := 0
    count := 0
    inorder(root, k, &count, &result)
    return result
}

func inorder(root *TreeNode, k int, count, result *int) {
    if root == nil {
        return
    }
    inorder(root.Left, k, count, result)
    *count++
    if *count == k {
        *result = root.Val
        return
    }
    inorder(root.Right, k, count, result)
}