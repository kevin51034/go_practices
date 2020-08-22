/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(inorder []int, postorder []int) *TreeNode {
    indexInorder := make(map[int]int, len(inorder))
    for i:=0; i<len(inorder); i++ {
        indexInorder[inorder[i]] = i
    }
    return buildTreeDFS(0, 0, len(postorder)-1, postorder, indexInorder)
}

func buildTreeDFS(inStart, postStart, postEnd int, postorder []int, indexInorder map[int]int) *TreeNode {
    if postStart > postEnd {
        return nil
    }
    root := &TreeNode{Val:postorder[postEnd]}
    rootIndex := indexInorder[postorder[postEnd]]
    leftLength := rootIndex - inStart
    root.Left = buildTreeDFS(inStart, postStart, postStart+leftLength-1, postorder, indexInorder)
    root.Right = buildTreeDFS(rootIndex+1, postStart+leftLength, postEnd-1, postorder, indexInorder)
    return root
}