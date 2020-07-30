/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return root
    }
    if root.Val < key {
        root.Right = deleteNode(root.Right, key)
    } else if root.Val > key {
        root.Left = deleteNode(root.Left, key)
    } else if root.Val == key {
        if root.Left == nil {
            return root.Right
        } else if root.Right == nil {
            return root.Left
        } else {
            cur := root.Right
            for cur.Left!=nil {
                cur = cur.Left
            }
            cur.Left = root.Left
            return root.Right
        }
    }
    return root
}