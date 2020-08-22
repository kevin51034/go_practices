/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(preorder []int, inorder []int) *TreeNode {
    inorderIndex := make(map[int]int, len(inorder))
    for i:=0; i<len(inorder); i++ {
        inorderIndex[inorder[i]] = i
    }
    return buildTreeDFS(0, 0, len(preorder)-1, preorder, inorderIndex)
}

func buildTreeDFS(inStart, preStart, preEnd int, preorder []int, inorderIndex map[int]int) *TreeNode {
    if preStart > preEnd {
        return nil
    }
    root := &TreeNode{Val:preorder[preStart]}
    rootIndex := inorderIndex[preorder[preStart]]
    leftLength := rootIndex - inStart
    root.Left = buildTreeDFS(inStart, preStart+1, preStart+leftLength, preorder, inorderIndex)
    root.Right = buildTreeDFS(rootIndex+1, preStart+leftLength+1, preEnd, preorder, inorderIndex)
    return root
}