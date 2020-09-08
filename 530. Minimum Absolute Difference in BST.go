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
func getMinimumDifference(root *TreeNode) int {
    sortList := make([]int, 0)
    inorder(root, &sortList)
    min := math.MaxInt64
    for i:=0; i<len(sortList)-1; i++ {
        if sortList[i+1] - sortList[i] < min {
            min = sortList[i+1] - sortList[i]
        }
    }
    return min
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
// prev need to use pointer or answer would be wrong
func getMinimumDifference(root *TreeNode) int {
    min := math.MaxInt64
    prev := -1
    inorder(root, &min, &prev)
    return min
}

func inorder(root *TreeNode, min *int, prev *int) {
    if root == nil {
        return
    }
    inorder(root.Left, min, prev)
    if *prev >= 0 {
        if root.Val - *prev < *min {
            *min = root.Val - *prev
        }
    }
    *prev = root.Val
    inorder(root.Right, min, prev)
    return
}