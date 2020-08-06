/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func generateTrees(n int) []*TreeNode {
    if n<1 {
        return nil
    }
    return generate(1,n)
}

func generate(start, end int) []*TreeNode {
    if start > end {
        return []*TreeNode{nil}
    }
    result := make([]*TreeNode, 0)
    
    for i:=start; i<=end; i++ {
        leftTree := generate(start, i-1)
        rightTree := generate(i+1, end)
        for _,ltree := range leftTree {
            for _, rtree := range rightTree {
                root := &TreeNode{Val:i, Left:ltree, Right:rtree}
                result = append(result, root)
            }
        }
    }
    return result
}