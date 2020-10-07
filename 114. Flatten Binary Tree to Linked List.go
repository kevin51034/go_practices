/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// iterative
// time complexity: O(n)
// space complexity: O(n)

func flatten(root *TreeNode)  {
    list := make([]*TreeNode, 0)
    preorder(root, &list)
    cur := root
    for i:=0; i<len(list)-1; i++ {
        cur.Left = nil
        cur.Right = &TreeNode{Val:list[i+1].Val}
        cur = cur.Right
    }
    return
}

func preorder(root *TreeNode, list *[]*TreeNode) {
    if root == nil {
        return
    }
    *list = append(*list, root)
    preorder(root.Left, list)
    preorder(root.Right, list)
    return
}

// recursive
// time complexity: O(n)
// space complexity: O(1)
// prev must use double pointer

func flatten(root *TreeNode)  {
    var prev *TreeNode
    helper(root, &prev)
}

func helper(root *TreeNode, prev **TreeNode) {
    if root == nil {
        return
    }
    helper(root.Right, prev)
    helper(root.Left, prev)
    root.Right = *prev
    root.Left = nil
    *prev = root
}