/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// solution1
// Time complexity: O(n)
// Space complexity: O(n)

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return validation(root, math.MinInt64, math.MaxInt64)
}

func validation(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	return root.Val > min && root.Val < max && validation(root.Left, min, root.Val) && validation(root.Right, root.Val, max)
}

// solution2
// list all value in order, and check if is increasing.
// because in BST numbers is increasing if print in inorder.

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	list := make([]int, 0)
	inorder(root, &list)
	for i := 0; i < len(list)-1; i++ {
		if list[i] >= list[i+1] {
			return false
		}
	}
	return true
}

func inorder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, list)
	*list = append(*list, root.Val)
	inorder(root.Right, list)

}